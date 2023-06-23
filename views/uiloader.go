package views

import (
	"bettor/controller"
	"bettor/models"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// also an idea is to create a bayesian model that generates random numbers between 1 and 5
// Can use the os.StartProcess function to start a new process check os/exec
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
		w.SetContent(createNewGame(w))
	})
	//width := but1.Size().Width
	//but1.Resize(fyne.NewSize(40, 40))
	but2 := widget.NewButtonWithIcon("register new Game", theme.ContentAddIcon(), func() {
		w.SetContent(registerGameFullScore(w))
	})
	but3 := widget.NewButtonWithIcon("search for game", theme.SearchIcon(), func() {
		w.SetContent(piSearch(w))
	})
	but4 := widget.NewButtonWithIcon("Register team for game", theme.FileTextIcon(), func() {
		w.SetContent(registerTeams(w))
	})
	grid := container.NewGridWithRows(5, but1, but2, but3, but4)

	return grid
}

// A tab for registering teams to already created games.
func registerTeams(w fyne.Window) fyne.CanvasObject {
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
func createNewGame(w fyne.Window) fyne.CanvasObject {
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
