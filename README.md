# personal-finance-tracker

A service for tracking what you earn and spend. Transactions are imported from a
bank CSV export or added by hand, categorised by user-defined rules, checked
against monthly budgets, and reported by month and category with exact totals.

Written in Go (standard library only). Listens on port 8080.

## Build and run

Go does not need to be installed locally; everything builds inside Docker.

```powershell
docker build --build-arg VERSION=0.1.0 --build-arg COMMIT=$(git rev-parse --short HEAD) -t personal-finance-tracker .
docker run --rm -p 8080:8080 personal-finance-tracker
```

`VERSION` and `COMMIT` are optional (they default to `dev` and `unknown`).
`BUILT_AT` defaults to the time of the build. All three are reported by `GET /version`.

## Test

```powershell
docker run --rm -v "${PWD}:/src" -w /src golang:1.26-alpine go test ./...
```

## Routes

| Route                  | Method | Status          |
|------------------------|--------|-----------------|
| `/health`              | GET    | done            |
| `/version`             | GET    | done            |
| `/`                    | GET    | done            |
| `/reset`               | POST   | done            |
| `/transactions`        | POST   | not implemented |
| `/transactions/import` | POST   | not implemented |
| `/transactions`        | GET    | not implemented |
| `/rules`               | PUT    | not implemented |
| `/budgets`             | PUT    | not implemented |
| `/reports/monthly`     | GET    | not implemented |

Routes that are not implemented yet answer `501` with the standard error body.
Unknown routes answer `404 {"error":"not_found","message":"..."}`.

## Layout

```
cmd/server/          entry point: config, HTTP server, graceful shutdown
internal/api/        routes, handlers, JSON and error responses
internal/store/      in-memory state (cleared by POST /reset)
internal/buildinfo/  version, commit and build time, stamped in with -ldflags
```
