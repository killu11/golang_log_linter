# 🔍 loglinter

**Статический анализатор для Go-логов, который проверяет стиль, язык и безопасность лог-сообщений.**

Проверки внутри линтера:
1. Лог-сообщения должны начинаться со строчной буквы
2. Лог-сообщения должны быть только на английском языке
3. Лог-сообщения не должны содержать спецсимволы или эмодзи
4. Лог-сообщения не должны содержать потенциально чувствительные данные

**Улучшения:**
Настроил линтер для внутренних аргументов классического zap логгера, чтобы линтер смотрел в Fields

**Поддерживаемые логгеры:**
● log/slog
● go.uber.org/zap

---

## 🧪 Запуск тестов

```bash
cd analyzer && go test -v
```

---

## 🔧 Сборка и использование

### Самостоятельный анализатор
```bash
cd cmd && go build -o loglinter
./log_checker /path/to/your/file.go
```

### Плагин для golangci-lint
```bash
golangci-lint custom
```
> ⚠️ Сборка проходит успешно, но при вызове возникает паника, хотя локально (без сборки в кастомный линтер) лог-линтер работает успешно

---

## 📄 Пример использования для foo.go

```go
package main

import (
    "context"
    "log/slog"
    "os"

    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

func main() {
    // Настраиваем slog
    slogLogger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }))

    // Настраиваем zap
    zapConfig := zap.NewProductionConfig()
    zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    zapLogger, _ := zapConfig.Build()
    defer zapLogger.Sync()

    // === ТЕСТОВЫЕ СЛУЧАИ ДЛЯ ЛИНТЕРА ===

    // 1. ПРОБЛЕМЫ С РЕГИСТРОМ
    slogLogger.Info("Starting server on port 8080") // ❌ С заглавной
    slogLogger.Info("shutting down gracefully")     // ✅ Со строчной

    zapLogger.Info("Failed to connect to database") // ❌ С заглавной
    zapLogger.Info("connection established")        // ✅ Со строчной

    // 2. РУССКИЙ ЯЗЫК
    slogLogger.Info("запуск сервера")            // ❌ Русский
    slogLogger.Info("ошибка подключения к базе") // ❌ Русский

    zapLogger.Info("успешный запрос")   // ❌ Русский
    zapLogger.Info("сервер остановлен") // ❌ Русский

    // 3. СПЕЦСИМВОЛЫ И ЭМОДЗИ
    slogLogger.Info("server started! 🚀")       // ❌ Эмодзи
    slogLogger.Info("connection failed!!!")    // ❌ Много !!!
    slogLogger.Info("warning... something...") // ❌ Многоточия

    zapLogger.Info("api call completed!!!???") // ❌ Смесь символов
    zapLogger.Info("✅✅✅ all good")             // ❌ Эмодзи

    // 4. ЧУВСТВИТЕЛЬНЫЕ ДАННЫЕ
    password := "secret123"
    apiKey := "sk-123456789"
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"
    creditCard := "4111-1111-1111-1111"

    slogLogger.Info("user password: " + password) // ❌ Пароль
    slogLogger.Info("api_key=" + apiKey)          // ❌ API ключ
    slogLogger.Debug("token: " + token)           // ❌ Токен

    zapLogger.Info("credit card",
        zap.String("card", creditCard)) // ❌ Карта
    zapLogger.Debug("auth header",
        zap.String("Authorization", "Bearer "+token)) // ❌ Токен

    // 5. РАЗНЫЙ УРОВЕНЬ ЛОГОВ
    slogLogger.Debug("debug message with password: " + password) // ❌ Пароль в debug
    slogLogger.Warn("⚠️ warning⚠️")                              // ❌ Эмодзи

    zapLogger.Error("fatal error!!! token=" + token) // ❌ Токен + символы

    // 6. КОНТЕКСТ И СТРУКТУРНЫЕ ЛОГИ
    ctx := context.Background()
    slogLogger.LogAttrs(ctx, slog.LevelInfo, "user logged in! 👤", // ❌ Эмодзи
        slog.String("user", "admin"),
        slog.String("password", password), // ❌ Пароль
    )

    zapLogger.With(
        zap.String("request_id", "12345"),
        zap.String("api_key", apiKey), // ❌ Ключ
    ).Info("api request completed successfully!!!") // ❌ !!!

    // 7. СМЕСЬ ВСЕГО
    weirdLogs := []string{
        "Starting server! 🚀 на порту 8080",                  // ❌ Эмодзи + русский
        "ошибка: connection failed!!! password=" + password, // ❌ Всё вместе
        "✅✅✅ API key " + apiKey + " is valid!!!",            // ❌ Кошмар
    }

    for _, msg := range weirdLogs {
        slogLogger.Info(msg)
    }

    // 8. НОРМАЛЬНЫЕ ЛОГИ
    slogLogger.Info("server started on port 8080")
    slogLogger.Info("request processed successfully")
    slogLogger.Debug("cache hit for key: users:123")

    zapLogger.Info("database connection established")
    zapLogger.Warn("high memory usage",
        zap.Float64("usage_percent", 85.5))
    zapLogger.Error("failed to process request",
        zap.String("error", "timeout"),
        zap.Int("retry_count", 3))
}
```

---

## ✅ Результат проверки

```bash
./loglinter ../foo.go
```

```
/home/egorsslv/Desktop/golang_log_linter/foo.go:27:18: first symbol of log msg shouldn't be in upper case
/home/egorsslv/Desktop/golang_log_linter/foo.go:30:17: first symbol of log msg shouldn't be in upper case
/home/egorsslv/Desktop/golang_log_linter/foo.go:34:18: log msg must be in english only
/home/egorsslv/Desktop/golang_log_linter/foo.go:35:18: log msg must be in english only
/home/egorsslv/Desktop/golang_log_linter/foo.go:37:17: log msg must be in english only
/home/egorsslv/Desktop/golang_log_linter/foo.go:38:17: log msg must be in english only
/home/egorsslv/Desktop/golang_log_linter/foo.go:41:18: log msg shouldn't have specifical symbols and emojis
/home/egorsslv/Desktop/golang_log_linter/foo.go:42:18: log msg shouldn't have specifical symbols and emojis
/home/egorsslv/Desktop/golang_log_linter/foo.go:43:18: log msg shouldn't have specifical symbols and emojis
/home/egorsslv/Desktop/golang_log_linter/foo.go:45:17: log msg shouldn't have specifical symbols and emojis
/home/egorsslv/Desktop/golang_log_linter/foo.go:46:17: log msg shouldn't have specifical symbols and emojis
/home/egorsslv/Desktop/golang_log_linter/foo.go:54:38: args shouldn't be had sensitive data: password
/home/egorsslv/Desktop/golang_log_linter/foo.go:55:31: args shouldn't be had sensitive data: apiKey
/home/egorsslv/Desktop/golang_log_linter/foo.go:64:53: args shouldn't be had sensitive data: password
/home/egorsslv/Desktop/golang_log_linter/foo.go:65:18: log msg must be in english only
/home/egorsslv/Desktop/golang_log_linter/foo.go:65:18: log msg shouldn't have specifical symbols and emojis
/home/egorsslv/Desktop/golang_log_linter/foo.go:67:18: log msg shouldn't have specifical symbols and emojis
/home/egorsslv/Desktop/golang_log_linter/foo.go:99:3: log msg shouldn't have specifical symbols and emojis
/home/egorsslv/Desktop/golang_log_linter/foo.go:102:3: log msg shouldn't have specifical symbols and emojis
```