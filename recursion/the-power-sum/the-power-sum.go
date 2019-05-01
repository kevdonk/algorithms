package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func sum(X, N, I int32) int32 {
	current := int32(math.Pow(float64(I), float64(N)))
	switch {
	case current < X:
		return sum(X, N, I+1) + sum(X-current, N, I+1)
	case current == X:
		return 1
	default:
		return 0
	}
}

// Complete the powerSum function below.
func powerSum(X int32, N int32) int32 {
	return sum(X, N, 1)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	XTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	X := int32(XTemp)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	result := powerSum(X, N)

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
