package controller

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func SaveButton(radio *widget.RadioGroup, ent ...*widget.Entry) *widget.Button {
	return widget.NewButton("Save", func() {
		HT := ent[0].Text
		AT := ent[1].Text
		HTS := ent[2].Text
		ATS := ent[3].Text
		values := []string{HT, AT, HTS, ATS}
		CheckWriter(radio.Selected, values)
		entryDel(ent...)
	})
}

func creator(text string) *widget.Label {
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
