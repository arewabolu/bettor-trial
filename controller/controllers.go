package controller

import (
	"bettor/models"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func CheckWriter(flagValue string, flagArgs []string) string {
	success := "successfully registered data"
	err := WriteMatchData(flagValue, flagArgs)
	models.CheckErr(err)
	return success
}

func CheckReader(flagValue string, flagArgs []string) (percentages, goals []float64, err error) {
	homeTeam := strings.ToUpper(flagArgs[0])
	awayTeam := strings.ToUpper(flagArgs[1])
	percentages, goals, err = ReadMatch(flagValue, homeTeam, awayTeam)
	return
}

func ReadMatch(gameType, homeTeam, awayTeam string) (percentageWin, goals []float64, err error) {
	err = models.CheckRegisteredTeams(homeTeam, awayTeam)
	models.CheckErr(err)
	games := models.GetGames(&gameType, homeTeam, awayTeam)
	err = models.CheckifReg(&gameType, &homeTeam, &awayTeam, games)
	models.CheckErr(err)
	err = models.CheckValidLen(&gameType, &homeTeam, &awayTeam, games)
	models.CheckErr(err)
	value := models.Trials(games)
	fmt.Println(value)
	percentageWin = models.PercentageWins(homeTeam, awayTeam, games)

	if gameType == "4x4" {
		goals = models.ScorePercentage4x4(homeTeam, awayTeam, games)
		return
	}
	goals = models.ScorePercentagePen(homeTeam, awayTeam, games)
	return
}

func WriteMatchData(gameType string, data2Reg []string) (err error) {
	homeTeam := strings.ToUpper(data2Reg[0])
	awayTeam := strings.ToUpper(data2Reg[1])
	homeScore := data2Reg[2]
	awayScore := data2Reg[3]
	err = models.CheckRegisteredTeams(homeTeam, awayTeam)
	models.CheckErr(err)

	file, err := os.OpenFile(models.FilePath[gameType], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	wr.Write([]string{homeTeam, awayTeam, homeScore, awayScore})
	return nil
}
