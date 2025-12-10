# Архітектура проекту Pica3

## Огляд

Проект Pica3 організований за стандартними принципами Go проектів з урахуванням специфіки Wails фреймворку.

## Структура директорій

### Коренева директорія
- **main.go** - точка входу в застосунок
  - Ініціалізує екземпляр застосунку
  - Завантажує конфігурацію
  - Запускає Wails runtime
  - Використовує `embed.FS` для вбудовування фронтенду

### internal/
Містить приватну логіку застосунку, яка НЕ може бути імпортована зовнішніми проектами.

#### internal/app/
- **app.go** - основна структура застосунку
  - `App` struct - зберігає стан застосунку
  - `Startup()` - життєвий цикл: запуск
  - `Shutdown()` - життєвий цикл: завершення
  - `DomReady()` - життєвий цикл: готовність DOM
  - Методи, які можна викликати з фронтенду

### pkg/
Містить публічні пакети, які МОЖУТЬ бути імпортовані іншими проектами.

#### pkg/config/
- **window.go** - конфігурація вікна
  - `WindowConfig` - структура налаштувань вікна
  - `GetMacOptions()` - налаштування для macOS
  - `GetWindowsOptions()` - налаштування для Windows
  - `GetBackgroundColor()` - прозорий колір фону

### frontend/
Містить веб-інтерфейс застосунку.

#### frontend/dist/
- **index.html** - HTML розмітка
- **style.css** - CSS з ефектом glassmorphism
  - Прозорий фон
  - Backdrop filter для ефекту скла

## Потік даних

```
main.go
   ↓
internal/app.NewApp()
   ↓
pkg/config (завантаження налаштувань)
   ↓
wails.Run()
   ↓
frontend/dist (відображення UI)
```

## Життєвий цикл застосунку

1. **main()** створює екземпляр застосунку
2. **Startup()** викликається при запуску
3. **DomReady()** викликається після завантаження UI
4. **Shutdown()** викликається при закритті

## Комунікація фронтенд-бекенд

Методи з `internal/app/app.go` автоматично експортуються на фронтенд через Wails:

```go
// Бекенд (Go)
func (a *App) Greet(name string) string {
    return "Привіт, " + name + "!"
}
```

```javascript
// Фронтенд (JavaScript)
import { Greet } from '../wailsjs/go/app/App'

const message = await Greet("Світ")
```

## Прозорість вікна

### Backend (Go)
1. `BackgroundColour` з альфа-каналом 0 (повністю прозорий)
2. `WebviewIsTransparent: true` - прозорий webview
3. `WindowIsTranslucent: true` - напівпрозоре вікно
4. `BackdropType` (опціонально) - системні ефекти розмиття

### Frontend (CSS)
1. `html, body { background: transparent !important; }`
2. `backdrop-filter: blur(20px)` - розмиття фону
3. `rgba()` кольори з альфа-каналом для прозорості

## Розширення проекту

### Додавання нового функціоналу

1. **Бізнес-логіка**: додайте до `internal/app/`
2. **Конфігурація**: додайте до `pkg/config/`
3. **Утиліти**: створіть новий пакет у `pkg/`

### Приклад: додавання нового сервісу

```go
// internal/services/window_service.go
package services

type WindowService struct {
    config *config.WindowConfig
}

func NewWindowService(cfg *config.WindowConfig) *WindowService {
    return &WindowService{config: cfg}
}

func (s *WindowService) Minimize() {
    // логіка мінімізації
}
```

```go
// main.go
import "pica3/internal/services"

func main() {
    windowService := services.NewWindowService(config.DefaultWindowConfig())
    // використання сервісу
}
```

## Принципи організації коду

1. **Separation of Concerns** - розділення відповідальності
   - `main.go` - тільки запуск
   - `internal/` - бізнес-логіка
   - `pkg/` - конфігурація та утиліти

2. **Dependency Direction** - напрямок залежностей
   - `main.go` → залежить від `internal/` та `pkg/`
   - `internal/` → може використовувати `pkg/`
   - `pkg/` → не залежить від `internal/`

3. **Testability** - тестування
   - Кожен пакет має свої тести
   - Мінімальні залежності між пакетами
   - Легко мокати інтерфейси

## Best Practices

1. Тримайте `main.go` мінімальним
2. Винесіть всю логіку в `internal/`
3. Конфігурацію тримайте в `pkg/config/`
4. Використовуйте інтерфейси для залежностей
5. Документуйте публічні API
