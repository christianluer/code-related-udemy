N ?=

.PHONY: learn check watch verify lint

learn:
	go run ./cmd/learn

check:
	go run ./cmd/learn check $(N)

watch:
	go run ./cmd/learn watch $(N)

verify:
	go test -tags solution -race -count=1 ./lessons/...

lint:
	@test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)
	go vet ./...
	go vet -tags solution ./...
