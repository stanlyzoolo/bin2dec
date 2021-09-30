package main

import (
	"math"
	"strconv"
)

const (
	maxLength  float64 = 8
	systemBase float64 = 2
)

func bin2dec(checked string) int {
	var decimal int

	degree := len(checked)

	for _, i := range checked {
		digit, _ := strconv.Atoi(string(i))

		degree--

		decimal += digit * (int(math.Pow(systemBase, float64(degree))))
	}

	return decimal
}
