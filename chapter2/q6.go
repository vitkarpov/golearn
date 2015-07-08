/**
 * Question 6
 *
 * Write a function that calculates the average of a float64 slice
 */

package main

import (
    "fmt"
)

func getAverage(input []float64) float64 {
    var sum float64 = 0
    var l int = len(input)

    for _, v := range input {
        sum += v
    }
    return sum / float64(l)
}

func main() {
    input := []float64{ 1, 2, 3, 4.5, 5 }
    average := getAverage(input)

    fmt.Printf("%v\n", average);
}