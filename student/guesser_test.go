package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

// Test estimateRange function with different inputs
func TestEstimatedRange(t *testing.T){
	y = []float64{0, 1}
	tests := []struct {
		name  string
		n     float64
		lower float64
		upper float64
	}{
		//test edge cases
		{
            name:  "test 0",
            n:     0,
            lower: -1.00,
            upper: 2.00,
        },
        {
            name:  "test 1",
            n:     100,
            lower: 16.00,
            upper: 184.00,
        },
        {
            name:  "test 2 - zero",
            n:     0,
            lower: -35.00,
            upper: 134.00,
        },
	}

	for _, tt := range tests {
		low, up := estimateRange(tt.n)
		expectedLower := tt.lower
		expectedUpper := tt.upper
		if !reflect.DeepEqual(low, expectedLower) {
			t.Errorf("%s failed. Wrong estimated range, got %.2f expected %.2f\n", tt.name, low, expectedLower)
			return
		}
		if !reflect.DeepEqual(up, expectedUpper) {
			t.Errorf("%s failed. Wrong estimated range, got %.2f expected %.2f\n", tt.name, up, expectedUpper)
			return
		}
	}
	y = []float64{}
}

// TestlinearRegression tests the linearRegression function.
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

// Test processInput with different inputs
func Test_processInput(t *testing.T) {
	y = []float64{0, 1}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// test 1
		{name: "test 1",
			args:    args{r: bytes.NewBufferString("100\n")},
			wantW:   "42.00 225.00\n",
			wantErr: false},
		// test 2
		{name: "test 2 - empty input",
			args:    args{r: bytes.NewBufferString("\n")},
			wantW:   "",
			wantErr: true},
		// test 3
		{name: "test 3 - invalid input",
			args:    args{r: bytes.NewBufferString("abc")},
			wantW:   "",
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := processInput(tt.args.r, w); (err != nil) != tt.wantErr {
				t.Errorf("processInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("processInput() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
	y = []float64{}
}

// pearsonsCorrelation calculates the Pearson correlation coefficient between two datasets.
func Test_pearsonsCorrelation(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test 1",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{1, 2, 3, 4, 5},
			},
			want: 1.0,
		},
		{
			name: "test 2",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{5, 4, 3, 2, 1},
			},
			want: -1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pearsonsCorrelation(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("pearsonsCorrelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

// test calculateVariance with various inputs
func Test_calculateVariance(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// test edge cases
		{
			name: "test 1 - zero variance",
			args: args{
				numbers: []float64{1, 1, 1, 1, 1},
			},
			want: 0.0,
		},
		{
			name: "test 2 - positive variance",
			args: args{
				numbers: []float64{1, 2, 3, 4, 5},
			},
			want: 2.0,
		},
		{
			name: "test 3 - negative variance",
			args: args{
				numbers: []float64{-1, -2, -3, -4, -5},
			},
			want: 2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateVariance(tt.args.numbers); got != tt.want {
				t.Errorf("calculateVariance() = %v, want %v", got, tt.want)
			}
		})
	}
}

// test calculateMean with various inputs
func Test_calculateMean(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// test edge cases
		{
            name: "test 1 - zero mean",
            args: args{
                numbers: []float64{1, 1, 1, 1, 1},
            },
            want: 1.0,
        },
        {
            name: "test 2 - positive mean",
            args: args{
                numbers: []float64{1, 2, 3, 4, 5},
            },
            want: 3.0,
        },
        {
            name: "test 3 - negative mean",
            args: args{
                numbers: []float64{-1, -2, -3, -4, -5},
            },
            want: -3.0,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMean(tt.args.numbers); got != tt.want {
				t.Errorf("calculateMean() = %v, want %v", got, tt.want)
			}
		})
	}
}
