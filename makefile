server-run:
	docker-compose up -d

local-run:
	go run main.go

server-stop:
	docker-compose stop

remove:
	docker-compose down