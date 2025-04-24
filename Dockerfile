FROM golang:1.24.2-bookworm AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o MonitoringBot

EXPOSE 8241

CMD ["/build/MonitoringBot"]