package eutil

import (
	"errors"
	"net"
)

func GetIP() (string, error) {
	return GetIPv4()
}

func GetIPv4() (string, error) {
	ip, err := getNetIP()
	if err != nil {
		return "", err
	}
	if len(ip) <= 0 {
		return "", errors.New("你好像没有接入局域网？")
	}
	return ip.To4().String(), nil
}

func GetIPv6() (string, error) {
	ip, err := getNetIP()
	if err != nil {
		return "", err
	}
	if len(ip) <= 0 {
		return "", errors.New("你好像没有接入局域网？")
	}
	return ip.To16().String(), nil
}

func getNetIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	// 取所有网卡IP
	var addrs []net.Addr
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err = iface.Addrs()
		if err != nil {
			return nil, err
		}
		break
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip == nil || ip.IsLoopback() {
			continue
		}
		return ip, nil
	}
	return nil, nil
}
