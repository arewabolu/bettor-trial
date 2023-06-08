package views

import (
	"bettor/controller"
	"bettor/models"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func RegisterGameFullScore(w fyne.Window) fyne.CanvasObject {
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

	fullCanvas := container.NewVBox(backButn, Select, HTHBox, HTSHBox, ATHBox, ATSHBox, submit)
	ent := []*widget.Entry{&HTEnt.Entry, &ATEnt.Entry, &ATSEnt.Entry, &HTSEnt.Entry}
	saveFunc := func() {
		values := []string{HTEnt.Text, ATEnt.Text, HTSEnt.Text, ATSEnt.Text}

		if Select.Selected == "" {
			dlog := dialog.NewError(errors.New("please select the game type"), w)
			dlog.Show()
			w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
				if ke.Name == fyne.KeyReturn {
					dlog.Hide()
					w.RequestFocus()
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
					dlog.Hide()
					w.RequestFocus()
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

// deprecated: not implemented for usage
func RegisterGamehalfScores(w fyne.Window) fyne.CanvasObject {
	HTEnt := newSubmitEntry(w)
	ATEnt := newSubmitEntry(w)
	HT1stHalf := newSubmitEntry(w)
	AT1stHalf := newSubmitEntry(w)
	HT2ndHalf := newSubmitEntry(w)
	AT2ndHalf := newSubmitEntry(w)

	HTLabel := widget.NewLabel("Home team:")
	ATLabel := widget.NewLabel("Away team:")
	HT1stHalfLabel := widget.NewLabel("Home teams firs-thalf Score")
	AT1stHalfLabel := widget.NewLabel("Away team first-half Score:")
	HT2ndHalfLabel := widget.NewLabel("Home teams second-half Score")
	AT2ndHalfLabel := widget.NewLabel("Away team second-half Score:")

	HTHBox := container.NewBorder(nil, nil, HTLabel, nil, HTEnt)
	ATHBox := container.NewBorder(nil, nil, ATLabel, nil, ATEnt)
	HT1stHalfHBox := container.NewBorder(nil, nil, HT1stHalfLabel, nil, HT1stHalf)
	AT1stHalfHBox := container.NewBorder(nil, nil, AT1stHalfLabel, nil, AT1stHalf)
	HT2ndHalfHBox := container.NewBorder(nil, nil, HT2ndHalfLabel, nil, HT2ndHalf)
	AT2ndHalfHBox := container.NewBorder(nil, nil, AT2ndHalfLabel, nil, AT2ndHalf)
	ent := []*widget.Entry{&HTEnt.Entry, &HT1stHalf.Entry, &HT2ndHalf.Entry, &ATEnt.Entry, &AT1stHalf.Entry, &AT2ndHalf.Entry}

	submit := SaveButton2(w, ent...)
	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})

	fullCanvas := container.NewVBox(backButn, HTHBox, ATHBox, HT1stHalfHBox, AT1stHalfHBox, HT2ndHalfHBox, AT2ndHalfHBox, submit)

	HTEnt.button = submit
	ATEnt.button = submit
	HT1stHalf.button = submit
	HT2ndHalf.button = submit
	AT1stHalf.button = submit
	AT2ndHalf.button = submit

	w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		switch ke.Name {
		case fyne.Key1:
			w.Canvas().Focus(HTEnt)
		case fyne.KeyF2:
			w.Canvas().Focus(HT1stHalf)
		case fyne.KeyF3:
			w.Canvas().Focus(HT2ndHalf)
		case fyne.KeyF4:
			w.Canvas().Focus(ATEnt)
		case fyne.KeyF5:
			w.Canvas().Focus(AT1stHalf)
		case fyne.KeyF6:
			w.Canvas().Focus(AT2ndHalf)
		case fyne.KeyReturn:
			w.Canvas().Focus(submit)
			test.Tap(submit)
		}

	})
	return fullCanvas
}
