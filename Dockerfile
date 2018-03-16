FROM alpine:3.6
USER root
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY pc-go-daemon .

CMD ["./pc-go-daemon"]