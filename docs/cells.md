## Cells

Cells are the name we’ve given to the cloud instances that Kip runs pods on.  The cells run Alpine Linux and a lightweight agent called [Itzo](https://github.com/elotl/itzo) that performs the task of running pods on the cells.  More information about cells can be found in the itzo repository.

To boot cells, unless an instance-type annotation is present in the pod, kip will choose the cheapest instance type that satisfies the resource requirements of the pod and will fall back to the `defaultInstanceType` in the `provider.yaml` file if no resources are specified.  Server certificates are passed to the cell via instance user data and those certificates allow the kip controller to connect to the cell.

Cells can be customized by specifying a [cloud-init](cloud-init.md) file that will be applied when the instance boots.

### Bring your Own Image

We maintain images that are optimized for cells and come with our tools pre-installed. However, if you need a custom-built image or you have a dependency only available for certain Linux distributions, you can either add our cell agent to your image builds, or use cloud-init to fetch it and start it at boot time.

You can check out our image build scripts and configuration files [here](https://github.com/elotl/kip-cell-image).

If you want to use cloud-init to download our cell agent, first update your provider config:

    cells:
      cloudInitFile: /etc/kip/cloudinit.yaml
      bootImageSpec:
        owners: "099720109477"
        filters: name=ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-*

Add a cloud-init section to the provider configmap, e.g. for ubuntu or debian:

    cloudinit.yaml: |
      runcmd:
        - apt-get update && apt-get install -y ipset iproute2 iptables
        - wget -O /usr/local/bin/itzo https://itzo-kip-download.s3.amazonaws.com/itzo-latest
        - wget -O /usr/local/bin/kube-router https://milpa-builds.s3.amazonaws.com/kube-router
        - wget -O /usr/local/bin/tosi https://tosi.s3.amazonaws.com/tosi
        - chmod 755 /usr/local/bin/*
        - mount --make-rprivate /
        - /usr/local/bin/itzo -v=2 > /var/log/itzo.log 2>&1 &

Finally, restart the provider:

    $ kubectl delete pod -n kube-system -l app=kip-provider
