// Objective: Find set of n (2 or 3) entries from the input entries in the file
// that sum to 2020, and print their product.
// Puzzle Link : https://adventofcode.com/2020/day/1

package main

import (
	"fmt"
	"log"

	inputreader "../../pkg"
)

func main() {

	// Checking for the sum of enteries to be 2020.
	sumverify := 2020

	// Get Numbers from the Input text file in an Int array.
	numbers := inputreader.GetLinesInIntArray("day1.txt")

	// Checking for the two entries the sum to 2020.
	ent1, ent2, prod := findtwoentries(numbers, sumverify)
	fmt.Printf("Found the fix: %d & %d\n", ent1, ent2)
	fmt.Printf("Product : %d\n", prod)

	// Checking for the three entries the sum to 2020.
	ent1, ent2, ent3, prod := findthreeentries(numbers, sumverify)
	fmt.Printf("Found the fix: %d, %d, %d\n", ent1, ent2, ent3)
	fmt.Printf("Product : %d\n", prod)
}

// findtwoentries function checks for the two entries the sum to the
// sumverify argument and returns the entries and their product.
func findtwoentries(numbers []int, sumverify int) (int, int, int) {
	for i, n := range numbers {
		for _, m := range numbers[i+1:] {
			if (n + m) == sumverify {
				log.Printf("Found the fix: %d & %d\n", n, m)
				prod := (n * m)
				log.Printf("Product : %d\n", prod)
				return n, m, prod
			}
		}
	}
	return 0, 0, 0
}

// findtwoentries function checks for the three entries the sum to the
// sumverify argument and returns the entries and their product.
func findthreeentries(numbers []int, sumverify int) (int, int, int, int) {
	for i, x := range numbers {
		for j, y := range numbers[i+1:] {
			for _, z := range numbers[j+1:] {
				if (x + y + z) == sumverify {
					log.Printf("Found the fix: %d, %d, %d", x, y, z)
					prod := (x * y * z)
					log.Printf("Product : %d", prod)
					return x, y, z, prod
				}
			}
		}
	}
	return 0, 0, 0, 0
}
