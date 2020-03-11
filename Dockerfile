FROM alpine

RUN apk add --update bash ca-certificates iptables

COPY virtual-kubelet /virtual-kubelet
RUN chmod 755 /virtual-kubelet
COPY milpactl /kipctl
RUN chmod 755 /kipctl

ENTRYPOINT ["/virtual-kubelet"]
