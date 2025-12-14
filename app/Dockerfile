# Build Stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./

RUN go build -o simpletimeservice

# Run Stage
FROM alpine:3.19

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/simpletimeservice .

USER appuser

EXPOSE 8080
CMD ["./simpletimeservice"]