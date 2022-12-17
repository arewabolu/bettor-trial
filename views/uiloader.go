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
	a := app.NewWithID("com.example.myid")
	w := a.NewWindow("Bettor")
	w.Resize(fyne.NewSize(600, 600))
	w.SetContent(uiLoader(w))
	w.ShowAndRun()
}

// TODO must be able to register teams at the creation of new category
func uiLoader(w fyne.Window) fyne.CanvasObject {
	listData := []string{"create new category", "register new Game", "search for game"}
	but1 := widget.NewButtonWithIcon(listData[0], theme.ContentAddIcon(), func() {
		w.SetContent(loadRightSide3(w))
	})
	width := but1.Size().Width
	but1.Resize(fyne.NewSize(width, 20))
	but2 := widget.NewButtonWithIcon(listData[1], theme.ContentAddIcon(), func() {
		w.SetContent(loadRightSide1(w))
	})
	but2.Resize(fyne.NewSize(width, 20))
	but3 := widget.NewButtonWithIcon(listData[2], theme.SearchIcon(), func() {
		w.SetContent(loadRightSide2(w))
	})
	but3.Resize(fyne.NewSize(width, 20))
	grid := container.NewAdaptiveGrid(len(listData), but1, but2, but3)
	return grid
}

func loadRightSide1(w fyne.Window) fyne.CanvasObject {
	HTEnt := new(widget.Entry)
	ATEnt := new(widget.Entry)
	HTSEnt := new(widget.Entry)
	ATSEnt := new(widget.Entry)

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
	vBox := container.NewVBox(HTHBox, HTSHBox, ATHBox, ATSHBox)

	submit := SaveButton(Select, w, HTEnt, ATEnt, HTSEnt, ATSEnt)
	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})

	rightSide := container.NewBorder(backButn, submit, nil, nil, container.NewBorder(Select, nil, nil, nil, vBox))
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
		GP, percentageWinorDraw, odds, err := controller.CheckReader(Select.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		//labels := groupie(percentageWinorDraw, values, rad)

		backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			Box.RemoveAll()
			w.SetContent(loadRightSide2(w))
		})

		w.SetContent(container.NewBorder(backButn, nil, nil, nil, tableRender(values, GP, percentageWinorDraw, odds)))

	}

	return Box
}

func tableRender(team []string, GP int, percentageWinorDraw, odds []float64) *widget.Table {
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
			case tci.Col == 4 && tci.Row == 0:
				label.SetText("XOdds")

			case tci.Col == 0 && tci.Row == 1:
				label.SetText(team[0])
			case tci.Col == 1 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%d", GP))
			case tci.Col == 2 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[0]))
			case tci.Col == 3 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[2]))
			case tci.Col == 4 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%.2f", odds[0]))

			case tci.Col == 0 && tci.Row == 2:
				label.SetText(team[1])
			case tci.Col == 1 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%d", GP))
			case tci.Col == 2 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[1]))
			case tci.Col == 3 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[2]))
			case tci.Col == 4 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%.2f", odds[1]))
			}

		})

	return table
}

func groupie(percentages []float64, Teams []string, rad *widget.RadioGroup) []*widget.Label {
	// Percentages
	perc1 := Creator(fmt.Sprintf("%s win percentage %.2f\n", Teams[0], percentages[0]))
	perc2 := Creator(fmt.Sprintf("%s win percentage %.2f\n", Teams[1], percentages[1]))
	perc3 := Creator(fmt.Sprintf("draw percentage %.2f\n", percentages[2]))

	// Goals
	//goalVal1 := new(widget.Label)
	//goalVal2 := new(widget.Label)
	//goalVal3 := new(widget.Label)
	//if rad.Selected == "fifa4x4Eng" {
	//	goalVal1.Text = fmt.Sprintf("There's a %.2f of both teams scoring over 6 goal(s)\n", goalPercentages[0])
	//	goalVal2.Text = fmt.Sprintf("There's a %.2f of both teams scoring over 7 goal(s)\n", goalPercentages[1])
	//	goalVal3.Text = fmt.Sprintf("There's a %.2f of both teams scoring over 8 goal(s)\n", goalPercentages[2])
	//} else {
	//	goalVal1.Text = fmt.Sprintf("There's a %.2f of both teams scoring 1 goal(s)\n", goalPercentages[0])
	//	goalVal2.Text = fmt.Sprintf("There's a %.2f of both teams scoring 2 goal(s)\n", goalPercentages[1])
	//	goalVal3.Text = fmt.Sprintf("There's a %.2f of both teams scoring 3 goal(s)\n", goalPercentages[2])
	//}
	return []*widget.Label{perc1, perc2, perc3}
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
