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

func pickingNumbers(a []int32) int32 {
	var highest int32
	counts := make(map[int32]int32)
	numbers := make([]int32, 0)
	for i := 0; i < len(a); i++ {
		counts[a[i]]++
	}
	for number, _ := range counts {
		numbers = append(numbers, number)
	}
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
	highest = counts[numbers[0]]
	for n := int32(1); n < int32(len(numbers)); n++ {
		var current int32
		diff := numbers[n] - numbers[n-1]
		if diff != int32(1) {
			// highest of the counts
			if counts[n] > counts[n-1] {
				current = counts[numbers[n]]
			} else {
				current = counts[numbers[n-1]]
			}
		} else {
			// sum of the counts
			current = counts[numbers[n]] + counts[numbers[n-1]]
		}
		if current > highest {
			highest = current
		}
	}
	return highest
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
