## API server access

Access to the Kubernetes API server is needed:
* To create the virtual node object, watch pods, and list various Kubernetes resources in Kip. 
* To allow Kip to create and manage its CRDs.
* By the network agent, running on cells, to be able to implement service proxying and manage network policies.

By default, Kubernetes service accounts are used for authentication and authorization. The Kip pod service account is used for the first two use cases, and another service account is used by the cell network agent. If you use the [example deployment manifests and configs](deploy/) provided, the necessary service accounts and cluster roles will be created and configured at setup time.

As another option, both Kip and the network agent can use a kubeconfig file (either the same, or separate ones).

You can use a ConfigMap to attach a kubeconfig file into the container:

    apiVersion: v1
    kind: ConfigMap
    metadata:
      name: kip-config-cfd95dk474
      namespace: kube-system
    data:
      kubeconfig: |
        apiVersion: v1
        clusters:
        - cluster:
            certificate-authority-data: <base64-encoded CA certificate>
            server: https://10.20.30.40:6443 # API server.
          name: kubernetes
        contexts:
        - context:
            cluster: kubernetes
            user: kubernetes-admin
          name: kubernetes-admin@kubernetes
        current-context: kubernetes-admin@kubernetes
        kind: Config
        preferences: {}
        users:
        - name: kubernetes-admin
          user:
            client-key-data: <base64-encoded client key>
            client-certificate-data: <base64-encoded client certificate>

Create a volume of the file in the ConfigMap:

    volumes:
    - configMap:
        defaultMode: 420
        items:
        - key: kubeconfig
          mode: 384
          path: kubeconfig
        name: kip-config-cfd95dk474
      name: kip-config

Attach the volume to the container:

    containers:
    - name: kip
      [...]
      volumeMounts:
      - mountPath: /etc/kip
        name: kip-config

To configure access via a kubeconfig file, make sure the kubeconfig file is available and set the `KUBECONFIG` environment variable for the Kip container:

    containers:
    - name: kip
      [...]
      env:
      - name: KUBECONFIG
        value: /etc/kip/kubeconfig

For the network agent, use the command line parameter `--network-agent-kubeconfig`, for example:

    containers:
    - name: kip
      command:
      - /kip
      - --provider
      - kip
      - --provider-config
      - /etc/kip/provider.yaml
      - --network-agent-kubeconfig
      - /etc/kip/kubeconfig
      env:
      - name: KUBECONFIG
        value: /etc/kip/kubeconfig
