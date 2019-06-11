# 编译发布版本
# make 平台 v=版本号 p=发布模式
# 例：make linux v=v1.0.1 p=dev(默认)
# 例：make mac v=v1.0.2 p=release

# Get the git commit
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_IMPORT=carp.cn/whale/version

GOLDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT)$(GIT_DIRTY) -X $(GIT_IMPORT).Version=$(v) -X $(GIT_IMPORT).VersionPrerelease=$(p)

export GOLDFLAGS

all: linux windows mac

linux:
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -race -ldflags '$(GOLDFLAGS)'

linux_cross:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '$(GOLDFLAGS)'

windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags '$(GOLDFLAGS)'

mac:
	@CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -race -ldflags '$(GOLDFLAGS)'

govendor:
	@govendor add +external
