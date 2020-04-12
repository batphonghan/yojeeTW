FROM scratch

# RUN apk update && apk upgrade && apk add --no-cache bash git

WORKDIR /app

COPY server server

ENTRYPOINT ["/app/server"]
