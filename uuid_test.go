package quuid

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("pkg", "quuid")
)

func TestQUUID(t *testing.T) {
	id1 := UUID()
	id2 := UUID()

	log.Infof("TestQUUID() id1=%s, id2=%s", id1, id2)
	assert.NotEqual(t, id1, id2)
}

func TestQUUID2(t *testing.T) {
	generator := New()
	id1 := generator.UUID()

	time.Sleep(time.Millisecond)
	id2 := generator.UUID()

	log.Infof("TestQUUID2() id1=%s, id2=%s", id1, id2)
	assert.NotEqual(t, id1, id2)
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UUID()
	}
}

func BenchmarkUUIDWithHWAddress(b *testing.B) {
	g := New(WithHWAddressPrefix)
	for i := 0; i < b.N; i++ {
		_ = g.UUID()
	}
}
