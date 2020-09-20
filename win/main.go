package main

import (
	"github.com/lxn/walk/declarative"
)

func main() {
	mainWindow := declarative.MainWindow{
		Title:   "金庸原版修改器",
		MinSize: declarative.Size{Width: 600, Height: 400},
		Layout:  declarative.VBox{},
	}
	mainWindow.Run()
}
