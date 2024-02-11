FROM golang:1.21-alpine as builder
LABEL authors="tutunak"

COPY . /app
WORKDIR /app
RUN go build -o dowloader .

FROM alpine:3.19 as production
LABEL authors="tutunak"
COPY --from=builder /app/dowloader /app/dowloader
RUN addgroup -S dowloader && adduser -S dowloader -G dowloader && \
    chown -R dowloader:dowloader /app
USER dowloader
WORKDIR /app
CMD ["./dowloader"]
