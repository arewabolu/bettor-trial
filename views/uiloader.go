package views

import (
	"bettor/controller"
	"bettor/models"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type App interface {
	Register()
	Retrieve()
}

func AppStart() {
	a := app.NewWithID("gitlab.com/arewabolu/game-averages")
	w := a.NewWindow("Bettor")
	w.Resize(fyne.NewSize(600, 600))
	w.SetContent(uiLoader(w))
	w.ShowAndRun()
}

// TODO must be able to register teams at the creation of new category
func leftContainer(freeContainer *fyne.Container, w fyne.Window) fyne.CanvasObject {
	listData := []string{"create new category", "register new Game", "search for game"}
	list := widget.NewList(
		func() int { return len(listData) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(listData[lii])
		})
	list.OnSelected = func(id widget.ListItemID) {
		if id == 0 {
			freeContainer.RemoveAll()
			freeContainer.Add(loadRightSide3())
		}
		if id == 1 {
			freeContainer.RemoveAll()
			freeContainer.Add(loadRightSide1())
		}
		if id == 2 {
			freeContainer.RemoveAll()
			freeContainer.Add(loadRightSide2(freeContainer, w))
		}
	}
	list.OnUnselected = func(id widget.ListItemID) {
		freeContainer.RemoveAll()
		freeContainer.Add(container.NewCenter(widget.NewLabel("Please Select an Option!")))
	}
	return list
}

func uiLoader(w fyne.Window) *container.Split {
	fsttext := container.NewCenter(widget.NewLabel("Please Select an Option!"))
	emptyCont := container.NewBorder(nil, nil, nil, nil, fsttext)

	l := leftContainer(emptyCont, w)

	simp := container.NewHSplit(l, emptyCont)

	return simp
}

func loadRightSide1() fyne.CanvasObject {
	HTEnt := new(widget.Entry)
	ATEnt := new(widget.Entry)
	HTSEnt := new(widget.Entry)
	ATSEnt := new(widget.Entry)

	radOptions := models.DirIterator(models.GetBase())

	rad := widget.NewRadioGroup(radOptions, func(s string) {
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
	submit := controller.SaveButton(rad, HTEnt, ATEnt, HTSEnt, ATSEnt)
	rightSide := container.NewBorder(rad, submit, nil, nil, vBox)
	return rightSide
}

func loadRightSide2(freeCont *fyne.Container, w fyne.Window) fyne.CanvasObject {
	HTEnt := new(widget.Entry)
	ATEnt := new(widget.Entry)
	HTLabel := widget.NewLabel("Home Team:")
	ATLabel := widget.NewLabel("Away Team:")

	radOptions := models.DirIterator(models.GetBase())
	rad := widget.NewRadioGroup(radOptions, func(s string) {
	})

	submit := widget.NewButton("Search", func() {
		HT := HTEnt.Text
		AT := ATEnt.Text
		values := []string{HT, AT}
		percentages, goals, err := controller.CheckReader(rad.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		labels := groupie(percentages, goals, radOptions, values, rad)
		freeCont.RemoveAll()
		backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			freeCont.RemoveAll()
			freeCont.Add(loadRightSide2(freeCont, w))
		})
		freeCont.Add(container.NewBorder(backButn, nil, nil, nil, container.NewVBox(labels[0], labels[1], labels[2], labels[3], labels[4], labels[5])))
	})

	HTHBox := container.NewBorder(nil, nil, HTLabel, nil, HTEnt)
	ATHBox := container.NewBorder(nil, nil, ATLabel, nil, ATEnt)
	vBox := container.NewVBox(HTHBox, ATHBox)
	return container.NewBorder(rad, submit, nil, nil, vBox)
}

func creator(text string) *widget.Label {
	label := new(widget.Label)
	label.Wrapping = fyne.TextWrapBreak
	label.SetText(text)
	return label
}

func groupie(percentages, goalPercentages []float64, radOptions, Teams []string, rad *widget.RadioGroup) []*widget.Label {
	// Percentages
	perc1 := creator(fmt.Sprintf("%s win percentage %.2f\n", Teams[0], percentages[0]))
	perc2 := creator(fmt.Sprintf("%s win percentage %.2f\n", Teams[1], percentages[1]))
	perc3 := creator(fmt.Sprintf("draw percentage %.2f\n", percentages[2]))

	// Goals
	goalVal1 := new(widget.Label)
	goalVal2 := new(widget.Label)
	goalVal3 := new(widget.Label)
	if rad.Selected == "fifa4x4Eng" {
		goalVal1.Text = fmt.Sprintf("There's a %.2f of both teams scoring over 6 goal(s)\n", goalPercentages[0])
		goalVal2.Text = fmt.Sprintf("There's a %.2f of both teams scoring over 7 goal(s)\n", goalPercentages[1])
		goalVal3.Text = fmt.Sprintf("There's a %.2f of both teams scoring over 8 goal(s)\n", goalPercentages[2])
	} else {
		goalVal1.Text = fmt.Sprintf("There's a %.2f of both teams scoring 1 goal(s)\n", goalPercentages[0])
		goalVal2.Text = fmt.Sprintf("There's a %.2f of both teams scoring 2 goal(s)\n", goalPercentages[1])
		goalVal3.Text = fmt.Sprintf("There's a %.2f of both teams scoring 3 goal(s)\n", goalPercentages[2])
	}
	return []*widget.Label{perc1, perc2, perc3, goalVal1, goalVal2, goalVal3}
}

func loadRightSide3() fyne.CanvasObject {
	gameType := new(widget.Entry)
	button := widget.NewButton("Create", func() {
		models.CreateFile(gameType.Text)
	})

	vBox := container.NewVBox(gameType, button)
	return vBox
}
