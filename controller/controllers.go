package controller

import (
	"bettor/models"
	"encoding/csv"
	"os"
	"strings"
)

func CheckWriter(flagValue string, flagArgs []string) string {
	success := "successfully registered data"
	err := WriteMatchData(flagValue, flagArgs)
	models.CheckErr(err)
	return success
}

func CheckReader(gameType string, gameValues []string) (percentages, goals []float64, err error) {
	homeTeam := strings.ToUpper(gameValues[0])
	awayTeam := strings.ToUpper(gameValues[1])
	percentages, goals, err = ReadMatch(gameType, homeTeam, awayTeam)
	return
}

func ReadMatch(gameType, homeTeam, awayTeam string) (percentageWin, goals []float64, err error) {
	err = models.CheckRegisteredTeams(homeTeam, awayTeam)
	if err != nil {
		return nil, nil, err
	}
	games := models.GetGames(gameType, homeTeam, awayTeam)
	err = models.CheckifReg(&gameType, &homeTeam, &awayTeam, games)
	if err != nil {
		return nil, nil, err
	}
	err = models.CheckValidLen(&gameType, &homeTeam, &awayTeam, games)
	if err != nil {
		return nil, nil, err
	}
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
	//should modify CheckRegisteredTeam??? should return error to verify if team exists
	//err = models.CheckRegisteredTeams(homeTeam, awayTeam)
	models.CheckErr(err)

	file, err := os.OpenFile("./database/"+gameType+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	wr.Write([]string{homeTeam, awayTeam, homeScore, awayScore})
	return nil
}
