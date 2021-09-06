FROM golang:latest


RUN go build -o main.go

CMD ["./bank/main"]