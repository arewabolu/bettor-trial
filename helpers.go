package main

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

func checkErr(err error) {
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

func checkRegisteredTeams(homeTeam, awayTeam string) error {
	teamArr := []string{avl, ars, bha, bre, bur, che, cry, eve, lei, liv, lu, mci, mu, nor, nu, sou, tot, wat, whu, wol, bar, bay, juv, rma, psg}
	if !slices.Contains(teamArr, homeTeam) || !slices.Contains(teamArr, awayTeam) {
		err := errors.New("one of the teams names is incorrect")
		return err
	}
	return nil
}
