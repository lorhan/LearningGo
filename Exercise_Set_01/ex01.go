/*

>> go mod init strings-exercise
Creates a new module for your project and writes a go.mod file
into the current directory.

The go.mod file records the module path (here strings-exercise)
and the Go version and will later record dependency requirements
(module versions) when you add imports from other modules.

>> go run main.go
Compiles the specified Go source file(s) into a temporary binary
and immediately runs that binary.

The compiled binary is ephemeral (kept in a temporary cache
directory) — go run does not leave a named executable in your
working directory.


>> go build -o strings-app
Compiles the package in the current directory (or the package(s)
you specify) and writes an executable file named strings-app
(because of -o) into the current directory.

>> ./strings-app
Executes the binary file strings-app from the current directory
on Unix-like systems. The ./ ensures you’re explicitly running
the file in the current directory rather than looking for it on
your PATH.


*/

/*
Declaring a main package, which is a way to group functions,
and it is made up of files in the same directory.
*/
package main

// fmt:
// used for formatting text, printing to console, etc.
// From Go's standard library
import (
	"fmt"
	"strings"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func getNumberOfVowels(myString string) int {
	cntVowels := 0
	for _, runeValue := range myString {
		if isVowel(runeValue) {
			cntVowels++
		}
	}
	return cntVowels
}

func getReversedString(myString string) string {
	runeSlice := []rune(myString)
	maxIdx := len(runeSlice) - 1
	for iA, iB := 0, maxIdx; iA < iB; iA, iB = iA+1, iB-1 {
		runeSlice[iA], runeSlice[iB] = runeSlice[iB], runeSlice[iA]
	}
	return string(runeSlice)
}

// A main function executes by default when you run the main package.
func main() {
	fmt.Println("Hello, world")

	// 1. Create a string with your full name.
	var myName string = "Fulano Silva"

	// 2. Print its length using len().
	fmt.Printf("Strings length: %d\n", len(myName))

	// 3. Access the first and last character (as bytes).
	fmt.Printf("First byte: %d. Last byte: %d.\n", myName[0], myName[len(myName)-1])

	// 4. Convert the string to uppercase using strings.ToUpper.
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(myName))

	// 5. Count how many vowels it contains.
	fmt.Printf("Number of vowels: %d\n", getNumberOfVowels(myName))

	// 6. Check if it contains the substring "a" or "A".
	if strings.Contains(myName, "a") || strings.Contains(myName, "A") {
		fmt.Printf("Substring found!\n")
	} else {
		fmt.Printf("Substring NOT FOUND!\n")
	}

	// 7. Extract the substring from index 2 to 5.
	idxStart := 2
	idxEnd := 5
	fmt.Printf("Substring from indexes [%d] to [%d]: %s\n", idxStart, idxEnd, myName[idxStart:idxEnd])

	// 8. Split the name into words using strings.Split.
	strSplit := strings.Split(myName, " ")
	fmt.Println("Splitting by empty-space:", strSplit)

	// 9. Join the words back using a hyphen.
	fmt.Println("Hyphen joined:", strings.Join(strSplit, "-"))

	// 10. Replace all vowel characters with '*'.
	strReplaced := myName
	for _, runeValue := range "aeiouAEIOU" {
		strReplaced = strings.ReplaceAll(strReplaced, string(runeValue), "*")
	}
	fmt.Println("Replacing vowels with '*': ", strReplaced)

	// 11. Iterate over the string as runes and print each rune + its index.
	fmt.Println("Iterating over string as runes")
	for idx, runeValue := range myName {
		fmt.Printf("[%d] %c\n", idx, runeValue)
	}

	// 12. Reverse the string (properly handling runes).
	fmt.Println("Reversing the string in a rune-safe way: ", getReversedString(myName))

}

// ```go
// ```
