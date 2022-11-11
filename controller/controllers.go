package controller

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func CheckWriter(flagValue string, flagArgs []string) string {
	success := "successfully registered data"
	err := WriteMatchData(flagValue, flagArgs)
	CheckErr(err)
	return success
}

func CheckReader(flagValue string, flagArgs []string) (percentages, goals []float64, err error) {
	homeTeam := strings.ToUpper(flagArgs[0])
	awayTeam := strings.ToUpper(flagArgs[1])
	percentages, goals, err = getAggregateVerbose(flagValue, homeTeam, awayTeam)
	return
}

func getAggregateVerbose(gameType, homeTeam, awayTeam string) (percentageWin, goals []float64, err error) {
	err = CheckRegisteredTeams(homeTeam, awayTeam)
	CheckErr(err)
	games := GetGames(&gameType, homeTeam, awayTeam)
	err = CheckifReg(&gameType, &homeTeam, &awayTeam, games)
	CheckErr(err)
	err = CheckValidLen(&gameType, &homeTeam, &awayTeam, games)
	CheckErr(err)
	value := Trials(games)
	fmt.Println(value)
	percentageWin = PercentageWins(homeTeam, awayTeam, games)

	if gameType == "4x4" {
		goals = ScorePercentage4x4(homeTeam, awayTeam, games)
		return
	}
	goals = scorePercentagePen(homeTeam, awayTeam, games)
	return
}

func WriteMatchData(gameType string, data2Reg []string) (err error) {
	homeTeam := strings.ToUpper(data2Reg[0])
	awayTeam := strings.ToUpper(data2Reg[1])
	homeScore := data2Reg[2]
	awayScore := data2Reg[3]
	err = CheckRegisteredTeams(homeTeam, awayTeam)
	checkErr(err)

	file, err := os.OpenFile(filePath[gameType], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	wr.Write([]string{homeTeam, awayTeam, homeScore, awayScore})
	return nil
}
