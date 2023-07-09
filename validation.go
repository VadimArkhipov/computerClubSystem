package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func validate(data []string) {
	datePattern := `^(0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]) (0[0-9]|1[0-9]|2[0-3]):([0-5][0-9])$`
	positiveInteger := `^[1-9]\d*$`
	ID1 := `^(0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]) 1 [0-9a-z_-]+$`
	ID2 := `^(0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]) 2 [0-9a-z_-]+ [1-9]\d*$`
	ID3 := `^(0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]) 3 [0-9a-z_-]+$`
	ID4 := `^(0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]) 4 [0-9a-z_-]+$`

	if len(data) < 1 {
		os.Exit(1)
	}
	tableCount := data[0]
	matched, _ := regexp.MatchString(positiveInteger, tableCount)
	if !matched {
		fmt.Println(tableCount)
		os.Exit(1)
	}

	if len(data) < 2 {
		os.Exit(1)
	}
	openCloseTime := data[1]
	matched, _ = regexp.MatchString(datePattern, openCloseTime)
	if !matched {
		fmt.Println(openCloseTime)
		os.Exit(1)
	}

	openTime, _ := time.Parse("15:04", strings.Split(openCloseTime, " ")[0])
	closeTime, _ := time.Parse("15:04", strings.Split(openCloseTime, " ")[1])
	if closeTime.Before(openTime) {
		fmt.Println(openCloseTime)
		os.Exit(1)
	}

	if len(data) < 3 {
		os.Exit(1)
	}
	cost := data[2]
	matched, _ = regexp.MatchString(positiveInteger, cost)
	if !matched {
		fmt.Println(cost)
		os.Exit(1)
	}

	tables, _ := strconv.Atoi(tableCount)

	if len(data) < 4 {
		os.Exit(1)
	}
	for i := 3; i < len(data); i++ {
		matched1, _ := regexp.MatchString(ID1, data[i])
		matched2, _ := regexp.MatchString(ID2, data[i])
		if matched2 {
			if num, _ := strconv.Atoi(strings.Split(data[i], " ")[3]); num > tables {
				fmt.Println(data[i])
				os.Exit(1)
			}
		}
		matched3, _ := regexp.MatchString(ID3, data[i])
		matched4, _ := regexp.MatchString(ID4, data[i])

		if !(matched1 || matched2 || matched3 || matched4) {
			fmt.Println(data[i])
			os.Exit(1)
		}

		if i >= 4 {
			prevTime, _ := time.Parse("15:04", data[i-1][0:5])
			currTime, _ := time.Parse("15:04", data[i][0:5])
			if currTime.Before(prevTime) {
				fmt.Println(data[i])
				os.Exit(1)
			}
		}
	}
}
