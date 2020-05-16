FROM debian:stable-slim

ENV DEBIAN_FRONTEND "noninteractive"

RUN apt-get update -y && \
        apt-get dist-upgrade -y && \
        apt-get install -y curl gettext-base iproute2 openssl

RUN mkdir -p /usr/local/bin

RUN curl -fL https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/linux/amd64/kubectl > /usr/local/bin/kubectl
RUN chmod +x /usr/local/bin/kubectl

RUN curl -fL https://github.com/ldx/token2kubeconfig/releases/download/v0.0.2/token2kubeconfig-amd64 > /usr/local/bin/token2kubeconfig
RUN chmod +x /usr/local/bin/token2kubeconfig

RUN mkdir -p /opt/csr
COPY csr /opt/csr

CMD /opt/csr/create.sh
