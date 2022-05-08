FROM golang:1.18.1-alpine3.15
RUN apk update
RUN apk add git
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /app/api

CMD ["/app/api"]