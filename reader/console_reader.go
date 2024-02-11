package reader

import (
	"fmt"
	"log"
)

type ConsoleReader struct {
}

func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{}
}

func (c *ConsoleReader) readInt() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal("could not read input number", err)
	}
	return n
}

func (c *ConsoleReader) readChar() rune {
	var r rune
	_, err := fmt.Scanf("%c", &r)
	if err != nil {
		log.Fatal("Could not read input char", err)
	}
	return r
}

func (c *ConsoleReader) readDouble() float64 {
	var f float64
	_, err := fmt.Scan(&f)
	if err != nil {
		log.Fatal("could not read input float", err)
	}
	return f
}

func (c *ConsoleReader) readString() string {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal("could not read input string", err)
	}
	return s
}
