FROM golang:1.27-alpine as builder

RUN apk update
RUN apk add git openssh tzdata build-base python3 net-tools

WORKDIR /app

COPY . .

RUN go install github.com/buu700/gin@latest
RUN go mod tidy

RUN go build -o user-service .

FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    apk --no-cache add curl && \
    mkdir /app

WORKDIR /app

EXPOSE 8001

COPY --from=builder /app/user-service /app/user-service
COPY --from=builder /app/migrations /app/migrations

ENTRYPOINT ["/app/user-service"]
CMD ["serve"]