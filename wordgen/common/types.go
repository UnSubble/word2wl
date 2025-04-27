package common

type WordMutatorFunc func(word string) (string, error)
type PathMutatorFunc func(paths []string) ([]string, error)

type Generator interface {
	Generate() ([]string, error)
}
