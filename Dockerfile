FROM golang:1.21 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o . ./...

FROM busybox:1.36

WORKDIR /app

COPY --from=build /app/inspektor ./

EXPOSE 8080

CMD ["./inspektor"]