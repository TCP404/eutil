package eutil

import (
	"errors"
	"net"
)

func GetLocalIP() (string, error) {
	return GetLocalIPv4()
}

func GetLocalIPv4() (string, error) {
	addrs, err := getAddrs()
	if err != nil {
		return "", errors.New("获取IP失败: " + err.Error())
	}
	for _, addr := range addrs {
		ip := getLocalIP(addr)
		if ip == nil {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}
		return ip.String(), nil
	}
	return "", errors.New("获取 IPv4 地址失败")
}

func GetLocalIPv6() (string, error) {
	addrs, err := getAddrs()
	if err != nil {
		return "", errors.New("获取IP失败: " + err.Error())
	}
	for _, addr := range addrs {
		ip := getLocalIP(addr)
		if ip == nil {
			continue
		}
		if ip = ip.To16(); ip == nil {
			continue // not an ipv6 address
		}
		return ip.String(), nil
	}
	return "", errors.New("获取 IPv6 地址失败")
}

func getAddrs() ([]net.Addr, error) {
	var addrss []net.Addr
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // 网卡没有开启
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // 这是个环回地址
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		addrss = append(addrss, addrs...)
	}
	if len(addrss) <= 0 {
		return nil, errors.New("你好像没有接入局域网？")
	}
	return addrss, nil
}

func getLocalIP(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	default:
		return nil
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	return ip
}
