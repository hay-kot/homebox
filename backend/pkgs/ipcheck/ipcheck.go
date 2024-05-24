// Package ipcheck provides helper functions to validate IP addresses against criteria
package ipcheck

import (
	"fmt"
	"net"
	"strings"
)

func ValidateAgainstList(ip string, comaSeparatedList string) bool {
	if comaSeparatedList == "" || ip == "" {
		return false
	}

	if net.ParseIP(ip) == nil {
		ip, _, _ = net.SplitHostPort(ip)
	}

	if ip == "" {
		return false
	}

	cidrs := strings.Split(comaSeparatedList, ",")
	testedIP, _, err := net.ParseCIDR(fmt.Sprintf("%s/32", ip))
	if err != nil {
		return false
	}

	for _, cidr := range cidrs {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err == nil && ipnet.Contains(testedIP) {
			return true
		}
	}

	return false
}
