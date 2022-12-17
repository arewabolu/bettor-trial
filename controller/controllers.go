package controller

import (
	"bettor/models"
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

func CheckWriter(flagValue string, flagArgs []string) error {
	err := WriteMatchData(flagValue, flagArgs)
	return err
}

func CheckReader(gameType string, gameValues []string) (GP int, percentageWinorDraw, odds []float64, err error) {
	homeTeam := strings.ToUpper(gameValues[0])
	awayTeam := strings.ToUpper(gameValues[1])
	GP, percentageWinorDraw, odds, err = ReadMatch(gameType, homeTeam, awayTeam)
	return
}

func ReadMatch(gameType, homeTeam, awayTeam string) (GP int, PercentageWinorDraw, odds []float64, err error) {
	//err = models.CheckRegisteredTeams(homeTeam, awayTeam)
	//if err != nil {
	//	return nil, err
	//}
	games := models.GetGames(&gameType, homeTeam, awayTeam)
	err = models.CheckifReg(&gameType, &homeTeam, &awayTeam, games)
	if err != nil {
		return 0, nil, nil, err
	}
	GP = len(games)
	err = models.CheckValidLen(&gameType, &homeTeam, &awayTeam, games)
	if err != nil {
		return 0, nil, nil, err
	}
	PercentageWinorDraw = models.PercentageWinorDraw(gameType, homeTeam, awayTeam, games)
	homeOdds := models.OddsCalc(PercentageWinorDraw[0], PercentageWinorDraw[2])
	awayOdds := models.OddsCalc(PercentageWinorDraw[1], PercentageWinorDraw[2])

	odds = []float64{homeOdds, awayOdds}
	//return len(games), win percentage, draw percentage, expected odds
	return
}

func WriteMatchData(gameType string, data2Reg []string) (err error) {
	homeTeam := strings.ToUpper(strings.TrimSpace(data2Reg[0]))
	awayTeam := strings.ToUpper(strings.TrimSpace(data2Reg[1]))
	homeScore := strings.TrimSpace(data2Reg[2])
	awayScore := strings.TrimSpace(data2Reg[3])
	if strings.Contains(gameType, "pen") && homeScore == awayScore {
		return errors.New("there are no draws in penalties")
	}
	if homeTeam == "" || homeScore == "" || awayTeam == "" || awayScore == "" {
		return errors.New("please fill all entries")
	}
	//should modify CheckRegisteredTeam??? should return error to verify if team exists
	//err = models.CheckRegisteredTeams(homeTeam, awayTeam)
	//models.CheckErr(err)

	file, err := os.OpenFile(models.GetBase()+gameType+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	wr.Write([]string{homeTeam, awayTeam, homeScore, awayScore})
	return nil
}
