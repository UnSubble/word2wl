package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"github.com/unsubble/word2wl/wordgen"
)

const PROJECT_NAME = "word2wl"
const VERSION = "1.0.1"
const THRESHOLD = 10 * 1024 * 1024 // 10MB threshold

func initCobra() *cobra.Command {
	var (
		datasetPath    string
		keyword        string
		level          int
		threads        int
		batchSize      int
		outputFile     string
		recursive      bool
		recursiveLevel int
		verbose        bool
		specialChars   string
	)

	rootCmd := &cobra.Command{
		Use:     PROJECT_NAME,
		Short:   "Wordlist Generator: Create word variations based on a dataset and keyword.",
		Version: VERSION,
		RunE: func(cmd *cobra.Command, args []string) error {

			if datasetPath == "" || keyword == "" {
				return fmt.Errorf("dataset and keyword are required")
			}

			data, err := os.ReadFile(datasetPath)
			if err != nil {
				return fmt.Errorf("failed to read dataset: %v", err)
			}

			initialWords := splitLines(string(data))
			if len(initialWords) == 0 {
				return fmt.Errorf("dataset is empty")
			}

			var writer io.Writer = os.Stdout
			if outputFile != "" {
				file, err := os.Open(outputFile)
				if err != nil && !os.IsNotExist(err) {
					return fmt.Errorf("an error occurred while opening the output file %s: %v", outputFile, err)
				}

				if err != nil && os.IsNotExist(err) {
					file, err = os.Create(outputFile)
					if err != nil {
						return fmt.Errorf("failed to create output file: %v", err)
					}
				}
				defer file.Close()

				fileInfo, err := os.Stat(outputFile)
				if err != nil {
					return fmt.Errorf("failed to stat output file: %v", err)
				}

				if fileInfo.Size() > THRESHOLD {
					if verbose {
						fmt.Println("Output file exceeds 10MB. Compressing to .gz format.")
					}

					f, err := os.Create(outputFile + ".gz")
					if err != nil {
						return fmt.Errorf("failed to create compressed output file: %v", err)
					}
					defer f.Close()

					writer = gzip.NewWriter(f)
				} else {
					f, err := os.Create(outputFile)
					if err != nil {
						return fmt.Errorf("failed to create output file: %v", err)
					}
					defer f.Close()
					writer = f
				}
			}

			words := initialWords
			cycles := 1
			if recursive {
				if recursiveLevel > 0 {
					cycles = recursiveLevel
				} else {
					cycles = level
				}
			}

			for cycle := 0; cycle < cycles; cycle++ {
				var wg sync.WaitGroup
				wordChan := make(chan []string)
				resultChan := make(chan []string)

				if verbose {
					fmt.Printf("Starting cycle %d of %d...\n", cycle+1, cycles)
				}

				bar := progressbar.NewOptions(len(words),
					progressbar.OptionSetWriter(os.Stderr),
					progressbar.OptionSetDescription(fmt.Sprintf("Cycle %d", cycle+1)),
					progressbar.OptionSetWidth(15),
				)

				specialCharMap := make(map[rune]struct{})
				for _, ch := range specialChars {
					specialCharMap[ch] = struct{}{}
				}

				for i := 0; i < threads; i++ {
					wg.Add(1)
					go func() {
						defer wg.Done()
						for batch := range wordChan {
							generator := wordgen.NewWordGenerator(batch, keyword, level, specialCharMap)
							results, err := generator.Generate()
							if err == nil {
								resultChan <- results
							}
						}
					}()
				}

				go func() {
					for i := 0; i < len(words); i += batchSize {
						end := i + batchSize
						if end > len(words) {
							end = len(words)
						}
						wordChan <- words[i:end]
					}
					close(wordChan)
				}()

				go func() {
					wg.Wait()
					close(resultChan)
				}()

				var nextWords []string

				for results := range resultChan {
					for _, result := range results {
						fmt.Fprintln(writer, result)
						bar.Add(1)
						if recursive {
							nextWords = append(nextWords, result)
						}
					}
				}

				if verbose {
					fmt.Printf("Cycle %d complete.\n", cycle+1)
				}

				if recursive {
					words = nextWords
					if len(words) == 0 {
						break
					}
				}
			}

			if gzipWriter, ok := writer.(*gzip.Writer); ok {
				gzipWriter.Flush()
				gzipWriter.Close()
			}

			return nil
		},
	}

	rootCmd.Flags().StringVarP(&datasetPath, "dataset", "d", "", "Path to the dataset file (e.g., rockyou.txt)")
	rootCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "Keyword to inject into the dataset")
	rootCmd.Flags().IntVarP(&level, "level", "l", 1, "Mutation power level (1 = basic, 5 = advanced)")
	rootCmd.Flags().IntVarP(&threads, "threads", "t", 4, "Number of concurrent threads")
	rootCmd.Flags().IntVarP(&batchSize, "batch-size", "b", 100, "Number of words per batch")
	rootCmd.Flags().StringVarP(&outputFile, "output-file", "o", "", "Output file path (default: stdout)")
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recursively mutate generated words")
	rootCmd.Flags().IntVarP(&recursiveLevel, "recursive-level", "R", 0, "Number of recursive mutation cycles (default: same as --level)")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.Flags().StringVarP(&specialChars, "special", "s", "", "Special characters that should be reserved and tokenized separately")
	rootCmd.Flags().BoolP("version", "V", false, "Show version")
	rootCmd.Flags().BoolP("help", "h", false, "Help for word2wl")

	return rootCmd
}

func splitLines(s string) []string {
	var lines []string
	current := ""
	for _, r := range s {
		if r == '\n' || r == '\r' {
			if current != "" {
				lines = append(lines, current)
				current = ""
			}
		} else {
			current += string(r)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func main() {
	cmd := initCobra()
	if err := cmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
