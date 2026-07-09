FROM golang:1.25 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cut-it

FROM scratch
COPY --from=build /app/cut-it .
EXPOSE 8080
ENTRYPOINT ["./cut-it"]