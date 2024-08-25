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
		z := And(x, y)
		y = Xor(y, x)
		x = z << 1
	}
	return y
}

// performs substraction between two numbers
func Substract(x int, y int) int {
	for y != 0 {
		z := And(Not(x), y)
		x = Xor(x, y)
		y = z << 1
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
	x = Xor(x, y)
	y = Xor(x, y)
	x = Xor(x, y)
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

// Get the index of most significant bit (MSB)
func MSB(x int) int {
	index := Log(x) / Log(2)
	return 1 << index
}

// get square root of a number
func SquareRoot(x int) int {
	z := MSB(x)
	y := 0
	for z != 0 {
		if ((y + z) * (y + z)) <= x {
			y += z
		}
		z >>= 1
	}
	return y
}

// calculate GCD (greatest common divisor)
func GCD(x int, y int) int {
	for y > 0 {
		z := y
		y = x % y
		x = z
	}
	return x
}

// calculate LCM (least common multiple)
func LCM(x int, y int) int {
	return (x * y) / GCD(x, y)
}
