GIT_COMMIT := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date +'%Y-%m-%d %H:%M:%S.%3N')
BUILD_TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "unknown")
VERSION_PKG := github.com/catplanet007/lit/version
LD_FLAGS := -X '$(VERSION_PKG).GitCommit=$(GIT_COMMIT)' -X '$(VERSION_PKG).BuildTime=$(BUILD_TIME)' -X '$(VERSION_PKG).BuildTag=$(BUILD_TAG)'

.PHONY: build

build:
	go build -ldflags="$(LD_FLAGS)" -o version ./main.go
