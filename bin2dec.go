package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	maxLength  float64 = 8
	systemBase float64 = 2
)

type sequence struct {
	binary  string
	err     error
	decimal int
}

func main() {
	var expression sequence

	fmt.Print("Enter a condition consisting of binary digits (0 or 1), up to 8 characters: ")
	fmt.Scan(&expression.binary)

	condition, err := checkInput(expression.binary)
	if err != nil {
		expression.err = err
	}

	expression.decimal = bin2dec(condition)
	fmt.Printf("Output:\ncondition: %s, conversion result: %d, error: %#v\n",
		expression.binary, expression.decimal, expression.err)
}

func bin2dec(condition string) int {
	var decimal int

	degree := len(condition)

	for _, i := range condition {
		digit, _ := strconv.Atoi(string(i))

		degree--

		decimal += digit * (int(math.Pow(systemBase, float64(degree))))
	}

	return decimal
}
