package govirtlib

import "fmt"

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
