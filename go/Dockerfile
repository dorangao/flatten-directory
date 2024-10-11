FROM golang:1.16

WORKDIR /app

COPY flatten_directory.go ./

RUN go build -o flatten_directory flatten_directory.go

ENTRYPOINT ["./flatten_directory"]
