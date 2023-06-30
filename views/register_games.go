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

	submit := writeSaveButton(Select, w, &HTEnt.Entry, &ATEnt.Entry, &HTSEnt.Entry, &ATSEnt.Entry)
	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})

	fullCanvas := container.NewVBox(backButn, Select, HTHBox, HTSHBox, ATHBox, ATSHBox, submit)
	ent := []*widget.Entry{&HTEnt.Entry, &ATEnt.Entry, &ATSEnt.Entry, &HTSEnt.Entry}
	saveFunc := func() {
		values := []string{HTEnt.Text, ATEnt.Text, HTSEnt.Text, ATSEnt.Text}

		if Select.Selected == "" {
			dlog := dialog.NewError(errors.New("please select the game type"), w)
			dlog.Show()
			w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
				if ke.Name == fyne.KeyReturn {
					defer w.Canvas().Focus(HTEnt)
					dlog.Hide()
				}
			})
			return
		}

		err := controller.CheckWriter(Select.Selected, values)
		if err != nil {
			dlog := dialog.NewError(err, w)
			dlog.Show()
			w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
				if ke.Name == fyne.KeyReturn {
					defer w.Canvas().Focus(HTEnt)
					dlog.Hide()
				}
			})
		}

		entryDel(ent...)
		fullCanvas.Refresh()
	}

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
			saveFunc()
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

	submit := prependSaveButton(Select, w, &HTEnt.Entry, &ATEnt.Entry, &HTSEnt.Entry, &ATSEnt.Entry)
	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})

	fullCanvas := container.NewVBox(backButn, Select, HTHBox, HTSHBox, ATHBox, ATSHBox, submit)
	ent := []*widget.Entry{&HTEnt.Entry, &ATEnt.Entry, &ATSEnt.Entry, &HTSEnt.Entry}
	saveFunc := func() {
		values := []string{HTEnt.Text, ATEnt.Text, HTSEnt.Text, ATSEnt.Text}

		if Select.Selected == "" {
			dlog := dialog.NewError(errors.New("please select the game type"), w)
			dlog.Show()
			w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
				if ke.Name == fyne.KeyReturn {
					defer w.Canvas().Focus(HTEnt)
					dlog.Hide()
				}
			})
			return
		}
		fmt.Println()
		err := controller.PrependMatchData(Select.Selected, values)
		if err != nil {
			dlog := dialog.NewError(err, w)
			dlog.Show()
			w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
				if ke.Name == fyne.KeyReturn {
					defer w.Canvas().Focus(HTEnt)
					dlog.Hide()
				}
			})
		}

		entryDel(ent...)
		fullCanvas.Refresh()
	}

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
			saveFunc()
		}

	})
	return fullCanvas
}
