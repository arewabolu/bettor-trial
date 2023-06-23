package views

import (
	"bettor/controller"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func SaveButton(Select *widget.Select, w fyne.Window, ent ...*widget.Entry) *widget.Button {
	return widget.NewButton("Save", func() {
		HT := ent[0].Text
		HT1stHalf := ent[1].Text
		HT2ndHalf := ent[2].Text
		AT := ent[3].Text
		values := []string{HT, HT1stHalf, HT2ndHalf, AT}

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
			dlog := dialog.NewError(err, w)
			dlog.Show()
			w.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
				if ke.Name == fyne.KeyReturn {
					dlog.Hide()
				}
			})
			return
		}
		entryDel(ent...)
	})
}

func entryDel(entries ...*widget.Entry) {
	for _, entry := range entries {
		entry.SetText("")
	}
}

func TabKey(w fyne.Window, butn *widget.Button) *fyne.KeyEvent {
	key := new(fyne.KeyEvent)
	if key.Name == fyne.KeyEscape {
		w.Canvas().Unfocus()
	}
	if key.Name == fyne.KeyReturn {
		test.Tap(butn)
	}
	return key
}
