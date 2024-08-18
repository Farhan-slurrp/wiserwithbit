package wiserwithbit

// performs the bitwise AND operation between two numbers.
func And(x, y uint) uint {
	return x & y
}

// performs the bitwise OR operation between two numbers.
func Or(x, y uint) uint {
	return x | y
}

// performs the bitwise XOR operation between two numbers.
func Xor(x, y uint) uint {
	return x ^ y
}

// performs the bitwise NOT operation on a number.
func Not(x uint) uint {
	return ^x
}

// performs addition between two numbers
func Add(x uint, y uint) uint {
	for x != 0 {
		c := And(x, y)
		y = Xor(y, x)
		x = c << 1
	}
	return y
}

// performs substraction between two numbers (x should be > y)
func Substract(x uint, y uint) uint {
	if x < y {
		panic("x cannot be less than y")
	}
	for y != 0 {
		c := And(Not(x), y)
		x = Xor(x, y)
		y = c << 1
	}

	return x
}

// get right most bit
func GetRightMostBit(x uint) uint {
	return And(x, -x)
}

// clear right most bit
func ClearRightMostBit(x uint) uint {
	return And(x, x-1)
}

// check if number is the power of 2
func IsPowerOfTwo(x uint) bool {
	return And(x, x-1) == 0
}

// swap two values
func Swap(x uint, y uint) (uint, uint) {
	x ^= y
	y ^= x
	x ^= y
	return x, y
}

// check if number is odd
func IsOdd(x uint) bool {
	return And(x, 1) == 1
}

// check if number is even
func IsEven(x uint) bool {
	return And(x, 1) == 0
}

// get the log of a number
func Log(x uint) uint {
	y := uint(0)
	for x > 1 {
		y += 1
		x >>= 1
	}
	return y
}

// get square root of a number
func SquareRoot(x uint) uint {
	msb := Log(x) / Log(2)
	c := uint(1 << msb)
	y := uint(0)
	for c != 0 {
		if ((y + c) * (y + c)) <= x {
			y += c
		}
		c >>= 1
	}
	return y
}
