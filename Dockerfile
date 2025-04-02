FROM golang:1.24-alpine

WORKDIR /app

COPY . .

# CMD ["tail", "-f", "/dev/null"]

RUN go mod tidy

RUN go build -o main main.go

EXPOSE 8000

CMD ["./main"]