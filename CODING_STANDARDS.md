# Coding standards

- The required `format` check runs gofmt from the Go version in `go.mod`. gofmt has no line-length setting, so Go code has no column limit. The required `lint` check runs golangci-lint at the exact version pinned in CI. Every dependency is pinned to an exact version.
- Deliver the issue's scope with the smallest change that meets it. Prefer deletion to new code.
- No new abstraction without two real call sites, and no caps, retries or fallbacks that no requirement names.
