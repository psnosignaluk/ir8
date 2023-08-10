package api

import (
	"fmt"
	"net"

	"github.com/google/gopacket/pcap"
	"github.com/seancfoley/ipaddress-go/ipaddr"
)

func cidr(mask net.IPMask) *ipaddr.PrefixBitCount {
	var prefix *ipaddr.PrefixBitCount

	if len(mask.String()) == 8 && mask.String() == "ffffff00" {
		// TODO (PS) A bug that needs to be squashed. We need to explode the ffffff00 value returned
		// as the IPv4 netmask into something that ipaddr understands and doesn't throw a seqfault over.
		prefix = ipaddr.NewIPAddressString("255.255.255.0").GetAddress().GetBlockMaskPrefixLen(true)
	} else {
		prefix = ipaddr.NewIPAddressString(mask.String()).GetAddress().GetBlockMaskPrefixLen(true)
	}

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
			fmt.Printf("\tName: %s\n", device.Name)
		}
	} else {
		for _, device := range devices {
			if device.Name == iface {
				fmt.Printf("Interface Name: %s\n", device.Name)
				for _, address := range device.Addresses {
					fmt.Printf("\tIP/Subnet: %s/%s\n", address.IP, cidr(address.Netmask))
				}
			}
		}
	}
}
