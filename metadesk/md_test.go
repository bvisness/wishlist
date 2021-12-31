package metadesk

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestStringRoundTrip(t *testing.T) {
	assert.Equal(t, "foo", goStr(mdStr(defaultInstance.a, "foo")))
}

func TestYouCanTrackThings(t *testing.T) {
	md := NewMetadesk()
	foo := 3
	ptr := unsafe.Pointer(&foo)
	md.track(foo, ptr)

	{
		existing, ok := md.go2c[foo]
		if assert.True(t, ok) {
			assert.Equal(t, ptr, existing)
		}
	}
	{
		existing, ok := md.c2go[ptr]
		if assert.True(t, ok) {
			assert.Equal(t, foo, existing)
		}
	}
}

func TestBlep(t *testing.T) {
	ParseWholeString("test", `
		pos: [x: 1, y: 2, z: 3],
		vel: (4, 5, 6),
	`)
}
