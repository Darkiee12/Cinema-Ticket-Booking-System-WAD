swag init

:: build for linux
set GOOS=linux
set GOARCH=arm64
set CGO_ENABLED=0
go build -a -o app-linux

:: build for windows
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=0
go build -a -o app-windows.exe

docker build -t cinema .
docker tag cinema:staging
