package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence(t *testing.T) {
	os.Remove(".test.lock")
	os.Setenv("QMP_MAC_GENERATOR_LOCK_FILE", ".test.lock")
	defer os.Unsetenv("QMP_MAC_GENERATOR_LOCK_FILE")

	mac, err := newMACAddress()
	assertValue(t, "52:54:00:ab:00:00", mac, err)

	mac, err = newMACAddress()
	assertValue(t, "52:54:00:ab:00:01", mac, err)
}

func TestRollover(t *testing.T) {
	assert.Equal(t, []byte{0x52, 0x54, 0x00, 0xAB, 0x00, 0x00}, nextMAC([]byte{0x52, 0x54, 0x00, 0xAB, 0xFF, 0xFF}))
}

func assertValue(t *testing.T, expected, actual string, err error) {
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
