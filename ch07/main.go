package main

import "fmt"

type byteCounter int

func (c *byteCounter) Write(p []byte) (n int, err error) {
	*c += byteCounter(len(p))
	return len(p), nil
}

func main() {
	var c byteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
