# Contributing to ebiten_extended

Thank you for your interest in contributing.

## Development setup

1. Clone the repository and ensure you have Go 1.21+ and [Ebiten v2](https://github.com/hajimehoshi/ebiten) installed.
2. Run tests from the module root:
   ```bash
   go test ./...
   ```

## Submitting changes

1. Make your changes in a branch.
2. Ensure `go build ./...` and `go test ./...` pass.
3. Follow standard Go style (`gofmt`, `go vet`). Doc comments for exported symbols help godoc and users.
4. Open a pull request with a clear description of the change.

## Documentation

Exported types, functions, and methods should have doc comments so they appear correctly on [pkg.go.dev](https://pkg.go.dev). The first sentence is used as the summary; start with the name of the symbol (e.g. "NewSprite creates a new SpriteNode...").
