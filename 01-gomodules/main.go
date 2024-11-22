package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("This is modules tutorial")

	_ = mux.Route{}

	fmt.Println("imported MUX library")

}
