package popcount

var pc [265]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Count1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 2.3
func Count2(x uint64) int {
	count := 0
	for i := uint8(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

// 2.4
func Count3(x uint64) int {
	count := uint64(0)
	for i := uint8(0); i < 64; i++ {
		count += (x >> i) & 1
	}
	return int(count)
}

// 2.5
func Count4(x uint64) int {
	count := 0
	for ; x != 0; x &= x - 1 {
		count++
	}
	return count
}

// おまけ
func Count5(x uint64) int {
	x = x&0x5555555555555555 + x>>1&0x5555555555555555
	x = x&0x3333333333333333 + x>>2&0x3333333333333333
	x = x&0x0f0f0f0f0f0f0f0f + x>>4&0x0f0f0f0f0f0f0f0f
	x = x&0x00ff00ff00ff00ff + x>>8&0x00ff00ff00ff00ff
	x = x&0x0000ffff0000ffff + x>>16&0x0000ffff0000ffff
	x = x&0x00000000ffffffff + x>>32&0x00000000ffffffff
	return int(x)
}
