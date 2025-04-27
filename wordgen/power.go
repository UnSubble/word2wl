package wordgen

import (
	"fmt"
	"math/rand"
	"strings"
)

func Capitalize(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	return strings.ToUpper(string(word[0])) + word[1:], nil
}

func ToUpper(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	return strings.ToUpper(word), nil
}

func ToLower(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	return strings.ToLower(word), nil
}

func ReverseWord(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

func LeetSpeak(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	replacer := strings.NewReplacer(
		"a", "@",
		"e", "3",
		"i", "1",
		"o", "0",
		"s", "$",
		"t", "7",
	)
	return replacer.Replace(word), nil
}

func RandomCase(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	var newWord string
	for _, r := range word {
		if rand.Intn(2) == 0 {
			newWord += strings.ToLower(string(r))
		} else {
			newWord += strings.ToUpper(string(r))
		}
	}
	return newWord, nil
}

func DuplicateWord(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	return word + word, nil
}

func AddRandomNumber(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	num := rand.Intn(10)
	return fmt.Sprintf("%s%d", word, num), nil
}

func ShuffleLetters(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	runes := []rune(word)
	n := len(runes)
	for i := range runes {
		j := rand.Intn(n)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes), nil
}

func InsertRandomSymbol(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	symbols := []string{"@", "#", "$", "%", "&", "*", "!"}
	symbol := symbols[rand.Intn(len(symbols))]
	pos := rand.Intn(len(word) + 1)

	return word[:pos] + symbol + word[pos:], nil
}

func StretchWord(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	var newWord strings.Builder
	for _, r := range word {
		repeat := rand.Intn(3) + 1
		for i := 0; i < repeat; i++ {
			newWord.WriteRune(r)
		}
	}
	return newWord.String(), nil
}

func AlternateCase(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	var newWord strings.Builder
	upper := true
	for _, r := range word {
		if upper {
			newWord.WriteString(strings.ToUpper(string(r)))
		} else {
			newWord.WriteString(strings.ToLower(string(r)))
		}
		upper = !upper
	}
	return newWord.String(), nil
}

func RemoveVowels(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	var newWord strings.Builder
	vowels := "aeiouAEIOU"
	for _, r := range word {
		if !strings.ContainsRune(vowels, r) {
			newWord.WriteRune(r)
		}
	}
	return newWord.String(), nil
}

func MemeCase(word string) (string, error) {
	if len(word) == 0 {
		return "", fmt.Errorf("word is empty")
	}
	var newWord strings.Builder
	for _, r := range word {
		if rand.Intn(2) == 0 {
			newWord.WriteString(strings.ToLower(string(r)))
		} else {
			newWord.WriteString(strings.ToUpper(string(r)))
		}
	}
	return newWord.String(), nil
}
