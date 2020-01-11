package api

func getAddressOfType(t NetworkAddressType, a []NetworkAddress) string {
	for i := 0; i < len(a); i++ {
		if a[i].Type == t {
			return a[i].Address
		}
	}
	return ""
}

func GetPublicIP(a []NetworkAddress) string {
	return getAddressOfType(PublicIP, a)
}

func GetPublicDNS(a []NetworkAddress) string {
	return getAddressOfType(PublicDNS, a)
}

func GetPrivateIP(a []NetworkAddress) string {
	return getAddressOfType(PrivateIP, a)
}

func GetPrivateDNS(a []NetworkAddress) string {
	return getAddressOfType(PrivateDNS, a)
}

func GetPodIP(a []NetworkAddress) string {
	return getAddressOfType(PodIP, a)
}

func NewNetworkAddresses(ip, dns string) []NetworkAddress {
	return []NetworkAddress{
		{
			Type:    PrivateIP,
			Address: ip,
		},
		{
			Type:    PrivateDNS,
			Address: dns,
		},
	}
}

func SetPodIP(ip string, a []NetworkAddress) []NetworkAddress {
	for i := 0; i < len(a); i++ {
		if a[i].Type == PodIP {
			a[i].Address = ip
			return a
		}
	}
	a = append(a, NetworkAddress{
		Type:    PodIP,
		Address: ip,
	})
	return a
}

func SetPrivateDNS(dns string, a []NetworkAddress) []NetworkAddress {
	for i := 0; i < len(a); i++ {
		if a[i].Type == PrivateDNS {
			a[i].Address = dns
			return a
		}
	}
	a = append(a, NetworkAddress{
		Type:    PrivateDNS,
		Address: dns,
	})
	return a
}

func SetPublicAddresses(ip, dns string, a []NetworkAddress) []NetworkAddress {
	needsDNS := len(dns) > 0
	needsIP := len(ip) > 0
	for i := 0; i < len(a); i++ {
		if a[i].Type == PublicIP {
			a[i].Address = ip
			needsIP = false
		} else if a[i].Type == PublicDNS {
			a[i].Address = dns
			needsDNS = false
		}
	}
	if needsIP {
		a = append(a, NetworkAddress{
			Type:    PublicIP,
			Address: ip,
		})
	}
	if needsDNS {
		a = append(a, NetworkAddress{
			Type:    PublicDNS,
			Address: dns,
		})
	}
	return a
}

func DeletePublicAddresses(a []NetworkAddress) []NetworkAddress {
	pruned := make([]NetworkAddress, len(a))
	for i := 0; i < len(a); i++ {
		if a[i].Type != PublicIP && a[i].Type != PublicDNS {
			pruned = append(pruned, a[i])
		}
	}
	return pruned
}

func CopyAddresses(a []NetworkAddress) []NetworkAddress {
	c := make([]NetworkAddress, len(a))
	copy(c, a)
	return c
}
