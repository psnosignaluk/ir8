package api

import (
	"fmt"
	"net"

	"github.com/google/gopacket/pcap"
)

func cidr(mask net.IPMask) int {
	subnet := net.IPMask(mask)
	prefix, _ := subnet.Size()

	return prefix
}

func DeviceList(iface string) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println(err.Error())
	}

	if iface == "all" || iface == "" {
		fmt.Println("Local Interfaces:")
		for _, device := range devices {
			for _, address := range device.Addresses {
				fmt.Printf("\tName: %s\t[%s/%d]\n", device.Name, address.IP, cidr(address.Netmask))
			}
		}
	} else {
		for _, device := range devices {
			if device.Name == iface {
				fmt.Printf("Interface Name: %s\n", device.Name)
				for _, address := range device.Addresses {
					fmt.Printf("\tIP/Subnet: %s/%d\n", address.IP, cidr(address.Netmask))
				}
			}
		}
	}
}

func DetectVPN(iface string) string {
	return iface
}
