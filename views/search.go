package views

import (
	"bettor/controller"
	"bettor/models"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// DEPRECATED: Use
/*
func Searchwith2Teams(w fyne.Window) fyne.CanvasObject {
	HTEnt := new(widget.Entry)
	ATEnt := new(widget.Entry)
	HTLabel := widget.NewLabel("Home Team:")
	ATLabel := widget.NewLabel("Away Team:")

	submit := new(widget.Button)
	submit.Text = "Search"

	radOptions := models.DirIterator(models.GetBase())
	Select := widget.NewSelect(radOptions, func(s string) {
	})

	backButn1 := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})
	HTHBox := container.NewBorder(nil, nil, HTLabel, nil, HTEnt)
	ATHBox := container.NewBorder(nil, nil, ATLabel, nil, ATEnt)
	vBox2 := container.NewVBox(HTHBox, ATHBox)

	Box := container.NewBorder(container.NewVBox(backButn1, Select), submit, nil, nil, vBox2)

	submit.OnTapped = func() {
		HT := HTEnt.Text
		AT := ATEnt.Text
		values := []string{HT, AT}
		if Select.Selected == "" {
			dialog.ShowError(errors.New("please select the game type"), w)
			return
		}
		//GP, percentageWinorDraw, odds,
		GP, even, odd, err := controller.CheckReader(Select.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		//labels := groupie(percentageWinorDraw, values, rad)

		backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			Box.RemoveAll()
			w.SetContent(Searchwith2Teams(w))
		})

		w.SetContent(container.NewBorder(backButn, nil, nil, nil, tableRender(values, GP, []float64{even, odd})))

	}

	return Box
}
*/
func PiSearch(w fyne.Window) fyne.CanvasObject {
	HTEnt := new(widget.Entry)
	ATEnt := new(widget.Entry)
	HTLabel := widget.NewLabel("Home Team:")
	ATLabel := widget.NewLabel("Away Team:")

	submit := new(widget.Button)
	submit.Text = "Search"

	radOptions := models.DirIterator(models.GetBase())
	Select := widget.NewSelect(radOptions, func(s string) {
	})

	backButn1 := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})
	HTHBox := container.NewBorder(nil, nil, HTLabel, nil, HTEnt)
	ATHBox := container.NewBorder(nil, nil, ATLabel, nil, ATEnt)
	vBox2 := container.NewVBox(HTHBox, ATHBox)

	Box := container.NewBorder(container.NewVBox(backButn1, Select), submit, nil, nil, vBox2)

	submit.OnTapped = func() {
		values := []string{HTEnt.Text, ATEnt.Text}
		if Select.Selected == "" {
			dialog.ShowError(errors.New("please select the game type"), w)
			return
		}
		//GP, percentageWinorDraw, odds,
		ratingH, ratingA, err := controller.CheckReader(Select.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		team1 := widget.NewLabel(fmt.Sprintf("%s:	%.4f", values[0], ratingH))
		team2 := widget.NewLabel(fmt.Sprintf("%s:	%.4f", values[1], ratingA))
		backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			Box.RemoveAll()
			w.SetContent(PiSearch(w))
		})

		w.SetContent(container.NewVBox(backButn, team1, team2))

	}

	return Box
}

func SearchWith1Team(w fyne.Window) fyne.CanvasObject {
	TeamEntry := new(widget.Entry)
	statusEntry := widget.NewSelect([]string{"home", "away"}, func(s string) {
	})

	TeamLabel := widget.NewLabel("Team:")

	submit := new(widget.Button)
	submit.Text = "Search"

	radOptions := models.DirIterator(models.GetBase())
	Select := widget.NewSelect(radOptions, func(s string) {
	})

	backButn1 := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})
	box := container.NewVBox(TeamEntry)
	TeamBox := container.NewBorder(nil, nil, TeamLabel, nil, box)

	Box := container.NewBorder(container.NewVBox(backButn1, Select, statusEntry), submit, nil, nil, TeamBox)

	submit.OnTapped = func() {
		team := TeamEntry.Text
		status := statusEntry.Selected

		if Select.Selected == "" {
			dialog.ShowError(errors.New("please select the game type"), w)
			return
		}

		//meanTeamGls, meanOppGoals, meanHomeGoals, meanAwayGoals := controller.Searcher(Select.Selected, team)
		xG, MAE := models.SearcherV2(Select.Selected, team, status)
		backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			Box.RemoveAll()
			w.SetContent(SearchWith1Team(w))
		})

		teamTable := widget.NewTable(
			func() (int, int) { return 2, 3 },
			func() fyne.CanvasObject { return widget.NewLabel("xxxxxxxxxxxx") },
			func(tci widget.TableCellID, co fyne.CanvasObject) {
				label := co.(*widget.Label)
				switch {
				case tci.Col == 0 && tci.Row == 0:
					label.SetText("Team")
				case tci.Col == 1 && tci.Row == 0:
					label.SetText("xG")
				case tci.Col == 2 && tci.Row == 0:
					label.SetText("MAE")

				case tci.Col == 0 && tci.Row == 1:
					label.SetText(team)
				case tci.Col == 1 && tci.Row == 1:
					label.SetText(fmt.Sprintf("%.2f", xG))
				case tci.Col == 2 && tci.Row == 1:
					label.SetText(fmt.Sprintf("%.2f", MAE))
				}
			})

		w.SetContent(container.NewBorder(backButn, nil, nil, nil, teamTable))

	}

	return Box
}
