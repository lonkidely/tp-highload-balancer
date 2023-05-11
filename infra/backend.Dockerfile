FROM golang:alpine AS build_image

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main.bin cmd/app/main.go

FROM alpine

WORKDIR /build

COPY --from=build_image /build/main.bin /build/main.bin
