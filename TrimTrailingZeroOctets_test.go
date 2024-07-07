package getNetworkAddress

import (
	"testing"
)

func TestTrimTrailingZeroOctets(t *testing.T) {
	testCases := []struct {
		input    string
		count    int
		expected string
	}{
		{"192.168.1.0", 1, "192.168.1"},
		{"10.0.0.0", 2, "10.0"},
		{"172.16.0.0.0", 3, "172.16"},
		{"8.8.8.8", 1, "8.8.8.8"},
		{"192.168.1.0", 0, "192.168.1.0"}, // Edge case where count is 0
	}

	for _, tc := range testCases {
		actual := TrimTrailingZeroOctets(tc.input, tc.count)

		if actual != tc.expected {
			t.Errorf("Expected '%s' for input '%s' with count %d, but got '%s'", tc.expected, tc.input, tc.count, actual)
		}
	}
}
