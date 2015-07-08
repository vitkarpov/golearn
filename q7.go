package golearn

func LessFirst(i, j int) (int, int) {
    if i > j {
        i, j = j, i
    }
    return i, j
}