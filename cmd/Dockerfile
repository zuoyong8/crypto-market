# build stage
FROM golang:alpine as build-env
RUN apk add build-base

ENV GO111MODULE=on
ADD . /src

RUN cd /src && \
  go build -o bin/server cmd/server.go
  

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/server /app/bin/server
#COPY --from=build-env /src/config/config.yml.production /app/config/config.yml

ENV TZ Asia/Shanghai

EXPOSE 8080

CMD ["/app/bin/server","run"]
