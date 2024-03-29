package dungeon_test

import (
	"testing"

	"github.com/notarock/dungeon/pkg/dungeon"
	"github.com/stretchr/testify/assert"
)

func TestInitPlayer(t *testing.T) {
	t.Run("negative x return an error", func(t *testing.T) {
		_, err := dungeon.InitPlayer(-3, 0)
		assert.Error(t, err)
	})
	t.Run("negative y return an error", func(t *testing.T) {
		_, err := dungeon.InitPlayer(0, -1)
		assert.Error(t, err)
	})
	t.Run("Init player at given positions", func(t *testing.T) {
		x := 50
		y := 71
		p, err := dungeon.InitPlayer(x, y)
		assert.Nil(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, x, p.GetX())
		assert.Equal(t, y, p.GetY())
	})

}
