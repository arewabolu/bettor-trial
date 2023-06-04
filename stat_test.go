package main

import (
	"bettor/models"
	"testing"

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
	games := []models.Data{
		{Home: "CRY", Away: "NOR", HomeScore: 5, AwayScore: 3},
		{Home: "CRY", Away: "MU", HomeScore: 3, AwayScore: 5},
		{Home: "LU", Away: "CHE", HomeScore: 5, AwayScore: 4},
		{Home: "ARS", Away: "SOU", HomeScore: 8, AwayScore: 2},
		{Home: "MCI", Away: "WHU", HomeScore: 4, AwayScore: 6},
		{Home: "BRE", Away: "TOT", HomeScore: 4, AwayScore: 3},
		{Home: "BHA", Away: "LEI", HomeScore: 3, AwayScore: 5},
		{Home: "AVL", Away: "WAT", HomeScore: 3, AwayScore: 2},
		{Home: "EVE", Away: "BUR", HomeScore: 3, AwayScore: 6},
		{Home: "LIV", Away: "WOL", HomeScore: 5, AwayScore: 3},
		{Home: "NOR", Away: "NU", HomeScore: 2, AwayScore: 4},
		{Home: "NU", Away: "LIV", HomeScore: 5, AwayScore: 1},
		{Home: "WOL", Away: "EVE", HomeScore: 4, AwayScore: 6},
		{Home: "BUR", Away: "AVL", HomeScore: 5, AwayScore: 9},
		{Home: "WAT", Away: "BHA", HomeScore: 3, AwayScore: 5},
		{Home: "LEI", Away: "BRE", HomeScore: 6, AwayScore: 4},
		{Home: "TOT", Away: "MCI", HomeScore: 5, AwayScore: 4},
		{Home: "WHU", Away: "ARS", HomeScore: 2, AwayScore: 3},
		{Home: "SOU", Away: "LU", HomeScore: 5, AwayScore: 5},
		{Home: "CHE", Away: "CRY", HomeScore: 2, AwayScore: 10},
	}

	for _, test := range games {
		hT := pi.Search(models.GetBase()+"ratingsfifa4x4Eng.csv", test.Home, "home")
		aT := pi.Search(models.GetBase()+"ratingsfifa4x4Eng.csv", test.Away, "away")
		ratingH := (hT.HomeRating + hT.AwayRating) / 2
		ratingA := (aT.HomeRating + aT.AwayRating) / 2
		t.Error(test.Home, ":", ratingH, test.HomeScore, test.Away, ":", ratingA, test.AwayScore, "\n")
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
	home := "TOT"
	awat := "BHA"
	hT := pi.Search(models.GetBase()+"ratingsfifa4x4Eng.csv", home, "home")
	aT := pi.Search(models.GetBase()+"ratingsfifa4x4Eng.csv", awat, "away")
	ratingH := (hT.HomeRating + hT.AwayRating) / 2
	ratingA := (aT.HomeRating + aT.AwayRating) / 2
	t.Error(home, ":", ratingH, awat, ":", ratingA, "\n")
}
