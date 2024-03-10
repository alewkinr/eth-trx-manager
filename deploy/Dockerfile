# builder image
FROM golang:1.21.0-alpine3.18 as builder

RUN mkdir /build

COPY . /build/
WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -a -o service cmd/main.go

# generate clean, final image for end users
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder /build/service /app

# executable
ENTRYPOINT [ "./service" ]

