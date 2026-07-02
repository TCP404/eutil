package eutil

import (
	"net"
	"testing"
)

func TestSelectLocalIPv4PrefersPrivateAddress(t *testing.T) {
	addrs := []net.Addr{
		mustIPNet(t, "192.18.2.3/24"),
		mustIPNet(t, "192.168.2.3/24"),
	}

	got, err := selectLocalIPv4(addrs)
	if err != nil {
		t.Fatalf("selectLocalIPv4 returned error: %v", err)
	}
	if got != "192.168.2.3" {
		t.Fatalf("selectLocalIPv4 = %q, want 192.168.2.3", got)
	}
}

func mustIPNet(t *testing.T, cidr string) *net.IPNet {
	t.Helper()
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		t.Fatalf("ParseCIDR(%q): %v", cidr, err)
	}
	ipNet.IP = ip
	return ipNet
}
