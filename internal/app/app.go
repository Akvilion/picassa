package app

import (
	"context"
)

// App структура для зберігання стану застосунку
type App struct {
	ctx context.Context
}

// NewApp створює новий екземпляр застосунку
func NewApp() *App {
	return &App{}
}

// Startup викликається при запуску застосунку
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Shutdown викликається при закритті застосунку
func (a *App) Shutdown(ctx context.Context) {
	// Тут можна виконати очищення ресурсів
}

// DomReady викликається після завантаження DOM
func (a *App) DomReady(ctx context.Context) {
	// Викликається після завантаження фронтенду
}

// Greet повертає привітання (приклад методу для фронтенду)
func (a *App) Greet(name string) string {
	return "Привіт, " + name + "!"
}
