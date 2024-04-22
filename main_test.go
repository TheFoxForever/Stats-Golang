package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func TestCalcLinReg(t *testing.T) {
	data := stats.Series{
		{X: 1, Y: 2},
		{X: 2, Y: 4},
		{X: 3, Y: 6},
	}

	expectedOutput := stats.Series{
		{X: 1, Y: 2},
		{X: 2, Y: 4},
		{X: 3, Y: 6},
	}

	resultCoordinates, err := calcLinReg("Test Dataset", data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(expectedOutput) != len(resultCoordinates) {
		t.Errorf("Incorrect number of item returned. Expected: %d, got %d", len(expectedOutput), len(resultCoordinates))
	}

	for i, coord := range resultCoordinates {
		if coord != expectedOutput[i] {
			t.Errorf("Expected coordinate %v, got %v at index %d", expectedOutput[i], coord, i)
		}
	}

}
