package counter_test

import (
	"fmt"
	"goBible/chapter7/exam7-1"
)

func ExampleByteCounter() {
	var c counter.ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	// Output:
	// 5
	// 12
}

func ExampleWordCounter() {
	var c counter.WordCounter
	c.Write([]byte("Hello, Word!"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "Hi, there %s", name)
	fmt.Println(c)

	// Output:
	// 2
	// 3
}

func ExampleLineCounter() {
	var c counter.LineCounter
	c.Write([]byte("Hello, Word!\nhh"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "Hi\n, there %s", name)
	fmt.Println(c)

	// Output:
	// 2
	// 2
}
