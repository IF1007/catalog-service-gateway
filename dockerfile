FROM golang:1.10

WORKDIR /home/gateway
# TODO: discover why this work and I dont need cd service
ADD service/ /home/gateway

RUN make build

EXPOSE 80

ENTRYPOINT ["./deploy/gateway.o", "v1"]