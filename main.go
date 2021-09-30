package main

import "fmt"

type task struct {
	input   string
	err     error
	decimal interface{}
}

func main() {
	var process task

	fmt.Print("Enter a sequence consisting of input digits (0 or 1), up to 8 characters: ")
	fmt.Scanf("%s", &process.input)

	if err := check(process.input); err != nil {
		process.err = err
		process.decimal = "failed"
		fmt.Printf("Entered sequence: %s, conversion result: %v, error: %#v\n",
			process.input, process.decimal, process.err)
	} else {
		process.decimal = bin2dec(process.input)
		fmt.Printf("Entered sequence: %s, conversion result: %v\n",
			process.input, process.decimal)
	}
}
