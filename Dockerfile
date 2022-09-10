FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com:ametanishant/restapi
RUN go get github.com/gorilla/mux
RUN cd /build && git clone github.com:ametanishant/restapi.git

RUN cd /build/restapi && go build && ./restapi

EXPOSE 8080

ENTRYPOINT ["/build/restapi/main/main"]
