package main

import "testing"

// With single function
var tests = []struct {
    name string
    divident float32
    divisor float32
    expected float32
    isErr bool
}{
    {"valid-data", 100.0, 10.0, 10.0, false},
    {"invalid-data", 100.0, 0.0, 0.0, true},
    {"expect-5", 50.0, 10.0, 5.0, false},
}

func TestDivision(t *testing.T) {
    for _, tt := range tests {
        got, err := divide(tt.divident, tt.divisor)
        if tt.isErr {
            if err == nil {
                t.Error("Did not get an error when we should have")
            }
        } else {
            if err != nil {
                t.Error("Got an error when we should not have")
            }
        }
        if got != tt.expected {
            t.Errorf("Expected %f but got %f",tt.expected,got)
        }
    }
}

// Function by function (this is not needed)
func TestDivide(t *testing.T) {
    _, err := divide(10.0, 1.0)
    if err != nil {
        t.Error("Got an error when we should not have")
    }
}

func TestBadDivision(t *testing.T) {
    _, err := divide(10.0, 0)
    if err == nil {
        t.Error("Did not get an error when we should have")
    }
}
