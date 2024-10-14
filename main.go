package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var y []float64
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
	if len(y) == 50 {
		y = y[1:]
	}
	y = append(y, n)
	if count < 50 {
		lower = n - 90
		upper = n + 90
		count++
	} else {
		x := make([]float64, 0, len(y))
		for i := range y {
			x[i] = float64(i)
		}
		slope, intercept := linearRegression(x, y)
		fmt.Println(slope, intercept)

	}
	fmt.Printf("%.2f %.2f\n", lower, upper)
}

func linearRegression(x, y []float64) (float64, float64) {
	n := float64(len(x))
	var sumX, sumXY, sumY, sumX2 float64
	for i := range x {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}
	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX * sumX)
	intercept := (sumY - slope*sumX) / n
	return slope, intercept
}