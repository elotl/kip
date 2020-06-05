## Troubleshooting

### Kubectl describe

The output of `kubectl describe` is helpful to see why a pod is stuck in Pending or why it has errored.

### Virtual Kubelet Logs

The a good place to look for more answers is the output of the kip pod.

```bash
./kubectl -nkube-system logs kip -c kip -f
```

### Logging into Cells via SSH

As an extreme measure, it might be necessary to enable ssh access to a Cell in order to debug the system. The easiest way to do this is to configure a cloud-init file that authorizes a user to log into a Cell. See the [cloud-init](cloud-init.md) section for the details on using cloud-init. Briefly, to enable ssh access to a cell, provider.yml will need to contain a pointer to a cloud-init file and that file should contain commands to either add a user and their ssh key or add an ssh key for the root user.  The cloud-init file will need to be mounted along with provider.yml into the provider.

1. Specify the path to the cloud-init file in provider.yaml.  This is the path to the cloudinit.yaml file inside the provider's pod.  The easiest thing to do is to put the cloudinit.yaml file in the same directory as provier.yaml

```yaml

# snippet of /etc/kip/provider.yml
cells:
  cloudInitFile: /etc/kip/cloudinit.yaml
```

2. Create provider.yaml and cloud-init.yaml in a ConfigMap:

```bash
kubectl create configmap kip-config --from-file=./provider.yaml --from-file=./cloudinit.yaml
```

3. Add cloudinit.yaml as an item in the kip-config ConfigMap volume for kip:

```yaml
spec:
  - name: kip-config
    configMap:
      name: kip-config
      items:
      - key: provider.yaml
        path: provider.yaml
        mode: 0600
      - key: cloudinit.yaml
        path: cloudinit.yaml
        mode: 0600
```

### Installing Additional Software on Cells

Sometimes it is necessary to dig deep and start using more advanced tools to poke around your infrastructure and Cells. Cells run a customized version of Alpine. Once you’re setup with ssh access to the cell, it’s possible to add additional software to the Cell to troubleshoot issues.

```bash
# steps to install and use curl on the Node
$ apk update
$ apk add curl bind-tools
$ curl www.myhost.com:80
```

### Viewing Kip's Internal State

If you're doing development on the kip provider, it's helpful to see the state of resources inside the kip process.  The kip image packages an executable `kipctl` alongside the kip executable to communicate with kip.  To enable kipctl to talk to kip, add the `--debug-server` flag to the kip's command line arguments and restart the kip pod.  After execing into the pod you can inspect the internal state of kip.

```bash
./kipctl get pods
# output of pods similar to kubectl

./kipctl get nodes
# output of kip nodes (cells) similar to kubectl
```
