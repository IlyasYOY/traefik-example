FROM golang:1.13.8-alpine as builder

WORKDIR /opt/build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine

COPY --from=builder /opt/build/app /opt/app

EXPOSE 80
ENTRYPOINT ["/opt/app"]