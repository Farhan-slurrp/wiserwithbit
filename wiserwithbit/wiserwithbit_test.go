package wiserwithbit

import (
	"testing"
)

func TestAnd(t *testing.T) {
	result := And(1, 2)
	if result != 0 {
		t.Errorf("And(1, 2) expected 0, but got %d", result)
	}
}

func TestOr(t *testing.T) {
	result := Or(1, 2)
	if result != 3 {
		t.Errorf("Or(1, 2) expected 3, but got %d", result)
	}
}

func TestXor(t *testing.T) {
	result := Xor(1, 3)
	if result != 2 {
		t.Errorf("Xor(1, 3) expected 2, but got %d", result)
	}
}

func TestNot(t *testing.T) {
	result := Not(1)
	if result != 18446744073709551614 {
		t.Errorf("Not(1) expected 18446744073709551614, but got %d", result)
	}
}

func TestAdd(t *testing.T) {
	result := Add(10, 20)
	if result != 30 {
		t.Errorf("Add(10, 30) expected 30, but got %d", result)
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(20, 10)
	if result != 10 {
		t.Errorf("Substract(20, 10) expected 10, but got %d", result)
	}
}

func TestGetRightMostsBit(t *testing.T) {
	result := GetRightMostBit(85)
	if result != 1 {
		t.Errorf("GetRightMostBit(85) expected 1, but got %d", result)
	}
}

func TestClearRightMostBit(t *testing.T) {
	result := ClearRightMostBit(85)
	if result != 84 {
		t.Errorf("ClearRightMostBit(85) expected 84, but got %d", result)
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	result1 := IsPowerOfTwo(85)
	if result1 != false {
		t.Errorf("IsPowerOfTwo(85) expected false, but got %v", result1)
	}

	result2 := IsPowerOfTwo(4)
	if result2 != true {
		t.Errorf("IsPowerOfTwo(4) expected true, but got %v", result2)
	}
}

func TestSwap(t *testing.T) {
	x, y := Swap(3, 1)
	if x != 1 && y != 3 {
		t.Errorf("Swap(3, 1) expected (1, 3), but got (%d, %d)", x, y)
	}
}

func TestIsOdd(t *testing.T) {
	result1 := IsOdd(4)
	if result1 != false {
		t.Errorf("IsOdd(4) expected false, but got %v", result1)
	}

	result2 := IsOdd(13)
	if result2 != true {
		t.Errorf("IsOdd(13) expected true, but got %v", result2)
	}
}

func TestIsEven(t *testing.T) {
	result1 := IsEven(4)
	if result1 != true {
		t.Errorf("IsEven(4) expected true, but got %v", result1)
	}

	result2 := IsEven(13)
	if result2 != false {
		t.Errorf("IsEven(13) expected false, but got %v", result2)
	}
}

func TestLog(t *testing.T) {
	result := Log(256)
	if result != 8 {
		t.Errorf("Log(256) expected 8, but got %v", result)
	}
}

func TestSquareRoot(t *testing.T) {
	result := SquareRoot(64)
	if result != 8 {
		t.Errorf("SquareRoot(64) expected 8, but got %v", result)
	}
}
