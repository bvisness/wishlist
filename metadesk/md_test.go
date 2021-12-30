package metadesk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringRoundTrip(t *testing.T) {
	assert.Equal(t, "foo", GoStr(Str("foo")))
}

func TestBlep(t *testing.T) {
	blep()
}
