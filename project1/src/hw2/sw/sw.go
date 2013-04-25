package sw

import "hw2/utils"

// Matrix computation for Smith-Waterman alignment
func ComputeMatrix(a string, b string, scoreMatrix map[string]int, gp int) [][]int {
    F := make([][]int, len(b) + 1)

    for i := 0; i < len(b) + 1; i++ {
        F[i] = make([]int, len(a) + 1)
    }

    for j := 0; j < len(a) + 1; j++ {
        F[0][j] = 0
    }

    for i := 0; i < len(b) + 1; i++ {
        F[i][0] = 0
    }

    for i := 1; i < len(b) + 1; i++ {
        for j := 1; j < len(a) + 1; j++ {
            match := F[i-1][j-1] + utils.Score(a[j-1], b[i-1], scoreMatrix)
            del := F[i-1][j] + gp
            insert := F[i][j-1] + gp
            F[i][j] = utils.GetMax(match, del, insert, 0) // max with 0
        }
    }

    return F
}

// utility function to get the rightmost and downmost starting point
func getStartingPoint(m [][]int) (int, int) {
    // find max in the entire matrix
    max := m[0][0]
    for i := 0; i < len(m); i++ {
        for j := 0; j < len(m[i]); j++ {
            if m[i][j] > max {
                max = m[i][j]
            }
        }
    }

    // now start looking
    row := len(m) - 1
    col := len(m[0]) - 1
    for col >= 0 {
        for i := row; i >= 0; i-- {
            if m[i][col] == max {
                return i, col
            }
        }
        col--
    }
    return -1, -1
}

// Traceback for Smith-Waterman
func GetTraceback(m [][]int, seqA string, seqB string, scoreMatrix map[string]int, gp int) (int, string, string) {
    alignmentA := ""
    alignmentB := ""
    i, j := getStartingPoint(m)
    I, J := i, j

    for (i > 0 || j > 0) {
        x := utils.GetMax(0, i - 1)
        y := utils.GetMax(0, j - 1)
        // stop if there's a 0 and no pointer that produced it...which only happens in the first row/column
        if m[i][j] == 0 && (i == 0 || j == 0) {
            break
        }
        if i > 0 && j > 0 && m[i][j] == (m[i-1][j-1] + utils.Score(seqA[y], seqB[x], scoreMatrix)) {
            alignmentA = string(seqA[y]) + alignmentA
            alignmentB = string(seqB[x]) + alignmentB
            i--
            j--
        } else if j > 0 && m[i][j] == (m[i][j-1] + gp) {
            alignmentA = string(seqA[y]) + alignmentA
            alignmentB = "-" + alignmentB
            j--
        } else if i > 0 && m[i][j] == (m[i-1][j] + gp) {
            alignmentA = "-" + alignmentA
            alignmentB = string(seqB[x]) + alignmentB
            i--
        }
    }
    return m[I][J], alignmentA, alignmentB
}
