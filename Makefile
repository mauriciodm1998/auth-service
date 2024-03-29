test-build-bake:
	docker build -t docker.io/mauricio1998/auth-service . -f build/Dockerfile

run-tests:
	go test ./... -coverprofile cover.out -tags=test && go tool cover -html=cover.out

docker-push:
	docker push docker.io/mauricio1998/auth-service

boilerplate: test-build-bake docker-push