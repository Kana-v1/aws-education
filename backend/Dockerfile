FROM golang:1.17.4-alpine3.15 as builder

ADD . /app/
WORKDIR /app

RUN go build main.go

FROM alpine:3.15 as runtime
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 5000

CMD ["./main"]

