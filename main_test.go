package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

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
			name: "test 2",
			n:    150,
			want: "60.00 240.00\n",
		},
		{
			name: "test 3",
			n:    200,
			want: "110.00 290.00\n",
		},
		{
			name: "test 4",
			n:    400,
			want: "310.00 490.00\n",
		},
	}

	for _, tt := range tests {
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
			name:  "test 2",
			x:     []float64{0, 1},
			y:     []float64{12.0, 12.0},
			wantM: 0,
			wantC: 12,
		},
		{
			name:  "test 3",
			x:     []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
			y:     []float64{12.0, 12.0, 14.0, 56.0, 34.0, 10.0, 12.0, 12.0, 14.0, 56.0, 34.0, 10.0, 12.0, 12.0, 14.0, 56.0, 34.0, 10.0, 12.0, 12.0, 14.0, 56.0, 34.0, 10.0, 12.0, 12.0, 14.0, 56.0, 34.0, 10.0, 12.0, 12.0, 14.0, 56.0, 34.0, 10.0, 12.0, 12.0, 14.0, 56.0, 34.0, 10.0, 12.0, 12.0, 14.0, 56.0, 34.0, 10.0, 12.0, 12.0},
			wantM: -0.013061224489795919,
			wantC: 22.88,
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
