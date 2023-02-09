ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine as backend_builder

WORKDIR /app

COPY Makefile go.mod go.sum main.go ./
COPY pkg ./pkg

RUN apk add --update make
RUN make build

FROM scratch

COPY --from=backend_builder /app/start .

WORKDIR /app

CMD ["./app"]

EXPOSE 8080