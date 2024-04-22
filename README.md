This repo was used to compare the precision of output between the https://github.com/montanaflynn/stats package and Python's numpy package. Comparissions were used against Anscombe's quartet where all statistical descriptions of the dataset are the same but plotting reveals how different they are. Simple linear regression was performed on all four datasets and points were returned or extracted for comparission. 

Due to polyfit returning only the slope and intercept while Go's stats package returned the coordinates from linear regression, I calculated datapoints based on the slope and intercept for points 4 through 13 to compare results between the two out to 6 decimal places. As far as the precision goes, both packages perform equally well and are both sufficient options in this instance. Dataset IV shows that returning only the slope and intercept can give the impression that there is an applicable regression model to fit even though it has perfect multicollinearity. Go returned NaN for this calculation and Python should return a slope of 0 to indicate in their own way that this data is not suitable to create a linear model.

Go's runtime was significantly faster than the Python implementation. I ran both programs from the commandline and attempted to keep the same number of functions calls similar but there is slightly more overhead in the Python version due to conversions to coordinates. This timing does not take into account the compiling process for the Go program or the import statements from Python.

Results:</br>
Go</br>
![Go Results](https://github.com/TheFoxForever/Stats-Golang/blob/main/Go_Screenshot.png)</br>
Python</br>
![Python Results](https://github.com/TheFoxForever/Stats-Golang/blob/main/Python_Screenshot.png)
