package metadesk

import (
	"fmt"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestStringRoundTrip(t *testing.T) {
	assert.Equal(t, "foo", goStr(mdStr(defaultInstance.a, "foo")))
}

func TestNodeRoundTrip(t *testing.T) {
	md := NewMetadesk()
	res := md.ParseWholeString("test", "foo: (1, 2, 3)")
	node := res.Node.ChildFromString("foo", 0)
	assert.Len(t, node.Children(), 3)

	assert.Equal(t, node, md.goNodeP(md.mdNodeP(node)))
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

func TestStuff(t *testing.T) {
	res := ParseWholeString("test", `
		pos: [x: 1, y: 2, z: 3],
		vel: (4, 5, 6),
	`)
	for _, msg := range res.Errors.Messages {
		fmt.Println(msg.String)
	}
	if res.Errors.MaxMessageKind >= MessageKind_Warning {
		t.Fatal("There were errors in the parse.")
	}

	pos := res.Node.ChildFromString("pos", 0)
	vel := res.Node.ChildFromString("vel", 0)
	assert.False(t, pos.IsNil())
	assert.False(t, vel.IsNil())

	assert.True(t, vel.Children()[0].Flags&NodeFlag_Numeric != 0)
	assert.True(t, vel.Children()[1].Flags&NodeFlag_Numeric != 0)
	assert.True(t, vel.Children()[2].Flags&NodeFlag_Numeric != 0)
}
