package models

import "sort"

func PercentageCalc(n, divider float64) float64 {
	multiplier := n * 100
	percentage := multiplier / divider
	return percentage
}

func AwayPercentCalc(HP, drawP float64) float64 {
	AVal := HP + drawP
	AP := 100 - AVal
	return AP
}

func Median(data []float64) float64 {
	// sort the data
	sort.Float64s(data)

	// get the length of the data
	n := len(data)

	// if the length of the data is odd, return the middle element
	if n%2 == 1 {
		return data[n/2]
	}

	// if the length of the data is even, return the average of the middle two elements
	return (data[n/2-1] + data[n/2]) / 2.0
}
