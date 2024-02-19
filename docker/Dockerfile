FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk --no-cache add bash make git gcc gettext musl-dev

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY . .
RUN go build -o ./ cmd/app/main.go

FROM alpine AS runner

COPY --from=builder /app .

CMD ["./main"]