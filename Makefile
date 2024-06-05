run:
	go mod tidy
	go run main.go

build:
	go mod tidy
	go build -o urlShortner

test:
	go test ./... -v -coverprofile=coverage.out 

