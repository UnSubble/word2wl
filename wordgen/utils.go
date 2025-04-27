package wordgen

func Unique(words []string) []string {
	seen := make(map[string]struct{})
	var result []string

	for _, w := range words {
		if _, ok := seen[w]; !ok {
			seen[w] = struct{}{}
			result = append(result, w)
		}
	}

	return result
}
