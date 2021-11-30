FROM golang:1.17-alpine

RUN apk add --no-cache git

WORKDIR /app/bupcket

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/bupcket .

EXPOSE 8090

CMD ["./out/bupcket", "server"]