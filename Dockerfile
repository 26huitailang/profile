FROM golang:1.13-alpine

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add make gcc musl-dev

ENV GOPROXY https://goproxy.io
ENV GO111MODULE=on

WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

RUN mkdir /code
WORKDIR /code
COPY . /code/
#RUN make init-env
