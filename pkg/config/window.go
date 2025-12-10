package config

import (
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

// WindowConfig містить налаштування для прозорого вікна
type WindowConfig struct {
	Title  string
	Width  int
	Height int
}

// DefaultWindowConfig повертає стандартні налаштування вікна
func DefaultWindowConfig() *WindowConfig {
	return &WindowConfig{
		Title:  "Pica3",
		Width:  800,
		Height: 600,
	}
}

// GetMacOptions повертає налаштування для macOS
func GetMacOptions() *mac.Options {
	return &mac.Options{
		TitleBar:             mac.TitleBarHiddenInset(),
		WebviewIsTransparent: true,
		WindowIsTranslucent:  true,
	}
}

// GetWindowsOptions повертає налаштування для Windows
func GetWindowsOptions() *windows.Options {
	return &windows.Options{
		WebviewIsTransparent:              true,
		WindowIsTranslucent:               true,
		DisableFramelessWindowDecorations: true,
		// BackdropType можна увімкнути для ефектів Windows:
		// BackdropType: windows.Acrylic, // Acrylic ефект
		// BackdropType: windows.Mica,    // Mica ефект (Windows 11)
	}
}

// GetBackgroundColor повертає прозорий колір фону
func GetBackgroundColor() *options.RGBA {
	return &options.RGBA{R: 0, G: 0, B: 0, A: 0}
}
