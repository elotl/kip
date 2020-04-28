FROM debian:stable-slim

RUN apt-get update && apt-get install -y \
        strongswan procps kmod iproute2 iptables gettext-base curl quagga-core
RUN curl -L https://github.com/osrg/gobgp/releases/download/v2.15.0/gobgp_2.15.0_linux_amd64.tar.gz | tar -xvzf - -C /usr/local/bin/ gobgp gobgpd
RUN curl -L https://github.com/kelseyhightower/confd/releases/download/v0.16.0/confd-0.16.0-linux-amd64 > /usr/local/bin/confd; chmod 755 /usr/local/bin/confd
COPY ./ipsec.conf.tmpl /etc/ipsec.conf.tmpl
COPY ./ipsec.secrets.tmpl /etc/ipsec.secrets.tmpl
COPY ./zebra.conf /etc/quagga/zebra.conf
COPY ./aws-updown.sh /etc/ipsec.d/aws-updown.sh
RUN chmod 0755 /etc/ipsec.d/aws-updown.sh
RUN mkdir -p /etc/confd/conf.d
RUN mkdir -p /etc/confd/templates
COPY ./gobgpd.toml /etc/confd/conf.d/gobgpd.toml
COPY ./gobgpd.conf.tmpl /etc/confd/templates/gobgpd.conf.tmpl
ENTRYPOINT ipsec start --nofork
