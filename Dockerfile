FROM golang:1.22.6-bullseye

LABEL maintainer-name="Hoplin"
LABEL maintainer-email="hoplin.dev@gmail.com"

EXPOSE 8080

COPY . /api
WORKDIR /api
RUN go mod download\
    && go build -o ota ./cmd/main.go

ENTRYPOINT ["./ota"]