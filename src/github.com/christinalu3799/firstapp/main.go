// SETTING UP A DEVELOPMENT ENVIRONMENT

// Declare a main package (a way to group functions, the package is made up of all the files in the same directory)
package main

// fmt is a popular package that contains functions for formatting text, including printing to the console
import (
	"fmt"
	"strconv"
)

// can group variables like this to show they're all related at the package level
var (
	actorName    string = "Elisabeth Sladen"
	companian    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

// ============================ TERMS & DEFINITIONS =============================
// shadowing: variables defined in the inner most scope takes precedence
// variables HAVE to be used if they are defined. otherwise, you will get an error
// naming convention: lowercase variables are scoped with the package
// --- uppercase variables are globally visible
// --- block scope: defined within a function
// --- keep acronyms as all UPPERCASE
// ==============================================================================
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
	fmt.Printf("%v, %T\n", j, j)
	// ======================= CONVERTING VARIABLE TYPES =========================
	var l float32
	l = float32(i)
	// before
	fmt.Printf("%v, %T\n", i, i)
	// after
	fmt.Printf("%v, %T\n", l, l)
	// ======================= CONVERTING INT TO ASCII ===========================
	var m string
	// need to import strconv package to properly do this!
	m = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", m, m)

}
