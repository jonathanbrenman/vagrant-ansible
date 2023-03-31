# Build stage
FROM golang:1.18.2-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy

WORKDIR /app/api
# Build the Go binary
RUN CGO_ENABLED=0 go build -o application .

# Final stage
FROM alpine:3.14
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/api/application .
EXPOSE 8080
CMD ["./application"]