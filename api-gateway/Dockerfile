FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/server ./api-gateway

CMD ["/app/server"]