package main

import "testing"



func TestMain(t *testing.T) {
    input = "1"
	output := EfficientToNormalRx(input)
	expectedOutput := 0
	if output != expectedOutput {
		t.Error("Output:\n %v \n\n---\n\n Expected Output:\n %v", output, expectedOutput)
	}
}
