package govirtlib

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

type goVirtError struct {
	name    string
	code    int
	message string
}

// convertLibvirtVersion converts Uint32 version into
// proper major minor macro
// The version is provided as an int following this formula:
// version * 1,000,000 + minor * 1000 + micro
// See src/libvirt-host.c # virConnectGetLibVersion
// https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetVersion
// https://libvirt.org/html/libvirt-libvirt-host.html#virConnectGetLibVersion
func convertLibvirtVersion(inputVer uint32) string {
	var ver int
	major := inputVer / 1000000
	ver %= 1000000
	minor := inputVer / 1000
	ver %= 1000
	micro := ver
	versionString := fmt.Sprintf("%d.%d.%d", major, minor, micro)
	return versionString
}

// stateToStatus func converts Domain state tu human readable format.
//
// see also https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectListAllDomainsFlags
func stateToStatus(state libvirt.DomainState) string {
	var status string
	switch state {
	case libvirt.DOMAIN_NOSTATE:
		status = nostate
	case libvirt.DOMAIN_RUNNING:
		status = running
	case libvirt.DOMAIN_BLOCKED:
		status = blocked
	case libvirt.DOMAIN_PAUSED:
		status = paused
	case libvirt.DOMAIN_SHUTDOWN:
		status = shutdown
	case libvirt.DOMAIN_CRASHED:
		status = crashed
	case libvirt.DOMAIN_PMSUSPENDED:
		status = pmSuspended
	case libvirt.DOMAIN_SHUTOFF:
		status = shutOff
	}
	return status
}

func gError(err goVirtError) string {
	return fmt.Sprintf(
		"govirtlibError(Name=%s, Code=%d, Message='%s')",
		err.name, err.code, err.message)
}
