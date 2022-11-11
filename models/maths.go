package models

func HomeAwayAvg(filePath *string, HT, AT string, games []Data) (homeAvg, awayAvg float64) {
	homeSum := float64(sumArray(HomeTeamScores(filePath, HT, AT)))
	awaySum := float64(sumArray(AwayTeamScores(filePath, HT, AT)))
	length := float64(len(games))
	homeAvg = homeSum / length
	awayAvg = awaySum / length
	return
}

func Trials(games []Data) string {
	var evens int
	var odd int
	for _, fixture := range games {
		matchScore := fixture.homeScore + fixture.awayScore
		if matchScore%2 == 0 {
			evens++
		} else {
			odd++
		}
	}
	if evens > odd {
		return "even"
	}
	return "odd"
}

func PercentageCalc(n, divider float64) float64 {
	multiplier := n * 100
	percentage := multiplier / divider
	return percentage
}

func AwayPercentCalc(HP, drawP float64) float64 {
	AVal := HP + drawP
	AP := 100 - AVal
	return AP
}
