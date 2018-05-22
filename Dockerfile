FROM golang:1.10

RUN go get -v github.com/sirupsen/logrus
RUN go get -v github.com/denisenkom/go-mssqldb
RUN go get -v github.com/jessevdk/go-flags
RUN go get -v github.com/mitchellh/gox

COPY . /go/src/github.com/odino/mssqldump
WORKDIR /go/src/github.com/odino/mssqldump

CMD bash
