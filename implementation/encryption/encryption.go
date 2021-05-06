package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "math"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */
// 0 1 2 3
// f e e d
// 4 5 6 7
// t h e d
// 8 9 10 11
// o g

// 0 4 8
// 1 5 9
// 2 6 (10)
// 3 7 (11)

// rows 3
// columns 4
// 0 1 2 3 4 5 6 7 8 9 10

func encryption(s string) string {
    length := float64(len(s))
    root := math.Sqrt(length)
    rows := math.Floor(root)
    columns := math.Ceil(root)
    if rows * columns < length {
        rows++
    }
    var stripped string
    stripped = strings.ReplaceAll(s, " ", "")
    var encrypted string
    for y := float64(0); y < columns; y++ {
        for x := float64(0); x + y < length; x+=columns {
            encrypted = encrypted + string(stripped[int(x+y)])
        }
        encrypted = encrypted + " "
    }
    return encrypted
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := encryption(s)

    fmt.Fprintf(writer, "%s\n", result)

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
