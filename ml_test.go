package main

import (
	"bettor/models"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/arewabolu/csvmanager"
)

type invests struct {
	balance float64
}

func (i *invests) bet(stake float64) {
	winning := stake * 1.35
	i.balance = i.balance - stake
	i.balance = i.balance + winning
}

func (i *invests) betLoss(stake float64) {
	i.balance = i.balance - stake
}

func TestClassEV(t *testing.T) {
	teams := []string{"ARS", "PSG", "BAY", "BAR", "RMA", "JUV", "LIV", "MCI"}
	game := "fifa22Pen"
	rds, err := csvmanager.ReadCsv(models.GetBase()+game+".csv", true)
	if err != nil {
		t.Error("failed to open game file", err)
	}

	for _, team := range teams {
		inv := &invests{
			balance: 500,
		}
		stake := 50.00
		var even, odd float64
		teamGoals := models.SearchTeam(team, rds)
		teamGoalsAgainst := models.SearchTeam3(team, rds)
		for i := 0; i < len(teamGoals); i++ { // len(teamGoals)
			sum := teamGoals[i] + teamGoalsAgainst[i]
			switch {
			case sum == 0:
				continue
			case sum%2 == 0 && inv.balance >= 50.0:
				even++
				inv.betLoss(stake)
			case sum%2 != 0 && inv.balance >= 50.0:
				odd++
				inv.bet(stake)
			}
		}
		evenprob := (even / (even + odd))
		oddProb := 1 - evenprob
		EV := (oddProb * 58) - (evenprob * 50)
		t.Error(team, ":", fmt.Sprintf("%.2f", EV), *inv)
	}
}

func TestClassEV2(t *testing.T) {
	teams := []string{"AVL", "ARS", "BHA", "BRE", "BUR", "CHE", "CRY", "EVE", "LEI", "LIV", "LU", "MCI", "MU", "NOR", "NU", "SOU", "TOT", "WAT", "WHU", "WOL"}
	game := "fifa4x4Eng"
	rds, _ := csvmanager.ReadCsv(models.GetBase()+game+".csv", true)
	for _, team := range teams {
		potentialPay := 30
		var even, odd float64
		teamGoals := models.SearchTeam(team, rds)
		teamGoalsAgainst := models.SearchTeam3(team, rds)
		for i := 0; i < 20; i++ {
			sum := teamGoals[i] + teamGoalsAgainst[i]
			switch {
			case sum == 0:
				continue
			case sum%2 == 0:
				even++
				potentialPay = potentialPay + 24
			default:
				odd++
				potentialPay = potentialPay - 30
			}
		}
		evenprob := (even / (even + odd))
		oddProb := 1 - evenprob
		EV := (evenprob * 36) - (oddProb * 30)
		t.Error(team, ":", EV, potentialPay)
	}
}

func TestClassEV3(t *testing.T) {
	teams := []string{"ARS", "PSG", "BAY", "BAR", "RMA", "JUV", "LIV", "MCI"}
	game := "fifa22Pen"

	for _, team := range teams {
		if team == models.Liv {
			continue
		}
		for _, team2 := range teams {
			if team == team2 {
				continue
			}
			games := models.GetGames(&game, team, team2)
			//if len(games) < 30 {
			//	continue
			//}
			potentialPay := 500
			var even, odd float64
			for _, game := range games {
				sum := game.HomeScore + game.AwayScore
				switch {
				case sum == 0:
					continue
				case sum%2 == 0 && potentialPay > 0:
					even++
					potentialPay = potentialPay - 50
				case sum%2 != 0 && potentialPay > 0:
					odd++
					potentialPay = potentialPay + 18

				}
			}
			evenprob := (even / (even + odd))
			oddProb := 1 - evenprob
			EV := (evenprob * 60) - (oddProb * 30)
			t.Error(team, team2, ":", fmt.Sprintf("%.2f", EV), potentialPay)
		}
	}
}

func TestClassEV4(t *testing.T) {
	team1 := models.Liv
	team2 := models.Bur
	game := "fifa4x4Eng" //"fifa22Pen" ""
	games := models.GetGames(&game, team1, team2)
	//if len(games) < 30 {
	//	t.Error("not enough games")
	//		t.FailNow()
	//}
	potentialPay := 100
	var even, odd float64
	for _, game := range games {
		sum := game.HomeScore + game.AwayScore
		switch {
		case sum == 0:
			continue
		case sum%2 == 0 && potentialPay > 0:
			even++
			potentialPay = potentialPay + 30
		default:
			odd++
			potentialPay = potentialPay - 30
		}
	}
	evenprob := (even / (even + odd))
	oddProb := 1 - evenprob
	EV := (evenprob * 60) - (oddProb * 30)
	t.Error(team1, team2, ":", fmt.Sprintf("%.2f", EV), potentialPay, even, odd)

}

func TestClassEV5(t *testing.T) {
	teams := []string{"AVL", "ARS", "BHA", "BRE", "BUR", "CHE", "CRY", "EVE", "LEI", "LIV", "LU", "MCI", "MU", "NOR", "NU", "SOU", "TOT", "WAT", "WHU", "WOL"}
	game := "fifa4x4Eng" //"fifa22Pen"
	//if len(games) < 30 {
	//	t.Error("not enough games")
	//		t.FailNow()
	//}
	file, _ := os.OpenFile("HvsA.csv", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	wr := csvmanager.WriteFrame{
		File:    file,
		Headers: []string{"team1", "team2", "evenPay", "oddPay"},
	}
	sl1 := make([]string, 0)
	sl2 := make([]string, 0)
	sl3 := make([]string, 0)
	sl4 := make([]string, 0)
	for _, team := range teams {
		for _, team2 := range teams {
			if team == team2 {
				continue
			}
			games := models.GetGames(&game, team, team2)
			//if len(games) < 30 {
			//	continue
			//}
			potentialPay1 := 500
			potentialPay2 := 500
			var even, odd float64
			for _, game := range games {
				sum := game.HomeScore + game.AwayScore
				switch {
				case sum == 0:
					continue
				case sum%2 == 0 && potentialPay1 > 0:
					even++
					potentialPay1 = potentialPay1 + 45
				case sum%2 != 0 && potentialPay2 > 0:
					odd++
					potentialPay2 = potentialPay2 + 45

				}
			}
			evenprob := (even / (even + odd))
			oddProb := 1 - evenprob
			EV := (evenprob * 60) - (oddProb * 30)

			ev := strconv.Itoa(potentialPay1)
			od := strconv.Itoa(potentialPay2)
			sl1 = append(sl1, team)
			sl2 = append(sl2, team2)
			sl3 = append(sl3, ev)
			sl4 = append(sl4, od)

			t.Error(team, team2, ":", fmt.Sprintf("%.2f", EV), potentialPay1, potentialPay2)
		}
	}
	wr.Columns = append(wr.Columns, sl1, sl2, sl3, sl4)
	wr.WriteCSV()
}
