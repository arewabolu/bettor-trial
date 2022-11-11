package models

import (
	"encoding/csv"
	"os"
	"strconv"
)

//func createFile() {
//	file, err := os.OpenFile("scoreRecords.csv", os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		panic("unable to create file to save records")
//	}
//	wr := csv.NewWriter(file)
//	defer wr.Flush()
//	wr.Write([]string{"homeTeam", "awayTeam", "homeScore", "awayScore"})
//}

// searches for fixtures that matches users request
func SplitData(homeTeam, awayTeam string, data []Data) (validGames []Data) {
	for _, game := range data {
		if homeTeam == game.home && awayTeam == game.away {
			validGames = append(validGames, game)
		}
		if game.away == homeTeam && game.home == awayTeam {
			validGames = append(validGames, game)
		}
	}
	return
}

// combiner for readRecords,splitRecords, and splitData
func GetGames(filePath *string, homeTeam, awayTeam string) (validGames []Data) {
	records := ReadRecords(filePath)
	data := splitRecords(records)
	validGames = SplitData(homeTeam, awayTeam, data)
	return
}

// goals per game
// goals per game against

// is as record of the number of goals scored between both teams

func HomeTeamScores(filePath *string, HT, AT string) []int {
	HTMatches := make([]int, 0)
	games := GetGames(filePath, HT, AT)
	for _, match := range games {
		if match.home == HT {
			HTMatches = append(HTMatches, match.homeScore)
		}
		if match.away == HT {
			HTMatches = append(HTMatches, match.awayScore)
		}
	}
	return HTMatches
}

func AwayTeamScores(filePath *string, HT, AT string) []int {
	ATMatches := make([]int, 0)
	games := GetGames(filePath, HT, AT)
	for _, match := range games {
		if match.away == AT {
			ATMatches = append(ATMatches, match.awayScore)
		}
		if match.home == AT {
			ATMatches = append(ATMatches, match.homeScore)
		}
	}
	return ATMatches
}

// Returns the average goals scored by both teams

func justOnce() (err error) {
	file, err := os.OpenFile("scoreRecords.csv", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	//home vs away = hs vs as
	wr := csv.NewWriter(file)
	defer wr.Flush()
	records := []Data{
		//{,,,},
	}

	//input err home!=away team
	wrData := make([][]string, len(records))
	for _, record := range records {
		row := []string{record.home, record.away, strconv.Itoa(record.homeScore), strconv.Itoa(record.awayScore)}
		wrData = append(wrData, row)
	}
	wr.WriteAll(wrData)
	return nil
}

// home team advantage vs away teams
func CompileAdvantage(data []Data) (homeAdv, awayAdv int) {
	for _, v := range data {
		if v.homeScore == v.awayScore {
			continue
		}
		if v.homeScore < v.awayScore {
			awayAdv++
		}
		if v.homeScore > v.awayScore {
			homeAdv++
		}
	}
	return
}

// returns the percentage of wins for each team
func PercentageWins(homeTeam, awayTeam string, games []Data) []float64 {
	//(games won/total games played) * 100
	var homeTeamWins float64
	var awayTeamWins float64
	var draws float64
	divider := float64(len(games))

	for _, fixture := range games {
		if homeTeam == fixture.home && fixture.homeScore > fixture.awayScore {
			homeTeamWins++
		}
		if homeTeam == fixture.away && fixture.awayScore > fixture.homeScore {
			homeTeamWins++
		}
		if awayTeam == fixture.home && fixture.awayScore > fixture.homeScore {
			awayTeamWins++
		}
		if awayTeam == fixture.home && fixture.homeScore > fixture.awayScore {
			awayTeamWins++
		}
		if fixture.homeScore == fixture.awayScore || fixture.awayScore == fixture.homeScore {
			draws++
		}
	}
	homeTeamPcnt := PercentageCalc(homeTeamWins, divider)
	drawPcnt := PercentageCalc(draws, divider)
	awayTeamPcnt := AwayPercentCalc(homeTeamPcnt, drawPcnt)
	return []float64{homeTeamPcnt, awayTeamPcnt, drawPcnt}
}

func ScorePercentage4x4(homeTeam, awayTeam string, games []Data) []float64 {
	var gameScore1 float64
	var gameScore2 float64
	var gameScore3 float64
	divider := float64(len(games))

	for _, fixture := range games {
		if fixture.homeScore+fixture.awayScore > 6 {
			gameScore1++
		}
		if fixture.homeScore+fixture.awayScore > 7 {
			gameScore2++
		}
		if fixture.homeScore+fixture.awayScore > 8 {
			gameScore3++
		}
	}

	gameScore1Per := PercentageCalc(gameScore1, divider)
	gameScore2Per := PercentageCalc(gameScore2, divider)
	gameScore3Per := PercentageCalc(gameScore3, divider)
	return []float64{gameScore1Per, gameScore2Per, gameScore3Per}
}

func ScorePercentagePen(homeTeam, awayTeam string, games []Data) []float64 {
	var gameScore1 float64
	var gameScore2 float64
	var gameScore3 float64
	divider := float64(len(games))

	for _, fixture := range games {
		if fixture.homeScore >= 1 && fixture.awayScore >= 1 {
			gameScore1++
		}
		if fixture.homeScore >= 2 && fixture.awayScore >= 2 {
			gameScore2++
		}
		if fixture.homeScore >= 3 && fixture.awayScore >= 3 {
			gameScore3++
		}
	}

	gameScore1Per := PercentageCalc(gameScore1, divider)
	gameScore2Per := PercentageCalc(gameScore2, divider)
	gameScore3Per := PercentageCalc(gameScore3, divider)
	return []float64{gameScore1Per, gameScore2Per, gameScore3Per}
}
