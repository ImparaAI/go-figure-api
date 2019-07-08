#!/bin/sh

if [ $DEVELOPMENT == true ]; then
  cd /go/src
  go get github.com/codegangsta/gin
  cd /go/src/app

  export GIN_BIN=/../../../tmp/gin-bin
  export GIN_PORT=8080
  export BIN_APP_PORT=8081
  export APP_PORT=8081

  touch /var/healthy

  gin run main.go
else
  touch /var/healthy

  supervisord -c '/etc/supervisor.d/supervisord.ini'
fi

