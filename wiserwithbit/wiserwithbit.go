package wiserwithbit

// performs the bitwise AND operation between two numbers.
func And(x, y int) int {
	return x & y
}

// performs the bitwise OR operation between two numbers.
func Or(x, y int) int {
	return x | y
}

// performs the bitwise XOR operation between two numbers.
func Xor(x, y int) int {
	return x ^ y
}

// performs the bitwise NOT operation on a number.
func Not(x int) int {
	return ^x
}

// performs addition between two numbers
func Add(x int, y int) int {
	for x != 0 {
		c := And(x, y)
		y = Xor(y, x)
		x = c << 1
	}
	return y
}

// performs substraction between two numbers
func Substract(x int, y int) int {
	for y != 0 {
		c := And(Not(x), y)
		x = Xor(x, y)
		y = c << 1
	}

	return x
}

// get right most bit
func GetRightMostBit(x int) int {
	return And(x, -x)
}

// clear right most bit
func ClearRightMostBit(x int) int {
	return And(x, x-1)
}

// check if number is the power of 2
func IsPowerOfTwo(x int) bool {
	return And(x, x-1) == 0
}

// swap two values
func Swap(x int, y int) (int, int) {
	x ^= y
	y ^= x
	x ^= y
	return x, y
}

// check if number is odd
func IsOdd(x int) bool {
	return And(x, 1) == 1
}

// check if number is even
func IsEven(x int) bool {
	return And(x, 1) == 0
}

// get the log of a number
func Log(x int) int {
	y := 0
	for x > 1 {
		y += 1
		x >>= 1
	}
	return y
}

// get square root of a number
func SquareRoot(x int) int {
	msb := Log(x) / Log(2)
	c := 1 << msb
	y := 0
	for c != 0 {
		if ((y + c) * (y + c)) <= x {
			y += c
		}
		c >>= 1
	}
	return y
}

// calculate GCD (greatest common divisor)
func GCD(x int, y int) int {
	for y > 0 {
		c := y
		y = x % y
		x = c
	}
	return x
}

// calculate LCM (least common multiple)
func LCM(x int, y int) int {
	return (x * y) / GCD(x, y)
}
