

reset: reset-dep
	rm ff

reset-dep:
	rm -rf ./Gopkg.lock ./Gopkg.toml ./vendor
	dep init

build-dev:
	go build -o ff ./cmd/search/main.go

build: reset
	go build -o ff ./cmd/search/main.go