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

func PrependMatchData(flagValue string, data []string) error {
	if strings.Contains(flagValue, "pen") && data[2] == data[3] {
		return errors.New("there are no draws in penalties")
	}
	if data[0] == "" || data[1] == "" || data[2] == "" || data[3] == "" {
		return errors.New("please fill all entries")
	}

	homeTeam := strings.ToUpper(strings.TrimSpace(data[0]))
	awayTeam := strings.ToUpper(strings.TrimSpace(data[1]))
	homeScore := strings.TrimSpace(data[2])
	awayScore := strings.TrimSpace(data[3])

	err := models.CheckifReg(flagValue, homeTeam, awayTeam)
	if err != nil {
		return err
	}

	err = csvmanager.PrependRow(models.GetBase()+flagValue+".csv", 0755, true, []string{homeTeam, awayTeam, homeScore, awayScore})
	if err != nil {
		return err
	}
	return nil
}

func ReadMatch(gameType, homeTeam, awayTeam string) (GP int, even, odd float64, err error) {
	games := models.GetGames(gameType, homeTeam, awayTeam)
	err = models.CheckifReg(gameType, homeTeam, awayTeam)
	if err != nil {
		return 0, 0, 0, err
	}
	GP = len(games)
	even, odd = models.SearchTeam4(homeTeam, awayTeam, gameType)
	return
}

func ReadPi(gameType string, teams []string) (float64, float64, error) {
	homeTeam := strings.ToUpper(strings.TrimSpace(teams[0]))
	awayTeam := strings.ToUpper(strings.TrimSpace(teams[1]))
	err := models.CheckifReg(gameType, homeTeam, awayTeam)
	if err != nil {
		return 0, 0, err
	}
	homeTeamGames, err := models.GetGamesV2(gameType, homeTeam)
	if err != nil {
		return 0, 0, err
	}
	homeTeamHomeRating := models.GenerateInstantPi(homeTeamGames, homeTeam).ProvisionalRating("home").HomeRating

	awayTeamGames, err := models.GetGamesV2(gameType, awayTeam)
	if err != nil {
		return 0, 0, err
	}
	awayTeamAwayRating := models.GenerateInstantPi(awayTeamGames, awayTeam).ProvisionalRating("away").AwayRating

	return homeTeamHomeRating, awayTeamAwayRating, nil
}

func ReadExpectedGoals(homeTeamHomeRating float64, awayTeamAwayRating float64) (float64, float64) {
	return pi.ExpectedGoalIndividual(homeTeamHomeRating), pi.ExpectedGoalIndividual(awayTeamAwayRating)
}

// Generate the Head to Head Pi Ratings for the 2 teams using their
// previous match results
func ReadH2HPi(gameType string, teams []string) (float64, float64, error) {
	home := strings.ToUpper(strings.TrimSpace(teams[0]))
	away := strings.ToUpper(strings.TrimSpace(teams[1]))
	games, err := models.GetGamesV3(gameType, home, away)
	if err != nil {
		return 0, 0, err
	}
	homeTeam := models.GenerateInstantPi(games, home).ProvisionalRating("home").HomeRating
	awayTeam := models.GenerateInstantPi(games, away).ProvisionalRating("away").AwayRating

	return homeTeam, awayTeam, nil
}

// ReadPiV2 IS UNUSED
func ReadPiV2(gameType, home, away string) (float64, float64, error) {
	hT, err := pi.Search(models.GetBaseGameType("ratings", gameType), home)
	if err != nil {
		return 0, 0, err
	}
	aT, err := pi.Search(models.GetBaseGameType("ratings", gameType), away)
	if err != nil {
		return 0, 0, err
	}
	hT = hT.ProvisionalRating("home")
	aT = aT.ProvisionalRating("away")
	return hT.HomeRating, aT.AwayRating, nil
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
	hT, aT, err := pi.UpdateTeamRatings(models.GetBaseGameType("ratings", gameType), homeTeam, awayTeam, homeScore, awayScore)
	if err != nil {
		return err
	}
	hT.WriteRatings(models.GetBaseGameType("ratings", gameType))
	aT.WriteRatings(models.GetBaseGameType("ratings", gameType))
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
