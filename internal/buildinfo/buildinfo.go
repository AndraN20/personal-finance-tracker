// Package buildinfo holds values stamped into the binary at build time.
//
// They are set with -ldflags, for example:
//
//	go build -ldflags "-X github.com/AndraN20/personal-finance-tracker/internal/buildinfo.Version=1.0.0"
package buildinfo

const App = "personal-finance-tracker"

var (
	Version = "dev"
	Commit  = "unknown"
	BuiltAt = "unknown"
)
