package bloom

type BloomFilter struct {
	BitArr []bool                                // bit array of Bloom Filter
	Stored int                                   // number of items stored in the bloom filter
	K      int                                   // number of Hash functions
	HashFn func(data string, k, sz int) []uint32 // Hash function which takes input and gives K indices to set the value on BitArr to 1
}

func (bf *BloomFilter) Add(data string) {
	hval := bf.HashFn(data, bf.K, len(bf.BitArr))
	for _, idx := range hval {
		bf.BitArr[idx] = true
	}
	// fmt.Printf("%v k-hash %v\n", data, hval)
	bf.Stored++
}

func (bf *BloomFilter) MaybePresent(data string) bool {
	hval := bf.HashFn(data, bf.K, len(bf.BitArr))
	// fmt.Printf("%v k-hash %v\n", data, hval)
	for _, idx := range hval {
		if bf.BitArr[idx] == false {
			return false
		}
	}

	return true
}

func New(size, k int, hashFn func(data string, k, sz int) []uint32) *BloomFilter {
	bitArr := make([]bool, size)
	for i := range bitArr {
		bitArr[i] = false
	}

	return &BloomFilter{
		BitArr: bitArr,
		Stored: 0,
		K:      k,
		HashFn: hashFn,
	}
}
