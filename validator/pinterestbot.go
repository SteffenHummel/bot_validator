package validator

import (
	"fmt"
	"net"
)

// https://help.pinterest.com/en/business/article/pinterest-crawler

func ValidatePinterestIpAdresses(ipToCheck net.IP) {
		_, ipNet, err := net.ParseCIDR("54.236.1.1/24")
		if err != nil {
			fmt.Println("Error parsing CIDR: ", err)
			return
		}
		if ipNet.Contains(ipToCheck) {
			fmt.Println("Pinterest: true")
			return
		}
	fmt.Println("Pinterest: false")
}
