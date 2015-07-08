/**
 * Question 6
 *
 * Write a function that calculates the average of a float64 slice
 */

package golearn

func GetAverage(input []float64) float64 {
    var sum float64 = 0
    var l int = len(input)

    if l == 0 {
        return 0
    }

    for _, v := range input {
        sum += v
    }
    return sum / float64(l)
}
