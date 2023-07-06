package views

import (
	"bettor/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func registerGameFullScore(w fyne.Window) fyne.CanvasObject {
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

	submit := widget.NewButton("Submit", func() {
		submitData(Select, w, &HTEnt.Entry, &ATEnt.Entry, &HTSEnt.Entry, &ATSEnt.Entry)
	})
	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})
	controlsCont := container.NewBorder(nil, nil, backButn, nil, Select)
	fullCanvas := container.NewVBox(controlsCont, HTHBox, HTSHBox, ATHBox, ATSHBox, submit)

	HTEnt.button = submit
	ATEnt.button = submit
	HTSEnt.button = submit
	ATSEnt.button = submit

	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch ke.Name {
		case fyne.Key1:
			w.Canvas().Focus(HTEnt)
		case fyne.KeyF2:
			w.Canvas().Focus(HTSEnt)
		case fyne.KeyF3:
			w.Canvas().Focus(ATEnt)
		case fyne.KeyF4:
			w.Canvas().Focus(ATSEnt)
		case fyne.KeyReturn:
			w.Canvas().Focus(submit)
			submitData(Select, w, &HTEnt.Entry, &ATEnt.Entry, &HTSEnt.Entry, &ATSEnt.Entry)
		}

	})
	return fullCanvas
}

func prependGames(w fyne.Window) fyne.CanvasObject {
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

	submit := widget.NewButton("Save", func() {
		prependSave(Select, w, &HTEnt.Entry, &ATEnt.Entry, &HTSEnt.Entry, &ATSEnt.Entry)
	})
	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})

	controlsCont := container.NewBorder(nil, nil, backButn, nil, Select)
	fullCanvas := container.NewVBox(controlsCont, HTHBox, HTSHBox, ATHBox, ATSHBox, submit)

	HTEnt.button = submit
	ATEnt.button = submit
	HTSEnt.button = submit
	ATSEnt.button = submit

	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch ke.Name {
		case fyne.Key1:
			w.Canvas().Focus(HTEnt)
		case fyne.KeyF2:
			w.Canvas().Focus(HTSEnt)
		case fyne.KeyF3:
			w.Canvas().Focus(ATEnt)
		case fyne.KeyF4:
			w.Canvas().Focus(ATSEnt)
		case fyne.KeyReturn:
			w.Canvas().Focus(submit)
			prependSave(Select, w, &HTEnt.Entry, &ATEnt.Entry, &HTSEnt.Entry, &ATSEnt.Entry)
		}
	})
	return fullCanvas
}
