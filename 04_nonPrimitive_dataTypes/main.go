package main

import "fmt"

//! defining struct

type StudentInfo struct {
	name    string
	class   int
	rollno  int
	address AddressOfTheStudent
}

type AddressOfTheStudent struct {
	lane1    string
	lane2    string
	pincode  int
	district string
}

//! defining Interface

type Figure interface {
	Area() float64
}

//! defining struct of figure
type Rectangle struct {
	//!declaring struct variables

	length float64
	width  float64
}

//! function to calculate area of Rectangle

func (rect Rectangle) Area() float64 {

	// Area of rectangle = l * b
	area := rect.length * rect.width
	return area
}

//! main function
func main() {

	//!: Method-1
	var avi StudentInfo
	avi.name = "avinash"
	avi.class = 9
	avi.address.lane1 = "endowment colony"
	avi.address.pincode = 500013

	//!: Method-2
	cnu := StudentInfo{
		name:   "seenu",
		rollno: 18,
		address: AddressOfTheStudent{
			district: "Warangal",
			pincode:  506002,
			lane2:    "ag colony",
		},
	}

	//! Interface
	// declaring a rectangle instance

	rectangle := Rectangle{
		length: 10.5,
		width:  12.25,
	}

	// The Figure interface can hold rectangle and square type as they both implements the interface

	var f1 Figure = rectangle
	fmt.Printf("Area of rectangle: %.3f unit sq.\n", f1.Area())

	//!----------------------------------------------

	val1 := 12
	val2 := "hello"
	val3 := true
	var interfaceExample interface{}

	interfaceExample = val1
	fmt.Println("Interface val1 :", interfaceExample)
	interfaceExample = val2
	fmt.Println("Interface val2 :", interfaceExample)
	interfaceExample = val3

	fmt.Println("Interface val3 :", interfaceExample)

	//!-----------------------------------------------------
	fmt.Printf("name= %s, class= %d, rollno= %d, lane1=%s , pincode =%d \n", avi.name, avi.class, avi.rollno, avi.address.lane1, avi.address.pincode)
	fmt.Printf("name = %s , class = %d, rollno = %d, district=%s, pincode=%d,   lane=%s \n", cnu.name, cnu.class, cnu.rollno, cnu.address.district, cnu.address.pincode, cnu.address.lane2)

}
