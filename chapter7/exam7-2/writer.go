/*
 * Usage:
 *		func main() {
 *			w, c := CountingWriter(os.Stdin)
 *			fmt.Printf("%T\n", w)
 *			fmt.Printf("%d\n", c)
 *		}
 */
package main

import (
	"fmt"
	"io"
	"os"
)

type Counting struct {
	w     io.Writer
	count int64
}

func (c *Counting) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.count += int64(n)
	return len(p), err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var c = Counting{w: w}
	return c.w, &c.count
}

func main() {
	w, c := CountingWriter(os.Stdin)
	fmt.Printf("%T\n", w)
	fmt.Printf("%d\n", c)
}
