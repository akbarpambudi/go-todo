ENTRY_POINT = ./app/cmd
BUILD_TARGET_DIR = ./dst/bin
APP_NAME = godig-practice
.PHONY=build
build:
	go build -o ${BUILD_TARGET_DIR}/${APP_NAME} ${ENTRY_POINT}/*.go
.PHONY=run
run:
	go run ${ENTRY_POINT}/*.go
.PHONY=docker
docker: build
	docker build -t ${APP_NAME} .
