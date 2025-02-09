dev:
	go run ./main.go dev


dev-run: dev-stop
	docker compose -f docker-compose.dev.yml up --build -d
dev-stop:
	docker compose -f docker-compose.dev.yml down


run: stop
	docker compose up --build -d
stop:
	docker compose down
