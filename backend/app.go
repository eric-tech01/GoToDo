package backend

import (
	"context"
	"embed"
	"io"
	"net/http"
	"netix-todo-app/backend/gateway"
	"netix-todo-app/backend/lib/utils"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	log "github.com/eric-tech01/simple-log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func (e *Engine) CreateApp(assets embed.FS, icon []byte) {
	user := gateway.NewUserController()
	imageApp := gateway.NewImageApp()
	todoApp := gateway.NewTodoController()
	settingApp := gateway.NewSettingController()
	// Create application with options
	err := wails.Run(&options.App{
		Title:     "今日任务",
		Width:     300,
		Height:    480,
		MinWidth:  250,
		MinHeight: 400,
		MaxWidth:  500,
		MaxHeight: 800,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewFileLoader(),
		},
		Fullscreen:       false,
		Frameless:        false,
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: func(ctx context.Context) {
			sw, _ := utils.GetSystemMetrics()
			runtime.WindowSetPosition(ctx, sw-500, 50)
			user.Startup(ctx)
			imageApp.Startup(ctx)
			todoApp.Startup(ctx)
			settingApp.Startup(ctx)
		},
		Bind: []interface{}{
			user,
			imageApp,
			todoApp,
			settingApp,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},

		Logger: NewDefaultLogger(),
		Windows: &windows.Options{
			ResizeDebounceMS: 1,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

type FileLoader struct {
	http.Handler
}

func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Debugf("file req url ", r.URL)

	filePath := r.URL.Path
	rootPath := filePath[0:3]
	fileDir := rootPath[1:2] + ":" + filePath[3:]

	log.Debug(fileDir)
	f, err := os.OpenFile(fileDir, os.O_RDONLY, 0)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	bs, err := io.ReadAll(f)
	if err != nil {
		log.Error(err)
	}
	w.Write(bs)
}

// DefaultLogger is a utility to log messages to a number of destinations
type DefaultLogger struct{}

// NewDefaultLogger creates a new Logger.
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

// Print works like Sprintf.
func (l *DefaultLogger) Print(message string) {
	log.Debug(message)
}

// Trace level logging. Works like Sprintf.
func (l *DefaultLogger) Trace(message string) {
	log.Info("TRA | " + message)
}

// Debug level logging. Works like Sprintf.
func (l *DefaultLogger) Debug(message string) {
	log.Debug("DEB | " + message)
}

// Info level logging. Works like Sprintf.
func (l *DefaultLogger) Info(message string) {
	log.Info("INF | " + message)
}

// Warning level logging. Works like Sprintf.
func (l *DefaultLogger) Warning(message string) {
	log.Warn(message)
}

// Error level logging. Works like Sprintf.
func (l *DefaultLogger) Error(message string) {
	log.Error(message)
}

// Fatal level logging. Works like Sprintf.
func (l *DefaultLogger) Fatal(message string) {
	log.Fatal(message)
	os.Exit(1)
}
