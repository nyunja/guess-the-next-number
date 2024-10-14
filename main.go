package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var data []float64
var count int

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
		estimateRange(num)
		// fmt.Println(lower, upper)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading input: ", err)
	}
}

func estimateRange(n float64) {
	var lower, upper float64
	if len(data) == 50 {
		data = data[1:]
	}
	data = append(data, n)
	if count < 50 {
		lower = n - 90
		upper = n + 90
		count++
	}
	fmt.Printf("%.2f %.2f\n", lower, upper)

}