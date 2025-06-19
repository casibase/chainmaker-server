FROM golang:1.23.6 AS back
WORKDIR /go/src/chainserver
COPY . .
RUN apt-get update && apt-get install -y bash gcc
RUN chmod +x ./build.sh
RUN ./build.sh

FROM debian:stable-slim AS standard
LABEL MAINTAINER="https://casibase.org/"
ARG USER=chainserver

RUN apt-get update && apt-get install -y sudo curl ca-certificates lsof && rm -rf /var/lib/apt/lists/*

RUN useradd -m -u 1000 $USER \
    && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
    && chmod 0440 /etc/sudoers.d/$USER \
    && mkdir /logs \
    && chown -R $USER:$USER /logs

USER 1000
WORKDIR /
COPY --from=back --chown=$USER:$USER /go/src/chainserver/server ./server
COPY --from=back --chown=$USER:$USER /go/src/chainserver/swagger ./swagger

ENTRYPOINT ["/server"]
