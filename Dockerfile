FROM alpine

RUN apk add --update bash ca-certificates iptables

COPY virtual-kubelet /virtual-kubelet
RUN chmod 755 /virtual-kubelet

ENTRYPOINT ["/virtual-kubelet"]
