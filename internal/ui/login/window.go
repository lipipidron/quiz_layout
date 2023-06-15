package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"game_with_Nikita/internal/fileWork"
	"game_with_Nikita/internal/ui/game"
	"log"
)

func StartWindow() {
	myApp := app.New()
	myWindow := myApp.NewWindow("qwe")
	myWindow.Resize(fyne.NewSize(500, 200))
	myApp.Settings().SetTheme(theme.DarkTheme())
	ic, err := fyne.LoadResourceFromPath("assets/logos/logo.png")
	if err != nil {
		log.Fatal(err)
	}
	myWindow.SetIcon(ic)

	loginEntry := widget.NewEntry()
	loginEntry.SetPlaceHolder("Enter login")
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	content := container.NewVBox()
	logIn := widget.NewButton("Log in", func() {
		loggedIn, loginName := LogIn(loginEntry, passwordEntry, myWindow, content)
		if loggedIn {
			game.FirstStageLayout(content, loginName)
		}
	})

	signUp := widget.NewButton("Sign up", func() {
		SignUp(loginEntry, passwordEntry, myWindow, content)
	})

	update := widget.NewButton("Update", func() {
		fileWork.Update()
	})

	emptyLabel := widget.NewLabel("")
	buttons := container.NewGridWithColumns(10, logIn, signUp)
	for i := 0; i < 7; i++ {
		buttons.Add(emptyLabel)
	}
	buttons.Add(update)
	content = container.NewVBox(loginEntry, passwordEntry, buttons)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
