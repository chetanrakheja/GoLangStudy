package main

import "fmt"

// testVar := "Test Variable" // This will not work outside of a function
var testVar string= "Test Variable" // This will work outside of a function

const LoginToken string = "asjdhfjhasdfkjhasdf" // Public
const privateConst string = "my private constant" // Private

func main() {
	var username string = "Chetan Rakheja"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T \n", username)

	var isLoggedIN bool = false
	fmt.Println(isLoggedIN)
	fmt.Printf("Variable is of type %T \n", isLoggedIN)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)

	var smallFloat float32 = 255.4551234567890
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type %T \n", smallFloat)

	var bigFloat float64 = 255.4551234567890123
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type %T \n", bigFloat)

	// default values and some aliases
	var anotherInt int
	fmt.Println(anotherInt)
	fmt.Printf("Variable is of type %T \n", anotherInt)

	// implicite type 

	var website = "chetanrakheja.com"
	fmt.Println(website)

	// no var style
	numerOfUsers := 300000.0
	fmt.Println(numerOfUsers)
	fmt.Printf("Variable is of type %T \n", numerOfUsers)
}
