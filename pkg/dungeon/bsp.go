package dungeon

import ()

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

func (bn *BspNode) Partition(depth int) {
	if depth <= 0 {
		return
	}

	bn.Back = &BspNode{
		minx: bn.minx,
		maxx: (bn.minx + bn.maxx/2),
		miny: bn.miny,
		maxy: (bn.miny + bn.maxy/2),
	}
	bn.Front = &BspNode{
		minx: bn.minx + bn.maxx/2,
		maxx: bn.maxx,
		miny: bn.miny + bn.maxy/2,
		maxy: bn.maxy,
	}

	bn.Back.Partition(depth - 1)
	bn.Front.Partition(depth - 1)

}
