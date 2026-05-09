package device

import (
	"context"
	"net"
	"testing"
)

func TestStatusReturnsFirstPrivateIPv4Address(t *testing.T) {
	device := &Device{
		interfaces: func() ([]networkInterface, error) {
			return []networkInterface{
				{
					flags: net.FlagUp | net.FlagLoopback,
					addrs: []net.Addr{
						mustCIDR(t, "127.0.0.1/8"),
					},
				},
				{
					flags: net.FlagUp,
					addrs: []net.Addr{
						mustCIDR(t, "fe80::1/64"),
						mustCIDR(t, "8.8.8.8/24"),
						mustCIDR(t, "192.168.1.42/24"),
					},
				},
			}, nil
		},
	}

	status, err := device.Status(context.Background())
	if err != nil {
		t.Fatalf("Status returned error: %v", err)
	}

	if status != "192.168.1.42" {
		t.Fatalf("status = %q, want %q", status, "192.168.1.42")
	}
}

func TestStatusReturnsErrorWhenNoLANIPAddressExists(t *testing.T) {
	device := &Device{
		interfaces: func() ([]networkInterface, error) {
			return []networkInterface{
				{
					flags: net.FlagUp | net.FlagLoopback,
					addrs: []net.Addr{
						mustCIDR(t, "127.0.0.1/8"),
						mustCIDR(t, "8.8.8.8/24"),
					},
				},
			}, nil
		},
	}

	_, err := device.Status(context.Background())
	if err == nil {
		t.Fatal("Status returned nil error, want missing IP error")
	}
}

func mustCIDR(t *testing.T, value string) net.Addr {
	t.Helper()

	ip, network, err := net.ParseCIDR(value)
	if err != nil {
		t.Fatalf("ParseCIDR(%q) returned error: %v", value, err)
	}

	network.IP = ip
	return network
}
