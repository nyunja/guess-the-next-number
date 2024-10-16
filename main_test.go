package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestInvalidInput(t *testing.T) {
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	// Test with non-numeric input
	file, _ := os.Open("test.txt")

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test with non-numeric input
	os.Stdin = file

	main()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "error") {
		t.Errorf("Expected error message for invalid input, got: %s and %v", output, *file)
	}
}

func TestEstimatedRange(t *testing.T) {
	tests := []struct {
		name string
		n    float64
		want string
	}{
		{
			name: "test 1",
			n:    100,
			want: "10.00 190.00\n",
		},
		{
			name: "test 2 - zero",
			n:    0,
			want: "-90.00 90.00\n",
		},
		{
			name: "test 3 - negative number",
			n:    -50,
			want: "-140.00 40.00\n",
		},
		{
			name: "test 4 - very large number",
			n:    1000000,
			want: "999910.00 1000090.00\n",
		},
	}

	for _, tt := range tests {
		count = 1
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		estimateRange(tt.n)
		w.Close()
		os.Stdout = oldStdout
		var buf bytes.Buffer
		io.Copy(&buf, r)
		output := buf.String()
		expectedOutput := tt.want
		if !reflect.DeepEqual(output, expectedOutput) {
			t.Errorf("%s failed. Wrong estimated range, got %s expected %s\n", tt.name, output, expectedOutput)
		}
		y = []float64{}
	}
}

func TestLinearRegression(t *testing.T) {
	tests := []struct {
		name  string
		x     []float64
		y     []float64
		wantM float64
		wantC float64
	}{
		{
			name:  "test 1",
			x:     []float64{0, 1, 2, 3, 4, 5},
			y:     []float64{12.0, 12.0, 14.0, 56.0, 34.0, 10.0},
			wantM: 2.8,
			wantC: 16,
		},
		{
			name:  "test 4 - single point",
			x:     []float64{1},
			y:     []float64{1},
			wantM: 0,
			wantC: 1,
		},
		{
			name:  "test 5 - perfect line",
			x:     []float64{1, 2, 3, 4, 5},
			y:     []float64{2, 4, 6, 8, 10},
			wantM: 2,
			wantC: 0,
		},
	}
	for _, tt := range tests {
		gotM, gotC := linearRegression(tt.x, tt.y)
		if gotM != tt.wantM {
			t.Errorf("%s failed. calculateLinearRegression() gotM = %v, want %v", tt.name, gotM, tt.wantM)
		}
		if gotC != tt.wantC {
			t.Errorf("%s failed. calculateLinearRegression() gotC = %v, want %v", tt.name, gotC, tt.wantC)
		}
	}
}
