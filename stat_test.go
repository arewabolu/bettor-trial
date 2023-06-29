package main

import (
	"bettor/models"
	"strconv"
	"testing"

	"github.com/arewabolu/csvmanager"
	"github.com/arewabolu/pi-rating"
)

// RESULTS: You can make positive returns if you always ignore Hometeams that underperform at home alone no matter what.
// In fact, if you bet on underperforming hometeams, you should win only 35% of the time
// Further tests reveals that  in the case of underperforming hometeams you can also win a bet
// Still in testing: confirm how many goals are scored by underperforming hometeams.
func TestSearchWithPi(t *testing.T) {
	inv := &invests{balance: 200, stake: 50, odds: 1.32}
	inv2 := &invests{balance: 200, stake: 50, odds: 1.17}
	games := []models.Data{
		{Home: "MCI", Away: "BHA", HomeScore: 1, AwayScore: 2},
		{Home: "WHU", Away: "BHA", HomeScore: 0, AwayScore: 1},
		{Home: "NU", Away: "WHU", HomeScore: 3, AwayScore: 1},
		{Home: "BRE", Away: "MCI", HomeScore: 0, AwayScore: 1},
		{Home: "BHA", Away: "AVL", HomeScore: 2, AwayScore: 3},
		{Home: "CRY", Away: "WAT", HomeScore: 3, AwayScore: 1},
		{Home: "CHE", Away: "NOR", HomeScore: 0, AwayScore: 0},
		{Home: "LEI", Away: "SOU", HomeScore: 1, AwayScore: 0},
		{Home: "LU", Away: "ARS", HomeScore: 2, AwayScore: 1},
		{Home: "EVE", Away: "MU", HomeScore: 2, AwayScore: 2},
		{Home: "BUR", Away: "TOT", HomeScore: 0, AwayScore: 2},
		{Home: "LIV", Away: "WOL", HomeScore: 1, AwayScore: 1},
		{Home: "LIV", Away: "NU", HomeScore: 3, AwayScore: 1},
		{Home: "WOL", Away: "BUR", HomeScore: 1, AwayScore: 4},
		{Home: "TOT", Away: "EVE", HomeScore: 2, AwayScore: 0},
		{Home: "MU", Away: "LU", HomeScore: 1, AwayScore: 1},
		{Home: "ARS", Away: "LEI", HomeScore: 0, AwayScore: 1},
		{Home: "SOU", Away: "CHE", HomeScore: 1, AwayScore: 3},
		{Home: "NOR", Away: "CRY", HomeScore: 1, AwayScore: 4},
		{Home: "WAT", Away: "BHA", HomeScore: 1, AwayScore: 1},
		{Home: "AVL", Away: "BRE", HomeScore: 2, AwayScore: 3},
		{Home: "MCI", Away: "WHU", HomeScore: 2, AwayScore: 0},
		{Home: "NU", Away: "MCI", HomeScore: 2, AwayScore: 1},
		{Home: "WHU", Away: "AVL", HomeScore: 1, AwayScore: 0},
		{Home: "BRE", Away: "WAT", HomeScore: 0, AwayScore: 1},
		{Home: "BHA", Away: "NOR", HomeScore: 3, AwayScore: 1},
		{Home: "CRY", Away: "SOU", HomeScore: 1, AwayScore: 3},
		{Home: "CHE", Away: "ARS", HomeScore: 1, AwayScore: 2},
		{Home: "LEI", Away: "MU", HomeScore: 0, AwayScore: 0},
	}
	var count, win, loss int
	for _, test := range games {
		hT, _ := pi.Search(models.GetBaseGameType("ratings", "fifa22Eng"), test.Home)
		aT, _ := pi.Search(models.GetBaseGameType("ratings", "fifa22Eng"), test.Away)
		ratingH := (hT.HomeRating + hT.AwayRating) / 2
		ratingA := (aT.HomeRating + aT.AwayRating) / 2
		//	if ratingH < 0 {
		//		continue
		//	}
		switch {
		case ratingH < 0 && float64(test.HomeScore) < 1.5:
			inv2.bet()
			count++
			t.Error("after win with O/U:", inv.balance, count)
		case ratingH > ratingA && test.HomeScore >= test.AwayScore:
			if ratingH < 0 {
				continue
			}
			win++
			count++
			inv.bet()
			t.Error("after win:", inv.balance, count)
		case ratingA > ratingH && test.AwayScore >= test.HomeScore:
			if ratingH < 0 {
				continue
			}
			win++
			count++
			inv.bet()
			t.Error("after win:", inv.balance, count)
		case ratingH > ratingA && test.AwayScore > test.HomeScore:
			if ratingH < 0 {
				continue
			}
			loss++
			inv.betLoss()
			count++
			t.Error("after loss:", inv.balance, count)
		case ratingA > ratingH && test.HomeScore > test.AwayScore:
			if ratingH < 0 {
				continue
			}
			loss++
			inv.betLoss()
			count++
			t.Error("after loss:", inv.balance, count)
		}
	}
	t.Error(win, loss, len(games), inv.balance, inv2.balance)
}

func TestWriteratings(t *testing.T) {
	data, err := csvmanager.ReadCsv(models.GetBase()+"fifa22Eng.csv", 0755, true)
	if err != nil {
		panic(err)
	}
	rows := data.Rows()
	for _, game := range rows {
		match := game.String()
		HomeTeam := &pi.Team{Name: match[0]}
		AwayTeam := &pi.Team{Name: match[1]}
		homeGoal, err := strconv.Atoi(match[2])
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		awayGoal, err := strconv.Atoi(match[3])
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		_, _, err = pi.UpdateTeamRatings(models.GetBase()+"ratingsfifa22Eng.csv", HomeTeam.Name, AwayTeam.Name, homeGoal, awayGoal)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	}

}

func TestWriteToPi(t *testing.T) {
	_, _, err := pi.UpdateTeamRatings(models.GetBase()+"ratingsfifa4x4Eng.csv", "NU", "WAT", 5, 1)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSingles(t *testing.T) {
	home := "ARS"
	awat := "BUR"
	hT, _ := pi.Search(models.GetBaseGameType("ratings", "fifa22Eng"), home)
	aT, _ := pi.Search(models.GetBaseGameType("ratings", "fifa22Eng"), awat)
	homexG := pi.ExpectedGoalIndividual(hT.HomeRating)
	awayxG := pi.ExpectedGoalIndividual(aT.AwayRating)
	ratingH := (hT.HomeRating + hT.AwayRating) / 2
	ratingA := (aT.HomeRating + aT.AwayRating) / 2

	t.Error(home, ":", ratingH, homexG, awat, ":", ratingA, awayxG, homexG-awayxG)
}

// RESULTS: Betting on expected Goal difference produced drastically lower returns than betting on rating difference
func TestSearchWithPixG(t *testing.T) {
	inv := &invests{balance: 200, stake: 50, odds: 1.35}

	games := []models.Data{
		{Home: models.Avl, Away: models.Wol, HomeScore: 2, AwayScore: 1},
		{Home: models.Wat, Away: models.Tot, HomeScore: 1, AwayScore: 4},
		{Home: models.Nor, Away: "MU", HomeScore: 3, AwayScore: 3},
		{Home: "SOU", Away: "ARS", HomeScore: 0, AwayScore: 5},
		{Home: "ARS", Away: "NU", HomeScore: 5, AwayScore: 0},
		{Home: "MU", Away: "SOU", HomeScore: 3, AwayScore: 1},
		{Home: "TOT", Away: "NOR", HomeScore: 2, AwayScore: 1},
		{Home: "WOL", Away: "WAT", HomeScore: 3, AwayScore: 1},
		{Home: "LIV", Away: "AVL", HomeScore: 3, AwayScore: 2},
		{Home: "BUR", Away: "MCI", HomeScore: 0, AwayScore: 4},
		{Home: "EVE", Away: "WHU", HomeScore: 4, AwayScore: 1},
		{Home: "LU", Away: "BRE", HomeScore: 1, AwayScore: 1},
		{Home: "LEI", Away: "BHA", HomeScore: 3, AwayScore: 1},
		{Home: "CHE", Away: "CRY", HomeScore: 1, AwayScore: 2},
		{Home: "NU", Away: "CRY", HomeScore: 1, AwayScore: 2},
		{Home: "BHA", Away: "CHE", HomeScore: 2, AwayScore: 3},
		{Home: "BRE", Away: "LEI", HomeScore: 1, AwayScore: 3},
		{Home: "WHU", Away: "LU", HomeScore: 2, AwayScore: 1},
		{Home: "MCI", Away: "EVE", HomeScore: 3, AwayScore: 1},
		{Home: "AVL", Away: "BUR", HomeScore: 1, AwayScore: 3},
		{Home: "WAT", Away: "LIV", HomeScore: 3, AwayScore: 2},
		{Home: "NOR", Away: "WOL", HomeScore: 2, AwayScore: 3},
		{Home: "SOU", Away: "TOT", HomeScore: 1, AwayScore: 2},
		{Home: "ARS", Away: "MU", HomeScore: 3, AwayScore: 2},
		{Home: "MU", Away: "NU", HomeScore: 4, AwayScore: 0},
		{Home: "TOT", Away: "ARS", HomeScore: 3, AwayScore: 0},
		{Home: "WOL", Away: "SOU", HomeScore: 1, AwayScore: 2},
		{Home: "LIV", Away: "NOR", HomeScore: 5, AwayScore: 1},
		{Home: "BUR", Away: "WAT", HomeScore: 2, AwayScore: 1},
		{Home: "EVE", Away: "AVL", HomeScore: 2, AwayScore: 1},
	}
	var win, loss int
	for _, test := range games {
		hT, _ := pi.Search(models.GetBaseGameType("ratings", "fifa22Eng"), test.Home)
		aT, _ := pi.Search(models.GetBaseGameType("ratings", "fifa22Eng"), test.Away)
		homexG := pi.ExpectedGoalIndividual(hT.HomeRating)
		awayxG := pi.ExpectedGoalIndividual(aT.AwayRating)
		xG := pi.ExpectedGoalDifference(homexG, awayxG)
		GD := test.HomeScore - test.AwayScore
		switch {
		case xG > 0 && float64(GD) >= xG:
			win++
			inv.bet()
			t.Error("after win:", inv.balance)
		case xG < 0 && float64(GD) < xG:
			win++
			inv.bet()
			t.Error("after win:", inv.balance)

		case xG > 0 && float64(GD) < xG:
			loss++
			inv.betLoss()
			t.Error("after loss:", inv.balance)
		case xG < 0 && float64(GD) > xG:
			loss++
			inv.betLoss()
			t.Error("after loss:", inv.balance)
		}
	}
	t.Error(win, loss, len(games), inv.balance)

}
