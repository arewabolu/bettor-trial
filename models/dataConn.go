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
)

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
	Home, Away           string
	HomeScore, AwayScore int
}

func CreateFile(name string) error {
	if name == "" {
		return errors.New("please state the name of the file")
	}
	_, err := os.Create(GetBase() + name + ".csv")
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

// reads fixtures from file in filepath into an array
func ReadRecords(gameType string) (records [][]string) {
	file, err := os.Open(GetBase() + gameType + ".csv")
	rdder := bufio.NewReaderSize(file, 400)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	rd := csv.NewReader(rdder)
	records, err = rd.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return records
}

// splits all fixtures into an array of match data
func splitRecords(records [][]string) []Data {
	rdData := make([]Data, 0)
	for index, record := range records {
		if index == 0 {
			continue
		}
		if record[0] == "" || record[1] == "" || record[2] == "" || record[3] == "" {
			continue
		}
		homeScores, err := strconv.Atoi(record[2])
		if err != nil {
			fmt.Println(err)
		}
		awayScores, err := strconv.Atoi(record[3])
		if err != nil {
			fmt.Println(err)
		}
		data := Data{
			Home:      record[0],
			Away:      record[1],
			HomeScore: homeScores,
			AwayScore: awayScores,
		}
		rdData = append(rdData, data)
	}
	return rdData
}
