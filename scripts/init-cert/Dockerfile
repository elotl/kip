FROM debian:stable-slim

ARG K8S_VERSION_MAJOR
RUN test -n "$K8S_VERSION_MAJOR"
ARG K8S_VERSION_MINOR
RUN test -n "$K8S_VERSION_MINOR"

ENV DEBIAN_FRONTEND "noninteractive"

RUN apt-get update -y && \
        apt-get dist-upgrade -y && \
        apt-get install -y curl gettext-base iproute2 openssl

RUN mkdir -p /usr/local/bin

RUN set -e; \
    echo "installing kubectl for kubernetes ${K8S_VERSION_MAJOR}.${K8S_VERSION_MINOR}"; \
    vsn=$(curl -fL https://storage.googleapis.com/kubernetes-release/release/stable-${K8S_VERSION_MAJOR}.${K8S_VERSION_MINOR}.txt); \
    curl -fL https://storage.googleapis.com/kubernetes-release/release/${vsn}/bin/linux/amd64/kubectl > /usr/local/bin/kubectl-${vsn}; \
    chmod 755 /usr/local/bin/kubectl-${vsn}; \
    ln -snf /usr/local/bin/kubectl-${vsn} /usr/local/bin/kubectl; \
    echo "installed kubectl ${vsn}"

RUN curl -fL https://github.com/ldx/token2kubeconfig/releases/download/v0.0.2/token2kubeconfig-amd64 > /usr/local/bin/token2kubeconfig
RUN chmod +x /usr/local/bin/token2kubeconfig

RUN mkdir -p /opt/csr
COPY csr /opt/csr

CMD /opt/csr/get-cert.sh
