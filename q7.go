/**
 * Write a function that returns its (two) parameters
 * in the right, numerica (ascending) order:
 * f(7,2) -> 2,7
 * f(2,7) -> 2,7
 */

package golearn

func LessFirst(i, j int) (int, int) {
    if i > j {
        i, j = j, i
    }
    return i, j
}