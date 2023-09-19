package main

import (
	"fmt"

	"github.com/Shikhar03Stark/username-checks-out/bloom"
)

func main() {
	myset := bloom.New(128, 7, bloom.SimpleHash)

	myset.Add("harshit")
	myset.Add("shruty")
	myset.Add("george")
	myset.Add("Alan")

	// fmt.Println("BloomFilter", myset.BitArr)
	testStr := []string{"rome", "shruty", "harshit", "queen", "lemme", "try", "BLR"}

	for i := range testStr {
		s := testStr[i]

		if myset.MaybePresent(s) {
			fmt.Printf("%v maybe present in the set\n", s)
		} else {
			fmt.Printf("%v is not present in the set\n", s)
		}
	}
}
