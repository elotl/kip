#!/usr/bin/env bash

set -euo pipefail

kubectl_dir_base=/opt/kubectl
default_kubectl=${kubectl_dir_base}/1.16/kubectl

version_doc="$($default_kubectl version --output=json | jq -r '.serverVersion')"

major_vsn=$(echo "$version_doc" | jq -r '.major' | sed -r 's/([0-9]+).*/\1/')
minor_vsn=$(echo "$version_doc" | jq -r '.minor' | sed -r 's/([0-9]+).*/\1/')

kubectl=$default_kubectl
if [[ -d /opt/kubectl/${major_vsn}.${minor_vsn} ]]; then
    kubectl=/opt/kubectl/${major_vsn}.${minor_vsn}/kubectl
else
    for kdir in ${kubectl_dir_base}/*; do
        echo $kdir
        vdir=$(basename ${kdir})
        kdir_major=$(echo ${vdir} | sed -r 's/([0-9]+)\.([0-9]+)/\1/')
        if [[ ${kdir_major} -ne ${major_vsn} ]]; then
            continue
        fi
        kdir_minor=$(echo ${vdir} | sed -r 's/([0-9]+)\.([0-9]+)/\2/')
        if [[ ${kdir_minor} -eq $((${minor_vsn}-1)) ]] || \
            [[ ${kdir_minor} -eq $((${minor_vsn}+1)) ]]; then
            kubectl=${kdir}/kubectl
        fi
    done
fi

echo "using $kubectl" 1>&2
$kubectl "$@"
