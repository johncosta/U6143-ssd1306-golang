package ipaddr

import (
	"fmt"
	"net"
)

type Address struct{}

// GetDisplayValueForInterface returns a string value for the specifived an ipv4 interface, even if one isn't available.
func (a Address) GetDisplayValueForInterface(interfaceName string) string {
	ipv4Addr, err := a.InterfaceAddressByName(interfaceName)
	if err != nil {
		return "xx.xx.xx.xx"
	}
	return ipv4Addr.String()
}

// InterfaceAddressByName returns the ipv4 Interface for the name specified.
func (a Address) InterfaceAddressByName(interfaceName string) (ipv4Addr net.IP, err error) {
	ief, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return nil, err
	}

	addrs, err := ief.Addrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			return ipv4Addr, nil
		}
	}

	return nil, fmt.Errorf("no address found for interface: %s", interfaceName)
}
