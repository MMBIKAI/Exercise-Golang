package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// function to generate the expected output
func generateExpectedFizzBuzz(n int) string {

	var result string

	for j := 1; j <= n; j++ {

		// first condition to check is when the number divided by 3 and 5, calculating the remainder
		if j%3 == 0 && j%5 == 0 {
			result += "FizzBuzz\n" // if the condition is met, append "FizzBuzz" to the count
		} else if j%3 == 0 || j == 3 { // second condition to check is when the number divided by or containing 3,
			result += "Fizz\n" // if the condition is met, append "Fizz" to the count
		} else if j%5 == 0 || j == 5 { // third condition to check is when the number divided by or containing 5,
			result += "Buzz\n" // if the condition is met, append "Buzz" to the count
		} else {
			// if the number did not meet any conditions
			result += fmt.Sprintf("%d\n", j) // append the number to the count
		}
	}
	return result //return the generated sequence
}

// function to test fizzBuzz
func TestFizzBuzz(t *testing.T) {
	// t.Run allows you to run subtests within a test function
	t.Run("testing FizzBuzz for 100 iterations", func(t *testing.T) {
		n := 100
		expected := generateExpectedFizzBuzz(n) // generate the expected fizzBuzz output sequence

		got := fizzBuzz(n) // call the fizzBuzz function

		if got != expected { // compare the expected and the actual output
			t.Errorf("For n = %d, expected:\n%s\nbut got:\n%s", n, expected, got) // if no match report the error
		}
	})

	// Test for n divisible by 15 special case
	t.Run("testing FizzBuzz for numbers divisible by both 3 and 5", func(t *testing.T) {
		for _, n := range []int{15, 30, 45, 60, 75, 90} {
			// Run a subtest for each number
			t.Run(fmt.Sprintf("testing FizzBuzz for n = %d", n), func(t *testing.T) {
				// Generate the expected output for the current number
				expected := generateExpectedFizzBuzz(n)
				// Call the fizzBuzz function for the current number
				got := fizzBuzz(n)
				// Check if the actual output matches the expected output
				if got != expected {
					// Report an error if the actual output doesn't match the expected output
					t.Errorf("For n = %d, expected:\n%q\nbut got:\n%q", n, expected, got)
				}
			})
		}
	})

	// Test for values more then 100
	t.Run("testing FizzBuzz output limit for 100", func(t *testing.T) {
		n := 100                             // Define the upper limit of the FizzBuzz sequence
		output := fizzBuzz(n)                // Generate the FizzBuzz sequence for the specified limit
		lines := strings.Split(output, "\n") // Split the output into lines

		for _, line := range lines { // Iterate through each line of the output
			num, err := strconv.Atoi(line) // Convert the line to an integer
			if err != nil {
				continue // Skip non-numeric lines
			}
			if num > n { // Check if the generated number exceeds the limit
				t.Errorf("Generated number %d exceeds the limit of %d", num, n) // Report an error if the number exceeds the limit
			}
		}
	})

	// Test for NaN generation
	t.Run("testing NaN generation", func(t *testing.T) {
		n := 0 // define input value for testing

		expected := "NaN\n" // expected output for the input

		got := fizzBuzz(n) // call fizzBuzz function

		if got != expected { // compare the expected and the actual output
			t.Errorf("For n = %d, expected:\n%s\nbut got:\n%s", n, expected, got) // if no match no error report
		}
	})

	// Test for negative numbers between 1 and 100
	t.Run("Testing negative numbers between 1 and 100", func(t *testing.T) {
		// Iterate through numbers from 1 to 100
		for num := 1; num <= 100; num++ {
			got := fizzBuzz(num) // Generate the FizzBuzz sequence for the current number

			// Check if the output contains "Negative numbers"
			if got == "Negative numbers\n" { // If the output contains the message "Negative numbers"
				t.Errorf("For num = %d, negative number generated: %s", num, got) // Report an error
			}
		}
	})

}
