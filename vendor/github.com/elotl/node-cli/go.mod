module github.com/elotl/node-cli

go 1.12

require (
	github.com/Microsoft/hcsshim v0.8.6 // indirect
	github.com/cloudflare/cfssl v0.0.0-20180726162950-56268a613adf // indirect
	github.com/codedellemc/goscaleio v0.0.0-20170830184815-20e2ce2cf885 // indirect
	github.com/coreos/rkt v1.30.0 // indirect
	github.com/d2g/dhcp4 v0.0.0-20170904100407-a1d1b6c41b1c // indirect
	github.com/d2g/dhcp4client v0.0.0-20170829104524-6e570ed0a266 // indirect
	github.com/google/certificate-transparency-go v1.0.21 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/heketi/rest v0.0.0-20180404230133-aa6a65207413 // indirect
	github.com/heketi/utils v0.0.0-20170317161834-435bc5bdfa64 // indirect
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jteeuwen/go-bindata v0.0.0-20151023091102-a0ff2567cfb7 // indirect
	github.com/kardianos/osext v0.0.0-20150410034420-8fef92e41e22 // indirect
	github.com/kr/fs v0.0.0-20131111012553-2788f0dbd169 // indirect
	github.com/mholt/caddy v0.0.0-20180213163048-2de495001514 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/pkg/sftp v0.0.0-20160930220758-4d0e916071f6 // indirect
	github.com/sigma/go-inotify v0.0.0-20181102212354-c87b6cf5033d // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.7
	github.com/spf13/pflag v1.0.5
	github.com/virtual-kubelet/virtual-kubelet v1.2.1-0.20200629225228-bd977cb22454
	github.com/vmware/photon-controller-go-sdk v0.0.0-20170310013346-4a435daef6cc // indirect
	github.com/xanzy/go-cloudstack v0.0.0-20160728180336-1e2cbf647e57 // indirect
	go.opencensus.io v0.21.0
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.17.6
	k8s.io/apimachinery v0.17.6
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/klog v1.0.0
	k8s.io/kubectl v0.17.6
	k8s.io/kubernetes v1.17.6
)

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.6

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.6

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.6

replace k8s.io/apiserver => k8s.io/apiserver v0.17.6

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.6

replace k8s.io/cri-api => k8s.io/cri-api v0.17.6

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.6

replace k8s.io/kubelet => k8s.io/kubelet v0.17.6

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.6

replace k8s.io/apimachinery => k8s.io/apimachinery v0.17.6

replace k8s.io/api => k8s.io/api v0.17.6

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.6

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.6

replace k8s.io/component-base => k8s.io/component-base v0.17.6

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.6

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.6

replace k8s.io/metrics => k8s.io/metrics v0.17.6

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.6

replace k8s.io/code-generator => k8s.io/code-generator v0.17.6

replace k8s.io/client-go => k8s.io/client-go v0.17.6

replace k8s.io/kubectl => k8s.io/kubectl v0.17.6
