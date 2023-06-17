# -x for tracing on output, -e for exit on error
set -xe

go mod tidy

GOOS=linux GOARCH=amd64 go build  -o ./build/main ./cmd/main.go