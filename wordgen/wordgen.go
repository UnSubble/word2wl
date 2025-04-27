package wordgen

import "fmt"

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
			Capitalize,
			ToUpper,
			ToLower,
		}
	case 2:
		return []MutatorFunc{
			Capitalize,
			ToUpper,
			ToLower,
			ReverseWord,
			LeetSpeak,
		}
	case 3:
		return []MutatorFunc{
			Capitalize,
			ToUpper,
			ToLower,
			ReverseWord,
			LeetSpeak,
			RandomCase,
			DuplicateWord,
			AddRandomNumber,
		}
	case 4:
		return []MutatorFunc{
			Capitalize,
			ToUpper,
			ToLower,
			ReverseWord,
			LeetSpeak,
			RandomCase,
			DuplicateWord,
			AddRandomNumber,
			ShuffleLetters,
			InsertRandomSymbol,
			StretchWord,
		}
	case 5:
		return []MutatorFunc{
			Capitalize,
			ToUpper,
			ToLower,
			ReverseWord,
			LeetSpeak,
			RandomCase,
			DuplicateWord,
			AddRandomNumber,
			ShuffleLetters,
			InsertRandomSymbol,
			StretchWord,
			AlternateCase,
			RemoveVowels,
			MemeCase,
		}
	default:
		return []MutatorFunc{
			Capitalize,
			ToUpper,
			ToLower,
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
