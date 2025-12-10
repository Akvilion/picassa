package main

import (
	"embed"

	"pica3/internal/app"
	"pica3/pkg/config"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Створюємо екземпляр застосунку
	application := app.NewApp()

	// Отримуємо конфігурацію вікна
	windowConfig := config.DefaultWindowConfig()

	// Запускаємо застосунок
	err := wails.Run(&options.App{
		Title:                    windowConfig.Title,
		Width:                    windowConfig.Width,
		Height:                   windowConfig.Height,
		Frameless:                true, // Вікно без верхньої панелі
		BackgroundColour:         config.GetBackgroundColor(),
		WindowStartState:         options.Normal,
		EnableDefaultContextMenu: false,
		Windows:                  config.GetWindowsOptions(),
		Mac:                      config.GetMacOptions(),
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup:  application.Startup,
		OnShutdown: application.Shutdown,
		OnDomReady: application.DomReady,
		Bind: []interface{}{
			application,
		},
	})

	if err != nil {
		println("Помилка:", err.Error())
	}
}
