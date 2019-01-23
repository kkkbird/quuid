package quuid

import (
	"fmt"
	"sync"
	"time"
)

var defaultUUID = &QUUID{}

// QUUID main structure
type QUUID struct {
	sync.Mutex
	first int64
	last  int64
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
	return fmt.Sprintf("%d-%d", d.first, d.last)
}

// New create a new quuid generator
func New() *QUUID {
	return &QUUID{}
}

// UUID generator the uuid with
func UUID() string {
	return defaultUUID.UUID()
}
