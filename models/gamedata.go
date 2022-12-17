package models

// searches for fixtures that matches users request
func SplitData(homeTeam, awayTeam string, data []Data) (validGames []Data) {
	for _, game := range data {
		if homeTeam == game.Home && awayTeam == game.Away {
			validGames = append(validGames, game)
		}
		if homeTeam == game.Away && awayTeam == game.Home {
			validGames = append(validGames, game)
		}
	}
	return
}

// combiner for readRecords,splitRecords, and splitData
func GetGames(gameType *string, homeTeam, awayTeam string) (validGames []Data) {
	records := ReadRecords(*gameType)
	data := splitRecords(records)
	validGames = SplitData(homeTeam, awayTeam, data)
	return
}

// goals per game
// goals per game against

// is as record of the number of goals scored by the hometeam
func HomeTeamScores(filePath string, HT, AT string) []int {
	HTMatches := make([]int, 0)
	games := GetGames(&filePath, HT, AT)
	for _, match := range games {
		if match.Home == HT {
			HTMatches = append(HTMatches, match.HomeScore)
		}
		if match.Away == HT {
			HTMatches = append(HTMatches, match.AwayScore)
		}
	}
	return HTMatches
}

// is as record of the number of goals scored by the awayteam
func AwayTeamScores(filePath string, HT, AT string) []int {
	ATMatches := make([]int, 0)
	games := GetGames(&filePath, HT, AT)
	for _, match := range games {
		if match.Home == AT {
			ATMatches = append(ATMatches, match.AwayScore)
		}
		if match.Away == AT {
			ATMatches = append(ATMatches, match.HomeScore)
		}
	}
	return ATMatches
}

// Returns the average goals scored by both teams

// returns the percentage of wins and the percentage draw for each team respectively
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
		if fixture.HomeScore == oneDif(fixture.AwayScore) || fixture.AwayScore == oneDif(fixture.HomeScore) && gametype == "fifa22Pen" || gametype == "fifa18Pen" {
			continue
		}
	}
	homeTeamWinPcnt := PercentageCalc(homeTeamWins, divider)
	drawPcnt := PercentageCalc(draws, divider)
	awayTeamWinPcnt := AwayPercentCalc(homeTeamWinPcnt, drawPcnt)
	return []float64{homeTeamWinPcnt, awayTeamWinPcnt, drawPcnt}
}
