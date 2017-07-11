FROM golang:1.7
# app specific env
ENV BUILDPATH /go/src/middleware-jwt
ENV TOOLS /var/exec
ENV GOBINARIES /go/bin

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y htop

# Install reflex
WORKDIR $GOBINARIES
RUN go get github.com/cespare/reflex

RUN go get github.com/codegangsta/negroni
RUN go get github.com/dgrijalva/jwt-go
RUN go get gopkg.in/check.v1
RUN go get github.com/garyburd/redigo/redis
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get github.com/rs/cors
RUN go get github.com/agtorre/gocolorize
RUN go get github.com/fatih/structs

# Send to GOROOT
RUN cd /go/src/ && mv github.com /usr/local/go/src/
RUN cd /go/src/ && mv gopkg.in /usr/local/go/src/

ENV PORT 6060
# dockeer/app port
EXPOSE $PORT

# Make directories and add files as needed
RUN mkdir -p $TOOLS
ADD build.sh $TOOLS
ADD reflex.conf $TOOLS
RUN chmod +x $TOOLS/build.sh

# Execute reflex.
WORKDIR $BUILDPATH
CMD ["reflex","-c","/var/exec/reflex.conf"]
