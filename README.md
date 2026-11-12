# Personal Finance Tracker

A web service for tracking what you earn and spend. Transactions come in from a
bank CSV export or are added by hand, get sorted into categories by rules you
write, are checked against monthly budgets, and come back out as a monthly
report with exact totals. It is meant for one person who wants to see where
their money went without trusting a spreadsheet to add up.

- **Student:** Nichifor Andra, group DMWD
- **Project:** 2, Personal Finance Tracker
- **Language:** Go

## What it does

The service keeps transactions, categorisation rules and budgets, and answers
JSON over HTTP on port 8080. Amounts are always JSON strings with two decimal
places (`"-12.50"`), negative for money going out, and are never stored as
floating-point numbers, so `"0.10"` plus `"0.20"` is exactly `"0.30"`.

It does not connect to a bank, it does not guess what a transaction is beyond
the rules you give it, and it does not keep data across restarts: everything
lives in memory and `POST /reset` clears it.

**Current state (week 7):** the four common routes, the `--version` flag, the
container build and CI work. The six project routes are registered but answer
`501 not_implemented`. See [docs/architecture.md](docs/architecture.md) for
the design and for what is still missing.

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

Errors use one shape, `{"error":"<code>","message":"<explanation>"}`, and an
unknown route answers `404 not_found`.

## How to run it

```
docker run --rm -p 8080:8080 ghcr.io/andran20/personal-finance-tracker:latest
```

To build and run without Docker:

```
go run ./cmd/app
go run ./cmd/app --version
```

To build the image locally:

```
docker build -t personal-finance-tracker .
docker run --rm -p 8080:8080 personal-finance-tracker
```

Tests:

```
go vet ./...
go test ./... -race
```

## What you should see

```
$ curl -i http://localhost:8080/health
HTTP/1.1 200 OK
Content-Type: application/json

{"status":"ok","uptime_seconds":3}
```

```
$ curl http://localhost:8080/version
{"app":"personal-finance-tracker","version":"0.0.0-dev","commit":"<commit>","built_at":"<time>"}
```

Opening <http://localhost:8080/> in a browser shows the identity page with the
student, group and project.

## Layout

```
cmd/app/             entry point: --version, HTTP server, graceful shutdown
internal/api/        routes, handlers, JSON and error responses
internal/store/      in-memory state (cleared by POST /reset)
internal/version/    app name, version, commit and build time
docs/                architecture note, and the report at the end
```
