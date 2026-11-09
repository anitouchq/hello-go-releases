# 🐹 hello-go-releases — Multi-platform binaries → GitHub Releases

Учебный CI/CD-проект для Go. На каждом push в `main` и pull request workflow выполняет проверки кода. При push тега вида `v1.0.0` приложение компилируется для пяти платформ, после чего бинарники прикладываются к GitHub Release.

## CI checks

- `gofmt` — форматирование;
- `go vet ./...` — статический анализ;
- `go test ./... -v` — тесты;
- `go build` — контрольная сборка.

## Release matrix

| ОС | Архитектура | Файл |
|---|---|---|
| Linux | amd64 | `hello-go-linux-amd64` |
| Linux | arm64 | `hello-go-linux-arm64` |
| macOS | amd64 | `hello-go-darwin-amd64` |
| macOS | arm64 | `hello-go-darwin-arm64` |
| Windows | amd64 | `hello-go-windows-amd64.exe` |

Для сборки используется кросс-компиляция Go через `GOOS` и `GOARCH`. CGO отключён, чтобы получить самостоятельные бинарные файлы без системных библиотек C.

## Как создать Release

Создайте тег и отправьте его в GitHub:

```bash
git tag v1.0.0
git push origin v1.0.0
```

GitHub Actions сначала выполняет CI. Если проверки прошли успешно, matrix job собирает пять бинарников, а отдельный job публикует их в GitHub Releases.

## Локальный запуск

```bash
go test ./... -v
go run .
```

## Permissions

Для обычных CI-проверок используется `contents: read`. Job публикации получает `contents: write` для создания Release с помощью встроенного `GITHUB_TOKEN`.

[Главный каталог CI/CD проектов](https://github.com/anitouchq/CI-CD-Pipelines)
