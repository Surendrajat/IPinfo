package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IpInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Asn      struct {
		Asn    string `json:"asn"`
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Route  string `json:"route"`
		Type   string `json:"type"`
	} `json:"asn"`
	Company struct {
		Name   string `json:"name"`
		Domain string `json:"domain"`
		Type   string `json:"type"`
	} `json:"company"`
	Abuse struct {
		Address string `json:"address"`
		Country string `json:"country"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Network string `json:"network"`
		Phone   string `json:"phone"`
	} `json:"abuse"`
	Vpn struct {
		Any     bool `json:"any"`
		Ipsec   bool `json:"ipsec"`
		Openvpn bool `json:"openvpn"`
		Proxy   bool `json:"proxy"`
		Tor     bool `json:"tor"`
	} `json:"vpn"`
}

func main() {

	var ipInfo IpInfo

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://ipinfo.io/widget", nil)
	req.Header.Set("referer", "https://ipinfo.io/")
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&ipInfo)

	fmt.Printf("IP		%s\n", ipInfo.IP)
	fmt.Printf("Geo		%s, %s, %s\n", ipInfo.City, ipInfo.Region, ipInfo.Country)
	fmt.Printf("ASN		[%s: %s]\n", ipInfo.Asn.Type, ipInfo.Asn.Name)
	fmt.Printf("Company [%s: %s]\n", ipInfo.Company.Type, ipInfo.Company.Name)
}
