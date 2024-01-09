FROM golang:1.21

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./...

EXPOSE 8080

CMD ["/app/inspektor"]