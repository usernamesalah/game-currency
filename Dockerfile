FROM golang:1.14-alpine as go-builder
WORKDIR /src
COPY . .
RUN go build -o bin/game-currency .

FROM alpine:3.12
COPY --from=go-builder /src/bin/game-currency .
COPY migrations migrations

CMD ["./game-currency"]
