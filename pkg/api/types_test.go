package api

import "testing"

func TestComparePorts(t *testing.T) {
	tests := []struct {
		a    ServicePort
		b    ServicePort
		less bool
	}{
		{
			a:    ServicePort{Port: 5},
			b:    ServicePort{Port: 7},
			less: true,
		},
		{
			a:    ServicePort{Port: 7},
			b:    ServicePort{Port: 5},
			less: false,
		},
		{
			a:    ServicePort{Port: 5, PortRangeSize: 1, Protocol: ProtocolUDP},
			b:    ServicePort{Port: 5, PortRangeSize: 3, Protocol: ProtocolTCP},
			less: true,
		},
		{
			a:    ServicePort{Port: 5, PortRangeSize: 4, Protocol: ProtocolUDP},
			b:    ServicePort{Port: 5, PortRangeSize: 3, Protocol: ProtocolTCP},
			less: false,
		},
		{
			a:    ServicePort{Port: 5, PortRangeSize: 3, Protocol: ProtocolUDP},
			b:    ServicePort{Port: 5, PortRangeSize: 3, Protocol: ProtocolTCP},
			less: false,
		},
	}
	for i, test := range tests {
		if lessPorts(test.a, test.b) != test.less {
			t.Errorf("Test %d failed: %v", i, test)
		}
	}
}
