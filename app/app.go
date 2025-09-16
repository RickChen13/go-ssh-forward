package app

import (
	"context"

	"github.com/getlantern/systray"
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

	go cls.tray()
}

func (cls *App) tray() {
	tray := NewTray(cls.icon)
	systray.Run(func() {
		tray.OnReady(cls.ctx)
	}, nil)
}
