package aghnet

import "fmt"

// CheckOtherDHCP tries to discover another DHCP server in the network.
func CheckOtherDHCP(ifaceName string) (ok4, ok6 bool, err4, err6 error) {
	fmt.Println("[->] internal/aghnet/dhcp.go: CheckOtherDHCP()")
	return checkOtherDHCP(ifaceName)
}
