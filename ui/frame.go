package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/ClarkGuan/jinyong/conf"
)

type frame struct {
	content fyne.CanvasObject
	form    *widget.Form

	resistanceEntry    *widget.Entry
	senseEntry         *widget.Entry
	moralityEntry      *widget.Entry
	poisonousEntry     *widget.Entry
	doubleAttackCheck  *widget.Check
	qualificationEntry *widget.Entry

	wugongs [10]*widget.Select
	jingyan [10]*widget.Entry
	friends [5]*widget.Select

	onClick func()
}

func newFrame() *frame {
	f := frame{
		resistanceEntry:    widget.NewEntry(),
		senseEntry:         widget.NewEntry(),
		moralityEntry:      widget.NewEntry(),
		poisonousEntry:     widget.NewEntry(),
		doubleAttackCheck:  widget.NewCheck("", nil),
		qualificationEntry: widget.NewEntry(),
	}

	form := widget.NewForm(
		widget.NewFormItem("毒抗:", f.resistanceEntry),
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

	for i := range f.friends {
		f.friends[i] = widget.NewSelect(conf.Friends, nil)
		form.Append(fmt.Sprintf("朋友%d", i+1), f.friends[i])
	}

	form.SubmitText = "修改"
	f.form = form

	container := widget.NewVScrollContainer(form)
	container.SetMinSize(fyne.NewSize(200, 350))
	f.content = container

	return &f
}

func (f *frame) update(p conf.Property) {
	f.resistanceEntry.SetText(fmt.Sprintf("%d", p.Resistance()))
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

	for i := range f.friends {
		f.friends[i].SetSelected(conf.Friends[p.Friend(i)])
	}

	if f.onClick != nil {
		f.form.OnSubmit = f.onClick
		f.form.Refresh()
	}
}

func (f *frame) save(p conf.Property) error {
	if i, err := conf.ParseInt16(f.resistanceEntry.Text); err != nil {
		return err
	} else {
		p.UpdateResistance(i)
	}

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

	for i := range f.friends {
		index := conf.FriendsID[f.friends[i].Selected]
		p.UpdateFriend(i, int16(index))
	}

	return nil
}
