package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"github.com/yangl900/nethealth/logs"
)

func main() {
	TestLocalNetwork()
	TestDNS()
	TestOutboundUDPDial()
	TestOutboundHTTP()
	TestOutboundIPAddress()
}

// TestDNS tests DNS by resolve bing.com
func TestDNS() {
	logs.StartTest("DNS")
	addrs, err := net.LookupHost("bing.com")
	if err != nil || len(addrs) == 0 {
		logs.Fail("DNS lookup error." + err.Error())
		return
	}

	logs.Succeed("Resolved bing.com to " + strings.Join(addrs, ";"))
}

// TestOutboundHTTP tests outbound http connection
func TestOutboundHTTP() {
	logs.StartTest("OutboundHTTP")
	resp, err := http.Get("http://bing.com/")
	if err != nil {
		logs.Fail("Outbound http connection error: " + err.Error())
		return
	}

	logs.Succeed(fmt.Sprintf("Received status %d and %d bytes from http://www.bing.com", resp.StatusCode, resp.ContentLength))
}

// TestOutboundUDPDial tests outbound UDP connection
func TestOutboundUDPDial() {
	logs.StartTest("OutboundUDPDial")
	remoteIP := "1.1.1.1:53"
	conn, err := net.Dial("udp", remoteIP)

	if err != nil {
		logs.Fail("Outbound UDP connection error: " + err.Error())
		return
	}

	defer conn.Close()
	logs.Succeed(fmt.Sprintf("Connection established between %s and localhost. Local: %s. Remote: %s.", remoteIP, conn.LocalAddr(), conn.RemoteAddr()))
}

// TestOutboundIPAddress detects the outbound IP address
func TestOutboundIPAddress() {
	logs.StartTest("OutboundIPAddress")
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		logs.Info("Failed to detect outbound IP address: " + err.Error())
		return
	}

	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logs.Info("Failed to detect outbound IP address: " + err.Error())
		return
	}

	logs.Info("Outbound IP address: " + string(ip))
}

// TestLocalNetwork list local IP addresses
func TestLocalNetwork() {
	logs.StartTest("NIC")
	ifaces, err := net.Interfaces()

	if err != nil {
		logs.Fail("Failed to list network interface: " + err.Error())
		return
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()

		if err != nil {
			logs.Fail("Failed to get address of interface " + i.Name)
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			logs.Info(fmt.Sprintf("Local IP: %s", ip.String()))
		}
	}
}
