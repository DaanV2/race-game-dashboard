set windows-shell := ["powershell.exe", "-c"]

default:
    just --list

documentation:
    go doc -all -u -http

dev:
	wails3 dev

test:
    go test -v ./... --cover -coverprofile=reports/coverage.out --covermode atomic --coverpkg=./...

show-coverage-report:
    go tool cover -html=reports/coverage.out

lint:
    go tool golangci-lint run -v --fix

format:
    go fmt ./...

fix:
    go fix ./...