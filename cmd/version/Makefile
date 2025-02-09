GIT_COMMIT := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date +'%Y-%m-%d %H:%M:%S.%3N')
BUILD_TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "unknown")
VERSION_PKG_PATH := github.com/catplanet007/lit
LD_FLAGS := -X '$(VERSION_PKG_PATH)/version.GitCommit=$(GIT_COMMIT)' -X '$(VERSION_PKG_PATH)/version.BuildTime=$(BUILD_TIME)' -X '$(VERSION_PKG_PATH)/version.BuildTag=$(BUILD_TAG)'

.PHONY: build

build:
	go build -ldflags="$(LD_FLAGS)" -o version ./main.go
