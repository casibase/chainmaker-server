FROM golang:1.23.6-alpine AS BACK
WORKDIR /go/src/chainserver
COPY . .
RUN apk add --no-cache gcc musl-dev
RUN chmod +x ./build.sh
RUN ./build.sh

FROM alpine:latest AS STANDARD
LABEL MAINTAINER="https://casibase.org/"
ARG USER=chainserver

RUN sed -i 's/https/http/' /etc/apk/repositories
RUN apk add --update sudo
RUN apk add curl
RUN apk add ca-certificates && update-ca-certificates

RUN adduser -D $USER -u 1000 \
    && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
    && chmod 0440 /etc/sudoers.d/$USER \
    && mkdir logs \
    && chown -R $USER:$USER logs

USER 1000
WORKDIR /
COPY --from=BACK --chown=$USER:$USER /go/src/chainserver/server ./server

ENTRYPOINT ["/server"]
