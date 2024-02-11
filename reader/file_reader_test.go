package reader

import (
	"log"
	"os"
	"testing"
)

const filePath = "test_file.dat"

func createFile() *os.File {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Couldn't open the file")
	}
	return file
}

func removeFile() error {
	return os.Remove(filePath)
}

func TestFileReader_readInt(t *testing.T) {
	removeFile()
	file := createFile()
	defer file.Close()
	defer removeFile()

	file.WriteString("5")
	fr := NewFileReader(filePath)
	expected := 5
	got := fr.readInt()

	if got != expected {
		t.Fatalf("not matching : Expected : %d, got : %v", expected, got)
	}

	file.WriteString("  100")
	got = fr.readInt()
	expected = 100
	if got != expected {
		t.Fatalf("not matching : Expected : %d, got : %v", expected, got)
	}

	file.WriteString("\n 1000")

	got = fr.readInt()
	expected = 1000
	if got != expected {
		t.Fatalf("not matching : Expected : %d, got : %v", expected, got)
	}
}

func TestFileReader_readDouble(t *testing.T) {
	removeFile()
	file := createFile()
	defer file.Close()
	defer removeFile()

	file.WriteString("5.39")
	fr := NewFileReader(filePath)
	expected := 5.39
	got := fr.readDouble()

	if got != expected {
		t.Fatalf("not matching : Expected : %f, got : %v", expected, got)
	}

	file.WriteString("  100.01")
	got = fr.readDouble()
	expected = 100.01
	if got != expected {
		t.Fatalf("not matching : Expected : %f, got : %v", expected, got)
	}

	file.WriteString("\n 1000.03")

	got = fr.readDouble()
	expected = 1000.03
	if got != expected {
		t.Fatalf("not matching : Expected : %f, got : %v", expected, got)
	}
}

func TestFileReader_readString(t *testing.T) {
	removeFile()
	file := createFile()
	defer file.Close()
	defer removeFile()
	
	file.WriteString("5.39")
	fr := NewFileReader(filePath)
	expected := "5.39"
	got := fr.readString()

	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}

	file.WriteString("  100.01")
	got = fr.readString()
	expected = "100.01"
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}

	file.WriteString("\n 1000.03")

	got = fr.readString()
	expected = "1000.03"
	if got != expected {
		t.Fatalf("not matching : Expected : %s, got : %v", expected, got)
	}
}

func TestFileReader_readChar(t *testing.T) {
	removeFile()
	file := createFile()
	defer file.Close()
	defer removeFile()

	file.WriteString("5.3\n")
	fr := NewFileReader(filePath)
	expected := "5"
	got := string(fr.readChar())

	if got != expected {
		t.Fatalf("not matching : Expected : %v, got : %v", expected, got)
	}
	expected = "."
	got = string(fr.readChar())

	if got != expected {
		t.Fatalf("not matching : Expected : %v, got : %v", expected, got)
	}

	expected = "3"
	got = string(fr.readChar())

	if got != expected {
		t.Fatalf("not matching : Expected : %v, got : %v", expected, got)
	}

	// abl to capture the escaped characters.
	expected = "\n"
	got = string(fr.readChar())
	if got != expected {
		t.Fatalf("not matching : Expected : %v, got : %v", expected, got)
	}
}

func TestFileReader_readMixType(t *testing.T) {
	removeFile()
	file := createFile()
	defer file.Close()
	defer removeFile()

	file.WriteString("5.3\n 100 \t\n 10000 1000b00\n ")
	fr := NewFileReader(filePath)
	got_a := fr.readDouble()         // 5.3
	got_b := fr.readInt()            // 100
	fr.readChar()                    // space
	got_tab := string(fr.readChar()) // "\n"
	got_c := fr.readString()         // 1000
	got_d := fr.readString()         // 1000b00 , it also shifts the next lines
	got_e := string(fr.readChar())   // " "

	expected_a := 5.3
	expected_b := 100
	expected_tab := "\n"
	expected_c := "10000"
	expected_d := "1000b00"
	expected_e := " "

	check := func(expected, got any, issue string) {
		if got != expected {
			t.Fatalf(issue, " not matching : Expected : %v, got : %v", expected, got)
		}
	}
	check(expected_a, got_a, "a")
	check(expected_b, got_b, "b")
	check(expected_c, got_c, "c")
	check(expected_d, got_d, "d")
	check(expected_e, got_e, "e")
	check(expected_tab, got_tab, "tab")
}
