package cbf

import "testing"

const testCBFSize uint32 = 1024 * 1024

func TestCountingBloomFilterAdd(t *testing.T) {
	b := CreateCBF(testCBFSize)
	b.Add("test")
	if !b.Test("test") {
		t.Fatal("CBF.Test(\"test\") expected to be true after adding")
	}
}

func TestCountingBloomFilterRemove(t *testing.T) {
	b := CreateCBF(testCBFSize)
	b.Add("test")
	if !b.Test("test") {
		t.Fatal("CBF.Test(\"test\") expected to be true after adding")
	}
	b.Remove("test")

	if b.Test("test") {
		t.Fatal("CBF.Test(\"test\") expected to be false after removal")
	}
}

func TestCountingBloomFilterTest(t *testing.T) {
	b := CreateCBF(testCBFSize)
	b.Add("test")
	if !b.Test("test") {
		t.Fatal("CBF.Test(\"test\") expected to be true after adding")
	}

	if b.Test("test2") {
		t.Fatal("CBF.Test(\"test2\") expected to be false since it has never been added")
	}
}
