package main

import (
	"bettor/controller"
	"bettor/models"
	"bettor/views"
	"flag"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

var (
	advantage bool
	register  string
	search    string
)

func makeDir() error {
	//test on windows
	home, homeDirErr := models.GetHome()
	if homeDirErr != nil {
		return homeDirErr
	}
	err := os.MkdirAll(home+"/bettor/database/", os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func init() {

	err := makeDir()
	if err != nil {
		os.Exit(1)
	}
	flag.BoolVar(&advantage, "adv", false, "To check what teams have an advantage")
	flag.StringVar(&register, "reg", "", "used to register all games(penalties and mathces)")
	flag.StringVar(&search, "search", "", "used to search for match results")
}

func main() {
	flag.Parse()
	//check how many distinct games exist and also how many repetitve games exist
	// implement flags today -s -n -agg -h
	// function that takes all games of one team and returns their gpg ratio
	//also function to return individual totals per game for home and away team
	args := flag.Args()
	//app :=
	flagValues := []string{"4x4", "pen18", "pen22"}
	switch {
	case len(args) == 0:
		views.AppStart()
	case slices.Contains(flagValues, register):

		retStr := controller.CheckWriter(register, args)
		fmt.Println(retStr)

	case slices.Contains(flagValues, search):
	//	_, even, odd, err := controller.CheckReader(search, args)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}

	//fmt.Printf("%s win percentage %.2f\n", flag.Arg(0), percentageWinorDraw[0])
	//	fmt.Printf("%s win percentage %.2f\n", flag.Arg(1), percentageWinorDraw[1])
	//	fmt.Printf("draw percentage %.2f\n", percentageWinorDraw[2])

	//	if search == flagValues[0] {
	//	fmt.Printf("There's a %.2f of both teams scoring over 6 goal(s)\n", goals[0])
	//	fmt.Printf("There's a %.2f of both teams scoring over 7 goal(s)\n", goals[1])
	//	fmt.Printf("There's a %.2f of both teams scoring over 8 goal(s)\n", goals[2])
	//	} else {
	//		fmt.Printf("There's a %.2f of both teams scoring 1 goal(s)\n", goals[0])
	//		fmt.Printf("There's a %.2f of both teams scoring 2 goal(s)\n", goals[1])
	//		fmt.Printf("There's a %.2f of both teams scoring 3 goal(s)\n", goals[2])
	//	}

	default:
		fmt.Println("the value of your flag is incorrect,please confirm")
	}

}

// potential to score over 1.5 or ... goals
// both teams to score x or more
// percentage of wins,draws and losses against current opponent
// percentage/likeliness of wins,draws and losses against other opponent
//goals for and goals against
//does home or away have the advantage for games
// future features registers for reallife games (scores, corners, goalkicks, throw-ins e.t.c)
//even or odd scores

//move database registrars and callers to another file
//seperate get aggregate verbose for diff files
//handle conditionals and return values(change to nil)
//if pen check the scores. it must not be equal
