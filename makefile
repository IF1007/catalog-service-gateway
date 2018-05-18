build:
	docker build -t csd-gateway .
run:
	docker run -p 4000:80 --name gateway csd-gateway &