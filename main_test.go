package main_test

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestEndToEnd(t *testing.T) {
	t.Run("10.1.2.3/24 (no trim)", func(t *testing.T) {
		const cmd = "go"
		var (
			args = []string{
				"run", "main.go", "-cidr", "10.1.2.3/24",
			}
			actual []byte
			err    error
		)
		expected := []byte{
			27, 91, 48, 109, 49, 48, 46, 49, 46, 50, 46, 48, 10,
		}

		if actual, err = exec.Command(cmd, args...).Output(); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("mismatch\n"+
				"  actual: %s (%v)\n"+
				"expected: %s (%v)",
				string(actual), actual, string(expected), expected)
		}
	})
	t.Run("10.1.2.3/24 (no trim)", func(t *testing.T) {
		const cmd = "go"
		var (
			args = []string{
				"run", "main.go", "-cidr", "10.1.2.3/24", "-trim",
			}
			actual []byte
			err    error
		)
		expected := []byte{
			27, 91, 48, 109, 49, 48, 46, 49, 46, 50, 10,
		}

		if actual, err = exec.Command(cmd, args...).Output(); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("mismatch\n"+
				"  actual: %s (%v)\n"+
				"expected: %s (%v)",
				string(actual), actual, string(expected), expected)
		}
	})
	t.Run("10.1.2.3/16 (no trim)", func(t *testing.T) {
		const cmd = "go"
		var (
			args = []string{
				"run", "main.go", "-cidr", "10.1.2.3/16",
			}
			actual []byte
			err    error
		)
		expected := []byte{
			27, 91, 48, 109, 49, 48, 46, 49, 46, 48, 46, 48, 10,
		}

		if actual, err = exec.Command(cmd, args...).Output(); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("mismatch\n"+
				"  actual: %s (%v)\n"+
				"expected: %s (%v)",
				string(actual), actual, string(expected), expected)
		}
	})

	t.Run("10.1.2.3/16 (no trim)", func(t *testing.T) {
		const cmd = "go"
		var (
			args = []string{
				"run", "main.go", "-cidr", "10.1.2.3/16", "-trim",
			}
			actual []byte
			err    error
		)
		expected := []byte{
			27, 91, 48, 109, 49, 48, 46, 49, 10,
		}

		if actual, err = exec.Command(cmd, args...).Output(); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(actual, expected) {
			t.Fatalf("mismatch\n"+
				"  actual: %s (%v)\n"+
				"expected: %s (%v)",
				string(actual), actual, string(expected), expected)
		}
	})
}

func TestMain(m *testing.M) {
	m.Run()
}
