package govirtlib

import (
	"libvirt.org/go/libvirt"
)

// NewConnection establishes communication with the specified libvirt
func NewConnection(uri string) (*Govirtlib, error) {
	if uri == "" {
		uri = QEMUSystem
	}
	conn, err := libvirt.NewConnect(uri)
	if err != nil {
		return &Govirtlib{uri, nil}, err
	}

	return &Govirtlib{uri, conn}, nil
}

// GetVersion func returns Version struct
//
// https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetVersion
//
// https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetLibVersion
func (g *Govirtlib) GetVersion() (Version, error) {
	intLibvirtVersion, err := g.Connection.GetLibVersion()
	if err != nil {
		return Version{}, err
	}
	libvirtVersion := convertLibvirtVersion(intLibvirtVersion)
	hypervisorVersion := convertLibvirtVersion(libvirt.VERSION_NUMBER)
	return Version{VERSION, hypervisorVersion, libvirtVersion}, nil
}
