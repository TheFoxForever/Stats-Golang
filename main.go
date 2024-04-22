package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

func calcLinReg(dataset string, data stats.Series) ([]stats.Coordinate, error) {

	coordinates, err := stats.LinearRegression(data)
	if err != nil {
		return nil, fmt.Errorf("linear Regression error for dataset %s %w", dataset, err)
	}
	fmt.Println("Dataset ", dataset)
	for i := 0; i < len(coordinates); i++ {
		fmt.Println(coordinates[i])
	}
	fmt.Println()
	return coordinates, nil
}

var first stats.Series = []stats.Coordinate{
	{X: 10, Y: 8.04},
	{X: 8, Y: 6.95},
	{X: 13, Y: 7.58},
	{X: 9, Y: 8.81},
	{X: 11, Y: 8.33},
	{X: 14, Y: 9.96},
	{X: 6, Y: 7.24},
	{X: 4, Y: 4.26},
	{X: 12, Y: 10.84},
	{X: 7, Y: 4.82},
	{X: 5, Y: 5.68},
}

var second stats.Series = []stats.Coordinate{
	{X: 10, Y: 9.14},
	{X: 8, Y: 8.14},
	{X: 13, Y: 8.74},
	{X: 9, Y: 8.77},
	{X: 11, Y: 9.26},
	{X: 14, Y: 8.10},
	{X: 6, Y: 6.13},
	{X: 4, Y: 3.10},
	{X: 12, Y: 9.13},
	{X: 7, Y: 7.26},
	{X: 5, Y: 4.74},
}

var third stats.Series = []stats.Coordinate{
	{10, 7.46},
	{8, 6.77},
	{13, 12.74},
	{9, 7.11},
	{11, 7.81},
	{14, 8.84},
	{6, 6.08},
	{4, 5.39},
	{12, 8.15},
	{7, 6.42},
	{5, 5.73},
}

var fourth stats.Series = []stats.Coordinate{
	{8, 6.58},
	{8, 5.76},
	{8, 7.71},
	{8, 8.84},
	{8, 8.47},
	{8, 7.04},
	{8, 5.25},
	{8, 12.50},
	{8, 5.56},
	{8, 7.91},
	{8, 6.89},
}

func main() {
	startTime := time.Now()

	calcLinReg("I", first)
	calcLinReg("II", second)
	calcLinReg("III", third)
	calcLinReg("IV", fourth)

	timeTaken := time.Since(startTime)
	fmt.Println("Time Taken: ", timeTaken)
}
