# githerd make file
project_name = githerd
project_name_canonical = githerd
tag_version=$(shell git tag --sort=-version:refname |grep 'v'| head -n 1)
tag_commit=$(shell git rev-parse --short -n 1 `git tag --sort=-version:refname |grep 'v'| head -n 1`)
commit=$(shell git rev-parse --short HEAD)

build:

	@echo "Commit: ${commit}"
	@echo "TagCommit: ${tag_commit}"
	@echo "TagVersion: ${tag_version}"
	@echo "Building ${project_name}"
	go build \
	-ldflags "-X main.Version=${tag_version} -X main.Commit=${commit} -X main.TagCommit=${tag_commit}" -o "./bin/githerd" "./cmd/githerd/main.go"
run: build
	./bin/githerd
version:
	./bin/githerd version

test:
	go test -v -coverprofile=./.coverage/coverage.out ./...
	go tool cover -html=./.coverage/coverage.out -o ./.coverage/coverage.html


