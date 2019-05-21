package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func max(a, b int32) int32 {
    if a > b {
        return a
    }
    return b
}
// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
    bribes := int32(0)
    for i, x := range q {
        // no one can move more than 2 spots forwards
        if (x - (int32(i)+1)) > int32(2) {
            fmt.Println("Too chaotic")
            return
        }
        // count how many bribes each person has recieved
        // max 2 bribes, so bribers of i can only get to x-1 (-1 for index)
        start := max(x-int32(2),0)
        for j := start; j < int32(i); j++ {
            if q[j] > x {
                bribes++
            }
        }
    }
    fmt.Println(bribes)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(readLine(reader), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
