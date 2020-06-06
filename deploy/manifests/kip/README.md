# Deployment Manifests

Here you can find kustomize templates for deploying Kip.

## Deploy

Use the base directory for a basic setup:

    $ kubectl apply -k base/

## Override Settings or Configuration

Templating makes it really easy to override certain parameters or configuration settings. For example, to provide your own provider configuration file for Kip:

    $ mkdir -p overlays/local-dev
    $ cp base/provider.yaml overlays/local-dev/
    $ vi overlays/local-dev/provider.yaml
    $ cat <<-EOF > overlays/local-dev/kustomization.yaml
    bases:
    - ../../base
    configMapGenerator:
    - name: kip-config
      namespace: kube-system
      behavior: merge
      files:
      - provider.yaml
    EOF
    $ kubectl apply -k overlays/local-dev/
