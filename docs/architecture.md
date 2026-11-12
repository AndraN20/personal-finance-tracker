# Architecture note

Personal Finance Tracker, project 2. State at week 7.

## The shape

One Go binary, standard library only, shipped as a container that listens on 8080.

- `cmd/app` starts everything: handles `--version`, creates the store and the HTTP server, and shuts down cleanly on SIGTERM.
- `internal/api` owns the routes, the handlers, JSON decoding and encoding, and the error body `{"error","message"}`.
- `internal/store` holds all state in memory behind one mutex.
- `internal/version` reports the app name, version, commit and build time. The Dockerfile stamps the commit in at build time.

A request goes: `net/http` accepts it, `Server.ServeHTTP` hands it to the `ServeMux`, which matches on method and path (`POST /transactions`, `GET /reports/monthly`). The handler decodes and validates the body, calls one store method (which takes the lock), and writes the answer with `writeJSON` or `writeError`. Anything that matches no route falls through to a catch-all that answers a JSON 404.

For the project routes I plan two pieces between the handlers and the store: a CSV importer that turns each row into either a transaction or a rejection `{line, reason}`, and a categoriser that runs the rules over a description.

## The data

Today the store holds nothing. What it will hold:

- **Transactions:** id (integer, assigned from 1), date (a calendar date, no time or time zone), amount in cents as `int64`, description, category.
- **Rules:** an ordered list of `{match, category}`. `PUT /rules` replaces the whole list and the order is kept exactly as sent.
- **Budgets:** a map from category to a monthly limit in cents.
- The next id to hand out.

All of it lives in process memory. Nothing survives a restart: restarting the container and `POST /reset` end in the same place.

## Three decisions

**1. Money is an `int64` count of cents, not `float64` and not a decimal library.** A float cannot hold 0.10 exactly, so sums drift. `math/big` or a decimal package would be exact too, but the spec fixes amounts at two decimal places, so whole cents are enough and adding them is plain integer addition. I will parse `"-12.50"` myself and reject anything with more than two decimals instead of rounding it. This is decided but not written yet.

**2. The standard `ServeMux` instead of a router like chi or gin.** Since Go 1.22 it matches on method and path, which covers ten routes. A framework would be a dependency for very little. The cost: I had to add my own JSON 404, and a wrong method on a known path gets that 404 instead of a 405.

**3. State in memory, not SQLite.** It keeps the container a single `docker run` with no volume, and reset is trivial. SQLite would need either cgo or a pure-Go driver, plus a volume for the data to actually persist. The cost is that everything is lost on restart, which will bother me if I keep using the tool after the semester.

## What is not built yet

All six project routes are registered but answer 501. In detail:

- **The money type.** Parsing, formatting back to `"-0.30"`, and exact totals. Nothing is coded.
- **`POST /transactions`.** Validation of date and amount, and categorising when no category is given.
- **CSV import.** I plan to use `encoding/csv` for quoted fields. Still to handle: a missing header row, decimal commas, more than one date format, and a clear rule for which line number gets reported (I intend the physical line in the file, header included).
- **Importing the same export twice.** Not decided. My current idea is to fingerprint each row by date, amount, description and how many times that triple has appeared so far in the file. The same file imported again would then add nothing, but two real identical coffees on one day would both be kept.
- **Rules.** What `match` means exactly (I lean towards case-insensitive substring), and whether a new rule set re-categorises transactions that are already stored.
- **Budgets and `over_budget`** in the report.
- **`GET /reports/monthly`.** Also, I do not yet understand how the parameter M = 12 applies to a route that takes one month. I need to ask.
- **`GET /transactions`** with the `from`, `to`, `category` and `q` filters, newest first.
- **Tests** beyond the four common routes and the version contract.
- **Persistence** across restarts.
