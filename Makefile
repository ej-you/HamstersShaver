info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


dev:
	go run ./main.go

migrate:
	go run ./main.go migrate

compile:
	go build -o ./HamstersShaver ./main.go

prod:
	@echo "Running main app..."
	/root/HamstersShaver >> $(info_log) 2>> $(error_log)

