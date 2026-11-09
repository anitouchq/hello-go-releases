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

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)

## Скриншоты

Добавьте изображения в каталог `screenshots/`:

- `screenshots/application.png` — запуск приложения;
- `screenshots/github-actions-release.png` — успешная матричная сборка;
- `screenshots/github-release.png` — опубликованный GitHub Release с пятью файлами.

![Пример работы приложения](screenshots/application.png)
![Сборка GitHub Actions](screenshots/github-actions-release.png)
![Опубликованный GitHub Release](screenshots/github-release.png)

## Как создать Release-сборку

Сначала убедитесь, что нужные изменения уже находятся в `main`. Затем создайте новый тег версии и отправьте его на GitHub:

```bash
git checkout main
git pull origin main
git tag v1.0.0
git push origin v1.0.0
```

Тег должен соответствовать шаблону `v*`, например `v1.0.0`. Это запускает workflow: сначала CI-проверки, затем сборка пяти платформ, загрузка артефактов и публикация GitHub Release.

Для следующего релиза используйте новый номер, например `v1.0.1`. Уже опубликованный тег повторно не используйте.

Проверить процесс: [GitHub Actions](https://github.com/anitouchq/hello-go-releases/actions). Скачать бинарники: [GitHub Releases](https://github.com/anitouchq/hello-go-releases/releases).
