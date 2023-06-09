package models

import (
	"os"

	"github.com/arewabolu/csvmanager"
)

//func CheckErr(err error) {
//	if err != nil {
//		fmt.Println("could not process request because", err)
//		os.Exit(1)
//	}
//}

func GetHome() (string, error) {
	home, err := os.UserHomeDir()
	return home, err
}

func GetBase() string {
	home, _ := GetHome()
	basedir := home + "/bettor/database/"
	return basedir
}

func GetBaseTeamNames() string {
	home, _ := GetHome()
	basedir := home + "/bettor/database/TeamName/"
	return basedir
}

func GetBaseImage() string {
	home, _ := GetHome()
	basedir := home + "/bettor/database/probability/distribution.png"
	return basedir
}

func GetBaseRating() string {
	home, _ := GetHome()
	basedir := home + "/bettor/database/ratings/"
	return basedir
}

func GetBaseGameType(defaultFolder, gametype string) string {
	basedir := GetBase() + defaultFolder + "/" + defaultFolder + gametype + ".csv"
	return basedir
}

func StatusAllocator(rds csvmanager.Frame, team string) []int {
	status := make([]int, 0)

	for _, game := range rds.Rows() {
		nwData := &Data{}
		game.Interface(nwData)

		if team == nwData.Home {
			status = append(status, 1)
		}
		if team == nwData.Away {
			status = append(status, 0)
		}
	}
	return status
}
