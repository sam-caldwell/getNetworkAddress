package getNetworkAddress

import (
	"testing"
)

func TestGetNetworkAddress(t *testing.T) {
	testCases := []struct {
		cidr         string
		expectedAddr string
		expectError  bool
	}{
		{"192.168.1.0/24", "192.168.1.0", false},
		{"10.0.0.0/8", "10.0.0.0", false},
		{"172.16.0.0/16", "172.16.0.0", false},
		{"invalid_cidr", "", true}, // Example of a test case with an invalid CIDR string
	}

	for _, tc := range testCases {
		addr, err := GetNetworkAddress(tc.cidr)

		if tc.expectError {
			if err == nil {
				t.Errorf("Expected error for CIDR %s, but got nil", tc.cidr)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for CIDR %s: %v", tc.cidr, err)
			}

			if addr != tc.expectedAddr {
				t.Errorf("Expected network address %s for CIDR %s, but got %s", tc.expectedAddr, tc.cidr, addr)
			}
		}
	}
}
