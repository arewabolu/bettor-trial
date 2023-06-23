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

func piSearch(w fyne.Window) fyne.CanvasObject {
	HTEnt := new(widget.Entry)
	ATEnt := new(widget.Entry)
	HTLabel := widget.NewLabel("Home Team:")
	ATLabel := widget.NewLabel("Away Team:")

	submit := new(widget.Button)
	submit.Text = "Search"

	radOptions := models.DirIterator(models.GetBase())
	Select := widget.NewSelect(radOptions, func(s string) {
	})

	backButn1 := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})
	HTHBox := container.NewBorder(nil, nil, HTLabel, nil, HTEnt)
	ATHBox := container.NewBorder(nil, nil, ATLabel, nil, ATEnt)
	vBox2 := container.NewVBox(HTHBox, ATHBox)

	Box := container.NewBorder(container.NewVBox(backButn1, Select), submit, nil, nil, vBox2)

	submit.OnTapped = func() {
		values := []string{HTEnt.Text, ATEnt.Text}
		if Select.Selected == "" {
			dialog.ShowError(errors.New("please select the game type"), w)
			return
		}
		//GP, percentageWinorDraw, odds,
		ratingH, ratingA, err := controller.CheckReader(Select.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		team1 := widget.NewLabel(fmt.Sprintf("%s:	%.4f", values[0], ratingH))
		team2 := widget.NewLabel(fmt.Sprintf("%s:	%.4f", values[1], ratingA))
		RD := widget.NewLabel(fmt.Sprintf("%s:	%.4f", "rating difference", ratingH-ratingA))
		backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
			Box.RemoveAll()
			w.SetContent(piSearch(w))
		})

		teamsInfo := container.NewVBox(backButn, team1, team2, RD)
		w.SetContent(container.NewBorder(teamsInfo, nil, nil, nil, makeImage(w)))

	}

	return Box
}
