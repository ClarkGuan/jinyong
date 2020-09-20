// +build windows

package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := MainWindow{
		Title: "金庸群侠传修改器",
		Layout: VBox{
			SpacingZero: false,
		},
		Children: []Widget{
			PushButton{Text: "存档一", OnClicked: func() {
				fmt.Printf("点击存档1\n")
			}},
			PushButton{Text: "存档二", OnClicked: func() {
				fmt.Printf("点击存档2\n")
			}},
			PushButton{Text: "存档三", OnClicked: func() {
				fmt.Printf("点击存档3\n")
			}},
		},
	}.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
