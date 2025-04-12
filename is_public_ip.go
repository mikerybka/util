package util

import "net"

func IsPublicIP(ip net.IP) bool {
	return ip != nil && !IsLoopbackIP(ip) && !IsPrivateIP(ip)
}
