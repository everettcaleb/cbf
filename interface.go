package cbf

// CountingBloomFilter is an interface for a Bloom Filter that supports removal through counting
type CountingBloomFilter interface {
	// Add hashes the input string and adds it to the filter by incrementing the entry at the hash % size
	Add(s string)
	// Remove hashes the input string and removes it from the filter by decrementing the entry at the hash % size
	Remove(s string)
	// Test hashes the input string and checks the entries at hash % size to say whether that string has possibly been added to the filter
	Test(s string) bool

	computeHashes(s string) []uint32
}
