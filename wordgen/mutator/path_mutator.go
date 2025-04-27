package mutator

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

func ShufflePaths(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}
	shuffled := make([]string, len(paths))
	copy(shuffled, paths)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled, nil
}

func ReversePaths(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}
	reversed := make([]string, len(paths))
	for i, path := range paths {
		reversed[len(paths)-1-i] = path
	}
	return reversed, nil
}

func RemoveRandomPath(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}
	index := rand.Intn(len(paths))
	return append(paths[:index], paths[index+1:]...), nil
}

func AddTrickyPaths(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}
	newPaths := make([]string, 0)
	times := rand.Intn(10) + 1
	for i := 0; i < times; i++ {
		newPaths = append(newPaths, "..")
	}
	return append(newPaths, paths...), nil
}

func NormalizePaths(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}
	normalized := make([]string, len(paths))
	for i, path := range paths {
		normalized[i] = filepath.Clean(path)
	}
	return normalized, nil
}

func DedupePaths(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}
	seen := make(map[string]struct{})
	deduped := make([]string, 0, len(paths))
	for _, path := range paths {
		if _, exists := seen[path]; !exists {
			seen[path] = struct{}{}
			deduped = append(deduped, path)
		}
	}
	return deduped, nil
}

func SortPathsByLength(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}
	sorted := make([]string, len(paths))
	copy(sorted, paths)
	sort.Slice(sorted, func(i, j int) bool {
		return len(sorted[i]) < len(sorted[j])
	})
	return sorted, nil
}

func ExtractFileNames(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}
	fileNames := make([]string, len(paths))
	for i, path := range paths {
		fileNames[i] = filepath.Base(path)
	}
	return fileNames, nil
}

func RandomizeCasing(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return nil, fmt.Errorf("paths list is empty")
	}

	randomized := make([]string, len(paths))
	for i, path := range paths {
		var sb strings.Builder
		for _, char := range path {
			if rand.Intn(2) == 0 && unicode.IsLetter(char) {
				sb.WriteRune(unicode.ToUpper(char))
			} else {
				sb.WriteRune(char)
			}
		}
		randomized[i] = sb.String()
	}
	return randomized, nil
}
