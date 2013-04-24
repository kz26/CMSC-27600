package nw

import "hw2/utils"


func ComputeMatrix(a string, b string, scoreMatrix map[string]int, gp int) [][]int {
    F := make([][]int, len(b) + 1)

    for i := 0; i < len(b) + 1; i++ {
        F[i] = make([]int, len(a) + 1)
    }

    for j := 0; j < len(a) + 1; j++ {
        F[0][j] = j * gp
    }

    for i := 0; i < len(b) + 1; i++ {
        F[i][0] = i * gp
    }

    for i := 1; i < len(b) + 1; i++ {
        for j := 1; j < len(a) + 1; j++ {
            match := F[i-1][j-1] + utils.Score(a[j-1], b[i-1], scoreMatrix)
            del := F[i-1][j] + gp
            ins := F[i][j-1] + gp
            F[i][j] = utils.GetMax(match, del, ins)
        }
    }
    return F
}

func GetTraceback(m [][]int, seqA string, seqB string, scoreMatrix map[string]int, gp int) (int, string, string) {
    alignmentA := ""
    alignmentB := ""
    i := len(seqB)
    j := len(seqA)

    for (i > 0 || j > 0) {
        x := utils.GetMax(0, i - 1)
        y := utils.GetMax(0, j - 1)
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
    return m[len(seqB)][len(seqA)], alignmentA, alignmentB
}

