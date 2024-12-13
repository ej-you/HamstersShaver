info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


dev:
	go run ./main.go dev

migrate:
	go run ./main.go migrate

compile:
	go build -o ./tg_bot ./main.go

prod:
	@echo "Running migrations..."
	/root/tg_bot migrate
	@echo "Running main app..."
	/root/tg_bot >> $(info_log) 2>> $(error_log)
