package views

import (
	"bettor/controller"
	"bettor/models"
	"errors"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func writeSaveButton(Select *widget.Select, w fyne.Window, ent ...*widget.Entry) *widget.Button {
	return widget.NewButton("Save", func() {
		HT := ent[0].Text
		AT := ent[1].Text
		HTGoals := ent[2].Text
		ATGoals := ent[3].Text
		values := []string{HT, AT, HTGoals, ATGoals}

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

func prependSaveButton(Select *widget.Select, w fyne.Window, ent ...*widget.Entry) *widget.Button {
	return widget.NewButton("Save", func() {
		HT := ent[0].Text
		AT := ent[1].Text
		HTGoals := ent[2].Text
		ATGoals := ent[3].Text
		values := []string{HT, AT, HTGoals, ATGoals}

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
		err := controller.PrependMatchData(Select.Selected, values)

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

func makeImage(w fyne.Window) fyne.CanvasObject {
	uri, err := storage.ParseURI("file://" + models.GetBaseImage())
	if err != nil {
		dialog.ShowError(fmt.Errorf("could not parse \""+uri.String()+"\""), w)
	}

	val, err := storage.CanRead(uri)

	if err != nil {
		dialog.ShowError(fmt.Errorf("No folder for \""+uri.Scheme()+"\""), w)
	}
	if !val {
		dialog.ShowError(fmt.Errorf("Unable to open file \""+uri.Name()+"\"", err), w)
	}

	read, err := storage.Reader(uri)
	if err != nil {
		dialog.ShowError(fmt.Errorf("Unable to open file \""+uri.Name()+"\"", err), w)
	}
	defer read.Close()

	res, err := storage.LoadResourceFromURI(read.URI())
	if err != nil {
		dialog.ShowError(fmt.Errorf("error loading image %s", uri.Name()), w)
		return canvas.NewRectangle(color.Black)
	}

	img := canvas.NewImageFromResource(res)
	img.FillMode = canvas.ImageFillContain
	return img
}
