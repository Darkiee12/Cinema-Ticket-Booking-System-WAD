run:
	GO111MODULE=off go run main.go
swagger:
	swag init
build: swagger
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -a -o app-linux
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -a -o app-windows.exe
