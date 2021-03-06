package main

import (
	"testing"
)

type TestCase struct {
	testCaseID int
	err        error
	input      string
	expected   int
}

func TestCheck(t *testing.T) {

	cases := []TestCase{
		{
			testCaseID: 1,
			err:        check("12"),
		},
		{
			testCaseID: 2,
			err:        check("0101010101101101010001101010"),
		},
		{
			testCaseID: 3,
			err:        check("11 11"),
		},
		{
			testCaseID: 4,
			err:        check("000/00"),
		},
	}

	for _, tc := range cases {

		err := tc.err

		if err == nil {
			t.Errorf("testCaseID %v failed with error -> %s", tc.testCaseID, tc.err)
		}

	}

}

func TestBin2Dec(t *testing.T) {

	cases := []TestCase{
		{
			testCaseID: 1,
			input:      "11111111",
			expected:   255,
		},
		{
			testCaseID: 2,
			input:      "00 00",
			expected:   0,
		},
	}

	for _, tc := range cases {
		got := bin2dec(tc.input)

		if got != tc.expected {
			t.Errorf("testCaseID %v failed", tc.testCaseID)
		}
	}
}
