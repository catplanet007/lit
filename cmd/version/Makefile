GIT_COMMIT := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date +'%Y-%m-%d %H:%M:%S.%3N')
BUILD_TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "unknown")
PKG_PATH := github.com/catplanet007/lit

.PHONY: build

build:
	go build -ldflags="-X '$(PKG_PATH)/version.GitCommit=$(GIT_COMMIT)'-X '$(PKG_PATH)/version.BuildTime=$(BUILD_TIME)' -X '$(PKG_PATH)/version.BuildTag=$(BUILD_TAG)'" -o version ./main.go
