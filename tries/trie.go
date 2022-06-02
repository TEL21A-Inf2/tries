package tries

import (
	"sort"

	"github.com/tel21a-inf2/tries/elements"
)

// Typ für einen Trie.
type Trie struct {
	root *elements.Element
}

// Konstruktor für einen Trie.
func NewTrie() Trie {
	return Trie{elements.NewRoot()}
}

// Fügt einen Wert hinzu.
func (trie *Trie) Add(value string) {
	trie.root.Add(value, value)
}

// Liefert alle Werte.
func (trie Trie) GetValues() []string {
	result := trie.root.GetValues()
	sort.Strings(result)
	return result
}
