package models

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/arewabolu/csvmanager"
	"golang.org/x/exp/slices"
)

// Avl
// minimum of 190 non repetitive games needed
const (
	Avl = "AVL"
	Ars = "ARS"
	Bha = "BHA"
	Bre = "BRE"
	Bur = "BUR"
	Che = "CHE"
	Cry = "CRY"
	Eve = "EVE"
	Lei = "LEI"
	Liv = "LIV"
	Lu  = "LU"
	Mci = "MCI"
	Mu  = "MU"
	Nor = "NOR"
	Nu  = "NU"
	Sou = "SOU"
	Tot = "TOT"
	Wat = "WAT"
	Whu = "WHU"
	Wol = "WOL"
	Psg = "PSG"
	Bay = "BAY"
	Bar = "BAR"
	Rma = "RMA"
	Juv = "JUV"
)

type Data struct {
	Home      string
	Away      string
	HomeScore int
	AwayScore int
}

type TeamData struct {
	GoalFor, GoalAgainst, Status []float64
}

func CreateFile(name string) error {
	if name == "" {
		return errors.New("please state the name of the file")
	}
	file, err := os.OpenFile(GetBase()+name+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return err
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	err = wr.Write([]string{"homeTeam", "awayTeam", "homeScore", "awayScore"})
	if err != nil {
		return err
	}
	return nil
}

func CreateTeamsFile(gameName string) error {
	if gameName == "" {
		return errors.New("please state the name of the file")
	}
	file, err := os.OpenFile(GetBaseTeamNames()+gameName+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return err
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	err = wr.Write([]string{"Teams"})
	if err != nil {
		return err
	}
	return nil
}

func CreateRatingFile(gameName string) error {
	if gameName == "" {
		return errors.New("please state the name of the file")
	}
	file, err := os.OpenFile(GetBaseGameType("ratings", gameName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return err
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	err = wr.Write([]string{"TeamName", "HomeRating", "AwayRating", "ContinuousHomePerformance", "ContinuousAwayPerformance"})
	if err != nil {
		return err
	}
	return nil
}

func AddToTeam(gameName, teamName string) error {
	teamName = strings.ToUpper(strings.TrimSpace(teamName))
	if teamName == "" {
		return errors.New("please state the name of the team")
	}
	_, err := os.Stat(GetBaseTeamNames() + gameName + ".csv")
	if errors.Is(err, os.ErrNotExist) {
		CreateTeamsFile(gameName)
	}
	file, err := os.OpenFile(GetBaseTeamNames()+gameName+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return err
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	err = wr.Write([]string{teamName})
	if err != nil {
		return err
	}
	return nil
}

func AddtoRating(gameName, teamName string) error {
	teamName = strings.ToUpper(strings.TrimSpace(teamName))
	if teamName == "" {
		return errors.New("please state the name of the team")
	}
	_, err := os.Stat(GetBaseRating() + "ratings" + gameName + ".csv")
	if errors.Is(err, os.ErrNotExist) {
		CreateRatingFile(gameName)
	}
	file, err := os.OpenFile(GetBaseRating()+"ratings"+gameName+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
	if err != nil {
		return err
	}
	defer file.Close()

	wr := csv.NewWriter(file)
	defer wr.Flush()

	err = wr.Write([]string{teamName, "0", "0", "0", "0"})
	if err != nil {
		return err
	}
	return nil
}

// check if a team is registered to prevent errors when entering data
func CheckifReg(gameType, homeTeam, awayTeam string) error {
	_, err := os.Stat(GetBaseTeamNames() + gameType + ".csv")
	if errors.Is(err, os.ErrNotExist) {
		CreateTeamsFile(gameType)
	}
	reader, readErr := csvmanager.ReadCsv(GetBaseTeamNames()+gameType+".csv", 0700, true, 20)
	if readErr != nil {
		return readErr
	}
	teamNames := reader.Col("Teams").String()
	if len(teamNames) < 1 {
		return errors.New("no teams currently registered. please update teams for league")
	}
	if !slices.Contains(teamNames, homeTeam) {
		return errors.New("invalid home team entered")
	}
	if !slices.Contains(teamNames, awayTeam) {
		return errors.New("invalid away team entered")
	}
	return nil
}

func DirIterator(basedir string) []string {
	folder, _ := os.ReadDir(basedir)
	nameSlice := make([]string, 0)
	for _, dirFile := range folder {
		if strings.HasSuffix(dirFile.Name(), ".csv") {
			name := strings.TrimSuffix(dirFile.Name(), ".csv")
			nameSlice = append(nameSlice, name)
		}
	}
	return nameSlice
}

func FloattoString(x []float64) []string {
	strFloatArr := make([]string, 0, len(x))
	for i := range x {
		strFloat := strconv.FormatFloat(x[i], 'f', 2, 64)
		strFloatArr = append(strFloatArr, strFloat)
	}
	return strFloatArr
}
