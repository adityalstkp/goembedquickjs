build-embed:
	(cd web/ && pnpm rollup:embed)

run-go:
	go run cmd/http/*.go
