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
        if i != c.input.i || j != c.input.j {
            t.Errorf("expected %v, but received %v", c.expected, pair{i, j})
        }
    }
}