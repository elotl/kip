FROM debian:stable-slim

RUN apt-get update && \
        apt-get install -y strongswan procps kmod iproute2 iptables gettext-base
COPY ./ipsec.conf /etc/ipsec.conf.tmpl
COPY ./ipsec.secrets /etc/ipsec.secrets.tmpl
COPY ./aws-updown.sh /etc/ipsec.d/aws-updown.sh
RUN chmod 0755 /etc/ipsec.d/aws-updown.sh
ENTRYPOINT ipsec start --nofork
