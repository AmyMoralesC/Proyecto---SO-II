# Etapa 1: compilar el binario
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o proxy .

# Etapa 2: imagen mínima
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/proxy .

RUN mkdir -p logs

EXPOSE 8080

CMD ["./proxy"]
