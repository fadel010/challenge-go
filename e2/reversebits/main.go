package piscine

func ReverseBits(oct byte) byte {
	var rev byte

	for i := 0; i < 8; i++ {
		bit := (oct >> i) & 1
		rev = (rev << 1) | bit
	}

	return rev
}
