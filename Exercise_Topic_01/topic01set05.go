/*

Topic: 1. Go Syntax & Basic Language Constructs
Exercise Set 5: Loops (for)

>> go run topic01set05.go

*/

package main

import (
	"fmt"
)

func main(){
	// 1. Print numbers 1 to 10.
	for i := 1; i<11; i++{
		fmt.Printf("\ni=%d",i)
	}

	// 2. Print even numbers until 20.
	for i := 0; i<21; i++{
		if iMod := i%2; iMod == 0{
			fmt.Printf("\nThis one is even: %d",i)
		} 
	}

	// 3. Sum numbers in a slice: []int{3, 7, 2, 9, 4}.
	myArray := []int{3,7,2,9,4}
	mySlice := myArray[1:3]
	mySliceSum := 0
	for _,elemSlice := range mySlice {
		mySliceSum += elemSlice 
	}
	fmt.Printf("\nSum of slice: %v", mySliceSum)

	// 4. Find the largest number in a slice.
	myLongerArray := []int{3,1,4,55,44,1,2,54,78,8,1,12,6,5,78}
	myLongerSlice := myLongerArray[2:len(myLongerArray)-1]
	myMax, myMin := myLongerSlice[0], myLongerSlice[0]
	for _, elemVal := range myLongerSlice {
		if(elemVal < myMin){
			myMin = elemVal
		}
		if(elemVal > myMax){
			myMax = elemVal
		}
	}
	fmt.Printf("\nMin = %d; Max = %d", myMin, myMax)

	// 5. Use range to print index + value.
	for idx, elemVal := range myLongerSlice {
		fmt.Printf("\n[%v] - %v", idx, elemVal)
	}
	
	// 6. Use continue to skip odd numbers.
	fmt.Println("EXERCISE 6")
	for idx, elemVal := range myLongerSlice {
		if(elemVal % 2 != 0){
			continue
		}
		fmt.Printf("\n[%v] - %v", idx, elemVal)
	}
	
	// 7. Use break to stop at a given value.
	fmt.Println("\nEXERCISE 7")
	for idx, elemVal := range myLongerSlice {
		if(elemVal == 8){
			fmt.Printf("[%v] - %v", idx, elemVal)
			break
		}
	}
	
	// 8. Create an infinite loop that stops when counter reaches 5.
	fmt.Println("\nEXERCISE 08")
	myCnt := 5
	for {
		fmt.Printf("\nmyCnt = %d",myCnt)
		if(myCnt == 5) {
			break
		}
	}

	
	// 9. Loop through a string and print each rune.
	fmt.Println("EXERCISE 9")
	myString := "Learning Go"
	myRuneSlice := []rune(myString)
	for idx, elemRune := range myRuneSlice{
		fmt.Printf("\nRune #%d: %c", idx, elemRune)
	}


	// 10. Nested loops: print multiplication table 1–5.
	for cntOuter := 1; cntOuter < 6; cntOuter += 1{
		fmt.Printf("\nMULT TABLE FOR %d\n", cntOuter)
		for cntInner := 1; cntInner < 10; cntInner += 1 {
			fmt.Printf("%d ", cntInner*cntOuter)
		}
	}
	

	// 11. Use labeled loops to break out of outer loop.
	LoopEx11:
	for cntOuter := 1; cntOuter < 10; cntOuter += 1{
		myBreakCnt := 0
		for{
			if(myBreakCnt > 10){
				fmt.Printf("\nBreaking out of LoopEx11: myBreakCnt=%d",myBreakCnt)
				break LoopEx11
			}
			myBreakCnt++
		}
	}


	// 12. Implement a loop to find the first divisible-by-7 number ≥ 1,000.
	fmt.Println("\nEXERCISE 12")
	myTargetDivident := 7
	myCntNr2 := 0
	for {
		if myAttempt := 1000 + myCntNr2; (myAttempt % myTargetDivident) == 0 {
			fmt.Printf("\nFound it: %d is divisible by %d",myAttempt, myTargetDivident)
			break
		}
		myCntNr2 += 1
	}
}


