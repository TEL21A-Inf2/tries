package elements

import "fmt"

func ExampleTrie() {
	e1 := NewElement()
	e1.isRoot = true

	e1.Add("a", "a")
	fmt.Println(e1.GetValues())

	e1.Add("b", "b")
	fmt.Println(e1.GetValues())

	e1.Add("c", "c")
	fmt.Println(e1.GetValues())

	e1.Add("ab", "ab")
	fmt.Println(e1.GetValues())

	e1.Add("abc", "abc")
	fmt.Println(e1.GetValues())

	e1.Add("bd", "bd")
	fmt.Println(e1.GetValues())

	e1.Add("bc", "bc")
	fmt.Println(e1.GetValues())

	// Output:
	// [a]
	// [a b]
	// [a b c]
	// [a ab b c]
	// [a ab abc b c]
	// [a ab abc b bd c]
	// [a ab abc b bc bd c]
}
