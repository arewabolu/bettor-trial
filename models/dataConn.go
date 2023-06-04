package models

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/arewabolu/csvmanager"
	"gonum.org/v1/gonum/stat"
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
	file, err := os.OpenFile(GetBase()+gameName+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
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

func AddTeam(gameName, teamName string) error {
	if teamName == "" {
		return errors.New("please state the name of the team")
	}
	file, err := os.OpenFile(GetBase()+gameName+".csv", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
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

// checkValidFixtures only logs fixtures are not registered or not enough
func CheckifReg(gameType, homeTeam, awayTeam *string, data []Data) error {
	if len(data) == 0 {
		str := "No registered games between " + *homeTeam + " & " + *awayTeam
		printErr := errors.New(str)
		path := "./unregistered/" + *gameType + ".txt"
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		w := bufio.NewWriter(file)
		fmt.Fprintf(w, "Match data for %s vs  %s does not currently exist\n", *homeTeam, *awayTeam)
		w.Flush()
		return printErr
	}
	return nil
}

func CheckValidLen(gameType, homeTeam, awayTeam *string, data []Data) error {
	if len(data) < 5 {
		str := "Not enough registered fixtures between " + *homeTeam + " & " + *awayTeam + "!"
		printErr := errors.New(str)
		//path := "./insufficient/" + *gameType + ".txt"
		//file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0700)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer file.Close()
		//w := bufio.NewWriter(file)
		//fmt.Fprintf(w, "only %v fixtures for %s vs  %s are availble\n", len(data), *homeTeam, *awayTeam)
		//w.Flush()
		return printErr
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

func TeamAvgs(f csvmanager.Frame) (float64, float64) {
	HMGl := SearchTeam("MCI", f)
	meanTeamC := stat.CircularMean(FloatCon(HMGl), nil)
	meanTeam := stat.Mean(FloatCon(HMGl), nil)
	//HS := rds.Col("homeScore").Float()
	//AS := rds.Col("awayScore").Float()
	//mean, _ := stat.MeanVariance(HS, nil)
	//mean2, _ := stat.MeanVariance(AS, nil)
	return meanTeam, meanTeamC
}

func FloatCon(slice []int) []float64 {
	floatSli := make([]float64, 0)
	for _, input := range slice {
		floatSli = append(floatSli, float64(input))
	}
	return floatSli
}

func MeanDiff(x []float64, mean float64) []float64 {
	diff := make([]float64, 0, len(x))
	for index := range x {
		diff = append(diff, x[index]-mean)
	}
	return diff
}

func FloattoString(x []float64) []string {
	strFloatArr := make([]string, 0, len(x))
	for i := range x {
		strFloat := strconv.FormatFloat(x[i], 'f', 2, 64)
		strFloatArr = append(strFloatArr, strFloat)
	}
	return strFloatArr
}
