FROM golang:1.23 AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o /dist/ ./cmd/gurps-bchest-be

FROM scratch AS bin
COPY --from=build /dist/ /
COPY ./import/ /import/
CMD [ "/gurps-bchest-be" ]
