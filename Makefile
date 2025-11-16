build :
	@go build -o bin/GO-APIs cmd/main.go

test :
	@go test -v ./...

run : build
	@./bin/GO-APIs