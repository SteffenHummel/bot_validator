package validator

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/SteffenHummel/bot_validator/utils"
)

type GoogleBotPrefix struct {
	Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
	Ipv4Prefix string `json:"ipv4Prefix,omitempty"`
}

type GoogleBot struct {
	CreationTime string            `json:"creationTime"`
	Prefixes     []GoogleBotPrefix `json:"prefixes"`
}

func ValidateGoogleBotIpAdresses(ipToCheck net.IP) {
	// Download all Bing IP addresses
	body, err := utils.DownloadFile("https://developers.google.com/search/apis/ipranges/googlebot.json")

	if err != nil {
		fmt.Println("Error downloading file for GoogleBot ", err)
		return
	}

	var googleBot GoogleBot
	err = json.Unmarshal(body, &googleBot)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return
	}

	bodySpecialCrawlers, err := utils.DownloadFile("https://developers.google.com/search/apis/ipranges/special-crawlers.json")
	if err != nil {
		fmt.Println("Error downloading file for GoogleBot Special Crawlers ", err)
		return
	}

	var googleBotSpecialCrawlers GoogleBot
	err = json.Unmarshal(bodySpecialCrawlers, &googleBotSpecialCrawlers)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return
	}

	bodyUserTriggerdCrawlers, err := utils.DownloadFile("https://developers.google.com/search/apis/ipranges/user-triggered-fetchers.json")
	if err != nil {
		fmt.Println("Error downloading file for GoogleBot Special Crawlers ", err)
		return
	}

	var googleBotUserTriggeredCrawlers GoogleBot
	err = json.Unmarshal(bodyUserTriggerdCrawlers, &googleBotUserTriggeredCrawlers)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return
	}

	googleBot.Prefixes = append(googleBot.Prefixes, googleBotSpecialCrawlers.Prefixes...)
	googleBot.Prefixes = append(googleBot.Prefixes, googleBotUserTriggeredCrawlers.Prefixes...)

	for _, prefix := range googleBot.Prefixes {
		if prefix.Ipv4Prefix != "" {
			_, ipNet, err := net.ParseCIDR(prefix.Ipv4Prefix)
			if err != nil {
				fmt.Println("Error parsing CIDR: ", err)
				return
			}

			if ipNet.Contains(ipToCheck) {
				fmt.Println("GoogleBot: true")
				return
			}
		}
	}
	fmt.Println("GoogleBot: false")
}
