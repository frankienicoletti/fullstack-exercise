FROM golang:1.12-alpine as builder
RUN apk add git

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=builder /app/server /app/
EXPOSE 4000
ENTRYPOINT ["/app/server"]