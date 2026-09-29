FROM golang:1.26

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /photohost

EXPOSE 2137

CMD ["/photohost"]
