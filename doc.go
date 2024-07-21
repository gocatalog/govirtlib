/*
Package govirtlib implements libvirt's functionality

TODO: govirtlib Overview

# Example usage

	func printVersion() {
		g, err := govirtlib.NewConnection(govirtlib.QEMUSystem)
		if err != nil {
			log.Fatalf("failed to connect: %v", err)
		}
		defer g.Connection.Close()
		version, _ := g.GetVersion()
		b, _ := json.Marshal(version)
		fmt.Println(version)
		fmt.Println(string(b))
	}

# List All vm

	func printVm() {
		g, err := govirtlib.NewConnection(govirtlib.QEMUSystem)
		if err != nil {
		log.Fatalf("failed to connect: %v", err)
		}
		defer g.Connection.Close()
		vmInfoList,_ := g.ListAllVM()
		b, _ := json.Marshal(vmInfoList)
		fmt.Println(vmInfoList)
		fmt.Println(string(b))
	}

# Start vm

	func startvm() {
		g, err := govirtlib.NewConnection(govirtlib.QEMUSystem)
		if err != nil {
		log.Fatalf("failed to connect: %v", err)
		}
		vmName:= "cirros"
		domain , _ := g.GetVM(govirtlib.VMOptName, vmName)
		g.VMPowerOn(domain)
	}

# Stop vm

	func stopvm() {
		g, err := govirtlib.NewConnection(govirtlib.QEMUSystem)
		if err != nil {
		log.Fatalf("failed to connect: %v", err)
		}
		vmName:= "cirros"
		domain , _ := g.GetVM(govirtlib.VMOptName, vmName)
		g.VMPowerOff(domain, false)
	}

# Toggle vm

	func ToggleVm() {
		g, err := govirtlib.NewConnection(govirtlib.QEMUSystem)
		if err != nil {
		log.Fatalf("failed to connect: %v", err)
		}
		vmName:= "cirros"
		domain , _ := g.GetVM(govirtlib.VMOptName, vmName)
		g.VMToggle(domain)
	}
*/
package govirtlib
