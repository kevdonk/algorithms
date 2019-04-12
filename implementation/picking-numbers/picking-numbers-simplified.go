package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func absDiff(a, b int32) int32 {
	diff := b - a
	if diff < 0 {
		return -diff
	}
	return diff
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func pickingNumbers(a []int32) int32 {
	var highest int32
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	start := a[0]
	current := int32(1)
	for n := int32(1); n < int32(len(a)); n++ {
		if absDiff(start, a[n]) <= 1 {
			current++
		} else {
			start = a[n]
			if current > highest {
				highest = current
			}
			current = 1
		}
	}
	return max(highest, current)
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

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
