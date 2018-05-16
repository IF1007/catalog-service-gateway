FROM golang:1.10

WORKDIR /gateway

# TODO: discover why this work and I dont need cd service
ADD service/ /gateway

RUN make build

EXPOSE 80

CMD ["./deploy/gateway.o", "'v1'"]