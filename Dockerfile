FROM golang:1.21-alpine3.17 as dev-env

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

FROM dev-env as builder

RUN go build -o report-search ./cmd/report-search/
RUN go build -o fill-data ./cmd/fill-data/

FROM alpine:3.17

COPY ./entrypoint.sh /usr/bin/entrypoint.sh
RUN chmod +x /usr/bin/entrypoint.sh
COPY --from=builder /app/fill-data /usr/bin/fill-data
COPY --from=builder /app/report-search /usr/bin/report-search

CMD ["/usr/bin/entrypoint.sh"]