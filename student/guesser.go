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
	y = []float64{}
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
		if len(y) > 1 {
			fmt.Fprintf(w, "%.2f %.2f\n", lower, upper)
		}
	}
	return scanner.Err()
}

// Estimate range of the next number from stdin
func estimateRange(n float64) (float64, float64) {
	var lower, upper float64
	if len(y) == 4 {
		y = y[1:]
	}
	y = append(y, n)
	x := make([]float64, len(y))
	for i := range y {
		x[i] = float64(i)
	}
	slope, intercept := linearRegression(x, y)
	// Predict range
	predictedValue := (slope * float64(len(x))) + intercept

	variance := calculateVariance(y)
	stdDev := math.Sqrt(variance)

	// Calculate pearson correlation coefficient
	r := pearsonsCorrelation(x, y)
	// range adjustment usign pearson correlation
	rangeAdjustment := 1 - math.Abs(r)
	// Adjust predicted value based on range adjustment
	confidence := 2.0
	lower = predictedValue - confidence*stdDev - rangeAdjustment
	upper = predictedValue + confidence*stdDev + rangeAdjustment
	return math.Floor(lower), math.Ceil(upper)
}

// Get slope and intercept using linear regression
// The slope represents the change in y for a one-unit change in x.
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

// Calculate the Pearson correlation coefficient between two slices of numbers.
// The Pearson correlation coefficient ranges from -1 to 1.
func pearsonsCorrelation(x, y []float64) float64 {
	sumDiffXY, sumDiffX2, sumDiffY2 := 0.0, 0.0, 0.0
	meanX := calculateMean(x)
	meanY := calculateMean(y)
	for i := 0; i < len(y); i++ {
		diffX := x[i] - meanX
		diffY := y[i] - meanY
		sumDiffXY += diffX * diffY
		sumDiffX2 += diffX * diffX
		sumDiffY2 += diffY * diffY
	}
	return sumDiffXY / math.Sqrt(sumDiffX2*sumDiffY2)
}

// Calculate the variance of a slice of numbers. 
// Variance is the average squared difference from the mean. 
// The standard deviation is the square root of the variance.  
// This function returns the variance
func calculateVariance(numbers []float64) float64 {
	mean := calculateMean(numbers)
	var sumSquaredDiff float64
	for _, num := range numbers {
		sumSquaredDiff += math.Pow(num-mean, 2)
	}
	return sumSquaredDiff / float64(len(numbers))
}

// Calculate the mean of a slice of numbers.
func calculateMean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}
