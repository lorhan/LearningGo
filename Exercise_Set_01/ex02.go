/*

**Exercise Set 2: Variables, Types & Type Conversion**
>> go run ex02.go

*/

package main

import (
	"fmt"
	"strconv"
)

func getIntRounded[T float32 | float64](fVal T) int {
	return int(fVal + 0.5)
}

func main() {
	// TODO:
	// 1. Declare variables of types: int, float32, float64, string, bool.
	var intA int = 12
	var floatA float32 = 47.89
	var floatB float64 = 78326.35
	var strA string = "hello"
	var boolA bool = true

	// 2. Print their types using fmt.Printf and %T.
	fmt.Printf("Type of intA: [%T]\n", intA)
	fmt.Printf("Type of floatA: [%T]\n", floatA)
	fmt.Printf("Type of floatB: [%T]\n", floatB)
	fmt.Printf("Type of strA: [%T]\n", strA)
	fmt.Printf("Type of boolA: [%T]\n", boolA)

	// 3. Use short declaration (:=) to create inferred types.
	intInf := 23
	strInf := "Heya"
	fmt.Printf("intInf = %d; strInf = %s\n", intInf, strInf)

	// 4. Declare multiple variables at once.
	iMultiDec, sMultiDec := 48, "Sun"
	fmt.Printf("iMultiDec = %d; sMultiDec = %s\n", iMultiDec, sMultiDec)

	// 5. Convert an int to float64.
	iToBeConv := 21
	fConverted := float64(iToBeConv)
	fmt.Printf("Converted: %f\n", fConverted)

	// 6. Convert a float64 to int (and observe rounding behavior).
	xFloatLow, xFloatHigh := 22.2, 22.8
	xIntLow, xIntHigh := int(xFloatLow), int(xFloatHigh)

	fmt.Printf("Truncation: Low: %d; High %d \n", xIntLow, xIntHigh)
	fmt.Printf("Rounding: Low: %d; High %d \n", getIntRounded(xFloatLow), getIntRounded(xFloatHigh))

	// 7. Convert a string containing a number ("42") to int using strconv.Atoi.
	numInt, err := strconv.Atoi("12")
	if err == nil {
		fmt.Printf("The int is: %d\n", numInt)
	} else {
		fmt.Printf("Error in the conversion\n")
	}

	// 8. Declare a zero-value variable of each basic type.
	var zvInt int
	var zvFl32 float32
	var zvFl64 float64
	var zvStr string
	var zvBool bool
	fmt.Println("Value of zvInt:", zvInt)
	fmt.Println("Value of zvFl32:", zvFl32)
	fmt.Println("Value of zvFl64:", zvFl64)
	fmt.Println("Value of zvStr:", zvStr)
	fmt.Println("Value of zvBool:", zvBool)

	// 9. Declare an untyped constant and print its type after assignment.
	const myCon = 23.2
	fmt.Printf("Type of myCon: %T\n", myCon)

	// 10. Use iota to create a simple enum-like constant list.
	// 11. Swap values of two integer variables without a temp variable.
	// 12. Show the difference between := and var in shadowing behavior.

}
