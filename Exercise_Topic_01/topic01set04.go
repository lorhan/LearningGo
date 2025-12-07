/*

Topic: 1. Go Syntax & Basic Language Constructs
Exercise Set 4: Conditionals (if, else-if, switch)

>> go run topic01set04.go

*/

package main

import (
	"fmt"
)

func main(){
	fmt.Printf("Hello there welcome to topic01set04")
	// 1. Read an integer from the user.
	var xRead int;
	fmt.Print("\nType an integer: ")
	fmt.Scanf("%v", &xRead)
	fmt.Printf("\nHere is what you typed: %d",xRead)
	
	
	// 2. Check if the number is positive, negative, or zero.
	if(xRead < 0){
		fmt.Print("\nxRead is less then zero")
	}
	if(xRead == 0){
		fmt.Print("\nxRead is equal to zero")
	}
	if(xRead > 0){
		fmt.Print("\nxRead is larger then zero")
	}


	// 3. Check if the number is even or odd.
	xMod := xRead % 2
	sOddOrEven := "EVEN"
	if (xMod > 0){
		sOddOrEven = "ODD"
	}
	fmt.Printf("\nxMod = %d: xRead is %s", xMod, sOddOrEven)
	
	
	// 4. Use if with a short declaration: if x := ...; x > 10.
	fmt.Println("\nChecking if odd or even again, this time with dec+if:")
	if xMod2 := xRead%2; xMod2 == 0{
		fmt.Printf("\nxMod2 = %d: xRead is %s", xMod, "EVEN")
	} else {
		fmt.Printf("\nxMod2 = %d: xRead is %s", xMod, "ODD")
	}


	// 5. Use logical operators (&&, ||, !).
	fGradeA, fGradeB := 7.0, 8.0
	fGradeAvg := (fGradeA+fGradeB)/2
	fMinGrade, fMinAvg := 4.0, 7.0
	if((fGradeA<fMinGrade) || (fGradeB<fMinGrade) || (fGradeAvg<fMinAvg)){
		fmt.Println("\nFAILED")
	} else {
		fmt.Println("\nPASSED")
	}

	// 6. Write a switch that categorizes numbers: 0, 1–5, 6–10, >10.
	iMerit := 77
	sResult := ""
	switch{
		case iMerit == 0:
			sResult = "FAILED WITHOUT RETRY"
		case (1 <= iMerit) && (iMerit <= 5):
			sResult = "FAILED"
		case (6 <= iMerit) && (iMerit <= 10):
			sResult = "PASSED"
		case (iMerit>10):
			sResult = "PASSED WITH HONOURS"
	}
	fmt.Printf("\nMerit: %v; Result: %v\n", iMerit, sResult)


	// 7. Write a switch on strings for detecting: "go", "java", "python".
	sLanguage := "go"
	switch sLanguage{
	case "go":
		fmt.Println("Coding with GO")
	case "java":
		fmt.Println("Coding with JAVA")
	case "python":
		fmt.Println("Coding with PYTHON")
	}

	// 8. Write a switch with no condition (like chained ifs).
	iMerit2 := 77
	sResult2 := ""
	switch{
		case iMerit2 == 0:
			sResult2 = "(again) FAILED WITHOUT RETRY"
		case (1 <= iMerit2) && (iMerit2 <= 5):
			sResult2 = "(again) FAILED"
		case (6 <= iMerit2) && (iMerit2 <= 10):
			sResult2 = "(again) PASSED"
		case (iMerit2>10):
			sResult2 = "(again) PASSED WITH HONOURS"
	}
	fmt.Printf("\nMerit: %v; Result: %v\n", iMerit2, sResult2)
	// 9. Use fallthrough intentionally and observe behavior.
	iFallthrough := 9
	switch iFallthrough {
	case 1:
		fmt.Println("\nWrote 1")
	case 2:
		fmt.Println("\nWrote 1")
	default:
		fmt.Println("\nDefault")
	}
	// 10. Nest conditionals to test multiple properties.
	iMultProp := 9
	switch iMultProp {
	case 1:
		fmt.Println("\nWrote 1")
	case 2:
		fmt.Println("\nWrote 1")
	default:
		if (iMultProp<10){
			fmt.Println("\nLess then 10")
		}
	}
}


