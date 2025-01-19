dev:
	go run ./main.go dev


dev-run: dev-stop
	docker compose -f docker-compose.dev.yml up --build
dev-stop:
	docker compose -f docker-compose.dev.yml down


run: stop
	docker compose up --build
stop:
	docker compose down
