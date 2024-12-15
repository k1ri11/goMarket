FROM golang:latest

COPY . .

RUN go build -C cmd/server -o go-market
CMD ["cmd/server/go-market"]