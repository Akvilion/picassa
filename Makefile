.PHONY: dev build clean test run help

# Змінні
BINARY_NAME=pica3
BUILD_DIR=build/bin

# Кольори для виводу
GREEN=\033[0;32m
NC=\033[0m # No Color

help: ## Показати довідку
	@echo "Доступні команди:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

dev: ## Запустити у режимі розробки з hot reload
	@echo "$(GREEN)Запуск у режимі розробки...$(NC)"
	wails dev

build: ## Зібрати production версію
	@echo "$(GREEN)Збірка production версії...$(NC)"
	wails build

build-debug: ## Зібрати debug версію
	@echo "$(GREEN)Збірка debug версії...$(NC)"
	wails build -debug

run: build ## Зібрати та запустити
	@echo "$(GREEN)Запуск застосунку...$(NC)"
	./$(BUILD_DIR)/$(BINARY_NAME)

clean: ## Очистити build директорію
	@echo "$(GREEN)Очищення...$(NC)"
	rm -rf $(BUILD_DIR)
	rm -rf build/

test: ## Запустити тести
	@echo "$(GREEN)Запуск тестів...$(NC)"
	go test -v ./...

test-coverage: ## Запустити тести з покриттям
	@echo "$(GREEN)Запуск тестів з покриттям...$(NC)"
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

fmt: ## Форматувати код
	@echo "$(GREEN)Форматування коду...$(NC)"
	go fmt ./...
	gofmt -s -w .

lint: ## Перевірити код лінтером
	@echo "$(GREEN)Лінтинг коду...$(NC)"
	golangci-lint run

install-deps: ## Встановити залежності
	@echo "$(GREEN)Встановлення залежностей...$(NC)"
	go mod download
	go mod tidy

generate: ## Згенерувати Wails біндінги
	@echo "$(GREEN)Генерація біндінгів...$(NC)"
	wails generate module

update-wails: ## Оновити Wails
	@echo "$(GREEN)Оновлення Wails...$(NC)"
	go install github.com/wailsapp/wails/v2/cmd/wails@latest

check: fmt lint test ## Запустити всі перевірки (fmt, lint, test)

all: clean build ## Очистити та зібрати
