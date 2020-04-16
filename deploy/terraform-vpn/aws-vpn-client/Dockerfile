FROM debian:stable-slim

RUN apt-get update && apt-get install -y \
        strongswan procps kmod iproute2 iptables gettext-base curl quagga-core
RUN curl -L https://github.com/osrg/gobgp/releases/download/v2.15.0/gobgp_2.15.0_linux_amd64.tar.gz | tar -xvzf - -C /usr/local/bin/ gobgp gobgpd
COPY ./ipsec.conf.tmpl /etc/ipsec.conf.tmpl
COPY ./ipsec.secrets.tmpl /etc/ipsec.secrets.tmpl
COPY ./gobgp.conf.tmpl /etc/gobgp.conf.tmpl
COPY ./zebra.conf /etc/quagga/zebra.conf
COPY ./aws-updown.sh /etc/ipsec.d/aws-updown.sh
RUN chmod 0755 /etc/ipsec.d/aws-updown.sh
ENTRYPOINT ipsec start --nofork
