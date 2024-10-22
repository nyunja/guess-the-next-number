package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

var (
	y     []float64 = make([]float64, 0, 50)
	count int
)

// Takes input from stdin and prints the estimated range of the next number to stdout
func processInput(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return fmt.Errorf("invalid input %v", err)
		}
		lower, upper := estimateRange(num)
		fmt.Fprintf(w, "%.2f %.2f\n", lower, upper)
	}
	return scanner.Err()
}

// Estimate range of the next number from stdin
func estimateRange(n float64) (float64, float64) {
	var lower, upper float64
	if len(y) == 50 {
		y = y[1:]
	}
	y = append(y, n)
	if count < 5 {
		lower = n - 90
		upper = n + 90
		count++
	} else {
		x := make([]float64, len(y))
		for i := range y {
			x[i] = float64(i)
		}
		slope, intercept := linearRegression(x, y)
		// Predict range
		predictedValue := (slope * float64(len(x))) + intercept
		lower = math.Max(predictedValue-99, n-99)
		upper = math.Min(predictedValue+100, n+100)
		// Ensure minimum range
		minRange := math.Max(2, math.Abs(n)*0.02)
		if (upper - lower) < minRange {
			mid := (lower + upper) / 2
			lower = mid - minRange/2
			upper = mid + minRange/2
		}
		// Ensure maximum range
		if (upper - lower) > 200 {
			mid := (upper + lower) / 2
			lower = mid - 100
			upper = mid + 100
		}

	}
	return lower, upper
}

// Get slope and intercept using linear regression
func linearRegression(x, y []float64) (float64, float64) {
	if len(x) != len(y) {
		fmt.Println("Error: x and y should have the same length")
		os.Exit(1)
	}
	n := float64(len(x))
	var sumX, sumXY, sumY, sumX2 float64
	for i := range x {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumX2 += x[i] * x[i]
	}
	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n
	return slope, intercept
}
