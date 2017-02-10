package main

import (
	"fmt"
	"strconv"
)

func convertToMilitary(time string) string {
	hour := time[0:2]
	minute := time[3:5]
	second := time[6:8]
	aorp := time[8:10]
	if aorp == "PM" {
		if hour != "12" {
			h, _ := strconv.Atoi(hour)
			hour = strconv.Itoa(h + 12)
		}
	} else {
		if hour == "12" {
			hour = "00"
		}
	}
	return fmt.Sprintf("%v:%v:%v", hour, minute, second)
}

func main() {
	time := ""
	fmt.Scan(&time)

	fmt.Println(convertToMilitary(time))
}
