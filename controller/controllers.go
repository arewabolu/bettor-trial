package controller

import (
	"bettor/models"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/arewabolu/csvmanager"
	"github.com/arewabolu/pi-rating"
)

func CheckWriter(flagValue string, flagArgs []string) error {
	homeTeam := strings.ToUpper(strings.TrimSpace(flagArgs[0]))
	awayTeam := strings.ToUpper(strings.TrimSpace(flagArgs[1]))
	homeScore := strings.TrimSpace(flagArgs[2])
	awayScore := strings.TrimSpace(flagArgs[3])
	err := WriteMatchData(flagValue, []string{homeTeam, awayTeam, homeScore, awayScore})
	return err
}

func CheckReader(gameType string, gameValues []string) (GP int, even, odd float64, err error) {
	homeTeam := strings.ToUpper(gameValues[0])
	awayTeam := strings.ToUpper(gameValues[1])
	GP, even, odd, err = ReadMatch(gameType, homeTeam, awayTeam)
	return
}

func ReadMatch(gameType, homeTeam, awayTeam string) (GP int, even, odd float64, err error) {
	//err = models.CheckRegisteredTeams(homeTeam, awayTeam)
	//if err != nil {
	//	return nil, err
	//} and
	games := models.GetGames(&gameType, homeTeam, awayTeam)
	err = models.CheckifReg(gameType, homeTeam, awayTeam)
	if err != nil {
		return 0, 0, 0, err
	}
	GP = len(games)
	even, odd = models.SearchTeam4(homeTeam, awayTeam, gameType)
	return
}

func GenRating(gameType string) error {
	data, err := csvmanager.ReadCsv(models.GetBase()+gameType+".csv", 0755, true)
	if err != nil {
		return errors.New("could not generate game info")
	}
	rows := data.Rows()

	for _, game := range rows {
		match := game.String()
		homeGoal, err := strconv.Atoi(match[2])
		if err != nil {
			return err
		}
		awayGoal, err := strconv.Atoi(match[3])
		if err != nil {
			return err
		}
		WritePi(gameType, match[0], match[1], homeGoal, awayGoal)
	}
	return nil
}

// WritePi is a wrapper around UpdateTeamRatings
func WritePi(gameType, homeTeam, awayTeam string, homeScore, awayScore int) error {
	err := pi.UpdateTeamRatings(models.GetBaseGameType("ratings", gameType), homeTeam, awayTeam, homeScore, awayScore)
	if err != nil {
		return err
	}
	return nil
}

func WriteMatchData(gameType string, data2Reg []string) (err error) {

	if strings.Contains(gameType, "pen") && data2Reg[2] == data2Reg[3] {
		return errors.New("there are no draws in penalties")
	}
	if data2Reg[0] == "" || data2Reg[1] == "" || data2Reg[2] == "" || data2Reg[3] == "" {
		return errors.New("please fill all entries")
	}

	err = models.CheckifReg(gameType, data2Reg[0], data2Reg[1])
	if err != nil {
		return err
	}
	homeScoreInt, err := strconv.Atoi(data2Reg[2])
	if err != nil {
		return errors.New("please fill correct entry")
	}
	awayScoreInt, err := strconv.Atoi(data2Reg[3])
	if err != nil {
		return errors.New("please fill correct entry")
	}
	err = WritePi("fifa4x4Eng", data2Reg[2], data2Reg[3], homeScoreInt, awayScoreInt)
	if err != nil {
		return errors.New("please fill correct entry")
	}
	file, err := os.OpenFile(models.GetBase()+gameType+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	wr.Write([]string{data2Reg[0], data2Reg[1], data2Reg[2], data2Reg[3]})
	return nil
}

// deprecated: no longer supported. Use WriteMatchData instead.
func WriteMatchDataHalfs(data2Reg []string) (err error) {
	homeTeam := strings.ToUpper(strings.TrimSpace(data2Reg[0]))
	home1stHalfScore := strings.TrimSpace(data2Reg[1])
	home2ndHalfScore := strings.TrimSpace(data2Reg[2])

	awayTeam := strings.ToUpper(strings.TrimSpace(data2Reg[3]))
	away1stHalfScore := strings.TrimSpace(data2Reg[4])
	away2ndHalfScore := strings.TrimSpace(data2Reg[5])
	if homeTeam == "" || home1stHalfScore == "" || home2ndHalfScore == "" || awayTeam == "" || away1stHalfScore == "" || away2ndHalfScore == "" {
		return errors.New("please fill all entries")
	}
	//should modify CheckRegisteredTeam??? should return error to verify if team exists
	//err = models.CheckRegisteredTeams(homeTeam, awayTeam)
	//models.CheckErr(err)

	file, err := os.OpenFile(models.GetBase()+"fifa4x4halfsEng.csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	wr.Write([]string{homeTeam, home1stHalfScore, home2ndHalfScore, awayTeam, away1stHalfScore, away2ndHalfScore})
	return nil
}
