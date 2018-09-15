////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Implements the testing of different kinds of checksums, so performance can be evaluated.
package checksum_test

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"testing"
)

// Function BenchmarkSha256Checksum determines performance values for various
// number of bytes using SHA256.
// Typical Results:
//   pkg: github.com/abitofhelp/Checksums/checksum
//   BenchmarkSha256Checksum/Testing_1_bytes-8         	 5000000	       282 ns/op
//   BenchmarkSha256Checksum/Testing_10_bytes-8        	 5000000	       292 ns/op
//   BenchmarkSha256Checksum/Testing_100_bytes-8       	 3000000	       487 ns/op
//   BenchmarkSha256Checksum/Testing_1000_bytes-8      	  500000	      3244 ns/op
//   BenchmarkSha256Checksum/Testing_10000_bytes-8     	   50000	     31932 ns/op
//   BenchmarkSha256Checksum/Testing_100000_bytes-8    	    5000	    306250 ns/op
//   BenchmarkSha256Checksum/Testing_1000000_bytes-8   	     500	   3123127 ns/op
//   BenchmarkSha256Checksum/Testing_10000000_bytes-8  	      50	  31122289 ns/op
//   BenchmarkSha256Checksum/Testing_100000000_bytes-8 	       5	 305421880 ns/op
func BenchmarkSha256Checksum(b *testing.B) {

	lengths := []uint64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("Testing_%d_bytes", length), func(b *testing.B) {
			doSha256Test(uint64(length), b)
		})
	}
}

// Function doSha256Test performs the actual test for SHA256.
func doSha256Test(length uint64, b *testing.B) {
	// Build a buffer of the correct length using random data.
	content := make([]byte, length)
	rand.Read(content)

	// Exclude the setup time in the performance metrics.
	b.ResetTimer()

	// Determine the benchmark.  The testing framework will determine
	// how many iterations are required.
	for i := 0; i < b.N; i++ {
		sha256.Sum256(content)
	}
}
