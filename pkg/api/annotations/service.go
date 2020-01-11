package annotations

// https://kubernetes.io/docs/concepts/services-networking/service/#internal-load-balancer

// ServiceLoadBalancerAccessLogEmitInterval is the annotation used to
// specify access log emit interval.
const ServiceLoadBalancerAccessLogEmitInterval = "service.elotl.co/aws-load-balancer-access-log-emit-interval"

// ServiceLoadBalancerAccessLogEnabled is the annotation used on the
// service to enable or disable access logs.
const ServiceLoadBalancerAccessLogEnabled = "service.elotl.co/aws-load-balancer-access-log-enabled"

// ServiceLoadBalancerAccessLogS3BucketName is the annotation used to
// specify access log s3 bucket name.
const ServiceLoadBalancerAccessLogS3BucketName = "service.elotl.co/aws-load-balancer-access-log-s3-bucket-name"

// ServiceLoadBalancerAccessLogS3BucketPrefix is the annotation used
// to specify access log s3 bucket prefix.
const ServiceLoadBalancerAccessLogS3BucketPrefix = "service.elotl.co/aws-load-balancer-access-log-s3-bucket-prefix"

// ServiceLoadBalancerConnectionDrainingEnabled is the annnotation
// used on the service to enable or disable connection draining.
const ServiceLoadBalancerConnectionDrainingEnabled = "service.elotl.co/aws-load-balancer-connection-draining-enabled"

// ServiceLoadBalancerConnectionDrainingTimeout is the annotation
// used on the service to specify a connection draining timeout.
const ServiceLoadBalancerConnectionDrainingTimeout = "service.elotl.co/aws-load-balancer-connection-draining-timeout"

// ServiceLoadBalancerConnectionIdleTimeout is the annotation used
// on the service to specify the idle connection timeout.
// Must be > 1
const ServiceLoadBalancerConnectionIdleTimeout = "service.elotl.co/aws-load-balancer-connection-idle-timeout"

// ServiceLoadBalancerCrossZoneLoadBalancingEnabled is the annotation
// used on the service to enable or disable cross-zone load balancing.
const ServiceLoadBalancerCrossZoneLoadBalancingEnabled = "service.elotl.co/aws-load-balancer-cross-zone-load-balancing-enabled"

// ServiceLoadBalancerCertificate is the annotation used on the
// service to request a secure listener. Value is a valid certificate ARN.
// For more, see http://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/elb-listener-config.html
// CertARN is an IAM or CM certificate ARN, e.g. arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012
const ServiceLoadBalancerCertificate = "service.elotl.co/aws-load-balancer-ssl-cert"

// ServiceLoadBalancerSSLPorts is the annotation used on the service
// to specify a comma-separated list of ports that will use SSL/HTTPS
// listeners. Defaults to '*' (all).
const ServiceLoadBalancerSSLPorts = "service.elotl.co/aws-load-balancer-ssl-ports"

// ServiceLoadBalancerBEProtocol is the annotation used on the service
// to specify the protocol spoken by the backend (pod) behind a secure listener.
// Only inspected when `aws-load-balancer-ssl-cert` is used.
// If `http` (default) or `https`, an HTTPS listener that terminates the
//  connection and parses headers is created.
// If set to `ssl` or `tcp`, a "raw" SSL listener is used.
const ServiceLoadBalancerBEProtocol = "service.elotl.co/aws-load-balancer-backend-protocol"

// ServiceLoadBalancerHCHealthyThreshold is the annotation used on the
// service to specify the number of successive successful health
// checks required for a backend to be considered healthy for
// traffic. AWS requires this value to be >= 2 && <= 10.
const ServiceLoadBalancerHCHealthyThreshold = "service.elotl.co/aws-load-balancer-healthcheck-healthy-threshold"

// ServiceLoadBalancerHCUnhealthyThreshold is the annotation used on
// the service to specify the number of unsuccessful health checks
// required for a backend to be considered unhealthy for traffic.  AWS
// requires this value to be >= 2 && <= 10.
const ServiceLoadBalancerHCUnhealthyThreshold = "service.elotl.co/aws-load-balancer-healthcheck-unhealthy-threshold"

// ServiceLoadBalancerHCTimeout is the annotation used on the
// service to specify, in seconds, how long to wait before marking a health
// check as failed.  AWS requires this value to be between 2 and 60 inclusive.
const ServiceLoadBalancerHCTimeout = "service.elotl.co/aws-load-balancer-healthcheck-timeout"

// ServiceLoadBalancerHCInterval is the annotation used on the
// service to specify, in seconds, the interval between health checks.
// AWS requires this value to be in the range [5, 300].
const ServiceLoadBalancerHCInterval = "service.elotl.co/aws-load-balancer-healthcheck-interval"

// ServiceLoadBalancerProxyProtocol is the annotation used on the
// service to enable the proxy protocol on an ELB. Right now we only accept the
// value "*" which means enable the proxy protocol on all ELB backends. In the
// future we could adjust this to allow setting the proxy protocol only on
// certain backends.
//const ServiceLoadBalancerProxyProtocol = "service.elotl.co/aws-load-balancer-proxy-protocol"

// Idle timeout for connections on the load balancer in minutes.
const ServiceAnnotationLoadBalancerIdleTimeout = "service.elotl.co/azure-load-balancer-tcp-idle-timeout"

// Probe interval for checking backends, in seconds.
const ServiceAnnotationLoadBalancerProbeInterval = "service.elotl.co/azure-load-balancer-probe-interval"

// Number of probes for checking backends.
const ServiceAnnotationLoadBalancerNumberOfProbes = "service.elotl.co/azure-load-balancer-number-of-probes"
