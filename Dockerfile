FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go install github.com/ametanishant/restapi@latest

RUN cd /build && git clone https://github.com/ametanishant/restapi.git

RUN cd /build/restapi && go build 

RUN chmod +x ./restapi

EXPOSE 8000

ENTRYPOINT ["./restapi"]
