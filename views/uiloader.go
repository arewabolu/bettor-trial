package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type App interface {
	Register()
	Retrieve()
}

func AppStart() {
	a := app.New()
	w := a.NewWindow("Bettor")
	w.Resize(fyne.NewSize(600, 600))
	w.SetContent(uiLoader())
	w.ShowAndRun()
}

func TeamEntry(text string) (*widget.Entry, binding.String) {
	teamEntry := widget.NewEntry()
	teamBinder := binding.NewString()
	teamEntry.SetText(text)
	teamEntry.OnChanged = func(s string) {
		teamBinder.Set(s)
		teamEntry.Bind(teamBinder)
		text = s
	}
	return teamEntry, teamBinder
}

func leftContainer(freeContainer *fyne.Container) fyne.CanvasObject {
	listData := []string{"register new Game", "search for game"}
	list := widget.NewList(
		func() int { return len(listData) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(listData[lii])
		})
	list.OnSelected = func(id widget.ListItemID) {}
	return list
}

func uiLoader() *container.Split {
	fsttext := container.NewCenter(widget.NewLabel("Please Select a book!"))
	emptyCont := container.NewBorder(nil, nil, nil, nil, fsttext)

	l := leftContainer(emptyCont)

	simp := container.NewHSplit(l, emptyCont)
	simp.Offset = 0.25

	return simp
}

func loadRightSide() fyne.CanvasObject {
	englandTeams := []string{"AVL", "ARS", "BHA", "BRE", "BUR", "CHE", "CRY", "EVE", "LEI", "LIV", "LU", "MCI", "MU", "NOR", "NU", "SOU", "TOT", "WAT", "WHU", "WOL"}
	penTeams := []string{"PSG", "BAY", "BAR", "RMA", "JUV", "MCI", "LIV", "ARS"}

	var sVal string
	sel := new(widget.Select)
	rad := []string{"4x4", "pen18", "pen22"}
	widget.NewRadioGroup(rad, func(s string) {
		sVal = s
	})
	if sVal == rad[0] {
		sel.Options = englandTeams
	} else if sVal == rad[1] || sVal == rad[2] {
		sel.Options = penTeams
	}

	HTLabel := widget.NewLabel("Home Team:")
	HTSLabel := widget.NewLabel("Home Teams Score")
	ATLabel := widget.NewLabel("Away Team:")
	ATSLabel := widget.NewLabel("Away Teams Score:")
	entry, entryBind := TeamEntry("")
	entry2, entryBind2 := TeamEntry("")
	HTHBox := container.NewHBox(HTLabel, sel)
	HTSHBox := container.NewHBox(HTSLabel, entry)
	ATHBox := container.NewHBox(ATLabel, sel)
	ATSHBox := container.NewHBox(ATSLabel, entry2)
	submit := widget.NewButton("Save", func() {
		HTS, _ := entryBind.Get()
		ATS, _ := entryBind2.Get()
		main.CheckWriter()
	})
	container.NewVBox(HTHBox, HTSHBox, ATHBox, ATSHBox)
	rightSide := container.NewBorder(nil, submit, nil, nil)
	return rightSide
}
