package api

import (
	"io"
	"net"
	"net/http"
	"strconv"

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

func GetIPInfoData(token string, ip string) (map[string]string, error) {
	data := make(map[string]string)

	client := ipinfo.NewClient(nil, nil, token)

	info, err := client.GetIPInfo(net.ParseIP(ip))
	if err != nil {
		return nil, err
	}

	data["country"] = info.Country
	data["countryName"] = info.CountryName
	data["countryFlag"] = info.CountryFlag.Emoji
	data["city"] = info.City
	data["ip"] = info.IP.String()
	data["isEU"] = strconv.FormatBool(info.IsEU)

	return data, err
}
