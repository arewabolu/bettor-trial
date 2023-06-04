package models

import (
	"fmt"

	"github.com/arewabolu/csvmanager"
)

// combiner for readRecords and splitData
func GetGames(gameType *string, homeTeam, awayTeam string) (validGames []Data) {
	file, err := csvmanager.ReadCsv(GetBase()+*gameType+".csv", 0755, true, 400)
	if err != nil {
		fmt.Println(err)
	}
	//	data := splitRecords(records)
	validGames = SplitData(homeTeam, awayTeam, file)
	return
}

func PercentageWinorDraw(gametype, homeTeam, awayTeam string, games []Data) []float64 {
	//(games won/total games played) * 100
	var homeTeamWins float64
	var awayTeamWins float64
	var draws float64
	divider := float64(len(games))
	oneDif := func(x int) int {
		return x - 1
	}

	for _, fixture := range games {
		if homeTeam == fixture.Home && fixture.HomeScore > fixture.AwayScore {
			homeTeamWins++
		}
		if homeTeam == fixture.Home && fixture.AwayScore > fixture.HomeScore {
			homeTeamWins++
		}
		if awayTeam == fixture.Home && fixture.AwayScore > fixture.HomeScore {
			awayTeamWins++
		}
		if awayTeam == fixture.Home && fixture.HomeScore > fixture.AwayScore {
			awayTeamWins++
		}
		if fixture.HomeScore == fixture.AwayScore || fixture.AwayScore == fixture.HomeScore {
			draws++
		}
		if fixture.HomeScore == oneDif(fixture.AwayScore) && gametype == "fifa18Pen" {
			continue
		}
		if fixture.HomeScore == oneDif(fixture.AwayScore) && gametype == "fifa22Pen" {
			continue
		}
		if fixture.AwayScore == oneDif(fixture.HomeScore) && gametype == "fifa18Pen" {
			continue
		}
		if fixture.AwayScore == oneDif(fixture.HomeScore) && gametype == "fifa22Pen" {
			continue
		}
		//|| fixture.AwayScore == oneDif(fixture.HomeScore)
		//gametype == "fifa22Pen"
	}
	homeTeamWinPcnt := PercentageCalc(homeTeamWins, divider)
	drawPcnt := PercentageCalc(draws, divider)
	awayTeamWinPcnt := AwayPercentCalc(homeTeamWinPcnt, drawPcnt)
	return []float64{homeTeamWinPcnt, awayTeamWinPcnt, drawPcnt}
}

// returns a single teams home and away goals
func SearchTeam(team string, data csvmanager.Frame) (validGoals []int) {
	for _, game := range data.Rows() {
		nwData := &Data{}
		game.Interface(nwData)
		if team == nwData.Home {
			validGoals = append(validGoals, nwData.HomeScore)
		}
		if team == nwData.Away {
			validGoals = append(validGoals, nwData.AwayScore)
		}
	}
	return
}

// here a teams home goals and away goals a returned seperately
func SearchTeam2(team string, data csvmanager.Frame) (homeGoals []int, awayGoals []int) {
	for _, game := range data.Rows() {
		nwData := &Data{}
		game.Interface(nwData)

		if team == nwData.Home {
			homeGoals = append(homeGoals, nwData.HomeScore)
		}
		if team == nwData.Away {
			awayGoals = append(awayGoals, nwData.AwayScore)
		}
	}
	return
}

// returns opponents goals for a single team
// to replace GetGames especially when testing
func SearchTeam3(team string, data csvmanager.Frame) (homeGoals []int) {
	for _, game := range data.Rows() {
		nwData := &Data{}
		game.Interface(nwData)

		switch {
		case team == nwData.Home:
			homeGoals = append(homeGoals, nwData.AwayScore)
		case team == nwData.Away:
			homeGoals = append(homeGoals, nwData.HomeScore)
		}
	}
	return
}

func SearchTeam4(team1, team2 string, game string) (float64, float64) {
	games := GetGames(&game, team1, team2)
	var even, odd float64
	for _, game := range games {
		sum := game.HomeScore + game.AwayScore
		switch {
		case sum == 0:
			continue
		case sum%2 == 0:
			even++

		case sum%2 != 0:
			odd++
		}
	}
	evenProb := (even / (even + odd)) * 100
	oddProb := 100 - evenProb

	return oddProb, evenProb
}

// goals per game
// goals per game against

// searches for fixtures that matches users request
func SplitData(homeTeam, awayTeam string, data csvmanager.Frame) (validGames []Data) {

	for _, game := range data.Rows() {
		nwData := &Data{}
		game.Interface(nwData)
		if homeTeam == nwData.Home && awayTeam == nwData.Away {
			validGames = append(validGames, *nwData)
		}
		if homeTeam == nwData.Away && awayTeam == nwData.Home {
			validGames = append(validGames, *nwData)
		}
	}

	return
}

// is as record of the number of goals scored by the both teams
func TeamGoals(filePath string, HT, AT string) (HTMatches, ATMatches []int) {
	games := GetGames(&filePath, HT, AT)
	for _, match := range games {
		if match.Home == HT {
			HTMatches = append(HTMatches, match.HomeScore)
		}
		if match.Away == HT {
			HTMatches = append(HTMatches, match.AwayScore)
		}
		if match.Home == AT {
			ATMatches = append(ATMatches, match.AwayScore)
		}
		if match.Away == AT {
			ATMatches = append(ATMatches, match.HomeScore)
		}
	}
	return
}

func Wins(team string, data csvmanager.Frame) (wins []int) {
	for _, game := range data.Rows() {
		nwData := &Data{}
		game.Interface(nwData)

		if team == nwData.Home && nwData.HomeScore > nwData.AwayScore {
			wins = append(wins, nwData.AwayScore)
		}
		if team == nwData.Away && nwData.AwayScore > nwData.HomeScore {
			wins = append(wins, nwData.HomeScore)
		}
	}
	return
}

func Loss(team string, data csvmanager.Frame) (loss []int) {
	for _, game := range data.Rows() {
		nwData := &Data{}
		game.Interface(nwData)

		if team == nwData.Home && nwData.HomeScore < nwData.AwayScore {
			loss = append(loss, nwData.AwayScore)
		}
		if team == nwData.Away && nwData.AwayScore < nwData.HomeScore {
			loss = append(loss, nwData.HomeScore)
		}
	}
	return
}

// Returns the average goals scored by both teams

// returns the percentage of wins and the percentage draw for each team respectively
