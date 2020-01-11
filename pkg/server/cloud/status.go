package cloud

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/elotl/cloud-instance-provider/pkg/util/timeoutmap"
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	unavailableDuration = 10 * time.Minute
)

var (
	subnetRefreshPeriod = 150 * time.Second // (2.5 minutes)
)

// Status is a structure that details the layout of the cloud
// environment and also includes info on transient state of instance
// availability Various parts of the system need to know what is
// available and where or whether they can create nodes in subnets.
// This is used to figure that out.  The node controller also catches
// errors from node starts and writes availability here.  Reads can
// come from any goroutine, subnets are protected by a mutex
// TimeoutMaps have mutexes built into them.
type StatusKeeper interface {
	Start()
	Stop()
	Dump() []byte
	SupportsAvailabilityZones() bool
	GetAllSubnets() []SubnetAttributes
	GetAvailableZones(instanceType string, spot bool, privateIP bool) []string
	GetAvailableSubnets(instanceType string, spot, privateIP bool) []SubnetAttributes
	AddUnavailableInstance(instanceType string, spot bool)
	AddUnavailableZone(instanceType string, spot bool, zone string)
	AddUnavailableSubnet(instanceType string, spot bool, subnetID string)
	IsUnavailableZone(instanceType string, spot, privateIP bool, az string) bool
}

type subnetPoller struct {
	sync.RWMutex
	client            CloudClient
	stopchan          chan struct{} // mostly (only?) used from tests
	subnets           []SubnetAttributes
	availabilityZones []string
}

type LinkedAZSubnetStatus struct {
	subnetPoller
	unavailableInstances *timeoutmap.TimeoutMap
}

type AZSubnetStatus struct {
	subnetPoller
	unavailableSubnets *timeoutmap.TimeoutMap
	unavailableZones   *timeoutmap.TimeoutMap
}

////////////////////////////////////////////////////////////////////////////////

func subnetSupportsAddressType(addressAffinity SubnetAddressAffinity, privateIP bool) bool {
	return (addressAffinity == AnyAddress) ||
		(privateIP && addressAffinity == PrivateAddress) ||
		(!privateIP && addressAffinity == PublicAddress)
}

func makeUnavailableKey(instanceType string, spot bool, param string) string {
	return fmt.Sprintf("%s/%t/%s", instanceType, spot, param)
}

////////////////////////////////////////////////////////////////////////////////
func (s *subnetPoller) runRefreshLoop() {
	ticker := time.NewTicker(subnetRefreshPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			subnets, err := s.client.GetSubnets()
			if err != nil {
				glog.Errorf("Error refreshing cloud subnet info: %s, continuing with cached data", err)
				continue
			}
			s.Lock()
			s.subnets = subnets
			s.Unlock()

			availabilityZones, err := s.client.GetAvailabilityZones()
			if err != nil {
				glog.Errorf("Error refreshing cloud availability zone info: %s, continuing with cached data", err)
				continue
			}
			s.Lock()
			s.availabilityZones = availabilityZones
			s.Unlock()

		case <-s.stopchan:
			return
		}
	}
}

func (s *subnetPoller) start() {
	s.stopchan = make(chan struct{})
}

func (s *subnetPoller) stop() {
	if s.stopchan != nil {
		s.stopchan <- struct{}{}
		s.stopchan = nil
	}
}

////////////////////////////////////////////////////////////////////////////////

func NewLinkedAZSubnetStatus(client CloudClient) (*LinkedAZSubnetStatus, error) {
	subnets, err := client.GetSubnets()
	if err != nil {
		return nil, err
	}
	if len(subnets) == 0 {
		return nil, fmt.Errorf("No subnets found")
	}

	s := &LinkedAZSubnetStatus{
		subnetPoller: subnetPoller{
			client:  client,
			subnets: subnets,
		},
		unavailableInstances: timeoutmap.New(true, nil),
	}
	return s, nil
}

func (s *LinkedAZSubnetStatus) Start() {
	s.subnetPoller.start()
	go s.unavailableInstances.Start(33 * time.Second)
	go s.runRefreshLoop()
}

func (s *LinkedAZSubnetStatus) Stop() {
	s.subnetPoller.stop()
	s.unavailableInstances.Stop()
}

func (s *LinkedAZSubnetStatus) Dump() []byte {
	s.RLock()
	defer s.RUnlock()

	dumpStruct := struct {
		Subnets     []SubnetAttributes
		Unavailable []string
	}{
		Subnets:     s.GetAllSubnets(),
		Unavailable: s.unavailableInstances.Keys(),
	}
	b, err := json.MarshalIndent(dumpStruct, "", "    ")
	if err != nil {
		glog.Errorln("Error dumping data from cloud.Status", err)
		return nil
	}
	return b
}

func (s *LinkedAZSubnetStatus) SupportsAvailabilityZones() bool {
	return true
}

func (s *LinkedAZSubnetStatus) GetAllSubnets() []SubnetAttributes {
	s.RLock()
	defer s.RUnlock()
	subs := make([]SubnetAttributes, len(s.subnets))
	copy(subs, s.subnets)
	return subs
}

// Used in validation to figure out if a pod has a valid spec
// -- only valid in Amazon Cloud
func (s *LinkedAZSubnetStatus) GetAllAZSubnets(zone string, privateIP bool) []SubnetAttributes {
	s.RLock()
	defer s.RUnlock()
	sn := make([]SubnetAttributes, 0, 1)
	for i, _ := range s.subnets {
		if zone == s.subnets[i].AZ && subnetSupportsAddressType(s.subnets[i].AddressAffinity, privateIP) {
			sn = append(sn, s.subnets[i])
		}
	}
	return sn
}

// Only used in AWS, not part of the StatusKeeper interface
func (s *LinkedAZSubnetStatus) GetAvailableAZSubnets(instanceType, zone string, spot, privateIP bool) []SubnetAttributes {
	allSubnets := s.GetAllAZSubnets(zone, privateIP)
	availableSubnets := make([]SubnetAttributes, 0, len(allSubnets))
	for i, _ := range allSubnets {
		if !s.IsUnavailableSubnet(instanceType, spot, allSubnets[i].ID) &&
			allSubnets[i].AvailableAddresses > 0 {
			availableSubnets = append(availableSubnets, allSubnets[i])
		}
	}
	return availableSubnets
}

func (s *LinkedAZSubnetStatus) GetAvailableZones(instanceType string, spot bool, privateIP bool) []string {
	s.RLock()
	defer s.RUnlock()
	azset := sets.NewString()
	for i, _ := range s.subnets {
		if subnetSupportsAddressType(s.subnets[i].AddressAffinity, privateIP) &&
			!s.IsUnavailableSubnet(instanceType, spot, s.subnets[i].ID) &&
			s.subnets[i].AvailableAddresses > 0 {
			azset.Insert(s.subnets[i].AZ)
		}
	}
	return azset.List()
}

func (s *LinkedAZSubnetStatus) GetAvailableSubnets(instanceType string, spot, privateIP bool) []SubnetAttributes {
	s.RLock()
	defer s.RUnlock()
	sn := make([]SubnetAttributes, 0, len(s.subnets))
	for i, _ := range s.subnets {
		if subnetSupportsAddressType(s.subnets[i].AddressAffinity, privateIP) &&
			!s.IsUnavailableSubnet(instanceType, spot, s.subnets[i].ID) &&
			s.subnets[i].AvailableAddresses > 0 {
			sn = append(sn, s.subnets[i])
		}
	}
	return sn
}

func (s *LinkedAZSubnetStatus) AddUnavailableInstance(instanceType string, spot bool) {
	s.RLock()
	defer s.RUnlock()
	for i, _ := range s.subnets {
		s.AddUnavailableSubnet(instanceType, spot, s.subnets[i].ID)
	}
}

func (s *LinkedAZSubnetStatus) AddUnavailableZone(instanceType string, spot bool, zone string) {
	glog.Infof("Adding unavailable zone %s for instance type %s", zone, instanceType)
	s.RLock()
	defer s.RUnlock()
	for i, _ := range s.subnets {
		if zone == s.subnets[i].AZ {
			s.AddUnavailableSubnet(instanceType, spot, s.subnets[i].ID)
		}
	}
}

func (s *LinkedAZSubnetStatus) AddUnavailableSubnet(instanceType string, spot bool, subnetID string) {
	glog.Infof("Adding unavailable subnet %s for instance type %s", subnetID, instanceType)
	key := makeUnavailableKey(instanceType, spot, subnetID)
	// only update the entry if it doesn't already exist.  It might be
	// tempting to always update the object but that could lead to a
	// situation where we are permanently updating the unavailability
	// of the subnet.  For example, if we try to boot an instance of a
	// particular type and lookup a subnet for that instanceType, and
	// see that they're all unavailable.  We then throw a
	// NoCapacityError.  That leads to updating the available
	// timestamp of the instance, we then try to boot a node for it
	// and the cycle continues.
	_, exists := s.unavailableInstances.Get(key)
	if !exists {
		s.unavailableInstances.Add(key, struct{}{}, unavailableDuration, timeoutmap.Noop)
	}
}

func (s *LinkedAZSubnetStatus) IsUnavailableSubnet(instanceType string, spot bool, subnetID string) bool {
	key := makeUnavailableKey(instanceType, spot, subnetID)
	_, exists := s.unavailableInstances.Get(key)
	return exists
}

func (s *LinkedAZSubnetStatus) IsUnavailableZone(instanceType string, spot, privateIP bool, az string) bool {
	subnets := s.GetAvailableSubnets(instanceType, spot, privateIP)
	if az == "" {
		return len(subnets) == 0
	}
	for i := range subnets {
		if subnets[i].AZ == az {
			return false
		}
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////

func NewAZSubnetStatus(client CloudClient) (*AZSubnetStatus, error) {
	subnets, err := client.GetSubnets()
	if err != nil {
		return nil, err
	}
	if len(subnets) == 0 {
		return nil, fmt.Errorf("No subnets found")
	}
	azs, err := client.GetAvailabilityZones()
	if err != nil {
		return nil, err
	}
	s := &AZSubnetStatus{
		subnetPoller: subnetPoller{
			client:            client,
			subnets:           subnets,
			availabilityZones: azs,
		},
		unavailableSubnets: timeoutmap.New(true, nil),
		unavailableZones:   timeoutmap.New(true, nil),
	}
	return s, nil
}

func (s *AZSubnetStatus) Start() {
	s.subnetPoller.start()
	go s.unavailableSubnets.Start(33 * time.Second)
	go s.unavailableZones.Start(33 * time.Second)
	go s.runRefreshLoop()
}

func (s *AZSubnetStatus) Stop() {
	s.subnetPoller.stop()
	s.unavailableSubnets.Stop()
	s.unavailableZones.Stop()
}

func (s *AZSubnetStatus) Dump() []byte {
	s.RLock()
	defer s.RUnlock()

	dumpStruct := struct {
		Subnets            []SubnetAttributes
		AvailabilityZones  []string
		UnavailableSubnets []string
		UnavailableZones   []string
	}{
		Subnets:            s.GetAllSubnets(),
		AvailabilityZones:  s.availabilityZones,
		UnavailableSubnets: s.unavailableSubnets.Keys(),
		UnavailableZones:   s.unavailableZones.Keys(),
	}
	b, err := json.MarshalIndent(dumpStruct, "", "    ")
	if err != nil {
		glog.Errorln("Error dumping data from cloud.Status", err)
		return nil
	}
	return b
}

func (s *AZSubnetStatus) GetAllSubnets() []SubnetAttributes {
	s.RLock()
	defer s.RUnlock()
	subs := make([]SubnetAttributes, len(s.subnets))
	copy(subs, s.subnets)
	return subs
}

func (s *AZSubnetStatus) GetAllAvailabilityZones() []string {
	s.RLock()
	defer s.RUnlock()
	azs := make([]string, len(s.availabilityZones))
	copy(azs, s.availabilityZones)
	return azs
}

func (s *AZSubnetStatus) SupportsAvailabilityZones() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.availabilityZones) > 0
}

func (s *AZSubnetStatus) GetAvailableZones(instanceType string, spot bool, privateIP bool) []string {
	s.RLock()
	defer s.RUnlock()
	azs := make([]string, 0, len(s.availabilityZones))
	for i, _ := range s.availabilityZones {
		if !s.IsUnavailableZone(instanceType, spot, privateIP, s.availabilityZones[i]) {
			azs = append(azs, s.availabilityZones[i])
		}
	}
	return azs
}

func (s *AZSubnetStatus) GetAvailableSubnets(instanceType string, spot, privateIP bool) []SubnetAttributes {
	s.RLock()
	defer s.RUnlock()
	sn := make([]SubnetAttributes, 0, len(s.subnets))
	for i, _ := range s.subnets {
		if subnetSupportsAddressType(s.subnets[i].AddressAffinity, privateIP) &&
			!s.IsUnavailableSubnet(instanceType, spot, s.subnets[i].ID) &&
			s.subnets[i].AvailableAddresses > 0 {
			sn = append(sn, s.subnets[i])
		}
	}
	return sn
}

func (s *AZSubnetStatus) AddUnavailableZone(instanceType string, spot bool, zone string) {
	glog.Infof("Adding unavailable zone %s for instance type %s", zone, instanceType)
	key := makeUnavailableKey(instanceType, spot, zone)
	_, exists := s.unavailableZones.Get(key)
	if !exists {
		s.unavailableZones.Add(key, struct{}{}, unavailableDuration, timeoutmap.Noop)
	}
}

func (s *AZSubnetStatus) AddUnavailableSubnet(instanceType string, spot bool, subnetID string) {
	glog.Infof("Adding unavailable subnet %s for instance type %s", subnetID, instanceType)
	key := makeUnavailableKey(instanceType, spot, subnetID)
	// only update the entry if it doesn't already exist.  It might be
	// tempting to always update the object but that could lead to a
	// situation where we are permanently updating the unavailability
	// of the subnet.  For example, if we try to boot an instance of a
	// particular type and lookup a subnet for that instanceType, and
	// see that they're all unavailable.  We then throw a
	// NoCapacityError.  That leads to updating the available
	// timestamp of the instance, we then try to boot a node for it
	// and the cycle continues.
	_, exists := s.unavailableSubnets.Get(key)
	if !exists {
		s.unavailableSubnets.Add(key, struct{}{}, unavailableDuration, timeoutmap.Noop)
	}
}

func (s *AZSubnetStatus) AddUnavailableInstance(instanceType string, spot bool) {
	for i, _ := range s.subnets {
		s.AddUnavailableSubnet(instanceType, spot, s.subnets[i].ID)
	}
	for i, _ := range s.availabilityZones {
		s.AddUnavailableZone(instanceType, spot, s.availabilityZones[i])
	}
}

func (s *AZSubnetStatus) IsUnavailableSubnet(instanceType string, spot bool, subnetID string) bool {
	key := makeUnavailableKey(instanceType, spot, subnetID)
	_, exists := s.unavailableSubnets.Get(key)
	return exists
}

// In
func (s *AZSubnetStatus) IsUnavailableZone(instanceType string, spot, privateIP bool, az string) bool {
	key := makeUnavailableKey(instanceType, spot, az)
	_, exists := s.unavailableZones.Get(key)
	return exists
}
