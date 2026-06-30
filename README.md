# yze-go-noinit

A [`yze`](https://github.com/gomatic/yze) analyzer (category `patterns`) that forbids a package-level `func init`, per the gomatic Go standard that favors explicit construction and dependency injection over implicit package initialization.

- **Rule:** `yze/noinit`
- **Library:** exports `Analyzer` (a standard `go/analysis` analyzer) and `Registration` for the [`yze`](https://github.com/gomatic/yze) aggregator and [`stickler`](https://github.com/gomatic/stickler) runner.
- **Binary:** `cmd/yze-go-noinit` runs it standalone (`text`/`-json`, and as a `go vet -vettool`).

Built on the [`go-yze`](https://github.com/gomatic/go-yze) framework.
