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

func getNetIP() (ip net.IP, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP, nil
			}
		}
	}
	return
}
