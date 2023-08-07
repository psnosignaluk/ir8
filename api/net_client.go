package api

import (
	"io"
	"net"
	"net/http"

	"github.com/ipinfo/go/v2/ipinfo"
)

func SelfGetPublicIP() (map[string]string, error) {
	m := make(map[string]string)

	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	m["public_ip"] = string(ip)

	return m, nil
}

func GetIPInfoData(token string, ip string) (*ipinfo.Core, error) {
	client := ipinfo.NewClient(nil, nil, token)

	info, err := client.GetIPInfo(net.ParseIP(ip))
	if err != nil {
		return nil, err
	}

	return info, err
}
