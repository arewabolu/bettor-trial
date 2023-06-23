package main

import (
	"bettor/models"
	"fmt"
	"os"
	"testing"

	"github.com/arewabolu/csvmanager"
	"gonum.org/v1/gonum/stat"
)

type invests struct {
	balance, stake, odds float64
}

func (i *invests) bet() {
	winning := i.stake * i.odds
	i.balance = i.balance - i.stake
	i.balance = i.balance + winning
}

func (i *invests) betLoss() {
	i.balance = i.balance - i.stake
}

func TestClassEV3(t *testing.T) {
	teams := []string{"ARS", "PSG", "BAY", "BAR", "RMA", "JUV", "LIV", "MCI"}
	game := "fifa22Pen"

	for _, team := range teams {
		for _, team2 := range teams {
			if team == team2 {
				continue
			}
			games := models.GetGames(&game, team, team2)
			//if len(games) < 30 {
			//	continue
			//}
			inv := &invests{
				balance: 500,
				stake:   100.00,
				odds:    1.3,
			}
			var even, odd float64
			for _, game := range games {
				sum := game.HomeScore + game.AwayScore
				switch {
				case sum == 0:
					continue
				case sum%2 == 0 && inv.balance >= 50:
					even++
					inv.betLoss()

				case sum%2 != 0 && inv.balance >= 50:
					odd++

					inv.bet()
				}
			}
			if inv.balance > 500 {
				evenprob := (even / (even + odd))
				oddProb := 1 - evenprob
				EV := (evenprob * 60) - (oddProb * 30)
				t.Error(team, team2, ":", fmt.Sprintf("%.2f", EV), inv.balance)
			}

		}
	}
}

func TestClassEV4(t *testing.T) {
	team1 := models.Rma
	team2 := models.Liv
	game := "fifa22Pen"
	games := models.GetGames(&game, team1, team2)

	inv := &invests{
		balance: 500,
		stake:   50.00,
		odds:    1.435,
	}
	//var even, odd float64
	for _, game := range games {
		sum := game.HomeScore + game.AwayScore
		switch {
		case sum == 0:
			continue
		case sum%2 == 0 && inv.balance > 100:
			//even++
			inv.betLoss()
			//inv.bet()
		case sum%2 != 0 && inv.balance > 100:
			//odd++
			//inv.betLoss()
			inv.bet()
		}
	}
	//	evenprob := (even / (even + odd))
	//	oddProb := 1 - evenprob
	//	EV := (evenprob * 60) - (oddProb * 30)
	t.Error(team1, team2, ":", inv.balance) //fmt.Sprintf("%.2f", EV), potentialPay, even, odd

}

func TestClassEV5(t *testing.T) {
	teams := []string{"AVL", "ARS", "BHA", "BRE", "BUR", "CHE", "CRY", "EVE", "LEI", "LIV", "LU", "MCI", "MU", "NOR", "NU", "SOU", "TOT", "WAT", "WHU", "WOL"}
	game := "fifa4x4Eng" //"fifa22Pen"
	//if len(games) < 30 {
	//	t.Error("not enough games")
	//		t.FailNow()
	//}

	for _, team := range teams {
		for _, team2 := range teams {
			if team == team2 {
				continue
			}
			games := models.GetGames(&game, team, team2)
			//if len(games) < 30 {
			//	continue
			//}
			inv := &invests{balance: 500, stake: 100, odds: 1.9}
			var even, odd float64
			for _, game := range games {
				sum := game.HomeScore + game.AwayScore
				switch {
				case sum == 0:
					continue
				case sum%2 == 0 && inv.balance >= 100:
					even++
					inv.betLoss()
				case sum%2 != 0 && inv.balance >= 100:
					odd++
					inv.bet()
				}
			}
			if inv.balance > 850 {
				t.Error(team, team2, ":", inv.balance)
			}
		}

	}
}

func TestMeanVariance2(t *testing.T) {
	game := "fifa4x4Eng" //
	rds, err := csvmanager.ReadCsv(models.GetBase()+game+".csv", 0755, true)
	if err != nil {
		t.Error("failed to open game file", err)
		t.FailNow()
	}
	hS := rds.Col("homeScore").Float()
	aS := rds.Col("awayScore").Float()

	sums := make([]float64, 0)

	for i := 0; i < len(hS); i++ {
		sum := hS[i] + aS[i]
		sums = append(sums, float64(sum))
	}
	hmMean, hmVar := stat.MeanStdDev(sums, nil)

	file, _ := os.OpenFile("sumgoalsm.csv", os.O_CREATE|os.O_RDWR, 0755)
	wr := csvmanager.WriteFrame{
		Headers: []string{"totalGoals"},
		Arrays:  [][]string{models.FloattoString(sums)},
		File:    file,
	}
	wr.WriteCSV()
	//awayMean, awVar := stat.MeanStdDev(models.FloatCon(aw), nil)
	t.Error(hmMean + hmVar)
	//t.Error(awayMean, awVar)
}

func TestSimulator(t *testing.T) {
	teams := []string{"AVL", "ARS", "BHA", "BRE", "BUR", "CHE", "CRY", "EVE", "LEI", "LIV", "LU", "MCI", "MU", "NOR", "NU", "SOU", "TOT", "WAT", "WHU", "WOL"}
	game := "fifa4x4Eng" //"fifa22Pen"
	//if len(games) < 30 {
	//	t.Error("not enough games")
	//		t.FailNow()
	//}
	inv := &invests{balance: 200, stake: 100, odds: 1.9}
	count := 0
	for _, team := range teams {
		for _, team2 := range teams {
			if team == team2 {
				continue
			}

			//if len(games) < 30 {
			//	continue
			//}

			switch {
			case team == models.Bre && team2 == models.Nu:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.betLoss()
						count++
					case sum%2 != 0 && inv.balance >= 50:
						inv.bet()
						count++
					}
				}
			case team == models.Eve && team2 == models.Lei:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.betLoss()
						count++
					case sum%2 != 0 && inv.balance >= 50:
						inv.bet()
						count++
					}
				}
			case team == models.Lu && team2 == models.Mci:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.betLoss()
						count++
					case sum%2 != 0 && inv.balance >= 50:
						inv.bet()
						count++
					}
				}
			case team == models.Nu && team2 == models.Bre:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.betLoss()
						count++
					case sum%2 != 0 && inv.balance >= 50:
						inv.bet()
						count++
					}
				}
			case team == models.Lei && team2 == models.Eve:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.betLoss()
						count++
					case sum%2 != 0 && inv.balance >= 50:
						inv.bet()
						count++
					}
				}
			case team == models.Mci && team2 == models.Lu:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.betLoss()
						count++
					case sum%2 != 0 && inv.balance >= 50:
						inv.bet()
						count++
					}
				}

			}
		}
	}
	t.Error(count, inv.balance)
}

func TestSimulatorPen(t *testing.T) {
	teams := []string{"ARS", "PSG", "BAY", "BAR", "RMA", "JUV", "LIV", "MCI"}
	game := "fifa22Pen"
	//if len(games) < 30 {
	//	t.Error("not enough games")
	//		t.FailNow()
	//}
	inv := &invests{balance: 300, stake: 100, odds: 2.3}
	inv2 := &invests{balance: 300, stake: 100, odds: 1.35}
	count := 0
	for _, team := range teams {
		for _, team2 := range teams {
			if team == team2 {
				continue
			}
			switch {
			case team == models.Bay && team2 == models.Liv:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv2.balance >= 100:
						inv2.betLoss()
					case sum%2 != 0 && inv2.balance >= 100:
						inv2.bet()
						count++
					}
				}
			case team == models.Mci && team2 == models.Rma:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv2.balance >= 100:
						inv2.betLoss()
					case sum%2 != 0 && inv2.balance >= 100:
						inv2.bet()
						count++
					}
				}
			case team == models.Bay && team2 == models.Ars:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.bet()
					case sum%2 != 0 && inv.balance >= 50:
						inv.betLoss()
					}
				}
			case team == models.Psg && team2 == models.Liv:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.bet()
					case sum%2 != 0 && inv.balance >= 50:
						inv.betLoss()

					}
				}
			case team == models.Liv && team2 == models.Bay:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv2.balance >= 100:
						inv2.betLoss()
					case sum%2 != 0 && inv2.balance >= 100:
						inv2.bet()
						count++
					}
				}
			case team == models.Rma && team2 == models.Mci:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv2.balance >= 100:
						inv2.betLoss()
					case sum%2 != 0 && inv2.balance >= 100:
						inv2.bet()
						count++
					}
				}
			case team == models.Ars && team2 == models.Bay:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.bet()

					case sum%2 != 0 && inv.balance >= 50:
						inv.betLoss()
					}
				}
			case team == models.Liv && team2 == models.Psg:
				games := models.GetGames(&game, team, team2)
				for _, game := range games {
					sum := game.HomeScore + game.AwayScore
					switch {
					case sum == 0:
						continue
					case sum%2 == 0 && inv.balance >= 50:
						inv.bet()
					case sum%2 != 0 && inv.balance >= 50:
						inv.betLoss()
					}
				}
			}
		}
	}
	t.Error(count, inv.balance, inv2.balance)
}
