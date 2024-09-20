package govirtlib

import (
	"errors"
	"fmt"

	"libvirt.org/go/libvirt"
)

// NetworkInfo Struct containing all network details
type NetworkInfo struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// CreateNetwork creates a new network using the provided XML
func (g *Govirtlib) CreateNetwork(xml string, autoStart bool) error {

	network, err := g.Connection.NetworkDefineXML(xml)
	if err != nil {
		return err
	}

	err = network.Create()
	if err != nil {
		return err
	}
	if autoStart {
		err = network.SetAutostart(true)
		if err != nil {
			return err
		}
	}
	return nil
}

// ListNetworks lists all networks
func (g *Govirtlib) ListNetworks() ([]NetworkInfo, error) {
	networks, err := g.Connection.ListAllNetworks(libvirt.CONNECT_LIST_NETWORKS_ACTIVE | libvirt.CONNECT_LIST_NETWORKS_INACTIVE)
	if err != nil {
		return nil, err
	}

	networkInfoList := []NetworkInfo{}
	for _, network := range networks {
		name, err := network.GetName()
		if err != nil {
			return nil, err
		}

		uuid, err := network.GetUUIDString()
		if err != nil {
			return nil, err
		}

		networkInfo := NetworkInfo{
			Name: name,
			UUID: uuid,
		}
		networkInfoList = append(networkInfoList, networkInfo)
	}

	return networkInfoList, nil
}

// DeleteNetwork deletes an existing network by name
func (g *Govirtlib) DeleteNetwork(name string) error {
	// we should not delete "default network"
	DEFAULT_NET := "default"
	if name == DEFAULT_NET {
		msg := fmt.Sprintf("Failed to delete Network: %s", DEFAULT_NET)
		return errors.New(gError(goVirtError{"NetworkDelete", 4, msg}))

	}
	network, err := g.Connection.LookupNetworkByName(name)
	if err != nil {
		return err
	}

	err = network.Destroy()
	if err != nil {
		return err
	}

	err = network.Undefine()
	if err != nil {
		return err
	}

	return nil
}
