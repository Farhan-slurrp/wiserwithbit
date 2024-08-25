package wiserwithbit

import "testing"

func TestBitMap(t *testing.T) {
	bitmap := NewBitmap(8)

	bitmap.SetBit(4)
	bit := bitmap.GetBit(4)
	if bit != 16 {
		t.Errorf("Expect bit at position 4 to be 16, but got %v", bit)
	}

	bitmap.ClearBit(4)
	bit = bitmap.GetBit(4)
	if bit != 0 {
		t.Errorf("Expect bit at position 4 to be 0, but got %v", bit)
	}
}
