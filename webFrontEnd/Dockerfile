FROM scratch
# RUN apk update && apk upgrade && apk add --no-cache bash git

WORKDIR /app

COPY webFrontEnd webFrontEnd
COPY tweet.html tweet.html

ENTRYPOINT ["/app/webFrontEnd"]
