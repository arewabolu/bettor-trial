package main

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
	home, away           string
	homeScore, awayScore int
}

// writematchdata registers new match records
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

// reads fixtures from file in filepath into an array
func readRecords(gameType *string) (records [][]string) {

	file, err := os.Open(filePath[*gameType])
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

// checkValidFixtures only logs fixtures are not registered or not enough
func checkifReg(gameType, homeTeam, awayTeam *string, data []Data) error {
	if len(data) == 0 {
		str := "No registered games between" + *homeTeam + " & " + *awayTeam
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

func checkValidLen(gameType, homeTeam, awayTeam *string, data []Data) error {
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

// splits all fixtures into an array of match data
func splitRecords(records [][]string) []Data {
	rdData := make([]Data, 0)
	for index, record := range records {
		if index == 0 {
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
			home:      record[0],
			away:      record[1],
			homeScore: homeScores,
			awayScore: awayScores,
		}
		rdData = append(rdData, data)
	}
	return rdData
}
