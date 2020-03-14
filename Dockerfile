FROM golang:alpine3.11

COPY . /app
WORKDIR /app

RUN go get ./...
ENV GOOS linux
RUN go build -o kv-app cmd/main.go
ENTRYPOINT [ "./kv-app" ]