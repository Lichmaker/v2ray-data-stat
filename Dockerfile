FROM golang:1.16.5-alpine3.13 as goBuilder
RUN apk add --no-cache bash
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN /bin/bash -c 'if [ ! -e "cmd/stat/.env" ]; then  echo "env file not found" ;  exit 1  ; else echo "success" ; exit 0; fi '
WORKDIR /app/cmd/stat
RUN go build -o v2ray-data-stat

FROM alpine:latest
RUN apk add tzdata
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' > /etc/timezone
RUN mkdir -p /v2ray-data-stat
WORKDIR /v2ray-data-stat
COPY --from=goBuilder /app/cmd/stat/v2ray-data-stat .
COPY --from=goBuilder /app/cmd/stat/.env .
ENTRYPOINT ["./v2ray-data-stat"]
