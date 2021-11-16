package main

import (
	"fmt"
	"net"
	"time"

	"github.com/Haba1234/go-artnet"
)

func main() {
	artSubnet := "192.168.6.0/24"
	_, cidrNet, _ := net.ParseCIDR(artSubnet)

	address, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("error getting ips: %s\n", err)
	}

	var ip net.IP

	for _, addr := range address {
		ip = addr.(*net.IPNet).IP
		if cidrNet.Contains(ip) {
			break
		}
	}

	log := artnet.NewDefaultLogger("debug")
	c := artnet.NewController("controller-1", ip, log)
	c.Start()

	go func() {
		time.Sleep(2 * time.Second)
		c.SendDMXToAddress([512]byte{0x00, 0xff, 0x00, 0xff, 0x00}, artnet.Address{Net: 0, SubUni: 1})
		time.Sleep(2 * time.Second)
		c.SendDMXToAddress([512]byte{0xff, 0x00, 0x00, 0x00, 0xff}, artnet.Address{Net: 0, SubUni: 1})
		time.Sleep(2 * time.Second)
		c.SendDMXToAddress([512]byte{0x00, 0x00, 0xff, 0xff, 0x00}, artnet.Address{Net: 0, SubUni: 1})
		time.Sleep(2 * time.Second)
		c.SendDMXToAddress([512]byte{}, artnet.Address{Net: 0, SubUni: 1})
		time.Sleep(2 * time.Second)
	}()

	for {
		time.Sleep(time.Second)
	}
}
