package getNetworkAddress

import "strings"

// TrimTrailingZeroOctets - delete any .0 octets
func TrimTrailingZeroOctets(input string, count int) string {
	for i := 0; i < count; i++ {
		input = strings.TrimSuffix(input, ".0")
	}
	return input
}
