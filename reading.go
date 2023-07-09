package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(path string) []string {
	var data []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}
