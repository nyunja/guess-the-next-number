package main

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestEstimatedRange(t *testing.T){
	tests := []struct {
		name  string
		n     float64
		lower float64
		upper float64
	}{
		{
			name:  "test 1",
			n:     100,
			lower: 10.00,
			upper: 190.00,
		},
		{
			name:  "test 2 - zero",
			n:     0,
			lower: -90.00,
			upper: 90.00,
		},
		{
			name:  "test 3 - negative number",
			n:     -50,
			lower: -140.00,
			upper: 40.00,
		},
		{
			name:  "test 4 - very large number",
			n:     1000000,
			lower: 999910.00,
			upper: 1000090.00,
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

func Test_processInput(t *testing.T) {
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
			wantW:   "10.00 190.00\n",
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
}
