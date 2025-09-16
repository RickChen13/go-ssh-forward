package main

import (
	"embed"
	"go-ssh-forward/app"
	"go-ssh-forward/app/bll"
	"go-ssh-forward/app/common/Const"
	"go-ssh-forward/app/common/sqlite3"
	"go-ssh-forward/app/dal/forwardRule"
	"go-ssh-forward/app/dal/sshServer"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS
	//go:embed build/windows/icon.ico
	icon []byte
)

func main() {
	// 初始化数据库
	err := sqlite3.Init(Const.BASE_PATH + "/config.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	frDal := forwardRule.NewForwardRuleDal(sqlite3.Db)
	frBll := bll.NewForwardRuleBll(frDal)

	ssDal := sshServer.NewSshServerDal(sqlite3.Db)
	ssBll := bll.NewSshServerBll(ssDal)

	forwardDal := bll.NewForwardDal()

	myApp := app.NewApp(icon)
	// Create application with options
	err = wails.Run(&options.App{
		Width:     1024,
		Height:    768,
		MinWidth:  400,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//BackgroundColour:  &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:         myApp.Startup,
		HideWindowOnClose: true,
		// 无边框
		Frameless: true,
		// CSSDragProperty:   "widows",
		// CSSDragValue:      "1",
		Bind: []any{
			frBll,
			ssBll,
			forwardDal,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
