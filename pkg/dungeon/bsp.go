package dungeon

import (
	"math/rand"
	"time"
)

type BspNode struct {
	Front *BspNode
	Back  *BspNode
	minx  int
	maxx  int
	miny  int
	maxy  int
}

func (bn BspNode) GetMinx() int { return bn.minx }
func (bn BspNode) GetMaxx() int { return bn.maxx }
func (bn BspNode) GetMiny() int { return bn.miny }
func (bn BspNode) GetMaxy() int { return bn.maxy }

func NewBspNode(minx, maxx, miny, maxy int) BspNode {
	bn := BspNode{
		minx: minx,
		maxx: maxx,
		miny: miny,
		maxy: maxy,
	}
	return bn
}

func (bn *BspNode) Partition(depth int, chancesPercent int) {
	if depth <= 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(100) + 1
	if chancesPercent <= roll {
		return
	}

	lenx := bn.maxx - bn.minx
	leny := bn.maxy - bn.miny

	if lenx > leny {
		bn.Back = &BspNode{
			minx: bn.minx,
			maxx: (bn.minx + lenx/2),
			miny: bn.miny,
			maxy: bn.maxy,
		}
		bn.Front = &BspNode{
			minx: bn.minx + lenx/2,
			maxx: bn.maxx,
			miny: bn.miny,
			maxy: bn.maxy,
		}
	} else {
		bn.Back = &BspNode{
			minx: bn.minx,
			maxx: bn.maxx,
			miny: bn.miny,
			maxy: (bn.miny + leny/2),
		}
		bn.Front = &BspNode{
			minx: bn.minx,
			maxx: bn.maxx,
			miny: bn.miny + leny/2,
			maxy: bn.maxy,
		}
	}

	bn.Back.Partition(depth-1, chancesPercent)
	bn.Front.Partition(depth-1, chancesPercent)

}
