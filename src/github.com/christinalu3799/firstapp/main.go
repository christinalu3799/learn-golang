// SETTING UP A DEVELOPMENT ENVIRONMENT

// Declare a main package (a way to group functions, the package is made up of all the files in the same directory)
package main

// fmt is a popular package that contains functions for formatting text, including printing to the console
import "fmt"

// implementing a main function to print a message to the console. a "main" function executes by default when you run the main package
func main() {
	// 3 ways to declare variables
	// 1. Used to declare a variable when you don't want to initialize it yet
	var i int
	i = 42
	// 2. all in 1 line, useful when go doesn't have enough info to assign the type you want
	var j int = 50
	// 3. allow go to figure out data type
	// infers whole numbers are type int. if you add a decimal, it will recognize the variable as type float64.
	k := 46
	fmt.Println(i, j, k)
	// string formatter - %v gives us value and %T gives us type
	fmt.Printf("%v, %T", j, j)
}
