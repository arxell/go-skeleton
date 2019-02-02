run: build 
	./go-skeleton

build: 
	go build -o go-skeleton src/main.go

dep:
	cd src && dep ensure

test: format
	cd src && go test ./...

format:
	cd src && gofmt -w -s .
	cd src && goimports -w .

unit: format
	cd src && ginkgo -r -cover -coverprofile=coverage.out .

coverage: unit
	cd src && go tool cover -func=helpers/coverage.out

docker-build:
	docker-compose build

docker-run:
	docker-compose run go-skeleton
