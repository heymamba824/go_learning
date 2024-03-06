package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100, 10, 10, false},
	{"valid-data", 100, 0, 10, true},
}

func testdivision(t testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expexted an error but did not get one")
			}

		} else {
			if err != nil {
				t.Error("did not expext an error", err.Error())
			}
		}
		if got != tt.expected {
			t.Errorf("expected")
		}

	}

}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("got an error when we should not have")
	}

}
