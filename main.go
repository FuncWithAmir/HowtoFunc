package main

import "fmt"

func main() {
	// + - * / %
	// < > == <= >= !=
	// () && ||

	myFirstVar, mySecondVar := 5, 2
	myThirdVar := "hello"

	result := (myFirstVar <= mySecondVar) || (myThirdVar == "goodbye")

	fmt.Println(result)
}
