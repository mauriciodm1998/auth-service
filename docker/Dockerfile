FROM golang as runner

RUN mkdir app
COPY dist/auth-service app/
RUN chmod +x app
COPY ./internal/config/config.yaml /app
WORKDIR app

ENTRYPOINT [ "./auth-service", "--config-dir", "." ]