

dep-reset:
	rm -rf ./Gopkg.lock ./Gopkg.toml ./vendor
	dep init

build-dev:
	go build -o ff ./cmd/search/main.go

build: dep-reset
	rm ff
	go build -o ff ./cmd/search/main.go