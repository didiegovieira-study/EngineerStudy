FROM golang:latest

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
#ENV GOPROXY=https://proxy.golang.org

RUN go mod download -modcacherw

COPY . .

RUN go build -o app

CMD ["./app"]
