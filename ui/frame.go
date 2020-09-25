package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/ClarkGuan/jinyong/conf"
)

type frame struct {
	content fyne.CanvasObject

	senseEntry         *widget.Entry
	moralityEntry      *widget.Entry
	poisonousEntry     *widget.Entry
	doubleAttackCheck  *widget.Check
	qualificationEntry *widget.Entry

	wugongs [10]*widget.Select
	jingyan [10]*widget.Entry
}

func newFrame(parent fyne.Window) *frame {
	f := frame{
		senseEntry:         widget.NewEntry(),
		moralityEntry:      widget.NewEntry(),
		poisonousEntry:     widget.NewEntry(),
		doubleAttackCheck:  widget.NewCheck("", nil),
		qualificationEntry: widget.NewEntry(),
	}

	form := widget.NewForm(
		widget.NewFormItem("武学常识:", f.senseEntry),
		widget.NewFormItem("道德:", f.moralityEntry),
		widget.NewFormItem("功夫带毒:", f.poisonousEntry),
		widget.NewFormItem("左右互搏:", f.doubleAttackCheck),
		widget.NewFormItem("资质:", f.qualificationEntry),
	)

	for i := range f.wugongs {
		f.wugongs[i] = widget.NewSelect(conf.Gongfu, nil)
		form.Append(fmt.Sprintf("武功%d", i+1), f.wugongs[i])
	}

	for i := range f.jingyan {
		f.jingyan[i] = widget.NewEntry()
		form.Append(fmt.Sprintf("武功%d经验", i+1), f.jingyan[i])
	}

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
	f.content = container

	return &f
}

func (f *frame) update(p conf.Property) {
	f.senseEntry.SetText(fmt.Sprintf("%d", p.Sense()))
	f.moralityEntry.SetText(fmt.Sprintf("%d", p.Morality()))
	f.poisonousEntry.SetText(fmt.Sprintf("%d", p.Poisonous()))
	f.doubleAttackCheck.SetChecked(p.DoubleAttack())
	f.qualificationEntry.SetText(fmt.Sprintf("%d", p.Qualification()))

	for i := range f.wugongs {
		f.wugongs[i].SetSelected(conf.Gongfu[p.Wugong(i)])
	}

	for i := range f.jingyan {
		f.jingyan[i].SetText(fmt.Sprintf("%d", p.Jingyan(i)))
	}
}

func (f *frame) save(p conf.Property) error {
	if i, err := conf.ParseInt16(f.senseEntry.Text); err != nil {
		return err
	} else {
		p.UpdateSense(i)
	}

	if i, err := conf.ParseInt16(f.moralityEntry.Text); err != nil {
		return err
	} else {
		p.UpdateMorality(i)
	}

	if i, err := conf.ParseInt16(f.poisonousEntry.Text); err != nil {
		return err
	} else {
		p.UpdatePoisonous(i)
	}

	p.UpdateDoubleAttack(f.doubleAttackCheck.Checked)

	if i, err := conf.ParseInt16(f.qualificationEntry.Text); err != nil {
		return err
	} else {
		p.UpdateQualification(i)
	}

	for i := range f.wugongs {
		index := conf.GongfuID[f.wugongs[i].Selected]
		p.UpdateWugong(i, int16(index))
	}

	for i := range f.jingyan {
		if v, err := conf.ParseInt16(f.jingyan[i].Text); err != nil {
			return err
		} else {
			p.UpdateJingyan(i, v)
		}
	}

	return nil
}
