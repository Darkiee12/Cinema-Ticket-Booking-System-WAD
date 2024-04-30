FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
ADD ./app-linux /app/
ADD ./local_db_env /app/

# ADD ./zoneinfo.zip /usr/lsocal/go/lib/time/

ENTRYPOINT ["./app-linux"]