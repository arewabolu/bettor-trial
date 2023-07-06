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

	image := makeImage(w)

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
	controlsCont := container.NewBorder(nil, nil, backButn1, nil, Select)

	Box := container.NewBorder(controlsCont, submit, nil, nil, vBox2)

	submit.OnTapped = func() {
		values := []string{HTEnt.Text, ATEnt.Text}
		if Select.Selected == "" {
			dialog.ShowError(errors.New("please select the game type"), w)
			return
		}

		ratingH, ratingA, err := controller.ReadPi(Select.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		H2HratingH, H2HratingA, err := controller.ReadH2HPi(Select.Selected, values)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		team1 := widget.NewLabel(fmt.Sprintf("%s: %.4f | Head-to-Head: %.4f", values[0], ratingH, H2HratingH))
		team2 := widget.NewLabel(fmt.Sprintf("%s: %.4f | Head-to-Head: %.4f ", values[1], ratingA, H2HratingA))
		RD := widget.NewLabel(fmt.Sprintf("RD: %.4f | H2H Difference: %.4f ", ratingH-ratingA, H2HratingH-H2HratingA))
		backButn := widget.NewButtonWithIcon("Search Again", theme.NavigateBackIcon(), func() {
			Box.RemoveAll()
			w.SetContent(piSearch(w))
		})
		backButn.IconPlacement = widget.ButtonIconLeadingText
		exit := widget.NewButtonWithIcon("Exit search", theme.CancelIcon(), func() {
			Box.RemoveAll()
			w.SetContent(uiLoader(w))
		})
		exit.IconPlacement = widget.ButtonIconLeadingText
		teamsInfo := container.NewBorder(nil, container.NewVBox(team1, team2, RD), backButn, exit)
		w.SetContent(container.NewBorder(teamsInfo, nil, nil, nil, container.NewPadded(image)))
	}

	return Box
}
