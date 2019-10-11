package quuid

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var defaultUUID = &QUUID{
	Format: "%d-%d",
}

// Returns the hardware address. copy from https://github.com/gofrs/uuid/blob/master/generator.go
func defaultHWAddrFunc() (net.HardwareAddr, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return []byte{}, err
	}
	for _, iface := range ifaces {
		if len(iface.HardwareAddr) >= 6 {
			return iface.HardwareAddr, nil
		}
	}
	return []byte{}, fmt.Errorf("uuid: no HW address found")
}

// QUUID main structure
type QUUID struct {
	sync.Mutex
	Format string
	first  int64
	last   int64
}

// UUID generator the uuid
func (d *QUUID) UUID() string {
	d.Lock()
	defer d.Unlock()
	nowTs := time.Now().UnixNano() / 1000000
	if nowTs == d.first {
		d.last++
	} else {
		d.first = nowTs
		d.last = 0
	}
	return fmt.Sprintf(d.Format, d.first, d.last)
}

// WithHWAddressPrefix add MAC before uuid
func WithHWAddressPrefix(e *QUUID) {
	var (
		hardwareAddr [8]byte
		prefix       int64
	)

	if hwAddr, err := defaultHWAddrFunc(); err != nil {
		rand.Read(hardwareAddr[2:])
	} else {
		copy(hardwareAddr[2:], hwAddr)
	}

	err := binary.Read(bytes.NewBuffer(hardwareAddr[:]), binary.BigEndian, &prefix)
	if err != nil {
		fmt.Println(err)
	}

	e.Format = strconv.FormatInt(prefix, 10) + "-%d-%d"
}

// New create a new quuid generator
func New(opts ...func(*QUUID)) *QUUID {
	e := &QUUID{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// UUID generator the uuid with
func UUID() string {
	return defaultUUID.UUID()
}
