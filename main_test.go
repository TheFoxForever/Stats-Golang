package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func TestCalcLinReg_WithEmptyDataset(t *testing.T) {
	data := stats.Series{}
	expectedOutput := stats.Series{}

	resultCoordinates, err := calcLinReg("Empty Dataset", data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(expectedOutput) != len(resultCoordinates) {
		t.Errorf("Incorrect number of items returned. Expected: %d, got %d", len(expectedOutput), len(resultCoordinates))
	}

	for i, coord := range resultCoordinates {
		if coord != expectedOutput[i] {
			t.Errorf("Expected coordinate %v, got %v at index %d", expectedOutput[i], coord, i)
		}
	}
}

func TestCalcLinReg_WithLargeDataset(t *testing.T) {
	data := stats.Series{
		{X: 1, Y: 2},
		{X: 2, Y: 4},
		{X: 3, Y: 6},
		{X: 4, Y: 8},
		{X: 5, Y: 10},
		{X: 6, Y: 12},
		{X: 7, Y: 14},
		{X: 8, Y: 16},
		{X: 9, Y: 18},
		{X: 10, Y: 20},
	}
	expectedOutput := stats.Series{
		{X: 1, Y: 2},
		{X: 2, Y: 4},
		{X: 3, Y: 6},
		{X: 4, Y: 8},
		{X: 5, Y: 10},
		{X: 6, Y: 12},
		{X: 7, Y: 14},
		{X: 8, Y: 16},
		{X: 9, Y: 18},
		{X: 10, Y: 20},
	}

	resultCoordinates, err := calcLinReg("Large Dataset", data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(expectedOutput) != len(resultCoordinates) {
		t.Errorf("Incorrect number of items returned. Expected: %d, got %d", len(expectedOutput), len(resultCoordinates))
	}

	for i, coord := range resultCoordinates {
		if coord != expectedOutput[i] {
			t.Errorf("Expected coordinate %v, got %v at index %d", expectedOutput[i], coord, i)
		}
	}
}
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
