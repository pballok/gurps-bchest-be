FROM golang:alpine AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o /dist/ ./cmd/gurps-bchest-be

FROM alpine AS bin
COPY --from=build /dist/ /
COPY --from=build /src/internal/database/mysql/migrations /migrations
CMD [ "/gurps-bchest-be" ]
