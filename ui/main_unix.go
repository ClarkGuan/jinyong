// +build !windows

package main

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/ClarkGuan/jinyong/conf"
)

func main() {
	a := app.NewWithID("金庸群侠传修改器")
	w := a.NewWindow("金庸群侠传修改器")
	w.Resize(fyne.Size{Width: 320, Height: 480})
	entry := widget.NewEntry()
	execPath, _ := conf.ExecutablePath()
	entry.SetText(execPath)
	form := widget.NewForm(
		widget.NewFormItem("安装位置：", entry),
	)
	form.SubmitText = "确定"
	form.OnSubmit = func() {
		path, err := conf.SavesPath(strings.TrimSpace(entry.Text))
		if err != nil {
			dialog.NewError(err, w).Show()
			return
		}
		fmt.Println(path)
	}
	w.SetContent(form)
	w.ShowAndRun()
}

func newPageContent(parent fyne.Window) fyne.CanvasObject {
	form := widget.NewForm(
		widget.NewFormItem("武学常识:", widget.NewEntry()),
		widget.NewFormItem("道德:", widget.NewEntry()),
		widget.NewFormItem("功夫带毒:", widget.NewEntry()),
		widget.NewFormItem("左右互搏:", widget.NewCheck("", nil)),
		widget.NewFormItem("资质:", widget.NewEntry()),
		widget.NewFormItem("武功1:", widget.NewSelect(conf.Gongfu, nil)),
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
