FROM busybox

EXPOSE 8080

WORKDIR /root/

COPY main app
COPY template template
COPY asset asset

CMD ["./app"]