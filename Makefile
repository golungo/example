clean:
	rm -rf ./app

# Dev DB 
db:
	docker-compose rm -f
	docker-compose up --build

# Dev Backend 
start:
	go run main.go --config scripts/default.config.yaml

build:
	go build -o start main.go