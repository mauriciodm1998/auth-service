run-app:
	go run cmd/auth/main.go

run-tests-in-docker:
	docker build -t auth-service . -f docker/Dockerfile-build-tests

run-tests:
	go test ./... -coverprofile cover.out -tags=test && go tool cover -html=cover.out

