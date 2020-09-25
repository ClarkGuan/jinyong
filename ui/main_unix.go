// +build !windows

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/ClarkGuan/jinyong/conf"
)

func main() {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path, _ = os.Executable()
	}

	a := app.NewWithID("金庸群侠传修改器")
	infiniteProgressBar := widget.NewProgressBarInfinite()

	w := a.NewWindow("金庸群侠传修改器")
	w.SetMaster()
	w.SetOnClosed(closeFunc)
	w.Resize(fyne.NewSize(400, 20))
	w.SetContent(infiniteProgressBar)
	w.Show()

	// 监听
	ch := make(chan os.Signal, 1)
	go func() {
		for range ch {
			dialog.NewConfirm("警告", "是否退出？", func(b bool) {
				if b {
					a.Quit()
				}
			}, w).Show()
		}
	}()
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL)

	// 异步队列
	go func() {
		savesPath, err := conf.SavesPath(path)
		if err != nil {
			w.SetContent(widget.NewLabel(fmt.Sprintf("对不起，没有找到指定的存档文件 %q", path)))
			return
		}
		w.SetContent(widget.NewHBox(
			widget.NewButton("存档1", savePathFunc(w, savesPath[0])),
			widget.NewButton("存档2", savePathFunc(w, savesPath[1])),
			widget.NewButton("存档3", savePathFunc(w, savesPath[2])),
		))
	}()

	a.Run()
}

func closeFunc() {

}

func savePathFunc(win fyne.Window, s string) func() {
	return func() {
		f, err := os.Open(s)
		if err != nil {
			showErrorDialog(win, fmt.Sprintf("打开文件失败 %v", err))
			return
		}
		defer f.Close()
		buf, err := conf.Mmap(f)
		if err != nil {
			showErrorDialog(win, fmt.Sprintf("mmap 失败 %v", err))
			return
		}

		pf := newFrame(win)
		pf.update(buf)
		win.SetContent(pf.content)
	}
}

func showErrorDialog(win fyne.Window, s string) {
	dg := dialog.NewInformation("警告", s, win)
	dg.SetOnClosed(fyne.CurrentApp().Quit)
	dg.Show()
}
