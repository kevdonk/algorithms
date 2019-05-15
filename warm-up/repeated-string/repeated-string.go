package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func countAs(s string) map[int]int64 {
    count := int64(0)
    countMap := make(map[int]int64, len(s))
    for i, c := range s {
        if c == 'a' {
            count++
        }
        countMap[i] = count
    }
    return countMap
}

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
    length := len(s)
    // integer division will result in int64 result
    multiple := n / int64(length)
    remainder := n % int64(length)
    counts := countAs(s)
    fmt.Fprint(os.Stderr, counts)
    return multiple * counts[len(s) - 1] + counts[int(remainder) - 1]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
