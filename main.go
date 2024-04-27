package main

import (
	"fmt"
)

func FizzBuzz (number int) string {
	switch {
		case number % 3 == 0 && number % 5 == 0:
			return "FizzBuzz"
		case number % 5 == 0:
			return "Buzz"
		case number % 3 == 0:
			return "Fizz"
		default:
			return fmt.Sprintf("%d", number)
	} 
}

func main () {

	var maxNumber int

	fmt.Print("Please enter the maximum number = ")
	fmt.Scan(&maxNumber) 

	for i := 1; i <= maxNumber; i++ {
		fmt.Println(FizzBuzz(i))
	} 

}