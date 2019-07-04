#build stage
FROM golang:1.12 AS builder

ENV GO111MODULE=on

WORKDIR /app
COPY app/go.mod .
COPY app/go.sum .
RUN go mod download
COPY app .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

#final stage
FROM alpine:latest
ENV TZ=Europe/Moscow
RUN apk update && apk add tzdata && cp -r -f /usr/share/zoneinfo/$TZ /etc/localtime
WORKDIR /app
COPY config.json .
COPY --from=builder /app/fantlab .
ENV CONFIG_FILE=config.json
COPY docker-entrypoint.sh .
RUN chmod +x ./docker-entrypoint.sh
ENTRYPOINT [ "./docker-entrypoint.sh", "./fantlab" ]
LABEL Name=flgo Version=0.0.1
