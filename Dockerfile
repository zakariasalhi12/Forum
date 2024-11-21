FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN apk update && apk add --no-cache bash gcc musl-dev

RUN go mod download

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]