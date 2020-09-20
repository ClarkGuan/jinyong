// +build !windows

package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("金庸群侠传修改器")
	w.SetContent(widget.NewVBox(
		widget.NewButton("存档 1", func() {
			dialog.ShowConfirm("提醒", "您点击了存档 1", nil, w)
		}),
		widget.NewButton("存档 2", func() {
			dialog.ShowConfirm("提醒", "您点击了存档 2", nil, w)
		}),
		widget.NewButton("存档 3", func() {
			dialog.ShowConfirm("提醒", "您点击了存档 3", nil, w)
		}),
	))
	w.ShowAndRun()
}
