package controller

import (
	"bettor/models"
	"encoding/csv"
	"errors"
	"os"
	"strings"

	"github.com/arewabolu/csvmanager"
)

func CheckWriter(flagValue string, flagArgs []string) error {
	err := WriteMatchData(flagValue, flagArgs)
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
	err = models.CheckifReg(&gameType, &homeTeam, &awayTeam, games)
	if err != nil {
		return 0, 0, 0, err
	}
	GP = len(games)
	err = models.CheckValidLen(&gameType, &homeTeam, &awayTeam, games)
	if err != nil {
		return 0, 0, 0, err
	}
	even, odd = models.SearchTeam4(homeTeam, awayTeam, gameType)
	return
}

func SearcherV2(gameType string, team, status string) (float64, float64) {
	rds, _ := csvmanager.ReadCsv(models.GetBase()+gameType+".csv", true, 400)
	team = strings.ToUpper(team)
	teamGoals := models.SearchTeam(team, rds)
	teamGoalsFloat := models.FloatCon(teamGoals)
	goalsAgainst := models.SearchTeam3(team, rds)
	goalsAgainstFLoat := models.FloatCon(goalsAgainst)
	multiStatus := models.StatusAllocator(rds, team)

	var statusInt float64
	if status == "home" { //, "away"
		statusInt = 1
	} else {
		statusInt = 0
	}

	TD := &models.TeamData{
		GoalFor:     teamGoalsFloat,
		GoalAgainst: goalsAgainstFLoat,
		Status:      models.FloatCon(multiStatus),
	}
	r, MAE := models.TrainAndTest(TD)
	xG := models.AveragexGFCalc(statusInt, r, TD, 30)
	if xG == 0 {
		MAE = 0
	}
	return xG, MAE

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
