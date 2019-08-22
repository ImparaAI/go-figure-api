FROM golang:1.12-alpine as builder

# Install packages
RUN echo 'http://dl-cdn.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories && \
  apk update && \
  apk upgrade && \
  apk --no-cache add \
    gcc \
    libc-dev \
    git

# Set up gin development tool
WORKDIR $GOPATH/src

RUN go get -v github.com/codegangsta/gin
ENV GIN_BIN=/../../../tmp/gin-bin
ENV GIN_PORT=8080
ENV BIN_APP_PORT=8081
ENV DEVELOPMENT=true

WORKDIR $GOPATH/src/app

COPY . .
COPY docker/start.sh /bin/original_start.sh

# Turn on Go 1.11 Modules and build
ENV GO111MODULE=on
RUN go build -o /bin/app

# Set up start script
RUN tr -d '\r' < /bin/original_start.sh > /bin/start.sh && \
    chmod -R 700 /bin/start.sh

# Set application port env var
ENV APP_PORT=8080

ENTRYPOINT ["/bin/start.sh"]

# Production build
FROM scratch AS final

COPY --from=builder /bin/start.sh /bin
COPY --from=builder /bin/app /bin

WORKDIR /

ENV APP_PORT=8080

EXPOSE 8080

ENTRYPOINT ["/bin/start.sh"]