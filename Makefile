init-env:
	go mod tidy
	go get -u github.com/swaggo/swag/cmd/swag
run:
	go run main.go
swag:
	swag init
rebuild-swag-run: swag run

test:
	go test -v -cover ./...