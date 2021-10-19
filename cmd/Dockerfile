# build stage
FROM golang:alpine as build-env
RUN apk add build-base

ENV GO111MODULE=on
#ENV GOPROXY="https://goproxy.cn"

WORKDIR /src

COPY . /src

RUN cd /src && \
  go mod download

RUN  go build -o bin/server cmd/server.go
  

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/server /app/server
COPY --from=build-env /src/config/config.toml /app/config/config.toml 

ENV TZ Asia/Shanghai

EXPOSE 8080



CMD ["/app/server","run"]
