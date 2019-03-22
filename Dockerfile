FROM golang:1.11.1 AS build
WORKDIR /mnt
COPY . .
RUN CGO_ENABLED=0 go build -o ./bin/server main.go
RUN CGO_ENABLED=0 go build -o ./bin/web web/main.go

FROM alpine:3.7
WORKDIR /opt
EXPOSE 1234
EXPOSE 8080
RUN apk add --no-cache ca-certificates
COPY --from=build /mnt/bin/* /usr/bin/
CMD ["server"]
