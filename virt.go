package govirtlib

import (
	"libvirt.org/go/libvirt"
)

// Version ...
type Version struct {
	AppVersion        string `json:"appVersion"`
	HypervisorVersion string `json:"hypervisorVersion"`
	LibvirtVersion    string `json:"libvirtVersion"`
}

// govirtlib implements libvirt's functionality
type govirtlib struct {
	uri        string
	Connection *libvirt.Connect
}

// VMInfo Struct containing all vm Details
type VMInfo struct {
	// ID
	//
	// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetID
	ID uint `json:"id"`
	// UUID
	//
	// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetUUIDString
	UUID string `json:"uuid"`
	// Name
	//
	// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetID
	Name string `json:"name"`
	// State
	//
	// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetState
	State int `json:"state"`
	// Status converted from state into human readable format
	Status string `json:"status"`
}

// NewConnection establishes communication with the specified libvirt
func NewConnection(uri string) (*govirtlib, error) {
	if uri == "" {
		uri = QEMUSystem
	}
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return &govirtlib{uri, nil}, err
	}

	return &govirtlib{uri, conn}, nil
}

// GetVersion func returns Version struct
//
// https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetVersion
//
// https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetLibVersion
func (g *govirtlib) GetVersion() (Version, error) {
	intLibvirtVersion, err := g.Connection.GetLibVersion()
	if err != nil {
		return Version{}, err
	}
	libvirtVersion := convertLibvirtVersion(intLibvirtVersion)
	hypervisorVersion := convertLibvirtVersion(libvirt.VERSION_NUMBER)
	return Version{VERSION, hypervisorVersion, libvirtVersion}, nil
}

// ListAllVM ...
//
// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectListAllDomains
func (g *govirtlib) ListAllVM() ([]VMInfo, error) {

	flags := libvirt.CONNECT_LIST_DOMAINS_ACTIVE
	// libvirt.ConnectListDomainsActive | libvirt.ConnectListDomainsInactive
	domains, err := g.Connection.ListAllDomains(flags)
	if err != nil {
		return nil, err
	}

	vmInfoList := []VMInfo{}

	for _, vm := range domains {

		vmID, err := vm.GetID()
		if err != nil {
			return nil, err
		}

		vmUUID, err := vm.GetUUIDString()
		if err != nil {
			return nil, err
		}

		vmName, err := vm.GetName()
		if err != nil {
			return nil, err
		}

		vmState, reason, err := vm.GetState()
		if err != nil {
			return nil, err
		}

		vmStatus := stateToStatus(vmState)
		vmInfo := VMInfo{
			ID:     vmID,
			UUID:   vmUUID,
			Name:   vmName,
			State:  reason,
			Status: vmStatus,
		}
		vmInfoList = append(vmInfoList, vmInfo)
	}
	return vmInfoList, nil
}
