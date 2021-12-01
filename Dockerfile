# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o /dist/ ./cmd/gurps-bchest-be

FROM scratch AS bin
COPY --from=build /dist/ /
