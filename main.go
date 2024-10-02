package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofrs/flock"
)

func getLockName() string {
	fName := os.Getenv("QMP_MAC_GENERATOR_LOCK_FILE")
	if len(fName) == 0 {
		fName = "/var/lock/qmp-mac-generator.lock"
	}
	return fName
}

func getLock(fName string) (*flock.Flock, error) {
	fileLock := flock.New(fName)
	err := fileLock.Lock()
	if err != nil {
		return nil, fmt.Errorf("failed acquiring lock on %s : %s", fName, err)
	}
	return fileLock, nil
}

func readLastMAC(fName string) ([]byte, error) {
	data, err := os.ReadFile(fName)
	if err != nil {
		return nil, fmt.Errorf("failed reading %s: %s", fName, err)
	}

	if len(data) != 6 {
		data = []byte{0x52, 0x54, 0x00, 0xAB, 0xFF, 0xFF}
	}
	return data, nil
}

func nextMAC(data []byte) []byte {
	data[5]++
	if data[5] == 0 {
		data[4]++
	}
	return data
}

func saveMAC(fName string, data []byte) error {
	err := os.WriteFile(fName, data, 0644)
	if err != nil {
		return fmt.Errorf("failed updating %s: %s", fName, err)
	}
	return nil
}

func newMACAddress() (string, error) {
	fName := getLockName()

	fileLock, err := getLock(fName)
	if err != nil {
		return "", err
	}
	defer fileLock.Unlock()

	data, err := readLastMAC(fName)
	if err != nil {
		return "", err
	}

	data = nextMAC(data)

	err = saveMAC(fName, data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", data[0], data[1], data[2], data[3], data[4], data[5]), nil
}

func main() {
	mac, err := newMACAddress()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mac)
}
