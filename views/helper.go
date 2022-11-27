package views

import (
	"bettor/controller"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func SaveButton(radio *widget.RadioGroup, w fyne.Window, ent ...*widget.Entry) *widget.Button {
	return widget.NewButton("Save", func() {
		HT := ent[0].Text
		AT := ent[1].Text
		HTS := ent[2].Text
		ATS := ent[3].Text
		values := []string{HT, AT, HTS, ATS}

		if radio.Selected == "" {
			dialog.ShowError(errors.New("please select the game type"), w)
			return
		}

		err := controller.CheckWriter(radio.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		entryDel(ent...)
	})
}

func Creator(text string) *widget.Label {
	label := new(widget.Label)
	label.Wrapping = fyne.TextWrapBreak
	label.SetText(text)
	return label
}

func entryDel(entries ...*widget.Entry) {
	for _, entry := range entries {
		entry.SetText("")
	}
}
