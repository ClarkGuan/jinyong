package main

import (
	"fmt"
	"os"

	. "github.com/lxn/walk/declarative"
)

func main() {
	_, err := MainWindow{
		Title:   "金庸群侠传修改器",
		MinSize: Size{Width: 400, Height: 480},
		Layout:  HBox{},
	}.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
