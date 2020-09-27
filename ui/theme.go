package main

import (
	"fyne.io/fyne"
	rice "github.com/GeertJohan/go.rice"
)

type themeWrap struct {
	fyne.Theme
	res fyne.Resource
}

func (t *themeWrap) TextFont() fyne.Resource {
	if t.res != nil {
		return t.res
	} else {
		return t.Theme.TextFont()
	}
}

func (t *themeWrap) TextBoldFont() fyne.Resource {
	if t.res != nil {
		return t.res
	} else {
		return t.Theme.TextBoldFont()
	}
}

func (t *themeWrap) TextItalicFont() fyne.Resource {
	if t.res != nil {
		return t.res
	} else {
		return t.Theme.TextItalicFont()
	}
}

func (t *themeWrap) TextBoldItalicFont() fyne.Resource {
	if t.res != nil {
		return t.res
	} else {
		return t.Theme.TextBoldItalicFont()
	}
}

func theme(t fyne.Theme) (fyne.Theme, error) {
	box, err := rice.FindBox("font")
	if err != nil {
		return nil, err
	}
	bytes, err := box.Bytes("YaHei.Consolas.1.12.ttf")
	if err != nil {
		return nil, err
	}
	resource := fyne.NewStaticResource("hanzi.ttf", bytes)
	return &themeWrap{t, resource}, nil
}
