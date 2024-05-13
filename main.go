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

func main() {
	startTime := time.Now()

	calcLinReg("I", first)
	calcLinReg("II", second)
	calcLinReg("III", third)
	calcLinReg("IV", fourth)

	timeTaken := time.Since(startTime)
	fmt.Println("Time Taken: ", timeTaken)
}
