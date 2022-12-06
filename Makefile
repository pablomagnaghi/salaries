## initializer
initializer-dataset:
	@go run initializer/main.go

## run
run-locally:
	@go run cmd/main.go

## tests
test-locally:
	@go test -v ./...

test:
	docker run salaries-app go test -v ./...

## docker compose
up:
	docker-compose up
down:
	docker-compose down --remove-orphans
clean:
	docker container prune -f
	docker image prune -f



