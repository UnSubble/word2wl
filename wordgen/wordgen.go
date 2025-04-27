package wordgen

import (
	"fmt"

	"github.com/unsubble/word2wl/wordgen/mutators"
)

type MutatorFunc func(word string) (string, error)

type WordGenerator struct {
	Dataset       []string
	Keyword       string
	Mutators      []MutatorFunc
	ReservedChars map[rune]struct{}
}

func NewWordGenerator(dataset []string, keyword string, power int, reservedChars map[rune]struct{}) *WordGenerator {
	return &WordGenerator{
		Dataset:       dataset,
		Keyword:       keyword,
		Mutators:      GetMutators(power),
		ReservedChars: reservedChars,
	}
}

func GetMutators(level int) []MutatorFunc {
	switch level {
	case 1:
		return []MutatorFunc{
			mutators.Capitalize,
			mutators.ToUpper,
			mutators.ToLower,
		}
	case 2:
		return []MutatorFunc{
			mutators.Capitalize,
			mutators.ToUpper,
			mutators.ToLower,
			mutators.ReverseWord,
			mutators.LeetSpeak,
		}
	case 3:
		return []MutatorFunc{
			mutators.Capitalize,
			mutators.ToUpper,
			mutators.ToLower,
			mutators.ReverseWord,
			mutators.LeetSpeak,
			mutators.RandomCase,
			mutators.DuplicateWord,
			mutators.AddRandomNumber,
		}
	case 4:
		return []MutatorFunc{
			mutators.Capitalize,
			mutators.ToUpper,
			mutators.ToLower,
			mutators.ReverseWord,
			mutators.LeetSpeak,
			mutators.RandomCase,
			mutators.DuplicateWord,
			mutators.AddRandomNumber,
			mutators.ShuffleLetters,
			mutators.InsertRandomSymbol,
			mutators.StretchWord,
		}
	case 5:
		return []MutatorFunc{
			mutators.Capitalize,
			mutators.ToUpper,
			mutators.ToLower,
			mutators.ReverseWord,
			mutators.LeetSpeak,
			mutators.RandomCase,
			mutators.DuplicateWord,
			mutators.AddRandomNumber,
			mutators.ShuffleLetters,
			mutators.InsertRandomSymbol,
			mutators.StretchWord,
			mutators.AlternateCase,
			mutators.RemoveVowels,
			mutators.MemeCase,
		}
	default:
		return []MutatorFunc{
			mutators.Capitalize,
			mutators.ToUpper,
			mutators.ToLower,
		}
	}
}

func (wg *WordGenerator) Generate() ([]string, error) {
	var results []string

	for _, word := range wg.Dataset {
		tokens := Tokenize(word, wg.ReservedChars)
		pattern := Pattern{Tokens: tokens}

		newWords := pattern.ApplyKeyword(wg.Keyword)

		results = append(results, newWords...)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results generated")
	}

	var mutated []string
	for _, word := range results {
		for _, m := range wg.Mutators {
			if newWord, err := m(word); err == nil && newWord != "" {
				mutated = append(mutated, newWord)
			}
		}
	}

	allWords := append(results, mutated...)

	finalWords := Unique(allWords)

	return finalWords, nil
}
