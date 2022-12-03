FROM golang:1.19-alpine

WORKDIR /usr/src/server

COPY ./ ./
RUN go mod download && \
    go build -v -o /usr/local/bin/main ./main.go 

CMD ["main"]







