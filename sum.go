package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

// Because the func starts with a lower case it cqan only be accessed from within this package.
// it is essentially private
func internalSum(number int) int {
	return number - 1
}

// Because the func starts with a capital it can be called from anywhere. it is Public
// Sum two integers
func Sum(num1, num2 int) int {
	return num1 + num2
}
