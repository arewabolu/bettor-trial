package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

type appController struct {
	application fyne.App
	appwindow   fyne.Window
}

type submitEntry struct {
	widget.Entry
	window fyne.Window
	button *widget.Button
}

func newSubmitEntry(w fyne.Window) *submitEntry {
	e := &submitEntry{}
	e.ExtendBaseWidget(e)
	e.window = w
	return e
}

func (s *submitEntry) TypedKey(k *fyne.KeyEvent) {
	if k.Name == fyne.KeyEscape {
		s.window.Canvas().Unfocus()
		return
	}
	if k.Name == fyne.KeyReturn {
		s.window.Canvas().Unfocus()
		test.Tap(s.button)
		return
	}

	s.Entry.TypedKey(k)
}
