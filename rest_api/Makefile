info_log = "/logs/info-log.log"
error_log = "/logs/error-log.log"


dev:
	go run ./main.go dev

test:
	go run ./test_main.go

compile:
	go build -o ./HamstersShaverRestAPI ./main.go

prod:
	@echo "Running main app..."
	/root/HamstersShaverRestAPI >> $(info_log) 2>> $(error_log)

swagger:
	@swag init
	@swag fmt
