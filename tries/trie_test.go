package tries

import "fmt"

func ExampleTrie() {
	e1 := NewTrie()

	e1.Add("a")
	fmt.Println(e1.GetValues())

	e1.Add("b")
	fmt.Println(e1.GetValues())

	e1.Add("c")
	fmt.Println(e1.GetValues())

	e1.Add("ab")
	fmt.Println(e1.GetValues())

	e1.Add("abc")
	fmt.Println(e1.GetValues())

	e1.Add("bd")
	fmt.Println(e1.GetValues())

	e1.Add("bc")
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
