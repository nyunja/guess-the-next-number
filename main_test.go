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
