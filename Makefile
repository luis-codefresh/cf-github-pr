TARGET      ?= darwin
ARCH        ?= amd64
SRC          = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
OUT          = build

export GO111MODULE=on

default: build

build:
	     @echo "== Build =="
	     @mkdir -p $(OUT)
			 CGO_ENABLED=0 GOOS=$(TARGET) GOARCH=$(ARCH) go build -o $(OUT)/check-pr -ldflags="-s -w" -v check/main.go
