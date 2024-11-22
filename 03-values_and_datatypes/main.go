package main

import "fmt"

func main() {

	var number int = 21
	const number2 = 31

	var addressOfNumber *int
	addressOfNumber = &number

	// const addressOfNumber2 *int
	// addressOfNumber2 = &number2
	var decimal float32 = 0.7

	// var addressOfFloat *float32
	// addressOfFloat = &decimal

	var yes bool = true

	// var addressOfBool *bool
	// addressOfBool = &yes

	var language string = "golang"

	// var addressOfString *string

	// addressOfString = &language

	fmt.Printf("number  = %d , decimal = %f, yes = %v, language = %s \n", number, decimal, yes, language)

	// fmt.Println("number address :", addressOfNumber)
	// fmt.Println("float address :", addressOfFloat)
	// fmt.Println("boolean address :", addressOfBool)
	// fmt.Println("string address :", addressOfString)

	fmt.Println("var num :", addressOfNumber)
	fmt.Println("const num :", number2)

}
