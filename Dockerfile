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
FROM scratch
WORKDIR /app
COPY config.json .
COPY --from=builder /app/fantlab .
ENV PORT=4242
ENV MYSQL_CS=root@tcp(host.docker.internal:3306)/fl?charset=utf8&parseTime=True&loc=Local
ENV CONFIG_FILE=config.json
ENV GIN_MODE=release
EXPOSE 4242
ENTRYPOINT [ "/app/fantlab" ]
LABEL Name=flgo Version=0.0.1
