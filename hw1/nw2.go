package main

import "fmt"
import "os"

const MATCH = 2
const MISMATCH = -1
const GAP_PENALTY = -2

func getMax(x ...int) int {
    max := x[0]
    for i := 0; i < len(x); i++ {
        if x[i] > max {
            max = x[i]
        }
    }
    return max
}

func score(a byte, b byte) int {
    if a == b {
        return MATCH
    }
    return MISMATCH
}
 
func computeMatrix(a string, b string) [][]int {
    F := make([][]int, len(b) + 1)

    for i := 0; i < len(b) + 1; i++ {
        F[i] = make([]int, len(a) + 1) 
    }

    for j := 0; j < len(a) + 1; j++ {
        F[0][j] = j * GAP_PENALTY
    }

    for i := 0; i < len(b) + 1; i++ {
        F[i][0] = i * GAP_PENALTY
    }

    for i := 1; i < len(b) + 1; i++ {
        for j := 1; j < len(a) + 1; j++ {
            match := F[i-1][j-1] + score(a[j-1], b[i-1])
            delete := F[i-1][j] + GAP_PENALTY
            insert := F[i][j-1] + GAP_PENALTY
            F[i][j] = getMax(match, delete, insert)
        }
    }

    return F
}

func printMatrix(m [][]int) {
    for i := 0; i < len(m); i++ {
        fmt.Println(m[i])
    }
}

func getTraceback(m [][]int, a string, b string, i int, j int, alignmentA string, alignmentB string) {
    if i == 0 && j == 0 {
        fmt.Println(alignmentA, alignmentB)
        return
    }
    x := getMax(0, i - 1)
    y := getMax(0, j - 1)
    if i > 0 && j > 0 && m[i][j] == (m[i-1][j-1] + score(a[y], b[x])) {
        fmt.Println(m[i-1][j-1], "DIAG")
        getTraceback(m, a, b, i-1, j-1, alignmentA + string(a[y]), alignmentB + string(b[x]))
    }
    if i > 0 && m[i][j] == (m[i-1][j] + GAP_PENALTY) {
        fmt.Println(m[i-1][j], "UP")
        getTraceback(m, a, b, i-1, j, string(a[y]) + alignmentA, "-" + alignmentB)
    }
    if j > 0 && m[i][j] == (m[i][j-1] + GAP_PENALTY) {
        fmt.Println(m[i][j-1], "LEFT")
        getTraceback(m, a, b, i, j-1, "-" + alignmentA, string(b[x]) + alignmentB)
    }
}

func main() {
   F := computeMatrix(os.Args[1], os.Args[2])
   printMatrix(F)
   getTraceback(F, os.Args[1], os.Args[2], len(os.Args[2]), len(os.Args[1]), "", "")
}
