#build stage
FROM golang:1.13 AS builder
ENV GO111MODULE=on
WORKDIR /app
COPY sources .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

#final stage
FROM alpine:latest
ENV TZ=Europe/Moscow
RUN apk update && apk add tzdata && cp -r -f /usr/share/zoneinfo/$TZ /etc/localtime
WORKDIR /app
COPY --from=builder /app/fantlab .
COPY docker-entrypoint.sh wait-for.sh ./
RUN chmod +x ./docker-entrypoint.sh ./wait-for.sh
ENTRYPOINT [ "./docker-entrypoint.sh", "./fantlab" ]
LABEL Name=FantLab
