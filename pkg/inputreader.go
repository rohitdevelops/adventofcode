package inputreader

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

// GetLinesInIntArray is function to get Lines in Input file
// as an Integer array.
func GetLinesInIntArray(filename string) []int {
	handle := readFile(filename)

	var numbers []int
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}

	return numbers
}

// GetLinesInStringArray is function to get Lines in Input file
// as a String array.
func GetLinesInStringArray(filename string) []string {
	handle := readFile(filename)

	var words []string
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words
}

func readFile(filename string) io.Reader {
	handle, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return handle
}
