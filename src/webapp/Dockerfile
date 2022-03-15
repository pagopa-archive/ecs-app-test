FROM golang:alpine AS builder
WORKDIR /go/src/hello
RUN apk add --no-cache gcc libc-dev
ADD app.go /go/src/hello/app.go
ADD go.mod /go/src/hello/go.mod
ADD go.sum /go/src/hello/go.sum
RUN GOOS=linux GOARCH=amd64 go build -tags=netgo app.go

FROM alpine
COPY --from=builder /go/src/hello/app /app
CMD ["/app"]