# env
init-env:
	go mod tidy
	go get -u github.com/swaggo/swag/cmd/swag
clean:
	rm profile
vendor:
	go mod vendor

# program
build:
	go build -o main .
run:
	./main
run-dev:
	go run main.go
swag:
	swag init
rebuild-swag-run-dev: swag run-dev

test:
	go test -v -cover ./...
