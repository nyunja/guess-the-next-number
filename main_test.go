package main

import (
	"bytes"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestEstimatedRange(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	estimateRange(100)
	w.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()
	expectedOutput := "10.00 190.00\n"
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Wrong estimated range, got %s expected %s\n", output, expectedOutput)
	}
}

func TestLinearRegression(t *testing.T) {
	tests := []struct{
		name string
		x []float64
		y []float64
		wantM float64
		wantC float64
	}{
		{
			name: "test 1",
			x: []float64{0, 1, 2, 3, 4, 5},
			y: []float64{12.0, 12.0, 14.0, 56.0, 34.0, 10.0},
			wantM: 2.8,
			wantC: 16,
		},
	}
	for _, tt := range tests {
		gotM,gotC := linearRegression(tt.x, tt.y)
		if gotM != tt.wantM {
			t.Errorf("calculateLinearRegression() gotM = %v, want %v", gotM, tt.wantM)
		}
		if gotC != tt.wantC {
			t.Errorf("calculateLinearRegression() gotC = %v, want %v", gotC, tt.wantC)
		}
	}
}