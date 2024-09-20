package govirtlib

import (
	"errors"
	"fmt"
	"log"

	"libvirt.org/go/libvirt"
)

// ListAllVM ...
// See also https://libvirt.org/html/libvirt-libvirt-domain.html#virConnectListAllDomains
func (g *Govirtlib) ListAllVM() ([]VMInfo, error) {

	// List all active domains (on VMs)
	activeDomains, err := g.Connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		return nil, err
	}

	// List all inactive domains (shut off VMs)
	inactiveDomains, err := g.Connection.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		return nil, err
	}

	domains := append(activeDomains, inactiveDomains...)
	vmInfoList := []VMInfo{}

	for _, vm := range domains {

		vmUUID, err := vm.GetUUIDString()
		if err != nil {
			return nil, err
		}

		vmName, err := vm.GetName()
		if err != nil {
			return nil, err
		}

		vmState, _, err := vm.GetState()
		if err != nil {
			return nil, err
		}
		vmStatus := stateToStatus(vmState)
		vmInfo := VMInfo{
			UUID:   vmUUID,
			Name:   vmName,
			State:  int(vmState),
			Status: vmStatus,
		}
		vmInfoList = append(vmInfoList, vmInfo)
	}
	return vmInfoList, nil
}

// GetVM func
func (g *Govirtlib) GetVM(by vmOption, val string) (*libvirt.Domain, error) {

	var domain *libvirt.Domain
	var err error

	switch by {
	case VMOptName:
		domain, err = g.Connection.LookupDomainByName(val)
		if err != nil {
			return nil, err
		}
	case VMOptUUID:
		domain, err = g.Connection.LookupDomainByUUIDString(val)
		if err != nil {
			return nil, err
		}
	}
	return domain, nil
}

// VMPowerOff func will start the vm . it will either take Vm Name or Vm uuid
func (g *Govirtlib) VMPowerOff(domain *libvirt.Domain, force bool) error {
	defer func() {
		if err := domain.Free(); err != nil {
			log.Fatalf("Error freeing domain: %v", err)
		}
	}()

	// Shutdown the VM gracefully
	err := domain.Shutdown()

	// if shutdown success
	if err == nil {
		return nil
	}

	// if shutdown is not success and force is false then return error
	// if force is false then return
	if !force && err != nil {
		msg := fmt.Sprintf("Failed to shutdown domain gracefully: %v", err)
		return errors.New(gError(goVirtError{"ShutDownError", 2, msg}))
	}

	// try with force
	// If shutdown fails, forcefully destroy the VM
	err = domain.Destroy()
	if err != nil {
		msg := fmt.Sprintf("Failed to destroy domain: %v", err)
		return errors.New(gError(goVirtError{"ShutDownError", 3, msg}))
	}
	// destroy success
	return nil
}

// VMPowerOn ...
func (g *Govirtlib) VMPowerOn(domain *libvirt.Domain) error {
	defer func() {
		if err := domain.Free(); err != nil {
			log.Fatalf("Error freeing domain: %v", err)
		}
	}()

	// Start the VM
	err := domain.Create()
	if err != nil {
		return err
	}
	return nil
}

// VMToggle func will toggle between on and off state of vm, will start
// if the vm is off or else off if vm is in on state.
// se also https://libvirt.org/html/libvirt-libvirt-domain.html#virDomainState
func (g *Govirtlib) VMToggle(domain *libvirt.Domain) error {
	vmState, _, err := domain.GetState()
	if err != nil {
		return err
	}
	switch vmState {
	case 1:
		// 1 is running
		return g.VMPowerOff(domain, false)
	case 5, 6, 7:
		// 5 is off
		return g.VMPowerOn(domain)
	}
	return nil
}
