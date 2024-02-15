# FizzBuzz Exercise

This project implements the classic FizzBuzz algorithm and provides comprehensive unit tests to ensure its correctness and robustness. The FizzBuzz algorithm is a simple program that prints numbers from 1 to N, replacing multiples of or contains 3 with "Fizz", multiples of or contains 5 with "Buzz", and multiples of both 3 and 5 with "FizzBuzz".

## Getting Started

### Prerequisites

Before running the code or tests, ensure that you have Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/doc/install).

### Installation

1. Clone the repository to your local machine:

   
   git clone <repository-url>

2. Navigate to the project directory:

    
    cd fizzbuzz-exercise

### Running the Code

To run the FizzBuzz program, execute the following command:

    
    go run fizzBuzz.go

This will execute the FizzBuzz algorithm for numbers from 1 to 100 by default and print the output to the console.

### Running the Tests

The project includes a comprehensive suite of unit tests to validate the FizzBuzz implementation.

To run the tests, execute the following command:

    
    go test

This will run all the test cases defined in the fizzBuzz_test.go file and report the results.

### Test Cases

The test suite covers various scenarios, including:

- Normal sequence from 1 to N.
- Special cases such as multiples of 3, 5, and 15.
- Output limit for the specified range.
- Handling of NaN (Not a Number) case.
- Prevention of negative numbers in the sequence.
    
