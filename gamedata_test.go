package main

import (
	"bettor/controller"
	"bettor/models"
	"strconv"
	"testing"
)

func TestCheckWriter(t *testing.T) {

	//homeTeam := strings.ToUpper(nu)
	//awayTeam := strings.ToUpper(che)
	//err := WriteMatchData(homeTeam, awayTeam, "4", "5")
	//if err != nil {
	//	t.Error("Unable to write file due to:", err)
	//}
	//{,,,},
	//what happens when i enter an incorrect string ? does it stop all operations or should it continue?
	records := []models.Data{{Home: "wat", Away: "ars", HomeScore: 3, AwayScore: 4}}

	for _, slice := range records {

		homeTeam := slice.Home
		awayTeam := slice.Away
		homeScore := strconv.Itoa(slice.HomeScore)
		awayScore := strconv.Itoa(slice.AwayScore)
		err := controller.CheckWriter("fifa4x4Eng", []string{homeTeam, awayTeam, homeScore, awayScore})
		if err != nil {
			t.FailNow()
			t.Error("unable to write data")
		}
		//	err := WriteMatchData(homeTeam, awayTeam, homeScore, awayScore)

	}
}
