package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var hhmmss []string
	var twelve string
	fmt.Scan(&twelve)
	hhmmss = strings.Split(twelve, ":")
	if hhmmss[2][2:4] == "PM" {
		hour, _ := strconv.Atoi(hhmmss[0])
		if hour == 12 {
			hhmmss[0] = "12"
		} else {
			hour += 12
			hhmmss[0] = strconv.Itoa(hour)
		}
	} else if hhmmss[0] == "12" {
		hhmmss[0] = "00"
	}
	hhmmss[2] = hhmmss[2][0:2]
	fmt.Println(strings.Join(hhmmss, ":"))
}
