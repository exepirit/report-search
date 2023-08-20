FROM golang:1.21-alpine3.17 as go-dev-env

WORKDIR /build

COPY go.* ./
RUN go mod download

COPY . ./

FROM go-dev-env as go-builder

RUN go build -o report-search ./cmd/report-search/
RUN go build -o fill-data ./cmd/fill-data/

FROM node:18.8.0 as js-builder

WORKDIR /build

COPY ./web/package.json ./web/package-lock.json ./
RUN npm install

COPY ./web ./
RUN npm run build

FROM alpine:3.17

WORKDIR /app

COPY ./entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh
COPY --from=go-builder /build/fill-data /app/fill-data
COPY --from=go-builder /build/report-search /app/report-search
COPY --from=js-builder /build/build /app/web/build

CMD ["/app/entrypoint.sh"]