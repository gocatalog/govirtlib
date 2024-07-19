/*
Package govirtlib implements libvirt's functionality

TODO: govirtlib Overview

# Example usage

	func printVersion() {
		g, err := govirtlib.NewConnection(govirtlib.QEMUSystem)
		if err != nil {
			log.Fatalf("failed to connect: %v", err)
		}
		version, _ := g.GetVersion()
		b, _ := json.Marshal(version)
		fmt.Println(version)
		fmt.Println(string(b))
	}
*/
package govirtlib
