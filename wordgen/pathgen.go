package wordgen

import (
	"fmt"
	"os"
	"strings"

	"github.com/unsubble/word2wl/wordgen/common"
	"github.com/unsubble/word2wl/wordgen/mutator"
)

type PathGenerator struct {
	Dataset []string
	Keyword string
	mutator []common.PathMutatorFunc
}

func NewPathGenerator(dataset []string, keyword string, power int) *PathGenerator {
	return &PathGenerator{
		Dataset: dataset,
		Keyword: keyword,
		mutator: getMutatorForPathGenerator(power),
	}
}

func getMutatorForPathGenerator(level int) []common.PathMutatorFunc {
	switch level {
	case 1:
		return []common.PathMutatorFunc{
			mutator.NormalizePaths,
			mutator.ShufflePaths,
		}
	case 2:
		return []common.PathMutatorFunc{
			mutator.NormalizePaths,
			mutator.ShufflePaths,
			mutator.RandomizeCasing,
			mutator.AddTrickyPaths,
		}
	case 3:
		return []common.PathMutatorFunc{
			mutator.NormalizePaths,
			mutator.ShufflePaths,
			mutator.RandomizeCasing,
			mutator.AddTrickyPaths,
			mutator.ExtractFileNames,
			mutator.SortPathsByLength,
		}
	case 4:
		return []common.PathMutatorFunc{
			mutator.NormalizePaths,
			mutator.ShufflePaths,
			mutator.RandomizeCasing,
			mutator.AddTrickyPaths,
			mutator.ExtractFileNames,
			mutator.SortPathsByLength,
			mutator.ReversePaths,
		}
	case 5:
		return []common.PathMutatorFunc{
			mutator.NormalizePaths,
			mutator.ShufflePaths,
			mutator.RandomizeCasing,
			mutator.AddTrickyPaths,
			mutator.ExtractFileNames,
			mutator.SortPathsByLength,
			mutator.ReversePaths,
			mutator.DedupePaths,
			mutator.RemoveRandomPath,
		}
	default:
		return []common.PathMutatorFunc{
			mutator.NormalizePaths,
			mutator.ShufflePaths,
		}
	}
}

func (pg *PathGenerator) Generate() ([]string, error) {
	var results []string

	for _, path := range pg.Dataset {
		tokens := TokenizePath(path)

		for i, token := range tokens {
			if len(token.Value) > 0 && i < len(tokens)-1 {
				tokens[i].Value = token.Value + string(os.PathSeparator)
			}
		}

		pattern := Pattern{Tokens: tokens}

		newWords := pattern.ApplyKeyword(pg.Keyword)

		results = append(results, newWords...)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results generated")
	}

	var mutated []string
	for _, path := range results {
		for _, m := range pg.mutator {
			tokens := strings.Split(path, string(os.PathSeparator))
			if newPath, err := m(tokens); err == nil && len(newPath) > 0 {
				mutated = append(mutated, strings.Join(newPath, string(os.PathSeparator)))
			}
		}
	}

	allWords := append(results, mutated...)

	finalWords := common.Unique(allWords)

	return finalWords, nil
}
