package main

import (
	"embed"
	"netix-todo-app/backend"
	"os"

	log "github.com/eric-tech01/simple-log"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	//启动系统初始化服务，按顺序初始化
	eng := backend.Engine{}
	err := eng.Startup(eng.InitLog, eng.LoadResource, eng.InitDB, eng.InitSystemConfig)
	if err != nil {
		log.Error("engine startup failed")
		log.Fatal(err)
		os.Exit(1)
	}
	//创建GUI应用
	eng.CreateApp(assets, icon)
}
