package golearn

import (
    "testing"
)

type Test struct {
    input []float64
    expected float64
}

func TestAverage(t *testing.T) {
    cases := []Test {
        {[]float64 {1, 2, 3, 4.5, 5}, 3.1},
        {[]float64 {0, -1, -2, 1, 2}, 0},
        {[]float64 {}, 0},
        {[]float64 {-1, -2, -3, -4}, -2.5},
    }
    for _, c := range cases {
        received := GetAverage(c.input)
        if received != c.expected {
            t.Errorf("expected %v, but received %v", c.expected, received)
        }
    }
}
