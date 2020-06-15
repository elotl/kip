FROM debian:stable-slim

ARG K8S_VERSIONS
RUN test -n "$K8S_VERSIONS"

ENV DEBIAN_FRONTEND "noninteractive"

RUN apt-get update -y && \
        apt-get dist-upgrade -y && \
        apt-get install -y curl gettext-base iproute2 jq openssl

RUN set -e; \
    for k8s_version in ${K8S_VERSIONS}; do \
        echo "installing kubectl for kubernetes ${k8s_version}"; \
        vdir=/opt/kubectl/${k8s_version}; \
        mkdir -p ${vdir}; \
        vsn=$(curl -fL https://storage.googleapis.com/kubernetes-release/release/stable-${k8s_version}.txt); \
        curl -fL https://storage.googleapis.com/kubernetes-release/release/${vsn}/bin/linux/amd64/kubectl > ${vdir}/kubectl; \
        chmod 755 ${vdir}/kubectl; \
        echo "installed kubectl ${vsn} in ${vdir}"; \
    done

RUN mkdir -p /usr/local/bin
COPY kubectl.sh /usr/local/bin/kubectl

RUN curl -fL https://github.com/ldx/token2kubeconfig/releases/download/v0.0.2/token2kubeconfig-amd64 > /usr/local/bin/token2kubeconfig
RUN chmod +x /usr/local/bin/token2kubeconfig

RUN mkdir -p /opt/csr
COPY csr /opt/csr

CMD /opt/csr/get-cert.sh
