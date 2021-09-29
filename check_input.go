package bin2dec

import "github.com/pkg/errors"

func lengthErr(length float64) error {
	return errors.Errorf("maxlength must be %v, you entered %v", maxLength, length)
}

func binaryErr(p int, v string) error {
	return errors.Errorf("unexpected binary digit or character {position: %v; value: %s}", p, v)
}

func checkInput(input string) error {
	var binaries = map[rune]rune{'0': '0', '1': '1'}

	for k, s := range input {
		if s != binaries[s] {
			return binaryErr(k, string(s))
		}
	}

	if float64(len(input)) > maxLength {
		return lengthErr(maxLength)
	}

	return nil
}
