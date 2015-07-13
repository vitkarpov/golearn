/**
 * Write a function with sums its arguments.
 * Function should be with variadic number of arguments.
 */

package golearn

func Sum(numbers... int) int {
    result := 0
    for _, i := range numbers {
        result += i
    }
    return result
}