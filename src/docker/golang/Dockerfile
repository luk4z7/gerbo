FROM golang:1.8

# Configure environment variables
ENV BUILDPATH /go/src/gerbo
ENV TOOLS /var/exec
ENV GOBINARIES /go/bin

# Install sqlite3 e libsqlite3 dev
RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y sqlite3 libsqlite3-dev
RUN apt-get install -y cron vim

# Compile in folder of GOBINARIES -> /go/bin
WORKDIR $GOBINARIES
RUN go get github.com/cespare/reflex

RUN go get gopkg.in/check.v1
RUN go get gopkg.in/mgo.v2
RUN go get github.com/agtorre/gocolorize
RUN go get github.com/jinzhu/inflection
RUN go get github.com/fatih/structs
RUN go get github.com/jinzhu/gorm
RUN go get rsc.io/sqlite

# Send to GOROOT
RUN cd /go/src/ && mv rsc.io /usr/local/go/src/
RUN cd /go/src/ && mv github.com /usr/local/go/src/
RUN cd /go/src/ && mv gopkg.in /usr/local/go/src/

RUN mkdir /data/
RUN mkdir -p $TOOLS
ADD build.sh $TOOLS
ADD reflex.conf $TOOLS
RUN chmod +x $TOOLS/build.sh
ADD twitter-movie-ratings.db /data/
ADD insert.sql /data/
ADD entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENV PORT 6060
EXPOSE $PORT

CMD ["/entrypoint.sh"]