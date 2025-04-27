package wordgen

type TokenType int

const (
	Alpha TokenType = iota
	Digit
	Special
	Reserved
)

type TokenNode struct {
	Type  TokenType
	Value string
}

type Pattern struct {
	Tokens []TokenNode
}

func (p *Pattern) ApplyKeyword(keyword string) []string {
	var results []string

	for i := 0; i <= len(p.Tokens); i++ {
		var newWord string

		for j, token := range p.Tokens {
			if j == i {
				newWord += keyword
			}
			newWord += token.Value
		}

		if i == len(p.Tokens) {
			newWord += keyword
		}

		results = append(results, newWord)
	}

	for i, token := range p.Tokens {
		if token.Type == Special {
			var newWord string
			for j, t := range p.Tokens {
				if j == i {
					newWord += keyword
				}
				newWord += t.Value
			}
			results = append(results, newWord)

			newWord = ""
			for j, t := range p.Tokens {
				newWord += t.Value
				if j == i {
					newWord += keyword
				}
			}
			results = append(results, newWord)
		}
	}

	return results
}
