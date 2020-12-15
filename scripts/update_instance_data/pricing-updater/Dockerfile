FROM debian:stable-slim

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
        apt-get upgrade -y

COPY pricing-updater /pricing-updater
RUN chmod 755 /pricing-updater

ENTRYPOINT ["/pricing-updater"]
