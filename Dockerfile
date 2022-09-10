FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go install github.com/ametanishant/restapi@latest

RUN cd /build && git clone https://github.com/ametanishant/restapi.git

RUN cd /build/restapi && go build && ./restapi

EXPOSE 8080

ENTRYPOINT ["/build/restapi/main]
