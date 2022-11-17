package models

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Println("could not process request because", err)
		os.Exit(1)
	}
}

func sumArray(arr []int) (sumMedium int) {
	for _, item := range arr {
		sumMedium = sumMedium + item
	}
	return sumMedium
}

func CheckRegisteredTeams(homeTeam, awayTeam string) error {
	teamArr := []string{Avl, Ars, Bha, Bre, Bur, Che, Cry, Eve, Lei, Liv, Lu, Mci, Mu, Nor, Nu, Sou, Tot, Wat, Whu, Wol, Bar, Bay, Juv, Rma, Psg}
	if !slices.Contains(teamArr, homeTeam) || !slices.Contains(teamArr, awayTeam) {
		err := errors.New("one of the teams names is incorrect")
		return err
	}
	return nil
}

func GetHome() (string, error) {
	home, err := os.UserHomeDir()
	return home, err
}

func GetBase() string {
	home, _ := GetHome()
	basedir := home + "/bettor/database/"
	return basedir
}
