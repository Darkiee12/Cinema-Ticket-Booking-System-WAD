echo "Building docker image..."
docker build -t cinema .
docker tag cinema:latest cinema:staging

echo "Running docker compose..."
docker-compose up -d
