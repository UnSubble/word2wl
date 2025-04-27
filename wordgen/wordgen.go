package wordgen

import (
	"fmt"

	"github.com/unsubble/word2wl/wordgen/common"
	"github.com/unsubble/word2wl/wordgen/mutator"
)

type WordGenerator struct {
	Dataset       []string
	Keyword       string
	mutator       []common.WordMutatorFunc
	ReservedChars map[rune]struct{}
}

func NewWordGenerator(dataset []string, keyword string, power int, reservedChars map[rune]struct{}) *WordGenerator {
	return &WordGenerator{
		Dataset:       dataset,
		Keyword:       keyword,
		mutator:       getMutatorForWordGenerator(power),
		ReservedChars: reservedChars,
	}
}

func getMutatorForWordGenerator(level int) []common.WordMutatorFunc {
	switch level {
	case 1:
		return []common.WordMutatorFunc{
			mutator.Capitalize,
			mutator.ToUpper,
			mutator.ToLower,
		}
	case 2:
		return []common.WordMutatorFunc{
			mutator.Capitalize,
			mutator.ToUpper,
			mutator.ToLower,
			mutator.ReverseWord,
			mutator.LeetSpeak,
		}
	case 3:
		return []common.WordMutatorFunc{
			mutator.Capitalize,
			mutator.ToUpper,
			mutator.ToLower,
			mutator.ReverseWord,
			mutator.LeetSpeak,
			mutator.RandomCase,
			mutator.DuplicateWord,
			mutator.AddRandomNumber,
		}
	case 4:
		return []common.WordMutatorFunc{
			mutator.Capitalize,
			mutator.ToUpper,
			mutator.ToLower,
			mutator.ReverseWord,
			mutator.LeetSpeak,
			mutator.RandomCase,
			mutator.DuplicateWord,
			mutator.AddRandomNumber,
			mutator.ShuffleLetters,
			mutator.InsertRandomSymbol,
			mutator.StretchWord,
		}
	case 5:
		return []common.WordMutatorFunc{
			mutator.Capitalize,
			mutator.ToUpper,
			mutator.ToLower,
			mutator.ReverseWord,
			mutator.LeetSpeak,
			mutator.RandomCase,
			mutator.DuplicateWord,
			mutator.AddRandomNumber,
			mutator.ShuffleLetters,
			mutator.InsertRandomSymbol,
			mutator.StretchWord,
			mutator.AlternateCase,
			mutator.RemoveVowels,
			mutator.MemeCase,
		}
	default:
		return []common.WordMutatorFunc{
			mutator.Capitalize,
			mutator.ToUpper,
			mutator.ToLower,
		}
	}
}

func (wg *WordGenerator) Generate() ([]string, error) {
	var results []string

	for _, word := range wg.Dataset {
		tokens := TokenizeWord(word, wg.ReservedChars)
		pattern := Pattern{Tokens: tokens}

		newWords := pattern.ApplyKeyword(wg.Keyword)

		results = append(results, newWords...)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results generated")
	}

	var mutated []string
	for _, word := range results {
		for _, m := range wg.mutator {
			if newWord, err := m(word); err == nil && newWord != "" {
				mutated = append(mutated, newWord)
			}
		}
	}

	allWords := append(results, mutated...)

	finalWords := common.Unique(allWords)

	return finalWords, nil
}
