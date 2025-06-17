mod:
	go list -m --versions

critic:
	gocritic check -enableAll -disable=unnamedResult,unlabelStmt,hugeParam,singleCaseSwitch,builtinShadow,typeAssertChain ./...

security:
	gosec -exclude-dir=mysql,psql -exclude=G103,G115,G401,G501,G404 ./...

vulncheck:
	govulncheck ./...

lint:
	golangci-lint run ./...

test:
	go test -v -timeout 30s ./...

test.cover:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -html=cover.out

all: critic security vulncheck lint test