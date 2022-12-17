package models

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

// for calculating odds
func OddsCalc(winProb, drawProb float64) float64 {
	rlWinorDraw := winProb + drawProb
	rlProb := rlWinorDraw / 100
	denum := 1 - rlProb
	odds := rlProb / denum
	return odds
}
