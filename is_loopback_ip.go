package util

import "net"

var loopbackCIDRs = []string{
	"127.0.0.0/8", // IPv4 loopback
	"::1/128",     // IPv6 loopback
}

var loopbackNets []*net.IPNet

func init() {
	for _, cidr := range loopbackCIDRs {
		_, block, _ := net.ParseCIDR(cidr)
		loopbackNets = append(loopbackNets, block)
	}
}

func IsLoopbackIP(ip net.IP) bool {
	for _, block := range loopbackNets {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
