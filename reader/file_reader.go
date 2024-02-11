package reader

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FileReader struct {
	filePath string
	file     io.Reader
}

func NewFileReader(filePath string) *FileReader {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("couldn't open the file")
	}
	return &FileReader{
		filePath: filePath,
		file:     file,
	}
}

func (fr *FileReader) readInt() int {
	var n int
	_, err := fmt.Fscan(fr.file, &n)
	if err != nil {
		log.Fatal("could not read input number", err)
	}
	return n
}

func (fr *FileReader) readChar() rune {
	var r rune
	_, err := fmt.Fscanf(fr.file, "%c", &r)
	if err != nil {
		log.Fatal("Could not read input char", err)
	}
	return r
}

func (fr *FileReader) readDouble() float64 {
	var f float64
	_, err := fmt.Fscan(fr.file, &f)
	if err != nil {
		log.Fatal("could not read input float", err)
	}
	return f
}

func (fr *FileReader) readString() string {
	var s string
	_, err := fmt.Fscan(fr.file, &s)
	if err != nil {
		log.Fatal("could not read input string", err)
	}
	return s
}
