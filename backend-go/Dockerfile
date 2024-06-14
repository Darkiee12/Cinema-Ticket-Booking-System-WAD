FROM golang:1.21.6-alpine AS build

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify && go install github.com/swaggo/swag/cmd/swag@latest
COPY . .
RUN swag init
# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app-linux .

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
COPY --from=build /build/app-linux .

# ADD ./zoneinfo.zip /usr/lsocal/go/lib/time/

CMD ["/app/app-linux"]