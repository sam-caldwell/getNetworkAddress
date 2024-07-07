package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/ansi"
	"net"
	"os"
	"strconv"
	"strings"
)

// TrimTrailingZeroOctets - delete any .0 octets
func trimTrailingZeroOctets(input, suffix string, count int) string {
	for i := 0; i < count; i++ {
		input = strings.TrimSuffix(input, suffix)
	}
	return input
}

// getNetworkAddress takes a CIDR string and returns the network address.
func getNetworkAddress(cidr string) (string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", err
	}

	// Get the network address from the parsed CIDR
	networkAddr := ip.Mask(ipnet.Mask).String()
	return networkAddr, nil
}

func main() {
	cidr := flag.String("cidr", "", "cidr block")
	trim := flag.Bool("trim", false, "optional flag to trim right hand .0 octets")
	flag.Parse()
	if *cidr == "" {
		ansi.Red().Printf("invalid cidr block (%s)", *cidr).LF().Fatal(1).Reset()
	}
	networkAddr, _ := getNetworkAddress(*cidr)
	sz, err := func() (int, error) {
		parts := strings.Split(*cidr, "/")
		if len(parts) < 2 {
			return 0, fmt.Errorf("invalid cidr address")
		}
		n, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		return n / 8, nil
	}()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if *trim {
		fmt.Println(trimTrailingZeroOctets(networkAddr, ".0", sz))
	} else {
		fmt.Println(networkAddr)
	}
}
