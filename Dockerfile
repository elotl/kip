FROM alpine

RUN apk add --update bash ca-certificates iptables

copy third_party /third_party

COPY kipctl /kipctl
RUN chmod 755 /kipctl
COPY virtual-kubelet /virtual-kubelet
RUN chmod 755 /virtual-kubelet

ENTRYPOINT ["/virtual-kubelet"]
