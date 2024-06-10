package main

import (
	"fmt"
	"goplaybook/service"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Println("Please provide valid input to run the program")
		return
	}
	service.Process(args)

}
