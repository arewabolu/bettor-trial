package views

import (
	"bettor/controller"
	"errors"
	"fmt"

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
		AT1stHalf := ent[4].Text
		AT2ndHalf := ent[5].Text

		values := []string{HT, HT1stHalf, HT2ndHalf, AT, AT1stHalf, AT2ndHalf}

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
			dlog := dialog.NewError(errors.New("please select the game type"), w)
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

func SaveButton2(w fyne.Window, ent ...*widget.Entry) *widget.Button {
	return widget.NewButton("Save", func() {
		HT := ent[0].Text
		HT1stHalf := ent[1].Text
		HT2ndHalf := ent[2].Text
		AT := ent[3].Text
		AT1stHalf := ent[4].Text
		AT2ndHalf := ent[5].Text

		values := []string{HT, HT1stHalf, HT2ndHalf, AT, AT1stHalf, AT2ndHalf}

		err := controller.WriteMatchDataHalfs(values)
		if err != nil {
			dlog := dialog.NewError(errors.New("please select the game type"), w)
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

func tableRender(team []string, GP int, percentageWinorDraw []float64) *widget.Table {
	table := widget.NewTable(
		func() (int, int) { return 3, 5 },
		func() fyne.CanvasObject { return widget.NewLabel("xxxxxx") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			switch {
			case tci.Col == 0 && tci.Row == 0:
				label.SetText("Team")
			case tci.Col == 1 && tci.Row == 0:
				label.SetText("GP")
			case tci.Col == 2 && tci.Row == 0:
				label.SetText("even %")
			case tci.Col == 3 && tci.Row == 0:
				label.SetText("odd%")

			case tci.Col == 0 && tci.Row == 1:
				label.SetText(team[0])
			case tci.Col == 1 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%d", GP))
			case tci.Col == 2 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[0]))
			case tci.Col == 3 && tci.Row == 1:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[1]))

			case tci.Col == 0 && tci.Row == 2:
				label.SetText(team[1])
			case tci.Col == 1 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%d", GP))
			case tci.Col == 2 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[0]))
			case tci.Col == 3 && tci.Row == 2:
				label.SetText(fmt.Sprintf("%.2f", percentageWinorDraw[1]))
			}

		})

	return table
}

/*func TabKey2(butn *widget.Button) *fyne.KeyEvent {
	fmt.Println("called2")
	key := new(fyne.KeyEvent)
	if key.Name == fyne.KeyReturn {
		test.Tap(butn)
	}
	return key
}
*/
