FROM alpine

EXPOSE 8080

WORKDIR /root/

RUN apk add --no-cache bash

COPY main app
COPY template template
COPY asset asset

CMD ["./app"]