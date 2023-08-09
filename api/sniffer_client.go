package api

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

func DeviceList() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, device := range devices {
		if device.Name == "en0" {
			fmt.Printf("Name: %s\n", device.Name)
			for _, address := range device.Addresses {
				fmt.Printf("\tIP Address/Subnet: %s/%s\n", address.IP, address.Netmask)
			}
		}
	}
}
