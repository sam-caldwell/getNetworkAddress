package main

import (
	"flag"
	"fmt"
	"github.com/sam-caldwell/ansi"
	"github.com/sam-caldwell/getNetworkAddress"
	"os"
	"strconv"
	"strings"
)

func main() {
	cidr := flag.String("cidr", "", "cidr block")
	trim := flag.Bool("trim", false, "optional flag to trim right hand .0 octets")
	flag.Parse()
	if *cidr == "" {
		ansi.Red().Printf("invalid cidr block (%s)", *cidr).LF().Fatal(1).Reset()
	}
	networkAddr, _ := getNetworkAddress.GetNetworkAddress(*cidr)
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
		fmt.Println(getNetworkAddress.TrimTrailingZeroOctets(networkAddr, sz))
	} else {
		fmt.Println(networkAddr)
	}
}
