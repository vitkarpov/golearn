/**
 * Create a simple stack which can hold a fixed number of ints.
 * It does not have to grow beyond this limit.
 * Define push – put something on the stack – and pop – retrieve something from the stack – functions.
 * The stack should be a LIFO (last in, first out) stack.
 *
 * Bonus: write a String method which converts the stach to a string representation.
 * This way you can print the stack using: `fmt.Printf("My stach %v'n", stack)`.
 * A stack could be represented: [0:m] [1:l] [2:k]
 */

package golearn

import (
    "fmt"
    "strconv"
)

const size int = 10

type stack struct {
    i int
    // we could store only 10 elements max
    // it's a very poor stack
    data [size] int
}

// it's very important to now working
// with copies of stack each time
func (s *stack) push(el int) {
    s.data[s.i] = el
    s.i++
}

func (s *stack) pop() int {
    s.i--
    return s.data[s.i]
}

func (s stack) String() (str string) {
    for i := 0; i < s.i; i++ {
        str += "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
    }
    return
}
