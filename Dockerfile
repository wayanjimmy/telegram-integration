# Build stage
FROM docker.io/golang:1.22-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o main bin/memogram/main.go

# Final stage
FROM docker.io/alpine:latest
WORKDIR /app
COPY --from=build /app/main /app/main
CMD ["/app/main"]
