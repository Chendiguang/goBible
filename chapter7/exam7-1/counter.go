package counter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	// strings.NewReader() returns: io.Reader
	in := bufio.NewScanner(strings.NewReader(string(p)))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		*c += 1
	}
	if err := in.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "WordCounter: %q\n", err)
		os.Exit(1)
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	in := bufio.NewScanner(strings.NewReader(string(p)))
	in.Split(bufio.ScanLines)
	for in.Scan() {
		*c += 1
	}
	if err := in.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "LineCounter: %q\n", err)
		os.Exit(1)
	}
	return len(p), nil
}
