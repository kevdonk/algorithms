package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the dayOfProgrammer function below.
func dayOfProgrammer(year int32) string {
	sumOfOtherMonths := int32(31+31+30+31+30+31+31)
	febDays := int32(28)
	switch {
	case year < 1918:
		if (year % 4) {
			febDays = int32(29)
		}
	case year == 1918:
		febDays = int32(29-14)
	case year > 1918:
		if (year % 400 == 0) || (year % 4 == 0 && year % 100 != 0) {
			febDays = int32(29)
		}
	}
	day := 256 - (febDays + sumOfOtherMonths)
	return fmt.Sprintf("%0d.09.%v", day, year)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    year := int32(yearTemp)

    result := dayOfProgrammer(year)

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
