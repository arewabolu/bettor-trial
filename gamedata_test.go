package main

import (
	"bettor/controller"
	"bettor/models"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/arewabolu/csvmanager"
	tradefuncs "github.com/arewabolu/trademath"
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

func TestCheckReader(t *testing.T) {
	//try to pass
	homeTeam := "ars"
	awayTeam := "mci"
	//	fawayTeam := "bor"
	_, _, _, err := controller.CheckReader("fifa4x4Eng", []string{homeTeam, awayTeam})

	if err != nil {
		t.Error(err)
	}
}

func TestMAchanges(t *testing.T) {
	rds, _ := csvmanager.ReadCsv(models.GetBase()+"fifa4x4Eng.csv", 0644, true)
	goals := models.SearchTeam(models.Sou, rds)
	nwGoals := goals
	med := models.Median(models.FloatCon(nwGoals))
	MA3 := tradefuncs.MA(models.FloatCon(goals), 3)
	//	median := stat.Quantile(0.95, stat.Empirical, models.FloatCon(nwGoals), nil)
	if len(MA3) > 0 {
		fmt.Println(MA3)
		//	fmt.Println(median)
		fmt.Println(med)
		t.Error()
	}

}

func TestWriter(t *testing.T) {
	rds, _ := csvmanager.ReadCsv(models.GetBase()+"fifa4x4Eng.csv", 0644, true)
	team := models.Mci
	goals := models.SearchTeam(team, rds)
	nwGoals := models.FloatCon(goals)
	goals2 := models.SearchTeam3(team, rds)
	nwGoals2 := models.FloatCon(goals2)

	status := make([]string, 0)

	for _, game := range rds.Rows() {
		nwData := &models.Data{}
		game.Interface(nwData)

		if team == nwData.Home {
			status = append(status, "1")
		}
		if team == nwData.Away {
			status = append(status, "0")
		}

	}
	MAStr := models.FloattoString(nwGoals)
	MAStr2 := models.FloattoString(nwGoals2)

	file, err := os.OpenFile("xGxGAStatus"+team+".csv", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	w := &csvmanager.WriteFrame{
		Headers: []string{"goals", "goalsAgainst", "status"},
		Arrays:  [][]string{MAStr, MAStr2, status},
		File:    file,
	}
	w.WriteCSV()
}

func TestWins(t *testing.T) {
	rds, _ := csvmanager.ReadCsv(models.GetBase()+"fifa4x4Eng.csv", 0644, true)
	team1 := models.Mu
	//team2 := models.Che
	goals := models.SearchTeam(team1, rds)
	nwGoals := models.FloatCon(goals)
	goals2 := models.SearchTeam3(team1, rds)
	nwGoals2 := models.FloatCon(goals2)
	status := make([]int, 0)

	for _, game := range rds.Rows() {
		nwData := &models.Data{}
		game.Interface(nwData)

		if team1 == nwData.Home {
			status = append(status, 1)
		}
		if team1 == nwData.Away {
			status = append(status, 0)
		}

	}

	balance := float64(1000)
	betAmt := float64(50)
	var home, away, losses int

	for i := len(nwGoals) - 1; i >= len(nwGoals)-30; i-- {
		if status[i] == 1 {
			goal := 2.7534 + 1*0.9523 + nwGoals2[i]*0.2765
			//941230322026.9298 + (float64(status[i]) * -941230322021.7841) + nwGoals2[i]*-0.1917
			//2.7534 + 1*0.9523 + nwGoals2[i]*0.2765 + 0.30
			//goal := 3.058963891854226 + (nwGoals2[i] * 0.32253680990613737)
			if nwGoals[i] > goal {
				//	t.Error("wh:", goal, nwGoals[i])
				balance = balance - betAmt
				balance = balance + betAmt*1.2
				home++
			}
			if nwGoals[i] <= goal {
				t.Error("lh:", goal, nwGoals[i])
				losses++
				balance = balance - betAmt
			}
		} else {
			goal := 1509306280859.1194 + float64(status[i])*15093062808642.6016 + nwGoals2[i]*-0.1990
			//2.7534 + 1*0.9523 + nwGoals2[i]*0.2765 + 0.30
			//goal := 3.058963891854226 + (nwGoals2[i] * 0.32253680990613737)
			if nwGoals[i] > goal {
				//t.Error("away:", goal)
				//t.Error("wa:", goal, nwGoals[i])
				away++
				balance = balance - betAmt
				balance = balance + betAmt*1.2
			}
			if nwGoals[i] <= goal {
				t.Error("la:", goal, nwGoals[i])
				balance = balance - betAmt
				losses++
			}
		}
	}
	t.Error(balance, home, away, losses)

	/*
	   file, err := os.OpenFile("xGStatus"+team1+".csv", os.O_CREATE|os.O_RDWR, 0755)

	   	if err != nil {
	   		panic(err)
	   	}

	   	w := &csvmanager.WriteFrame{
	   		Headers: []string{"goals", "status"},
	   		Rows:    [][]string{models.FloattoString(nwGoals), status},
	   		File:    file,
	   	}

	   w.WriteNewCSV()
	*/
}

func TestSearcherV2(t *testing.T) {

	teams := []string{"AVL",
		"ARS",
		"BHA",
		"BRE",
		"BUR",
		"CHE",
		"CRY",
		"EVE",
		"LEI",
		"LIV",
		"LU",
		"MCI",
		"MU",
		"NOR",
		"NU",
		"SOU",
		"TOT",
		"WAT",
		"WHU",
		"WOL"}
	game := "fifa4x4Eng"
	//team1 := models.Bay

	for _, team := range teams {
		xG, _ := controller.SearcherV2(game, team, "home")
		t.Error(team, " : ", xG)
	}
}
