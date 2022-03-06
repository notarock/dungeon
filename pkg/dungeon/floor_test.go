package dungeon

import (
	"gotest.tools/assert"
	"testing"
)

func TestGenerateCanvas(t *testing.T) {
	canvas := GenerateCanvas()
	lenX := len(canvas)
	lenY := len(canvas[0])
	assert.Assert(t, lenX != 0)
	assert.Assert(t, lenY != 0)

	assert.Equal(t, canvas[0][0], Wall)
	assert.Equal(t, canvas[lenX-1][lenY-1], Wall)
}
