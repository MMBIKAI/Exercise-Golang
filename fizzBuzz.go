package main

import "fmt"

// function fizzBuzz generates the iteration sequence for numbers and returns a string when conditions are met
func fizzBuzz(n int) string {
	// defining a counter string to store the iteration sequence
	var count string

	// Check if n is equal to 0
	if n == 0 {
		return "NaN\n" // Return "NaN" for zeros
	}

	// Check if n is negative
	if n < 0 {
		return "Negative numbers\n" // Return message for negative numbers
	}

	// for loop to iterate between 1 and n
	for i := 1; i <= n; i++ {
		// first condition to check is when the number divided by 3 and 5, calculating the remainder
		if i%3 == 0 && i%5 == 0 {
			count += "FizzBuzz\n" // if the condition is met, append "FizzBuzz" to the count
		} else if i%3 == 0 || i == 3 { // second condition to check is when the number divided by or containing 3
			count += "Fizz\n" // if the condition is met, append "Fizz" to the count
		} else if i%5 == 0 || i == 5 { // third condition to check is when the number divided by or containing 5
			count += "Buzz\n" // if the condition is met, append "Buzz" to the count
		} else {
			// if the number did not meet any conditions
			count += fmt.Sprintf("%d\n", i) // append the number to the count
		}
	}
	return count // return the generated sequence
}

func main() {
	n := 100 // defining n with the desired value
	// call the fizzBuzz function with n and print the result
	fmt.Println(fizzBuzz(n))
}
