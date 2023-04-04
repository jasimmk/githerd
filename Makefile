# githerd make file
project_name = githerd
project_name_canonical = githerd

build:
	go build -ldflags "-X main.Version=`git tag --sort=-version:refname | head -n 1`" -o "./bin/githerd" "./cmd/githerd/main.go"
run: build
	./bin/githerd
version:
	./bin/githerd version

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o .coverage/coverage.html