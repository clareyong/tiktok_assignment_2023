FROM golang:1.18
WORKDIR /app
COPY ./go.* .
RUN go mod download
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build go build -o main
EXPOSE 8080
CMD ["./main"]
