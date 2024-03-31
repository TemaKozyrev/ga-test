FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main ./cmd/app

FROM alpine

RUN apk --no-cache add curl

WORKDIR /build

COPY --from=builder /build/main /build/main

EXPOSE 80

HEALTHCHECK --interval=5s --timeout=2s --retries=5 \
      CMD curl -f http://localhost:80 || exit 1

ENTRYPOINT ["./main"]