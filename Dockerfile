FROM ubuntu:18.04

EXPOSE 8080

WORKDIR /root/

COPY main app

CMD ["./app"]