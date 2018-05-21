############################# Makefile ##########################
clean:
	@rm -rf deploy/
build:
	@go get -v -u github.com/gorilla/mux
	@make clean
	@mkdir deploy
	@go build -o "deploy/gateway.o" .
run:
	@make build
	@./release/gateway.o :80