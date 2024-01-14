package main

import (
	"fmt"
	"liner-stats/pkg"
	"math"
	"os"
	"strings"
)

func main() {
	// get args from termial
	argement := os.Args[1:]
	if len(argement) != 1 {
		fmt.Println("Argument givin is not equal to one file! Plz try again.")
		return
	}
	// take the txt file from args
	file := argement[0]
	datafile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	datastr := strings.Split(string(datafile), "\n")
	data, err := pkg.StringToInt(datastr)
	if err != nil {
		fmt.Println("Error in data givin!")
		return
	}
	var array_y []int = data
	var array_x []int
	for i := range data {
		array_x = append(array_x, i)
	}
	// average for x array
	mean_x := pkg.Mean(array_x)
	// average for y array
	mean_y := pkg.Mean(array_y)
	// variance for x array
	variance_x := pkg.Variance(array_x, mean_x)
	// variance for y array
	variance_y := pkg.Variance(array_y, mean_y)
	covariance := pkg.Covariance(mean_x, mean_y, array_x, array_y)
	// calculate b for Linear Regression equation
	slope := covariance / variance_x
	// calculate a for Linear Regression equation
	intercept := mean_y - (slope * mean_x)
	// calculate Pearson Correlation Coefficient
	r := covariance / math.Sqrt(variance_x*variance_y)
	// print the result
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
