# Counting Bloom Filter (CBF)

[![Build Status](https://travis-ci.com/everettcaleb/cbf.svg?branch=master)](https://travis-ci.com/everettcaleb/cbf)
[![Coverage Status](https://coveralls.io/repos/github/everettcaleb/cbf/badge.svg?branch=master)](https://coveralls.io/github/everettcaleb/cbf?branch=master)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/56228e5748a3493e89a1296dc2c3d2f5)](https://www.codacy.com/app/everettcaleb/cbf?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=everettcaleb/cbf&amp;utm_campaign=Badge_Grade)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](LICENSE)

This is a project that implements a CBF in Go that uses 1-byte per entry. It uses the following hash functions, but could be re-written easily to use any hash function(s):

-  Adler32 (from [hash/adler32](https://golang.org/pkg/hash/adler32/))
-  FNV32 (from [hash/fnv](https://golang.org/pkg/hash/fnv/))
-  Murmur3-32 (from [github.com/spaolacci/murmur3](https://github.com/spaolacci/murmur3))

## Usage
Here's the interface:

    type CountingBloomFilter interface {
      Add(s string)
      Remove(s string)
      Test(s string) bool
    }

You can create an instance using `CreateCBF(size uint32)`. You can wrap an instance with logging via `CreateLoggingCBF(cbf CountingBloomFilter)`. Here's an example:

    package main

    import (
      "fmt"

      "github.com/everettcaleb/cbf"
    )

    func main() {
      f := cbf.CreateLoggingCBF(cbf.CreateCBF(1024 * 1024)) // 1 MB w/ logging
      f.Add("test")
      f.Add("test2")
      fmt.Printf("Test for 'test': %t\n", f.Test("test"))
      fmt.Printf("Test for 'test2': %t\n", f.Test("test2"))
      fmt.Printf("Test for 'test3': %t\n", f.Test("test3"))
      f.Remove("test2")
      fmt.Printf("Test for 'test': %t\n", f.Test("test"))
      fmt.Printf("Test for 'test2': %t\n", f.Test("test2"))
      fmt.Printf("Test for 'test3': %t\n", f.Test("test3"))
    }

## License
MIT License

Copyright 2018 Caleb Everett

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
