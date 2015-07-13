package golearn

import (
    "testing"
)

func TestStack(t *testing.T) {
    var stack Stack

    stack.push(1)
    got := stack.pop()

    if (got != 1) {
        t.Errorf("expected %v, but received %v", 1, got)
    }
}
