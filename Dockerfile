FROM debian:stable-slim

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
        apt-get upgrade -y && \
        apt-get install -y ca-certificates iptables

COPY kipctl /kipctl
RUN chmod 755 /kipctl
COPY virtual-kubelet /virtual-kubelet
RUN chmod 755 /virtual-kubelet

ENTRYPOINT ["/virtual-kubelet"]
