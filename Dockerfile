# Build stage
FROM golang:1.22-2-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o main bin/memogram/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/main /app/main
CMD ["/app/main"]
