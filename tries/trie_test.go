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

func ExampleTrie_Complete() {
	e1 := NewTrie()

	e1.Add("Haus")
	e1.Add("Halle")
	e1.Add("Hand")
	e1.Add("Halt")
	e1.Add("Haufen")

	fmt.Println(e1.Complete("H"))
	fmt.Println(e1.Complete("Ha"))
	fmt.Println(e1.Complete("Hal"))
	fmt.Println(e1.Complete("Hau"))
	fmt.Println(e1.Complete("Han"))
	fmt.Println(e1.Complete("Hanf"))

	// Output:
	// [Halle Halt Hand Haufen Haus]
	// [Halle Halt Hand Haufen Haus]
	// [Halle Halt]
	// [Haufen Haus]
	// [Hand]
	// []
}
