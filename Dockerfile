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
# Turn on Go 1.11 Modules and build
WORKDIR $GOPATH/src
ENV GIN_BIN=/../../../tmp/gin-bin
ENV GIN_PORT=8080
ENV BIN_APP_PORT=8081
ENV APP_PORT=8081
ENV GO111MODULE=on
ENV CGO_ENABLED=0
RUN go get -v github.com/codegangsta/gin

# Copy code to image
WORKDIR $GOPATH/src/app
COPY . .

RUN go build -o /bin/app

CMD gin run main.go


# Production build
FROM alpine AS final

COPY --from=builder /bin/app /bin/app

WORKDIR /

ENV APP_PORT=8080

EXPOSE 8080

CMD /bin/app