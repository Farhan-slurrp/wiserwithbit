package wiserwithbit

type Bitmap struct {
	bits []int
}

type IBitmap interface {
	SetBit(int)
	ClearBit(int)
	GetBit(int) int
}

func NewBitmap(size int) IBitmap {
	return &Bitmap{
		bits: make([]int, ((size + 31) / 32)),
	}
}

func (bm *Bitmap) SetBit(pos int) {
	bm.bits[pos/32] |= (1 << (pos % 32))
}

func (bm *Bitmap) ClearBit(pos int) {
	bm.bits[pos/32] &= ^(1 << (pos % 32))
}

func (bm *Bitmap) GetBit(pos int) int {
	return bm.bits[pos/32] & (1 << (pos % 32))
}
