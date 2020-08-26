FROM alpine:3.12

EXPOSE 8080

WORKDIR /root/

COPY main app
COPY template template
COPY asset asset

CMD ["./app"]