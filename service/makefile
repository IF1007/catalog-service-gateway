############################# Makefile ##########################
clean:
	@rm -rf deploy/
	@rm -rf secret.txt
build:
	@go get -v -u github.com/gorilla/mux
	@go get -v -u github.com/globalsign/mgo
	@go get -v -u github.com/satori/go.uuid
	@make clean
	@mkdir deploy
	@go build -o "deploy/gateway.o" .
run:
	@make build
	@./deploy/gateway.o :80