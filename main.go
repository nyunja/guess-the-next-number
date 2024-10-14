package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run main.go")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("Unable to convert %s\n", scanner.Text())
		}
		// estimateRange(num)
		fmt.Println(num)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input: ", err)
	}
}
