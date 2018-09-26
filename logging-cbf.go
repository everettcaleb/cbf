package cbf

import "fmt"

type loggingCBF struct {
	cbf CountingBloomFilter
}

func (f *loggingCBF) Add(s string) {
	fmt.Printf("CBF.Add(%s)\n", s)
	f.cbf.Add(s)
}

func (f *loggingCBF) Remove(s string) {
	fmt.Printf("bloomFilter.Remove(%s)\n", s)
	f.cbf.Remove(s)
}

func (f *loggingCBF) Test(s string) bool {
	fmt.Printf("bloomFilter.Test(%s)\n", s)
	return f.cbf.Test(s)
}

func (f *loggingCBF) computeHashes(s string) []uint32 {
	fmt.Printf("CBF.computeHashes(%s)\n (mod size):", s)
	hs := f.cbf.computeHashes(s)
	for _, h := range hs {
		fmt.Println(h)
	}
	return hs
}

// CreateLoggingCBF wraps an existing CBF with logging to STDOUT
func CreateLoggingCBF(cbf CountingBloomFilter) CountingBloomFilter {
	return &loggingCBF{cbf}
}
