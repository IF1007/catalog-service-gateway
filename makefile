build:
	docker build -t csd-gateway .
run:
	docker run -p 80:80 --name gateway csd-gateway &