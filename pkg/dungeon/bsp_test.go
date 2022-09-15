package dungeon_test

import (
	"testing"

	"github.com/notarock/dungeon/pkg/dungeon"
	"github.com/stretchr/testify/assert"
)

func TestPartition(t *testing.T) {
	bspNode := dungeon.NewBspNode(5, 100, 5, 100)

	t.Run("Depth of 1", func(t *testing.T) {
		bspNode.Partition(1, 100)

		assert.NotNil(t, bspNode.Back)
		assert.NotNil(t, bspNode.Front)

		assert.Equal(t, 5, bspNode.Back.GetMinx())
		assert.Equal(t, 5, bspNode.Back.GetMiny())
		assert.Equal(t, 100, bspNode.Back.GetMaxx())
		assert.Equal(t, 52, bspNode.Back.GetMaxy())
		assert.Equal(t, 5, bspNode.Front.GetMinx())
		assert.Equal(t, 52, bspNode.Front.GetMiny())
		assert.Equal(t, 100, bspNode.Front.GetMaxx())
		assert.Equal(t, 100, bspNode.Front.GetMaxy())
	})
}
