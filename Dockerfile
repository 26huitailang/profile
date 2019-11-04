FROM golang:1.13-alpine
RUN mkdir /code
RUN apt install make -y
WORKDIR /code
COPY . /code/
RUN make init-env
RUN make build