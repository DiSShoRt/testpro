FROM golang:latest

RUN mkdir /zadanie
RUN chmod +x .
WORKDIR /zadanie
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY ./ ./
RUN chmod +x .

RUN go build -o main .
CMD ["./main"]