// +build !windows

package main

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.NewWithID("金庸群侠传修改器")
	w := a.NewWindow("金庸群侠传修改器")
	w.SetContent(widget.NewTabContainer(
		widget.NewTabItem("存档 1", newPageContent(w)),
		widget.NewTabItem("存档 2", newPageContent(w)),
		widget.NewTabItem("存档 3", newPageContent(w)),
	))
	w.ShowAndRun()
}

func newPageContent(parent fyne.Window) fyne.CanvasObject {
	form := widget.NewForm(
		widget.NewFormItem("武学常识:", widget.NewEntry()),
		widget.NewFormItem("道德:", widget.NewEntry()),
		widget.NewFormItem("功夫带毒:", widget.NewEntry()),
		widget.NewFormItem("左右互搏:", widget.NewCheck("", nil)),
		widget.NewFormItem("资质:", widget.NewEntry()),
		widget.NewFormItem("武功1:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功1经验:", widget.NewEntry()),
	)
	form.SubmitText = "修改"
	form.OnSubmit = func() {
		progressInfiniteDialog := dialog.NewProgressInfinite("金庸群侠传修改器", "修改中，请等待……", parent)
		progressInfiniteDialog.Show()
		time.AfterFunc(5*time.Second, func() {
			progressInfiniteDialog.Hide()
		})
	}
	container := widget.NewVScrollContainer(form)
	container.SetMinSize(fyne.NewSize(200, 350))
	return container
}
