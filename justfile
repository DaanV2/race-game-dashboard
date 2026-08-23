set windows-shell := ["powershell.exe", "-c"]

default:
    just --list

documentation:
    go doc -all -u -http

dev:
	wails3 dev

build:
	go build ./...

[group('test')]
test:
    go test -v ./... --cover -coverprofile=reports/coverage.out --covermode atomic --coverpkg=./...

[group('test')]
show-coverage-report:
    go tool cover -html=reports/coverage.out

[group('checks')]
lint:
    go tool golangci-lint run -v --fix

[group('checks')]
format:
    go fmt ./...

[group('checks')]
fmt:
    go fmt ./...

[group('checks')]
generate:
    go generate -v ./...

[group('checks')]
fix:
    go fix ./...

[group('checks')]
checks: generate build fix fmt test lint