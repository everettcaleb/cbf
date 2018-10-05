package cbf

import "testing"

const testLoggingCBFSize uint32 = 1024 * 1024

func TestLoggingCountingBloomFilterAdd(t *testing.T) {
	b := CreateLoggingCBF(CreateCBF(testLoggingCBFSize))
	b.Add("test")
	if !b.Test("test") {
		t.Fatal("CBF.Test(\"test\") expected to be true after adding")
	}
}

func TestLoggingCountingBloomFilterRemove(t *testing.T) {
	b := CreateLoggingCBF(CreateCBF(testLoggingCBFSize))
	b.Add("test")
	if !b.Test("test") {
		t.Fatal("CBF.Test(\"test\") expected to be true after adding")
	}
	b.Remove("test")

	if b.Test("test") {
		t.Fatal("CBF.Test(\"test\") expected to be false after removal")
	}
}

func TestLoggingCountingBloomFilterTest(t *testing.T) {
	b := CreateLoggingCBF(CreateCBF(testLoggingCBFSize))
	b.Add("test")
	if !b.Test("test") {
		t.Fatal("CBF.Test(\"test\") expected to be true after adding")
	}

	if b.Test("test2") {
		t.Fatal("CBF.Test(\"test2\") expected to be false since it has never been added")
	}
}

func TestLoggingCountingBloomFilterComputeHashes(t *testing.T) {
	b := &loggingCBF{CreateCBF(testLoggingCBFSize)}
	hashes := b.computeHashes("test")
	if len(hashes) != 3 {
		t.Fatal("CBF.computeHashes expected to return 3 values")
	}
}
