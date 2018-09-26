package cbf

import (
	"hash/adler32"
	"hash/fnv"

	"github.com/spaolacci/murmur3"
)

type countingBloomFilter struct {
	size  uint32
	slice []byte
}

func (b *countingBloomFilter) Add(s string) {
	hs := b.computeHashes(s)

	for _, h := range hs {
		b.slice[h]++
	}
}

func (b *countingBloomFilter) Remove(s string) {
	hs := b.computeHashes(s)

	for _, h := range hs {
		b.slice[h]--
	}
}

func (b *countingBloomFilter) Test(s string) bool {
	hs := b.computeHashes(s)
	sum := 0

	for _, h := range hs {
		if b.slice[h] > 0 {
			sum++
		}
	}

	return sum == len(hs)
}

func (b *countingBloomFilter) computeHashes(s string) []uint32 {
	h := make([]uint32, 3)
	d := []byte(s)

	a := adler32.New()
	a.Write(d)
	h[0] = a.Sum32() % b.size

	f := fnv.New32()
	f.Write(d)
	h[1] = f.Sum32() % b.size

	m := murmur3.New32()
	m.Write(d)
	h[2] = m.Sum32() % b.size

	return h
}

// CreateCBF creates an implementation of CountingBloomFilter that uses 1-byte entries
func CreateCBF(size uint32) CountingBloomFilter {
	return &countingBloomFilter{
		size:  size,
		slice: make([]byte, size),
	}
}
