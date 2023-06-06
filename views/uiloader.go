package views

import (
	"bettor/controller"
	"bettor/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func AppStart() {
	app := &appController{
		application: app.NewWithID("com.example.myid"),
	}
	app.appwindow = app.application.NewWindow("Bettor")

	//app.appwindow.Resize(fyne.NewSize(550, 500))
	app.appwindow.SetContent(uiLoader(app.appwindow))
	app.appwindow.ShowAndRun()
}

func uiLoader(w fyne.Window) fyne.CanvasObject {
	but1 := widget.NewButtonWithIcon("create new category", theme.ContentAddIcon(), func() {
		w.SetContent(CreateNewGame(w))
	})
	//width := but1.Size().Width
	//but1.Resize(fyne.NewSize(width, 20))
	but2 := widget.NewButtonWithIcon("register new Game", theme.ContentAddIcon(), func() {
		w.SetContent(RegisterGameFullScore(w))
	})
	but3 := widget.NewButtonWithIcon("search for game", theme.SearchIcon(), func() {
		w.SetContent(Searchwith2Teams(w))
	})
	but4 := widget.NewButtonWithIcon("search for Team Data", theme.SearchIcon(), func() {
		w.SetContent(SearchWith1Team(w))
	})
	but5 := widget.NewButtonWithIcon("Register team for game", theme.FileTextIcon(), func() {
		w.SetContent(RegisterTeams(w))
	})
	grid := container.NewAdaptiveGrid(5, but1, but2, but3, but4, but5)

	return grid
}

// A tab for registering teams to already created games.
func RegisterTeams(w fyne.Window) fyne.CanvasObject {
	radOptions := models.DirIterator(models.GetBase())
	Select := widget.NewSelect(radOptions, func(s string) {
	})

	competitors := widget.NewEntry()
	competitors.MultiLine = false
	competitors.SetPlaceHolder("Add new team")
	saveCompetitors := widget.NewButton("Add team", func() {
		models.AddToTeam(Select.Selected, competitors.Text)
		models.AddtoRating(Select.Selected, competitors.Text)
		entryDel(competitors)
	})
	genRating := widget.NewButton("Generating Pi-ratings", func() {
		controller.GenRating(Select.Selected)
	})
	info := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		infoDialog := dialog.NewInformation("", "Should be used after entering all teams for the game.", w)
		infoDialog.Show()
		infoDialog.SetOnClosed(func() { w.RequestFocus() })
	})
	info.Resize(fyne.NewSize(10, 10))
	cont := container.NewBorder(nil, nil, info, nil, genRating)
	exit := widget.NewButton("exit", func() {
		w.SetContent(uiLoader(w))
	})

	hZ := container.NewBorder(nil, exit, nil, nil, competitors, saveCompetitors)
	Vcont := container.NewVBox(Select, hZ, cont)
	return Vcont
}

// A tab to create a new game and register teams to compete in such a game
func CreateNewGame(w fyne.Window) fyne.CanvasObject {
	gameType := new(widget.Entry)
	store := &gameType.Text

	competitors := widget.NewEntry()
	competitors.MultiLine = false
	competitors.SetPlaceHolder("Add new team")
	saveCompetitors := widget.NewButton("Add team", func() {
		models.AddToTeam(*store, competitors.Text)
		models.AddtoRating(*store, competitors.Text)
		entryDel(competitors)
	})
	exit := widget.NewButton("exit", func() {
		w.SetContent(uiLoader(w))
	})
	hZ := container.NewBorder(nil, exit, nil, saveCompetitors, competitors)
	button := widget.NewButtonWithIcon("Create", theme.NavigateNextIcon(), func() {
		models.CreateFile(gameType.Text)
		models.CreateTeamsFile(gameType.Text)
		models.CreateRatingFile(gameType.Text)
		w.SetContent(hZ)
	})

	backButn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		w.SetContent(uiLoader(w))
	})

	vBox := container.NewVBox(backButn, gameType, button)
	return vBox
}
