package main

import (
	"fmt"
	"net"
	"os"

	"github.com/SteffenHummel/bot_validator/validator"
)

func main() {
	if len(os.Args) > 1 {
		ip := net.ParseIP(os.Args[1])
		fmt.Println("Bot Validator: IP Address: ", ip)
		validator.ValidateBingIpAdresses(ip)
		validator.ValidateGoogleBotIpAdresses(ip)
	} else {
		fmt.Println("Bot Validator: Please provide an IP address as an argument.")
	}
}
