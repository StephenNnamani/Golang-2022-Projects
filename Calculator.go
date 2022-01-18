// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var number1, number2 float64
	var operator string
	fmt.Println("Your first number is: ")
	fmt.Scanln(&number1)
	fmt.Println("Your second number is: ")
	fmt.Scanln(&number2)
	fmt.Println("Type in your operator (+ - * / %): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%.3f %s %.3f", number1, operator, number2, number1+number2)
	case "-":
		fmt.Printf("%.3f %s %.3f", number1, operator, number2, number1-number2)
	case "*":
		fmt.Printf("%.3f %s %.3f", number1, operator, number2, number1*number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Zero can't be a denomenator")
		} else {
			fmt.Printf("%.3f %s %.3f", number1, operator, number2, number1/number2)
		}
	case "%":
		var numb1, numb2 int
		numb1 = int(number1)
		numb2 = int(number2)
		if numb2 == 0 {
			fmt.Println("Zero can't be a denomenator")
		} else {
			fmt.Printf("%.3f %s %.3f", numb1, operator, numb2, numb1%numb2)
		}
	default:
		fmt.Println("That is not an acceptable operator")
	}
}
