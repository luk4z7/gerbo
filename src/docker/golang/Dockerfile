FROM golang:1.8

ENV BUILDPATH /go/src/gerbo
ENV TOOLS /var/exec
ENV GOBINARIES /go/bin

WORKDIR $GOBINARIES
RUN go get github.com/cespare/reflex
RUN go get gopkg.in/check.v1
RUN go get github.com/mxk/go-sqlite/sqlite3
RUN go get gopkg.in/mgo.v2
RUN go get github.com/agtorre/gocolorize
RUN go get github.com/fatih/structs
RUN go get github.com/jinzhu/gorm

RUN mkdir -p $TOOLS
ADD build.sh $TOOLS
ADD reflex.conf $TOOLS
RUN chmod +x $TOOLS/build.sh

WORKDIR $BUILDPATH
CMD ["reflex","-c","/var/exec/reflex.conf"]