echo "Downloading packages..."
go mod download
echo "Compiling..."
GOARCH=arm64 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app-linux

echo "Building docker image..."
docker build -t cinema .
docker tag cinema:latest cinema:staging

echo "Running docker compose..."
docker-compose up -d
