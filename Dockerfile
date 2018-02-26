FROM alpine
RUN apk add --no-cache curl
ADD bin /opt/resource/
