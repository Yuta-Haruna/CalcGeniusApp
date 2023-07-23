package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("GoodNight")

	// ボタンを作成
	button := widget.NewButton("Click Here", func() {
		// ボタンがクリックされたときに実行する処理をここに書きます
		// 例：ウィンドウにメッセージを表示する
		myWindow.SetContent(widget.NewLabel("Button Click!"))
	})

	// ボタンをウィンドウに配置
	myWindow.SetContent(container.NewVBox(button))

	// ウィンドウを表示
	myWindow.ShowAndRun()
}
