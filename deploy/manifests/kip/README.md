# Deployment Manifests

Here you can find kustomize templates for deploying Kip.

## Deploy

You need [kustomize](https://github.com/kubernetes-sigs/kustomize) and [kubectl](https://github.com/kubernetes/kubectl). Use the base directory for a basic setup:

    $ kustomize build base/ | kubectl apply -f -

## Override Settings or Configuration

Templating makes it really easy to override certain parameters or configuration settings. For example, to provide your own provider configuration file for Kip:

    $ mkdir -p overlays/local-dev
    $ cp base/provider.yaml overlays/local-dev/
    $ vi overlays/local-dev/provider.yaml
    $ cat <<-EOF > overlays/local-dev/kustomization.yaml
    bases:
    - ../../base
    configMapGenerator:
    - name: config
      behavior: merge
      files:
      - provider.yaml
    EOF
    $ kustomize build overlays/local-dev/ | kubectl apply -f -
