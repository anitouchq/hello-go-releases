# hello-go-releases

[![CI](https://github.com/anitouchq/hello-go-releases/actions/workflows/ci.yml/badge.svg)](https://github.com/anitouchq/hello-go-releases/actions)

Учебный CI/CD-проект на Go. При push в `main` и pull request выполняются проверки кода. При отправке тега вида `v1.0.0` приложение собирается для пяти платформ, а готовые бинарники прикладываются к GitHub Release.

## CI checks

- `gofmt` — форматирование;
- `go vet ./...` — статический анализ;
- `go test ./... -v` — модульные тесты;
- `go build` — контрольная сборка.

## Release matrix

| ОС | Архитектура | Файл |
|---|---|---|
| Linux | amd64 | `hello-go-linux-amd64` |
| Linux | arm64 | `hello-go-linux-arm64` |
| macOS | amd64 | `hello-go-darwin-amd64` |
| macOS | arm64 | `hello-go-darwin-arm64` |
| Windows | amd64 | `hello-go-windows-amd64.exe` |

Для сборки используется кросс-компиляция Go через `GOOS` и `GOARCH`. CGO отключён, поэтому бинарники не требуют системных библиотек C.

## Как создать Release

```bash
git tag v1.0.0
git push origin v1.0.0
```

После успешных CI-проверок matrix job собирает пять файлов, а отдельный job публикует их в [GitHub Releases](https://github.com/anitouchq/hello-go-releases/releases).

## Локальный запуск

```bash
go test ./... -v
go run .
```

## Структура проекта

| Файл или каталог | Назначение |
|---|---|
| `.github/workflows/ci.yml` | CI, матричная сборка и публикация Release |
| `greeting/greeting.go` | Функции приветствия и вспомогательная логика |
| `greeting/greeting_test.go` | Тесты пакета `greeting` |
| `main.go` | Точка входа CLI-приложения |
| `go.mod` | Имя Go-модуля и версия Go |
| `.gitignore` | Локальные файлы, исключаемые из Git |

## Результаты

- [GitHub Actions](https://github.com/anitouchq/hello-go-releases/actions)
- [Опубликованные релизы](https://github.com/anitouchq/hello-go-releases/releases)


## Скриншоты

### Пример запуска CLI-приложения

![Пример запуска CLI-приложения](screenshots/application.png)

### Успешная матричная сборка

![Успешная матричная сборка](screenshots/github-actions-release.png)

### GitHub Release с пятью бинарными файлами

![GitHub Release с пятью бинарными файлами](screenshots/github-release.png)

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
