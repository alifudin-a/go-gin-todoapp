pgbash:
	docker exec -it pq-arsip-surat-unggulan bash

pgip:
	docker exec -i pq-arsip-surat-unggulan hostname -i | awk '{print $1}'

run:
	go run cmd/main/main.go

build:
	cd cmd/main; go build -o ../../bin/go-todoapp

exec:
	./bin/go-todoapp

startapp: build exec

tidy:
	go mod tidy