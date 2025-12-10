# Приклади використання Pica3

## Приклад 1: Додавання нового методу для фронтенду

### Backend (Go)

Відкрийте `internal/app/app.go` та додайте метод:

```go
package app

import (
	"context"
	"fmt"
	"time"
)

// GetCurrentTime повертає поточний час
func (a *App) GetCurrentTime() string {
	return time.Now().Format("15:04:05")
}

// GetSystemInfo повертає інформацію про систему
func (a *App) GetSystemInfo() map[string]string {
	return map[string]string{
		"os":      "Windows 11",
		"version": "1.0.0",
		"status":  "running",
	}
}

// CalculateSum додає два числа
func (a *App) CalculateSum(a int, b int) int {
	return a + b
}
```

### Frontend (JavaScript)

Створіть файл `frontend/dist/app.js`:

```javascript
// Імпорт згенерованих Wails методів
// (після запуску wails dev вони будуть доступні)

// Приклад виклику методів
async function updateTime() {
    const time = await window.go.app.App.GetCurrentTime()
    document.getElementById('time').textContent = time
}

async function loadSystemInfo() {
    const info = await window.go.app.App.GetSystemInfo()
    console.log('System info:', info)
}

async function calculate() {
    const result = await window.go.app.App.CalculateSum(5, 10)
    console.log('5 + 10 =', result) // 15
}

// Оновлюємо час кожну секунду
setInterval(updateTime, 1000)
```

## Приклад 2: Зміна стилю glassmorphism

### Світла тема

```css
/* frontend/dist/style.css */
.glass-card {
    background: rgba(255, 255, 255, 0.3); /* Більше білого */
    backdrop-filter: blur(15px);
    border: 1px solid rgba(255, 255, 255, 0.4);
    color: #222;
}
```

### Темна тема

```css
.glass-card {
    background: rgba(0, 0, 0, 0.3); /* Більше чорного */
    backdrop-filter: blur(20px) brightness(0.9);
    border: 1px solid rgba(255, 255, 255, 0.1);
    color: #fff;
}
```

### Кольорова тема (синя)

```css
.glass-card {
    background: rgba(59, 130, 246, 0.2); /* Синій відтінок */
    backdrop-filter: blur(25px) saturate(200%);
    border: 1px solid rgba(59, 130, 246, 0.3);
    color: #1e40af;
}
```

## Приклад 3: Додавання анімації

```css
/* frontend/dist/style.css */
.glass-card {
    /* ... існуючі стилі ... */

    /* Анімація появи */
    animation: fadeIn 0.5s ease-in;

    /* Плавний перехід при наведенні */
    transition: all 0.3s ease;
}

.glass-card:hover {
    background: rgba(255, 255, 255, 0.25);
    transform: translateY(-5px);
    box-shadow: 0 12px 40px 0 rgba(0, 0, 0, 0.4);
}

@keyframes fadeIn {
    from {
        opacity: 0;
        transform: scale(0.9);
    }
    to {
        opacity: 1;
        transform: scale(1);
    }
}
```

## Приклад 4: Динамічна зміна прозорості

### Backend

```go
// internal/app/app.go
package app

import "context"

type App struct {
    ctx context.Context
    transparency float64
}

func NewApp() *App {
    return &App{
        transparency: 0.15, // 15% за замовчуванням
    }
}

// SetTransparency змінює рівень прозорості
func (a *App) SetTransparency(value float64) {
    if value >= 0 && value <= 1 {
        a.transparency = value
    }
}

// GetTransparency повертає поточний рівень прозорості
func (a *App) GetTransparency() float64 {
    return a.transparency
}
```

### Frontend

```javascript
// frontend/dist/app.js
async function changeTransparency(value) {
    await window.go.app.App.SetTransparency(value)

    // Оновлюємо CSS
    const card = document.querySelector('.glass-card')
    card.style.background = `rgba(255, 255, 255, ${value})`
}

// Додаємо слайдер в HTML
// <input type="range" min="0" max="1" step="0.05" value="0.15"
//        onchange="changeTransparency(this.value)">
```

## Приклад 5: Конфігурація для різних ОС

### Розширена конфігурація Windows

```go
// pkg/config/window.go
func GetWindowsOptions() *windows.Options {
    return &windows.Options{
        WebviewIsTransparent: true,
        WindowIsTranslucent:  true,

        // Різні варіанти BackdropType:

        // 1. Acrylic - класичний ефект розмиття
        BackdropType: windows.Acrylic,

        // 2. Mica - новий матеріал Windows 11
        // BackdropType: windows.Mica,

        // 3. Tabbed - для вкладок (Windows 11)
        // BackdropType: windows.Tabbed,

        // 4. Без ефекту (повна прозорість)
        // BackdropType: windows.None,

        // Додаткові налаштування
        DisableWindowIcon: false,
        DisableFramelessWindowDecorations: false,
        WebviewBrowserPath: "", // Шлях до Edge WebView2
    }
}
```

## Приклад 6: Контекстне меню

### Backend

```go
// internal/app/app.go
import "github.com/wailsapp/wails/v2/pkg/menu"

func (a *App) CreateMenu() *menu.Menu {
    appMenu := menu.NewMenu()

    fileMenu := appMenu.AddSubmenu("Файл")
    fileMenu.AddText("Відкрити", nil, func(_ *menu.CallbackData) {
        // Логіка відкриття
    })
    fileMenu.AddSeparator()
    fileMenu.AddText("Вихід", nil, func(_ *menu.CallbackData) {
        a.ctx.Done()
    })

    return appMenu
}
```

## Приклад 7: Робота з файлами

### Backend

```go
// internal/app/app.go
import (
    "io/ioutil"
    "github.com/wailsapp/wails/v2/pkg/runtime"
)

// SelectFile відкриває діалог вибору файлу
func (a *App) SelectFile() (string, error) {
    file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
        Title: "Виберіть файл",
        Filters: []runtime.FileFilter{
            {DisplayName: "Текстові файли", Pattern: "*.txt"},
            {DisplayName: "Всі файли", Pattern: "*.*"},
        },
    })
    return file, err
}

// ReadFile читає файл
func (a *App) ReadFile(path string) (string, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
```

## Приклад 8: Тестування

### Unit тест для app

```go
// internal/app/app_test.go
package app

import "testing"

func TestGreet(t *testing.T) {
    app := NewApp()

    result := app.Greet("Тест")
    expected := "Привіт, Тест!"

    if result != expected {
        t.Errorf("Очікувалось %s, отримано %s", expected, result)
    }
}

func TestCalculateSum(t *testing.T) {
    app := NewApp()

    tests := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {10, 20, 30},
        {-5, 5, 0},
    }

    for _, tt := range tests {
        result := app.CalculateSum(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("CalculateSum(%d, %d) = %d; очікувалось %d",
                tt.a, tt.b, result, tt.expected)
        }
    }
}
```

Запуск тестів:
```bash
make test
```

## Приклад 9: Логування

```go
// internal/app/app.go
import (
    "log"
    "github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) LogInfo(message string) {
    log.Printf("[INFO] %s", message)
    runtime.LogInfo(a.ctx, message)
}

func (a *App) LogError(message string) {
    log.Printf("[ERROR] %s", message)
    runtime.LogError(a.ctx, message)
}
```

## Приклад 10: Повідомлення та діалоги

```go
// internal/app/app.go
import "github.com/wailsapp/wails/v2/pkg/runtime"

// ShowMessage показує інформаційне повідомлення
func (a *App) ShowMessage(title, message string) {
    runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
        Type:    runtime.InfoDialog,
        Title:   title,
        Message: message,
    })
}

// ShowConfirm показує діалог підтвердження
func (a *App) ShowConfirm(title, message string) (bool, error) {
    result, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
        Type:          runtime.QuestionDialog,
        Title:         title,
        Message:       message,
        Buttons:       []string{"Так", "Ні"},
        DefaultButton: "Так",
    })
    return result == "Так", err
}
```

## Додаткові ресурси

- Більше прикладів: [Wails Examples](https://wails.io/docs/examples/showcase)
- API Reference: [Wails API](https://wails.io/docs/reference/runtime/intro)
- Go документація: [golang.org](https://golang.org/doc/)
