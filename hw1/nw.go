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
            F[i][j] = int(getMax(match, delete, insert))
        }
    }

    return F
}

func printMatrix(m [][]int) {
    for i := 0; i < len(m); i++ {
        fmt.Println(m[i])
    }
}

func getTraceback(m [][]int, a string, b string, i int, j int) (string, string) {
    alignmentA := ""
    alignmentB := ""
    for (i > 0 || j > 0) {
        x := getMax(0, i - 1)
        y := getMax(0, j - 1)
        if i > 0 && j > 0 && m[i][j] == (m[i-1][j-1] + score(a[y], b[x])) {
            fmt.Println("DIAG")
            alignmentA = alignmentA + string(a[y])
            alignmentB = alignmentB + string(b[x])
            i--
            j--
        } else if i > 0 && m[i][j] == (m[i-1][j] + GAP_PENALTY) {
            fmt.Println("LEFT")
            alignmentA = string(a[y]) + alignmentA
            alignmentB = "-" + alignmentB
            i--
        } else if j > 0 && m[i][j] == (m[i][j-1] + GAP_PENALTY) {
            fmt.Println("UP")
            alignmentA = "-" + alignmentA
            alignmentB = string(b[x]) + alignmentB
            j--
        }
    }
    return alignmentA, alignmentB
}

func main() {
   F := computeMatrix(os.Args[1], os.Args[2])
   printMatrix(F)
   alignmentA, alignmentB := getTraceback(F, os.Args[1], os.Args[2], len(os.Args[2]), len(os.Args[1]))
   fmt.Println(alignmentA, alignmentB)
}
