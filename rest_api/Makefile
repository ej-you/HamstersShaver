info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


dev:
	go run ./main.go dev

test:
	go run ./test_main.go

compile:
	go build -o ./HamstersShaver ./main.go

prod:
	@echo "Running migrations..."
	/root/HamstersShaver
	@echo "Running main app..."
	/root/HamstersShaver >> $(info_log) 2>> $(error_log)


swagger:
	@swag init
	@swag fmt
