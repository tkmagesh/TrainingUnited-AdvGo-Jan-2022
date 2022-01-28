package utils

import "testing"

/* func TestIsPrime(t *testing.T) {
	//arrange
	input := 97
	expected := true

	//act
	actual := IsPrime(input)

	//assert
	if actual != expected {
		t.Errorf("IsPrime(%d) = %t, want %t", input, actual, expected)
	}
} */

type TestCase struct {
	name     string
	no       int
	expected bool
	actual   bool
}

func Test_IsPrime(t *testing.T) {
	testCases := []TestCase{
		TestCase{name: "Test_IsPrime_1", no: 1, expected: false},
		TestCase{name: "Test_IsPrime_2", no: 2, expected: true},
		TestCase{name: "Test_IsPrime_7", no: 7, expected: true},
		TestCase{name: "Test_IsPrime_9", no: 9, expected: false},
		TestCase{name: "Test_IsPrime_11", no: 11, expected: true},
		TestCase{name: "Test_IsPrime_12", no: 12, expected: false},
		TestCase{name: "Test_IsPrime_13", no: 13, expected: true},
		TestCase{name: "Test_IsPrime_15", no: 15, expected: false},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.actual = IsPrime(testCase.no)
			if testCase.actual != testCase.expected {
				t.Errorf("IsPrime(%d) = %t, want %t", testCase.no, testCase.actual, testCase.expected)
			}
		})
	}
}
