FROM golang:1.10

WORKDIR /home/gateway
ADD service/ /home/gateway

RUN make build

EXPOSE 80

ENTRYPOINT ["./deploy/gateway.o", ":80"]