install: go.sum
	go build -o $(HOME)/go/bin/executor -ldflags="-X 'github.com/celer-network/im-executor/cmd.Version=$(VERSION)'" ./cmd/main

build:
	DOCKER_BUILDKIT=1 docker build -f Dockerfile --build-arg GH_TOKEN=$(GH_TOKEN) --tag celer-network/executor .