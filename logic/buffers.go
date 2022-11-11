package main

import (
	"strings"
)

func CheckWriter(flagValue string, flagArgs []string) string {
	success := "successfully registered data"
	err := WriteMatchData(flagValue, flagArgs)
	checkErr(err)
	return success
}

func CheckReader(flagValue string, flagArgs []string) (percentages, goals []float64, err error) {
	homeTeam := strings.ToUpper(flagArgs[0])
	awayTeam := strings.ToUpper(flagArgs[1])
	percentages, goals, err = getAggregateVerbose(flagValue, homeTeam, awayTeam)
	return
}