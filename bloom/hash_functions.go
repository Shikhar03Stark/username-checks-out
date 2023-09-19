package bloom

import "math/rand"

func strhash(data string) uint32 {
	src := rand.New(rand.NewSource(3))
	limit := int(1e9 + 7)
	val := 0

	for idx := range data {
		c := int64(src.Int())
		c %= int64(limit)
		c *= (int64(data[idx]) + 1)
		c %= int64(limit)
		val += int(c)
		val %= limit
	}

	return uint32(val)
}

func rotateCapture(val, rotation, mask uint32, k int) []uint32 {
	values := make([]uint32, 0)
	bais := 1

	for i := 0; i < k; i++ {
		v := (val & mask)
		values = append(values, v)
		val = (val >> rotation)

		v = (v << (32 - rotation))
		val = (val | v)
		val += uint32(bais)
	}

	return values
}

func SimpleHash(data string, k, sz int) []uint32 {
	aux := data + "#"

	hashVal := strhash(aux)
	mask := uint32(1)
	for mask <= uint32(sz) {
		mask = (mask << 1)
	}
	khash := rotateCapture(hashVal, 5, mask-1, k)

	idx := make([]uint32, k)
	for i := range khash {
		idx[i] = khash[i] % uint32(sz)
	}

	return idx
}
