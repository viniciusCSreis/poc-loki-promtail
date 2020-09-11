FROM alpine:3.10.2

COPY bin/my-app-2 /server/

WORKDIR /server

RUN chmod +x /server/my-app-2

ENTRYPOINT /server/my-app-2