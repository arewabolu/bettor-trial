package views

import (
	"bettor/controller"
	"bettor/models"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func AppStart() {
	app := &appController{
		application: app.NewWithID("com.example.myid"),
	}
	app.appwindow = app.application.NewWindow("Bettor")

	//app.appwindow.Resize(fyne.NewSize(550, 500))
	app.appwindow.SetContent(uiLoader(app.appwindow))
	app.appwindow.ShowAndRun()
}

// TODO must be able to register teams at the creation of new category
func uiLoader(w fyne.Window) fyne.CanvasObject {
	listData := []string{"create new category", "register new Game", "search for game", "search for Team Data"}
	//GFHome, GFAway Avgs,
	but1 := widget.NewButtonWithIcon(listData[0], theme.ContentAddIcon(), func() {
		w.SetContent(loadRightSide3(w))
	})
	//width := but1.Size().Width
	//but1.Resize(fyne.NewSize(width, 20))
	but2 := widget.NewButtonWithIcon(listData[1], theme.ContentAddIcon(), func() {
		w.SetContent(loadRightSide1(w))
	})
	but3 := widget.NewButtonWithIcon(listData[2], theme.SearchIcon(), func() {
		w.SetContent(loadRightSide2(w))
	})
	but4 := widget.NewButtonWithIcon(listData[3], theme.SearchIcon(), func() {
		w.SetContent(loadRightSide4(w))
	})
	grid := container.NewAdaptiveGrid(len(listData), but1, but2, but3, but4)

	return grid
}

func loadRightSide1(w fyne.Window) fyne.CanvasObject {
	HTEnt := newSubmitEntry(w)
	ATEnt := newSubmitEntry(w)
	HTSEnt := newSubmitEntry(w)
	ATSEnt := newSubmitEntry(w)

	radOptions := models.DirIterator(models.GetBase())

	Select := widget.NewSelect(radOptions, func(s string) {
	})
	HTLabel := widget.NewLabel("Home Team:")
	HTSLabel := widget.NewLabel("Home Teams Score")
	ATLabel := widget.NewLabel("Away Team:")
	ATSLabel := widget.NewLabel("Away Teams Score:")
	HTHBox := container.NewBorder(nil, nil, HTLabel, nil, HTEnt)
	ATHBox := container.NewBorder(nil, nil, ATLabel, nil, ATEnt)
	HTSHBox := container.NewBorder(nil, nil, HTSLabel, nil, HTSEnt)
	ATSHBox := container.NewBorder(nil, nil, ATSLabel, nil, ATSEnt)

	submit := SaveButton(Select, w, &HTEnt.Entry, &ATEnt.Entry, &HTSEnt.Entry, &ATSEnt.Entry)
	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})

	rightSide := container.NewVBox(backButn, Select, HTHBox, HTSHBox, ATHBox, ATSHBox, submit)
	ent := []*widget.Entry{&HTEnt.Entry, &ATEnt.Entry, &ATSEnt.Entry, &HTSEnt.Entry}
	saveFunc := func() {
		values := []string{HTEnt.Text, ATEnt.Text, HTSEnt.Text, ATSEnt.Text}

		if Select.Selected == "" {
			dlog := dialog.NewError(errors.New("please select the game type"), w)
			dlog.Show()
			w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
				if ke.Name == fyne.KeyReturn {
					dlog.Hide()
				}
			})
			return
		}

		err := controller.CheckWriter(Select.Selected, values)
		if err != nil {
			dlog := dialog.NewError(errors.New("please select the game type"), w)
			dlog.Show()
			w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
				if ke.Name == fyne.KeyReturn {
					dlog.Hide()
				}
			})
		}

		entryDel(ent...)
		rightSide.Refresh()
	}

	HTEnt.button = submit
	ATEnt.button = submit
	HTSEnt.button = submit
	ATSEnt.button = submit

	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch ke.Name {
		case fyne.KeyF1:
			w.Canvas().Focus(HTEnt)
		case fyne.KeyF2:
			w.Canvas().Focus(HTSEnt)
		case fyne.KeyF3:
			w.Canvas().Focus(ATEnt)
		case fyne.KeyF4:
			w.Canvas().Focus(ATSEnt)
		case fyne.KeyReturn:
			w.Canvas().Focus(submit)
			saveFunc()
		}

	})
	return rightSide
}

func loadRightSide2(w fyne.Window) fyne.CanvasObject {
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
		GP, percentageWinorDraw, err := controller.CheckReader(Select.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		//labels := groupie(percentageWinorDraw, values, rad)

		backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			Box.RemoveAll()
			w.SetContent(loadRightSide2(w))
		})

		w.SetContent(container.NewBorder(backButn, nil, nil, nil, tableRender(values, GP, percentageWinorDraw)))

	}

	return Box
}

func loadRightSide3(w fyne.Window) fyne.CanvasObject {
	gameType := new(widget.Entry)
	button := widget.NewButton("Create", func() {
		models.CreateFile(gameType.Text)
	})
	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})
	vBox := container.NewVBox(backButn, gameType, button)
	return vBox
}

func loadRightSide4(w fyne.Window) fyne.CanvasObject {
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
		xG, MAE := controller.SearcherV2(Select.Selected, team, status)
		backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			Box.RemoveAll()
			w.SetContent(loadRightSide4(w))
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
					//	case tci.Col == 3 && tci.Row == 0:
					//		label.SetText("Home GPG")
					//	case tci.Col == 4 && tci.Row == 0:
					//		label.SetText("AwayGPG")
				case tci.Col == 0 && tci.Row == 1:
					label.SetText(team)
				case tci.Col == 1 && tci.Row == 1:
					label.SetText(fmt.Sprintf("%.2f", xG))
				case tci.Col == 2 && tci.Row == 1:
					label.SetText(fmt.Sprintf("%.2f", MAE))
					//	case tci.Col == 3 && tci.Row == 1:
					//		label.SetText(fmt.Sprintf("%.2f", meanHomeGoals))
					//	case tci.Col == 4 && tci.Row == 1:
					//		label.SetText(fmt.Sprintf("%.2f", meanAwayGoals))
				}
			})

		w.SetContent(container.NewBorder(backButn, nil, nil, nil, teamTable))

	}

	return Box
}

func tableRender(team []string, GP int, percentageWinorDraw []float64) *widget.Table {
	table := widget.NewTable(
		func() (int, int) { return 3, 5 },
		func() fyne.CanvasObject { return widget.NewLabel("xxxxxx") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			switch {
			case tci.Col == 0 && tci.Row == 0:
				label.SetText("Team")
			case tci.Col == 1 && tci.Row == 0:
				label.SetText("GP")
			case tci.Col == 2 && tci.Row == 0:
				label.SetText("W%")
			case tci.Col == 3 && tci.Row == 0:
				label.SetText("D%")

			case tci.Col == 0 && tci.Row == 1:
				label.SetText(team[0])
			case tci.Col == 1 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%d", GP))
			case tci.Col == 2 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[0]))
			case tci.Col == 3 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[2]))

			case tci.Col == 0 && tci.Row == 2:
				label.SetText(team[1])
			case tci.Col == 1 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%d", GP))
			case tci.Col == 2 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[1]))
			case tci.Col == 3 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[2]))
			}

		})

	return table
}
