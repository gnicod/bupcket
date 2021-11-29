FROM golang


WORKDIR /app

COPY go.mod ./
COPY go.sum ./
WORKDIR /app
RUN go mod download
CMD [ "go", "run", "main.go" ]

EXPOSE 8090