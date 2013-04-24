package utils

import "fmt"

func GetMax(x ...int) int {
    max := x[0]
    for i := 0; i < len(x); i++ {
        if x[i] > max {
            max = x[i]
        }
    }
    return max
}

func GetMaxOfList(x []int) int {
    max := x[0]
    for i := 0; i < len(x); i++ {
        if x[i] > max {
            max = x[i]
        }
    }
    return max
}

func Score(a byte, b byte, scoreMatrix map[string]int) int {
    return scoreMatrix[string(a) + string(b)]
}

func PrintMatrix(m [][]int) {
    for i := 0; i < len(m); i++ {
        fmt.Println(m[i])
    }
}
