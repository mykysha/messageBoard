status-check:
	curl http://localhost:8080/status

cl-start:
	go run main.go

docker-build:
	docker-compose up -d

docker-delete:
	docker-compose down --remove-orphans --volumes
	docker image rm messageBroker_server

docker-rebuild:
	make docker-delete
	make docker-build

docker-start:
	docker container start server

docker-stop:
	docker container stop server

update-gomod:
	make download-gomod
	make gomod-tidy

download-gomod:
	go get -u "github.com/spf13/viper"
	go get -u "github.com/gorilla/mux"

gomod-tidy:
	go mod tidy -go=1.18.2 -compat=1.18.2
