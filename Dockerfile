FROM golang:1.12 as builder

WORKDIR /app
COPY app/go.mod .
COPY app/go.sum .
RUN go mod download
COPY app .

ENV GO111MODULE=on
ENV PORT=4242
ENV MYSQL_CS=root@/fl?charset=utf8&parseTime=True&loc=Local
ENV GIN_MODE=release

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# *****************

FROM scratch
WORKDIR /app
COPY --from=builder /app/fantlab .
EXPOSE 4242
ENTRYPOINT ["/app/fantlab"]
