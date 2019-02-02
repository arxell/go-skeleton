run: 
	go run src/main.go

build: 
	go build -o hello src/main.go

dep:
	cd src && dep ensure

tests:
	cd src && go test ./...
