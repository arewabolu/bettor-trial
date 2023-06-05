package main

import (
	"bettor/models"
	"strconv"
	"testing"

	"github.com/arewabolu/csvmanager"
	"github.com/arewabolu/pi-rating"
	"gonum.org/v1/gonum/stat"
)

func TestDesribe(t *testing.T) {
	team1 := models.Psg
	team2 := models.Ars
	game := "fifa22Pen"
	games := models.GetGames(&game, team1, team2)
	sumsOfGoals := make([]int, 0)
	for _, game := range games {
		sumOfGoals := game.HomeScore + game.AwayScore
		sumsOfGoals = append(sumsOfGoals, sumOfGoals)
	}
	nwSumsofGoals := models.FloatCon(sumsOfGoals)
	mean, stdDev := stat.MeanStdDev(nwSumsofGoals, nil)
	skew := stat.Skew(nwSumsofGoals, nil)
	kurt := stat.ExKurtosis(nwSumsofGoals, nil)
	mode, count := stat.Mode(nwSumsofGoals, nil)
	t.Error("mean :", mean, "stdDev :", stdDev, "kurt :", kurt, "skew :", skew)
	t.Error(mode, count)
}

func TestSearchWithPi(t *testing.T) {
	inv := &invests{balance: 200, stake: 50, odds: 1.19}

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
	var count, win, loss int
	for _, test := range games {
		hT := pi.Search(models.GetBase()+"ratingsfifa22Eng.csv", test.Home, "home")
		aT := pi.Search(models.GetBase()+"ratingsfifa22Eng.csv", test.Away, "away")
		ratingH := (hT.HomeRating + hT.AwayRating) / 2
		ratingA := (aT.HomeRating + aT.AwayRating) / 2
		if ratingH-ratingA < 0.3 && ratingH-ratingA > -0.3 {
			continue
		}
		switch {
		case ratingH > ratingA && test.HomeScore >= test.AwayScore:
			win++
			count++
			inv.bet()
			t.Error("after win:", inv.balance, count)
		case ratingA > ratingH && test.AwayScore >= test.HomeScore:
			win++
			count++
			inv.bet()
			t.Error("after win:", inv.balance, count)
		case ratingH > ratingA && test.AwayScore > test.HomeScore:
			loss++
			inv.betLoss()
			count++
			t.Error("after loss:", inv.balance, count)
		case ratingA > ratingH && test.HomeScore > test.AwayScore:
			loss++
			inv.betLoss()
			count++
			t.Error("after loss:", inv.balance, count)
		}
		if ratingH-ratingA < 0.4 && ratingH-ratingA > -0.4 {
			//	t.Error("Possible loss", win, loss)
		}
	}
	t.Error(win, loss, len(games), inv.balance)
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
		err = pi.UpdateTeamRatings(models.GetBase()+"ratingsfifa22Eng.csv", HomeTeam.Name, AwayTeam.Name, homeGoal, awayGoal)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
	}

}

func TestWriteToPi(t *testing.T) {
	err := pi.UpdateTeamRatings(models.GetBase()+"ratingsfifa4x4Eng.csv", "NU", "WAT", 5, 1)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestSingles(t *testing.T) {
	home := "BRE"
	awat := "SOU"
	hT := pi.Search(models.GetBase()+"ratingsfifa22Eng.csv", home, "home")
	aT := pi.Search(models.GetBase()+"ratingsfifa22Eng.csv", awat, "away")
	ratingH := (hT.HomeRating + hT.AwayRating) / 2
	ratingA := (aT.HomeRating + aT.AwayRating) / 2

	t.Error(home, ":", ratingH, awat, ":", ratingA, ratingH-ratingA)
}
