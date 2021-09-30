package main

import (
	"errors"
	"fmt"
)

const (
	maxLength = 8
)

var (
	errMaxLength  = errors.New("unexpected length")
	errWrongInput = errors.New("unexpected input")
)

type input string

func (in input) maxLength() error {
	if len(in) > maxLength {
		return fmt.Errorf("%v: %v, but max length is: ", errMaxLength, len(in), maxLength)
	}
	return nil
}

func (in input) valid() error {
	for i, bit := range in {
		if (bit != '0') && (bit != '1') {
			return fmt.Errorf("%v: %v at position %v", errWrongInput, bit, i)
		}
	}
	return nil
}

func (in input) convert() (uint8, error) {
	if err := in.maxLength(); err != nil {
		return 0, err
	}
	if err := in.valid(); err != nil {
		return 0, err
	}
	var out uint8
	last := len(in) - 1
	for i, bit := range in {
		if bit == '1' {
			out |= 1 << (last - i)
		}
	}
	return out, nil
}

func main() {

	var in input // = input("10")

	fmt.Print("Enter a sequence consisting of input digits (0 or 1), up to 8 characters: ")
	fmt.Scanf("%s", &in)

	out, err := in.convert()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(" %s converted to: %d\n", in, out)

}
