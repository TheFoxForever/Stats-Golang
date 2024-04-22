package main

import (
	"fmt"
	"time"

	"github.com/montanaflynn/stats"
)

func calcLinReg(dataset string, data stats.Series) ([]stats.Coordinate, error) {

	coordinates, err := stats.LinearRegression(data)
	if err != nil {
		return nil, fmt.Errorf("Linear Regression error for dataset %s %w", dataset, err)
	}
	fmt.Println("Dataset ", dataset)
	for i := 0; i < len(coordinates); i++ {
		fmt.Println(coordinates[i])
	}
	fmt.Println()
	return coordinates, nil
}

var first = []stats.Coordinate{
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

var second = []stats.Coordinate{
	{10, 9.14},
	{8, 8.14},
	{13, 8.74},
	{9, 8.77},
	{11, 9.26},
	{14, 8.10},
	{6, 6.13},
	{4, 3.10},
	{12, 9.13},
	{7, 7.26},
	{5, 4.74},
}

var third = []stats.Coordinate{
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

var fourth = []stats.Coordinate{
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
