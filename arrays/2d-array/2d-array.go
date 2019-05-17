package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func sumFromStartPoint(arr [][]int32, x, y int32) int32 {
    sum := int32(0)
    sum += arr[y][x] + arr[y][x+1] + arr[y][x+2]
    sum += arr[y+1][x+1]
    sum += arr[y+2][x] + arr[y+2][x+1] + arr[y+2][x+2]
    return sum
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
    var max int32
    for x := int32(0); x < 4; x++ {
        for y := int32(0); y < 4; y++ {
            sum := sumFromStartPoint(arr, x, y)
            if (x == int32(0) && y == int32(0)) || sum > max {
                max = sum
            }
        }
    }
    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

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
