info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


dev:
	go run ./main.go dev

migrate:
	go run ./main.go migrate

compile:
	go build -o ./HamstersShaverBot ./main.go

prod:
	@echo "Running migrations..."
	/root/HamstersShaverBot
	@echo "Running main app..."
	/root/HamstersShaverBot >> $(info_log) 2>> $(error_log)
