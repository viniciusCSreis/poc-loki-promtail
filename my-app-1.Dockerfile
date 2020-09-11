FROM alpine:3.10.2

COPY bin/my-app-1 /server/

WORKDIR /server

RUN chmod +x /server/my-app-1

ENTRYPOINT /server/my-app-1