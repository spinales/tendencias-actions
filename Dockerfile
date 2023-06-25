FROM golang:1.17

ENV GO111MODULE=on

ENV APP_HOME /go/src/tendencias-actions
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"
COPY . .
EXPOSE 8080
CMD ["go", "run", "main.go"]
