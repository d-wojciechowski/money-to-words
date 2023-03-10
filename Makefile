run:
	go run main.go
fmt:
	go mod tidy
	go fmt ./...
bench:
	go test -v -bench . -run=^# ./...
cover:
	go test -v -cover ./...
bench5:
	go test -v -bench . -count 5 -run=^# ./...
test:
	go test -v ./...
build:
	go build  -ldflags '-w -s' -a -installsuffix cgo -o app.exe
docker-image:
	docker build -t unlucky/currency-converter .
docker-run:
	docker build -t unlucky/currency-converter . && docker run --rm -p 8081:8081 --name currency-converter unlucky/currency-converter