package govirtlib

import (
	"time"

	"libvirt.org/go/libvirt"
)

const (
	// VERSION govirtlib version number
	VERSION string = "0.0.1-dev"
	// QEMUSystem connects to a QEMU system mode daemon
	QEMUSystem string = "qemu:///system"

	// disconnectedTimeout is how long to wait for disconnect cleanup to
	// complete
	disconnectTimeout = 5 * time.Second
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
