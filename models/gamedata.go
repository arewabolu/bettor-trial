package models

//func createFile() {
//	file, err := os.OpenFile("scoreRecords.csv", os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//		panic("unable to create file to save records")
//	}
//	wr := csv.NewWriter(file)
//	defer wr.Flush()
//	wr.Write([]string{"homeTeam", "awayTeam", "HomeScore", "AwayScore"})
//}

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
func GetGames(gameType string, homeTeam, awayTeam string) (validGames []Data) {
	records := ReadRecords(gameType)
	data := splitRecords(records)
	validGames = SplitData(homeTeam, awayTeam, data)
	return
}

// goals per game
// goals per game against

// is as record of the number of goals scored between both teams

func HomeTeamScores(filePath string, HT, AT string) []int {
	HTMatches := make([]int, 0)
	games := GetGames(filePath, HT, AT)
	for _, match := range games {
		if match.Home == HT {
			HTMatches = append(HTMatches, match.HomeScore)
		}
		if match.Home == HT {
			HTMatches = append(HTMatches, match.AwayScore)
		}
	}
	return HTMatches
}

func AwayTeamScores(filePath string, HT, AT string) []int {
	ATMatches := make([]int, 0)
	games := GetGames(filePath, HT, AT)
	for _, match := range games {
		if match.Home == AT {
			ATMatches = append(ATMatches, match.AwayScore)
		}
		if match.Home == AT {
			ATMatches = append(ATMatches, match.HomeScore)
		}
	}
	return ATMatches
}

// Returns the average goals scored by both teams

// Home team advantage vs Home teams
func CompileAdvantage(data []Data) (homeAdv, awayAdv int) {
	for _, v := range data {
		if v.HomeScore == v.AwayScore {
			continue
		}
		if v.HomeScore < v.AwayScore {
			awayAdv++
		}
		if v.HomeScore > v.AwayScore {
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
		if fixture.HomeScore+fixture.AwayScore > 6 {
			gameScore1++
		}
		if fixture.HomeScore+fixture.AwayScore > 7 {
			gameScore2++
		}
		if fixture.HomeScore+fixture.AwayScore > 8 {
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
		if fixture.HomeScore >= 1 && fixture.AwayScore >= 1 {
			gameScore1++
		}
		if fixture.HomeScore >= 2 && fixture.AwayScore >= 2 {
			gameScore2++
		}
		if fixture.HomeScore >= 3 && fixture.AwayScore >= 3 {
			gameScore3++
		}
	}

	gameScore1Per := PercentageCalc(gameScore1, divider)
	gameScore2Per := PercentageCalc(gameScore2, divider)
	gameScore3Per := PercentageCalc(gameScore3, divider)
	return []float64{gameScore1Per, gameScore2Per, gameScore3Per}
}
