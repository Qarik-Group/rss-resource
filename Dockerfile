FROM alpine
RUN apk add --no-cache ca-certificates
ADD assets/ /opt/resource/
