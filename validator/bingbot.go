package validator

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/SteffenHummel/bot_validator/utils"
)

type BingBot struct {
	CreationTime string   `json:"creationTime"`
	Prefixes     []Prefix `json:"prefixes"`
}

type Prefix struct {
	Ipv4Prefix string `json:"ipv4Prefix"`
}

func ValidateBingIpAdresses(ipToCheck net.IP) {
	// Download all Bing IP addresses
	body, err := utils.DownloadFile("https://www.bing.com/toolbox/bingbot.json")
	if err != nil {
		fmt.Println("Error downloading bing Bot file: ", err)
		return
	}

	// Unmarshal the JSON data into the BingBot struct
	var bingBot BingBot
	err = json.Unmarshal(body, &bingBot)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return
	}

	for _, prefix := range bingBot.Prefixes {
		_, ipNet, err := net.ParseCIDR(prefix.Ipv4Prefix)
		if err != nil {
			fmt.Println("Error parsing CIDR: ", err)
			return
		}
		if ipNet.Contains(ipToCheck) {
			fmt.Println("Bingbot: true")
			return
		}
	}
	fmt.Println("Bingbot: false")
}
