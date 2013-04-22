package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

import "hw2/nw"

func main() {
    file, _ := os.Open(os.Args[1])
    reader := bufio.NewReader(file)

    line, _ := reader.ReadString('\n')
    seq1 := strings.TrimSpace(line)

    line, _ = reader.ReadString('\n')
    seq2 := strings.TrimSpace(line)

    line, _ = reader.ReadString('\n')
    alignmentMode, _ := strconv.ParseInt(strings.TrimSpace(line), 0, 0)

    line, _ = reader.ReadString('\n')
    gp, _ := strconv.ParseInt(strings.TrimSpace(line), 0, 0)
    gapPenalty := int(gp)

    line, _ = reader.ReadString('\n')
    alphabet := strings.TrimSpace(line)

    scoreMatrix := make(map[string]int)
    row := 0
    for {
        line, err := reader.ReadString('\n')
        if err != nil { break }

        line = strings.TrimSpace(line)
        vals := strings.Split(line, " ")
        for col := 0; col < len(vals); col++ {
            key := string(alphabet[row]) + string(alphabet[col])
            score, _ := strconv.ParseInt(vals[col], 0, 0)
            scoreMatrix[key] = int(score)

        }
        row++
    }

    //fmt.Println(seq1, seq2, alignmentMode, gapPenalty, alphabet, scoreMatrix)

    if alignmentMode == 0 {
        F := nw.ComputeMatrix(seq1, seq2, scoreMatrix, gapPenalty)
        bestScore, alignmentA, alignmentB := nw.GetTraceback(F, seq1, seq2, scoreMatrix, gapPenalty)
        fmt.Println(bestScore)
        fmt.Println(alignmentA)
        fmt.Println(alignmentB)
    }
}
