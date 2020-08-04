ENTRY_POINT = ./cmd
BUILD_TARGET_DIR = ./build/bin
APP_NAME = todo
.PHONY=build
build:
	go build -o ${BUILD_TARGET_DIR}/${APP_NAME} ${ENTRY_POINT}/*.go
.PHONY=run
run:
	go run ${ENTRY_POINT}/*.go
.PHONY=docker
docker: build
	docker build -t ${APP_NAME} .
