FROM alpine
RUN apk add --no-cache curl
ADD assets/* /opt/resource/
