This repo was used to compare the precision of output between the https://github.com/montanaflynn/stats package and Python's numpy package. Comparissions were used against Anscombe's quartet where all statistical descriptions of the dataset are the same but plotting reveals how different they are. Simple linear regression was performed on all four datasets and points were returned or extracted for comparission. 

Due to polyfit returning only the slope and intercept while Go's stats package returned the coordinates from linear regression, I calculated datapoints based on the slope and intercept for points 4 through 13 to compare results between the two out to 6 decimal places. As far as the precision goes, both packages perform equally well and are both sufficient options in this instance. Dataset IV shows that returning only the slope and intercept can give the impression that there is an applicable regression model to fit even though it has perfect multicollinearity. Go returned NaN for this calculation and Python should return a slope of 0 to indicate in their own way that this data is not suitable to create a linear model.

Go's runtime was significantly faster than the Python implementation. I ran both programs from the commandline and attempted to keep the same number of functions calls similar but there is slightly more overhead in the Python version due to conversions to coordinates. This timing does not take into account the compiling process for the Go program or the import statements from Python.

Results:</br>
Go&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;</t>Python</br>
![Go Results](https://github.com/TheFoxForever/Stats-Golang/blob/main/Go_Screenshot.png)![Python Results](https://github.com/TheFoxForever/Stats-Golang/blob/main/Python_Screenshot.png)</br>

Overall Outcome:
Pthon has a more mature data science and AI/ML environment with many packages built out than most languages do not and while there are resources such as Jupyter Notebooks for initial data exploration, Go is a superior approach when implementing solutions in a production environment. I would suggest that Python solutions can be a good fit for proof of concept and specific implementations, the brunt of the workload should be done in Go whenever possible. The speed different will only amplify as the amount of data grows.


Copilot adjustments / queries:

Query 1 :
Can I store the large variables in another file and reference them in main.go?

I did not accept this change since it was the same as what I already had.
also tried: 
    move all of the first, second, third, and fourth variables into another file.
but generated the same code as the previous query.

Suggested change:
package data

import "github.com/montanaflynn/stats"

var First stats.Series = []stats.Coordinate{
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

var Second stats.Series = []stats.Coordinate{
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

var Third stats.Series = []stats.Coordinate{
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

var Fourth stats.Series = []stats.Coordinate{
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

query 2:
Could you make these print statements better formatted? (calcLinReg function)

Generated the same code I already had typed. Discarded.

query 3:
Could you create more unit tests for this function?

Notes: I accepted this, but it fails this test with empty dataset. It calls calcLinReg and expects no error, but it should generate an error since it is passing an empty dataset. Leaving for awareness, but noting here.

Generated and accepted code:
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

ChatGPT 3.5:
Pasted my entire program to add into context window for chat.

# Query 1:
could I add another file and place the variables first second third and fourth in that file and just reference them in main?

# ANSWER:

Yes, you can certainly do that. Separating your variables into another file can help keep your code organized and modular. You can define your variables in a separate Go file and then import and use them in your main function.

Here's how you can achieve that:

    Create a new Go file, let's say datasets.go.
    Define your variables first, second, third, and fourth in this file.
    Import this file into your main.go file.
    Use these variables in your main function as before.

Here's an example of how your datasets.go file might look:

go

package main

import "github.com/montanaflynn/stats"

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

// Define third and fourth variables similarly

And then, your main.go file would import and use these variables:

go

package main

import (
    "fmt"
    "time"

    "github.com/montanaflynn/stats"
)

// Import the variables from datasets.go

func main() {
    startTime := time.Now()

    calcLinReg("I", first)
    calcLinReg("II", second)
    calcLinReg("III", third)
    calcLinReg("IV", fourth)

    timeTaken := time.Since(startTime)
    fmt.Println("Time Taken: ", timeTaken)
}

# Query 2:
can you show me how to import those variables in main?

# Answer:
Certainly! To import variables from another file in Go, you need to ensure that the package containing those variables is exported. In this case, since the first, second, third, and fourth variables are defined in the main package in datasets.go, they are already exported and can be imported into other files within the same package.

Here's how you can import and use those variables in your main.go file:

go

package main

import (
    "fmt"
    "time"

    "github.com/montanaflynn/stats"
)

// Import the variables from datasets.go
// You don't need to explicitly import them; they are automatically available since they are in the same package.

func main() {
    startTime := time.Now()

    calcLinReg("I", first)
    calcLinReg("II", second)
    calcLinReg("III", third)
    calcLinReg("IV", fourth)

    timeTaken := time.Since(startTime)
    fmt.Println("Time Taken: ", timeTaken)
}

In this example, you don't need to explicitly import first, second, third, and fourth variables since they are in the same package (main). Go automatically makes them available for use in any file within the main package. So, you can directly use them in your main function as if they were defined in the same file.


Notes:
Ran main.exe after implementing changes suggested, copied the same variable declaration for third and fourth dataset in the same manner. Suggested code worked on first try after rebuilding.