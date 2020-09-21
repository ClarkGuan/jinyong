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
		widget.NewFormItem("体力:", widget.NewEntry()),
		widget.NewFormItem("生命:", widget.NewEntry()),
		widget.NewFormItem("生命Max:", widget.NewEntry()),
		widget.NewFormItem("内力:", widget.NewEntry()),
		widget.NewFormItem("内力Max:", widget.NewEntry()),
		widget.NewFormItem("内力属性:", widget.NewSelect([]string{"阴", "阳", "阴阳调和"}, nil)),
		widget.NewFormItem("武器:", widget.NewEntry()),
		widget.NewFormItem("防具:", widget.NewEntry()),
		widget.NewFormItem("攻击:", widget.NewEntry()),
		widget.NewFormItem("轻功:", widget.NewEntry()),
		widget.NewFormItem("防御:", widget.NewEntry()),
		widget.NewFormItem("医疗:", widget.NewEntry()),
		widget.NewFormItem("用毒:", widget.NewEntry()),
		widget.NewFormItem("解毒:", widget.NewEntry()),
		widget.NewFormItem("抗毒:", widget.NewEntry()),
		widget.NewFormItem("拳掌:", widget.NewEntry()),
		widget.NewFormItem("御剑:", widget.NewEntry()),
		widget.NewFormItem("耍刀:", widget.NewEntry()),
		widget.NewFormItem("特殊兵器:", widget.NewEntry()),
		widget.NewFormItem("暗器:", widget.NewEntry()),
		widget.NewFormItem("武学常识:", widget.NewEntry()),
		widget.NewFormItem("道德:", widget.NewEntry()),
		widget.NewFormItem("功夫带毒:", widget.NewEntry()),
		widget.NewFormItem("左右互搏:", widget.NewCheck("", nil)),
		widget.NewFormItem("声望:", widget.NewEntry()),
		widget.NewFormItem("资质:", widget.NewEntry()),
		widget.NewFormItem("修炼道具:", widget.NewEntry()),
		widget.NewFormItem("修炼经验:", widget.NewEntry()),
		widget.NewFormItem("武功1:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功2:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功3:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功4:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功5:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功6:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功7:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功8:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功9:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功10:", widget.NewSelect([]string{"野球拳", "普通攻击"}, nil)),
		widget.NewFormItem("武功1经验:", widget.NewEntry()),
		widget.NewFormItem("武功2经验:", widget.NewEntry()),
		widget.NewFormItem("武功3经验:", widget.NewEntry()),
		widget.NewFormItem("武功4经验:", widget.NewEntry()),
		widget.NewFormItem("武功5经验:", widget.NewEntry()),
		widget.NewFormItem("武功6经验:", widget.NewEntry()),
		widget.NewFormItem("武功7经验:", widget.NewEntry()),
		widget.NewFormItem("武功8经验:", widget.NewEntry()),
		widget.NewFormItem("武功9经验:", widget.NewEntry()),
		widget.NewFormItem("武功10经验:", widget.NewEntry()),
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
