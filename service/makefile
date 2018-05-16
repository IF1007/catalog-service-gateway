############################# Makefile ##########################
clean:
	@rm -rf deploy/
build:
	@make clean
	@mkdir deploy
	@go build -o "deploy/gateway.o" .
run:
	@make build
	@./release/gateway.o config.json