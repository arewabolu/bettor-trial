package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type App interface {
	appStart()
	register()
	retrieve()
}

func appStart() {
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
