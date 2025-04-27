package wordgen

import (
	"unicode"
)

func Tokenize(word string, reservedRunes map[rune]struct{}) []TokenNode {
	var tokens []TokenNode
	var currentToken TokenNode

	for i, c := range word {
		var tType TokenType

		if _, ok := reservedRunes[c]; ok {
			tType = Reserved
		} else if unicode.IsLetter(c) {
			tType = Alpha
		} else if unicode.IsDigit(c) {
			tType = Digit
		} else {
			tType = Special
		}

		if i == 0 {
			currentToken = TokenNode{Type: tType, Value: string(c)}
			continue
		}

		if currentToken.Type == tType {
			currentToken.Value += string(c)
		} else {
			tokens = append(tokens, currentToken)
			currentToken = TokenNode{Type: tType, Value: string(c)}
		}
	}

	if currentToken.Value != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}
