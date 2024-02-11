package reader

type IReader interface {
	readInt() int
	readDouble() float64
	readString() string
	readChar() rune
}
