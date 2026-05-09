package device

import (
	"context"
	"errors"
	"fmt"
	"net"
)

type Device struct {
	interfaces func() ([]networkInterface, error)
}

type networkInterface struct {
	flags net.Flags
	addrs []net.Addr
}

var errLANIPAddressUnavailable = errors.New("lan ip address unavailable")

func New() *Device {
	return &Device{interfaces: systemNetworkInterfaces}
}

func (d *Device) Status(context.Context) (string, error) {
	interfaces := d.interfaces
	if interfaces == nil {
		interfaces = systemNetworkInterfaces
	}

	networkInterfaces, err := interfaces()
	if err != nil {
		return "", fmt.Errorf("list network interfaces: %w", err)
	}

	return lanIPAddress(networkInterfaces)
}

func systemNetworkInterfaces() ([]networkInterface, error) {
	systemInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	networkInterfaces := make([]networkInterface, 0, len(systemInterfaces))
	for _, systemInterface := range systemInterfaces {
		addrs, err := systemInterface.Addrs()
		if err != nil {
			continue
		}

		networkInterfaces = append(networkInterfaces, networkInterface{
			flags: systemInterface.Flags,
			addrs: addrs,
		})
	}

	return networkInterfaces, nil
}

func lanIPAddress(interfaces []networkInterface) (string, error) {
	for _, iface := range interfaces {
		if !activeLANInterface(iface.flags) {
			continue
		}

		if ip := ipv4Address(iface.addrs); ip != "" {
			return ip, nil
		}
	}

	return "", errLANIPAddressUnavailable
}

func activeLANInterface(flags net.Flags) bool {
	return flags&net.FlagUp != 0 && flags&net.FlagLoopback == 0
}

func ipv4Address(addrs []net.Addr) string {
	for _, addr := range addrs {
		ip := addrIP(addr)
		if ip == nil || ip.IsLoopback() {
			continue
		}

		if ipv4 := ip.To4(); ipv4 != nil && ipv4.IsPrivate() {
			return ipv4.String()
		}
	}

	return ""
}

func addrIP(addr net.Addr) net.IP {
	switch value := addr.(type) {
	case *net.IPNet:
		return value.IP
	case *net.IPAddr:
		return value.IP
	default:
		return nil
	}
}
