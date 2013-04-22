package utils

func getMax(x ...int) int {
    max := x[0]
    for i := 0; i < len(x); i++ {
        if x[i] > max {
            max = x[i]
        }
    }
    return max
}

func getMaxOfList(x []int) int {
    max := x[0]
    for i := 0; i < len(x); i++ {
        if x[i] > max {
            max = x[i]
        }
    }
    return max
}

func score(a byte, b byte, scoreMatrix map[string]int) int {
    return scoreMatrix[string(a) + string(b)]
}

