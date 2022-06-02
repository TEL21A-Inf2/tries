package elements

import "sort"

// Datentyp für Elemente eines Tries.
type Element struct {
	key      string
	value    string
	children map[string]*Element
	isRoot   bool
}

// Konstruktor für Trie-Elemente.
func NewElement() *Element {
	return &Element{"", "", make(map[string]*Element, 0), false}
}

// Konstruktor für eine Wurzel.
func NewRoot() *Element {
	result := NewElement()
	result.isRoot = true
	return result
}

// Liefert true, falls das Element leer ist.
func (element *Element) IsEmpty() bool {
	return !element.isRoot && element.key == ""
}

// Liefert das Kind mit dem angegebenen Key, falls es existiert.
// Liefert ansonsten nil.
func (element *Element) GetChild(key string) *Element {
	child, exists := element.children[key]
	if exists {
		return child
	}
	return nil
}

// Liefert true, falls das Element ein Kind mit dem angegebenen Key hat.
func (element *Element) HasChild(key string) bool {

	return element.GetChild(key) != nil
}

// Fügt ein neues leeres Kindelement mit dem angegebenen Key ein,
// falls noch keines mit diesem Key existiert.
func (element *Element) tryCreateChild(key string) {
	if !element.HasChild(key) {
		element.children[key] = NewElement()
	}
}

// Liefert true, falls das Element einen Wert gespeichert hat.
func (element *Element) HasValue() bool {
	return !element.isRoot && !element.IsEmpty() && element.value != ""
}

// Fügt einen Wert hinzu.
func (element *Element) Add(key, value string) {
	if key == "" {
		return
	}
	head, tail := key[:1], key[1:]
	if element.isRoot {
		element.tryCreateChild(head)
		element.children[head].Add(key, value)
	} else {
		if element.IsEmpty() {
			element.key = head
		}
		if tail == "" {
			element.value = value
			return
		}
		head = tail[:1]
		element.tryCreateChild(head)
		element.children[head].Add(tail, value)
	}
}

// Liefert alle Werte, die in diesem Baum gespeichert sind.
func (element *Element) GetValues() []string {
	result := make([]string, 0)
	if element.IsEmpty() {
		return result
	}
	if element.HasValue() {
		result = append(result, element.value)
	}
	for _, child := range element.children {
		result = append(result, child.GetValues()...)
	}
	sort.Strings(result)
	return result
}

// Liefert alle Completions für den Key.
func (element *Element) Complete(key string) []string {
	if key == "" {
		return element.GetValues()
	}
	head, tail := key[:1], key[1:]
	if element.isRoot {
		if element.HasChild(head) {
			return element.children[head].Complete(tail)
		}
	} else {
		if element.HasChild(head) {
			return element.children[head].Complete(tail)
		}
	}
	return make([]string, 0)
}
