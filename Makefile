mod:
	go list -m --versions

test:
	go test -v

test.cover:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -html=cover.out

critic:
	gocritic check -enableAll -disable=unnamedResult,unlabelStmt,hugeParam,singleCaseSwitch,builtinShadow,typeAssertChain ./...

security:
	gosec -exclude-dir=examples -exclude=G103,G401,G501 ./...

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run ./...

all: critic security vulncheck lint test