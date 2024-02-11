package reader

import (
	"fmt"
	"os"
	"testing"
)

func TestConsoleReader_readInt(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdin = r

	cr := NewConsoleReader()

	expected_1 := 5
	expected_2 := 100
	expected_3 := 1000
	w.WriteString(fmt.Sprint(expected_1))
	w.WriteString(fmt.Sprintf("  %d", expected_2))
	w.WriteString(fmt.Sprintf("\n %d", expected_3))
	w.Close()

	got := cr.readInt()
	if got != expected_1 {
		t.Fatalf("not matching : Expected : %v, got : %v", expected_1, got)
	}

	got = cr.readInt()
	expected_2 = 100
	if got != expected_2 {
		t.Fatalf("not matching : Expected : %d, got : %v", expected_2, got)
	}

	got = cr.readInt()
	expected_3 = 1000
	if got != expected_3 {
		t.Fatalf("not matching : Expected : %d, got : %v", expected_3, got)
	}
}

func TestConsoleReader_readDouble(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdin = r

	cr := NewConsoleReader()

	w.WriteString("5.39")
	w.WriteString("  100.01")
	w.WriteString("\n 1000.03")
	w.Close()

	expected := 5.39
	got := cr.readDouble()

	if got != expected {
		t.Fatalf("not matching : Expected : %f, got : %v", expected, got)
	}

	got = cr.readDouble()
	expected = 100.01
	if got != expected {
		t.Fatalf("not matching : Expected : %f, got : %v", expected, got)
	}

	got = cr.readDouble()
	expected = 1000.03
	if got != expected {
		t.Fatalf("not matching : Expected : %f, got : %v", expected, got)
	}
}

func TestConsoleReader_readString(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdin = r

	cr := NewConsoleReader()

	w.WriteString("5.39")
	w.WriteString("  100.01")
	w.WriteString("\n 1000.03")
	w.Close()

	expected := "5.39"
	got := cr.readString()

	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}

	got = cr.readString()
	expected = "100.01"
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}

	got = cr.readString()
	expected = "1000.03"
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
}

func TestConsoleReader_readChar(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdin = r

	cr := NewConsoleReader()

	w.WriteString("5.3\t\n p")
	w.Close()

	expected := "5"
	got := string(cr.readChar())

	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
	got = string(cr.readChar())
	expected = "."
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
	got = string(cr.readChar())
	expected = "3"
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
	got = string(cr.readChar())
	expected = "\t"
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
	got = string(cr.readChar())
	expected = "\n"
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
	got = string(cr.readChar())
	expected = " "
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
	got = string(cr.readChar())
	expected = "p"
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
}
