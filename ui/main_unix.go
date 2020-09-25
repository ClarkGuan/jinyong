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
	w.Resize(fyne.NewSize(400, 20))
	w.SetContent(infiniteProgressBar)
	w.Show()

	// 监听
	ch := make(chan os.Signal, 1)
	go func() {
		for range ch {
			dialog.ShowConfirm("警告", "是否退出？", func(b bool) {
				if b {
					a.Quit()
				}
			}, w)
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

		frames := make([]*frame, 3)
		props := make([]conf.Property, 3)

		for i := range savesPath {
			if frames[i], props[i], err = savePathFunc(w, savesPath[i]); err != nil {
				showErrorDialog(w, fmt.Sprintf("打开文件或 mmap 失败 %v", err))
				return
			}
		}

		w.SetContent(widget.NewTabContainer(
			widget.NewTabItem("存档1", frames[0].content),
			widget.NewTabItem("存档2", frames[1].content),
			widget.NewTabItem("存档3", frames[2].content),
		))
	}()

	a.Run()
}

func unixMmap(path string) ([]byte, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return conf.Mmap(f)
}

func savePathFunc(win fyne.Window, s string) (*frame, conf.Property, error) {
	if buf, err := unixMmap(s); err != nil {
		return nil, nil, err
	} else {
		pf := newFrame(win)
		pf.onClick = func() {
			//infinite := dialog.NewProgressInfinite("注意", "数据存储中……", win)
			//infinite.Show()
			fmt.Println(pf.save(buf))
			conf.Munmap(buf)
			//infinite.Hide()
			fmt.Println("点击了")
		}
		pf.update(buf)
		return pf, buf, nil
	}
}

func showErrorDialog(win fyne.Window, s string) {
	dg := dialog.NewInformation("警告", s, win)
	dg.SetOnClosed(fyne.CurrentApp().Quit)
	dg.Show()
}
