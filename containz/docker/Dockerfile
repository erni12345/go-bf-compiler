FROM golang:1.22-bullseye

WORKDIR /app

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build main.go

CMD ["./main"]
