package models

import (
	"fmt"

	"github.com/arewabolu/csvmanager"
	"github.com/arewabolu/pi-rating"
)

// goals per game
// goals per game against
func GenerateInstantPi(games []Data, teamName string) pi.Team {
	team := pi.Newteam(teamName)
	for _, game := range games {
		switch {
		case game.Home == teamName:
			team = pi.BuildPiforHometeam(team, game.Away, game.HomeScore, game.AwayScore)
		case game.Away == teamName:
			team = pi.BuildPiforAwayteam(team, game.Home, game.HomeScore, game.AwayScore)
		}
	}
	return team
}

// searches for fixtures that matches users request
func splitData(homeTeam, awayTeam string, data csvmanager.Frame) (validGames []Data) {

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

func splitDataV2(team string, data csvmanager.Frame) (validGames []Data) {
	for _, game := range data.Rows() {
		nwData := &Data{}
		game.Interface(nwData)
		if team == nwData.Home || team == nwData.Away {
			validGames = append(validGames, *nwData)
		}
	}

	return
}

func splitDataV3(homeTeam, awayTeam string, data csvmanager.Frame) (validGames []Data) {
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

// combiner for readRecords and splitData
func GetGames(gameType string, homeTeam, awayTeam string) (validGames []Data) {
	file, err := csvmanager.ReadCsv(GetBase()+gameType+".csv", 0755, true, 400)
	if err != nil {
		fmt.Println(err)
	}

	validGames = splitData(homeTeam, awayTeam, file)
	return
}

func GetGamesV2(gameType string, team string) ([]Data, error) {
	file, err := csvmanager.ReadCsv(GetBase()+gameType+".csv", 0755, true, 400)
	if err != nil {
		return nil, err
	}

	validGames := splitDataV2(team, file)
	return validGames, nil
}

func GetGamesV3(gameType string, homeTeam, awayTeam string) ([]Data, error) {
	file, err := csvmanager.ReadCsv(GetBase()+gameType+".csv", 0755, true, 400)
	if err != nil {
		return nil, err
	}

	validGames := splitDataV3(homeTeam, awayTeam, file)
	return validGames, nil
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

func SearchTeam4(team1, team2 string, game string) (float64, float64) {
	games := GetGames(game, team1, team2)
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

// is as record of the number of goals scored by the both teams
func TeamGoals(filePath string, HT, AT string) (HTMatches, ATMatches []int) {
	games := GetGames(filePath, HT, AT)
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
