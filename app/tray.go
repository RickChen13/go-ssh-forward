package app

import (
	"context"
	_ "embed"

	"github.com/getlantern/systray"
	wRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

type Tray struct {
	icon []byte
}

// NewApp creates a new App application struct
func NewTray(icon []byte) *Tray {
	return &Tray{
		icon: icon,
	}
}

func (cls *Tray) OnReady(ctx context.Context) {
	// 设置托盘图标
	systray.SetIcon(cls.icon)
	systray.SetTitle("My Wails App")
	systray.SetTooltip("这是一个 Wails 应用的系统托盘")

	// 创建菜单项
	mShow := systray.AddMenuItem("打开主窗口", "打开主窗口")
	mQuit := systray.AddMenuItem("退出", "退出应用")
	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				wRuntime.WindowShow(ctx)
			case <-mQuit.ClickedCh:
				wRuntime.Quit(ctx)
				return
			}
		}
	}()
}
