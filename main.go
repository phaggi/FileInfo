package main

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	myApp := app.NewWithID("MyNewApp")

	MyWindow := myApp.NewWindow("FileInfo")
	openFileButton := widget.NewButton("Выбрать файл", func() {
		dialog.ShowFileOpen(
			func(file fyne.URIReadCloser, err error) {
				handlerFileRead(file, MyWindow, err)
			},
			MyWindow,
		)
	})
	closeFileButton := widget.NewButton("Закрыть", func() {
		MyWindow.Close()
	})
	setDarkTheme := widget.NewButton("Поставить светлую тему", func() {
		myApp.Settings().SetTheme(theme.LightTheme())
	})

	setLightTheme := widget.NewButton("Поставить темную тему", func() {
		myApp.Settings().SetTheme(theme.DarkTheme())
	})
	content := container.NewVBox(openFileButton, setDarkTheme, setLightTheme, closeFileButton)
	MyWindow.SetContent(content)
	MyWindow.Resize(fyne.NewSize(400, 200))
	MyWindow.ShowAndRun()
}

func handlerFileRead(file fyne.URIReadCloser, MyWindow fyne.Window, err error) {

	if file == nil {
		if err != nil {
			dialog.ShowError(err, MyWindow)
			return
		}
		return
	}
	info, err := os.Stat(file.URI().Path())
	if err != nil {
		dialog.ShowError(err, MyWindow)
		return
	}

	message := "Название: " + info.Name() + "\n" + "Размер: " + formatSize(info.Size()) + "\n" + "Тип: " + detectType(info.Name())

	dialog.ShowInformation("Информация о файле", message, MyWindow)
}
func formatSize(size int64) string {
	return fmt.Sprintf("%d bytes %d KB %d MB", size, size/1024, size/1024/1024)
}

func detectType(name string) string {
	result := strings.Split(name, ".")
	return result[len(result)-1]
}
