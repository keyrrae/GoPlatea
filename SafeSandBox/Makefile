IMAGE_NAME ?= go-playground/safesandbox

docker:
	go get -d ./... && \
	GOOS=linux GOARCH=amd64 go build -o app_linux_amd64 && \
	docker build -t $(IMAGE_NAME) .

run:
	docker run -d -p 3000:3000 -name go-playground -t $(IMAGE_NAME)