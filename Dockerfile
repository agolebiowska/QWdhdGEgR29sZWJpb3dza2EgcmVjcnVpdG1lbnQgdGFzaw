## First stage
FROM golang:alpine AS builder

ENV CGO_ENABLED 0

WORKDIR /go/src/app
COPY . .

RUN apk add --no-cache git
RUN go get ./...

RUN GOOS=linux go build -a -installsuffix cgo -o server ./cmd/main.go

## Second stage
FROM alpine

WORKDIR /app
COPY --from=builder /go/src/app/ /app/
RUN chmod +x ./server

CMD ["./server"]