package govirtlib

import "libvirt.org/go/libvirt"

// Version ...
type Version struct {
	AppVersion        string `json:"appVersion"`
	HypervisorVersion string `json:"hypervisorVersion"`
	LibvirtVersion    string `json:"libvirtVersion"`
}

// Govirtlib implements libvirt's functionality
type Govirtlib struct {
	uri        string
	Connection *libvirt.Connect
}

// VMInfo Struct containing all vm Details
type VMInfo struct {
	// UUID
	// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetUUIDString
	UUID string `json:"uuid"`
	// Name
	// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetID
	Name string `json:"name"`
	// State
	// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainGetState
	State int `json:"state"`
	// Status converted from state into human readable format
	Status string `json:"status"`
}
