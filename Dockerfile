FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /APP

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp cmd/app/main.go


CMD ["./myapp"]
