package app

import (
	"context"
	"fmt"
	"go-ssh-forward/app/common/mitt"

	"github.com/getlantern/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx  context.Context
	icon []byte
}

// NewApp creates a new App application struct
func NewApp(icon []byte) *App {
	return &App{
		icon: icon,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (cls *App) Startup(ctx context.Context) {
	cls.ctx = ctx

	mitt.Mitt.On("log", cls.log)
	go cls.tray()
}

func (cls *App) tray() {
	tray := NewTray(cls.icon)
	systray.Run(func() {
		tray.OnReady(cls.ctx)
	}, nil)
}

func (cls *App) log(data any) {
	jsCode := fmt.Sprintf("window.wailsApi.log('%v')", data)

	// 2. 执行 JavaScript 代码
	runtime.WindowExecJS(cls.ctx, jsCode)
}
