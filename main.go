package main

import (
	"bettor/controller"
	"bettor/models"
	"bettor/views"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/arewabolu/csvmanager"
	"golang.org/x/exp/slices"
)

var (
	register  string
	search    string
	listTeams string
	list      bool
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
	err = os.MkdirAll(home+"/bettor/database/TeamName/", os.ModePerm)
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

	help := fmt.Sprintln("First argument to any of the flags must be a registered game type e.g fifa22Eng, fifa18Pen")
	flag.StringVar(&register, "r", "", help+"used to register all games(penalties and mathces)")
	flag.StringVar(&search, "s", "", "used to search for match results")
	flag.StringVar(&listTeams, "lt", "", "list all teams of a game")
	flag.BoolVar(&list, "l", false, "list all registered games")
}

func main() {
	flag.Parse()
	args := flag.Args()
	flagValues := models.DirIterator(models.GetBase())

	if list {
		fmt.Fprintln(os.Stdout, flagValues)
		return
	}
	switch {
	case slices.Contains(flagValues, register):
		err := controller.CheckWriter(register, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

	case slices.Contains(flagValues, search):
		homeRating, AwayRating, err := controller.CheckReader(search, args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Fprintf(os.Stdout, "home: %.4f\naway:  %.4f\n", homeRating, AwayRating)
	case slices.Contains(flagValues, listTeams):
		_, err := os.Stat(models.GetBaseTeamNames() + listTeams + ".csv")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		reader, readErr := csvmanager.ReadCsv(models.GetBaseTeamNames()+listTeams+".csv", 0700, true, 20)
		if readErr != nil {
			fmt.Fprintln(os.Stderr, readErr)
			return
		}
		teamNames := reader.Col("Teams").String()
		if len(teamNames) < 1 {
			fmt.Fprintln(os.Stderr, errors.New("no teams currently registered. please update teams for league"))
		}
		fmt.Fprintln(os.Stdout, teamNames)
	case len(args) == 0:
		views.AppStart()
	default:
		fmt.Println("the value of your flag is incorrect,please confirm")
	}

}

// potential to score over 1.5 or ... goals
// both teams to score x or more
// percentage of wins,draws and losses against current opponent
// percentage/likeliness of wins,draws and losses against other opponent
// does home or away have the advantage for games
// future features registers for reallife games (scores, corners, goalkicks, throw-ins e.t.c)
//even or odd scores
