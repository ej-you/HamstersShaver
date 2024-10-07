info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


test:
	go run ./test_main.go

dev:
	go run ./main.go dev

compile:
	go build -o ./HamstersShaver ./main.go

prod:
	@echo "Running migrations..."
	/root/HamstersShaver
	@echo "Running main app..."
	/root/HamstersShaver >> $(info_log) 2>> $(error_log)

