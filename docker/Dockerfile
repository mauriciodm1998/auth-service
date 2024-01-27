## mauricio1998/auth-service

FROM golang:latest as builder

RUN mkdir app
COPY . app
WORKDIR app
RUN go test ./... -coverprofile cover.out -tags=test && go tool cover -html=cover.out
RUN CGO_ENABLED=0 go build -v -o dist/auth-service cmd/auth/main.go

FROM alpine:latest as baker

RUN mkdir -p app/bin
COPY --from=builder /go/app/dist/auth-service /app/bin
COPY ./internal/config/config.yaml /app/bin
RUN chmod +x app
WORKDIR app/bin

ENTRYPOINT [ "./auth-service", "--config-dir", "." ]