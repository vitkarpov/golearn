package golearn

import (
    "testing"
)

type pair struct {
    i int
    j int
}

type caseTest struct {
    input pair
    expected pair
}

func TestLessFirst(t *testing.T) {
    cases := []caseTest {
        {pair{1, 2}, pair{1, 2}},
        {pair{2, 1}, pair{1, 2}},
        {pair{-1, 2}, pair{-1, 2}},
        {pair{-1, -2}, pair{-2, -1}},
    }

    for _, c := range cases {
        i, j := LessFirst(c.input.i, c.input.j)
        received := pair{i, j}

        if received != c.expected {
            t.Errorf("expected %v, but received %v", c.expected, received)
        }
    }
}