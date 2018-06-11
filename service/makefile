############################# Makefile ##########################
clean:
	@rm -rf deploy/
build:
	@go get -v -u github.com/gorilla/mux
	@go get -v -u go get gopkg.in/mgo.v2
	@make clean
	@mkdir deploy
	@go build -o "deploy/gateway.o" .
run:
	@make build
	@./deploy/gateway.o :80