////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Implements the testing of different kinds of checksums, so performance can be evaluated.
package checksum_test

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"testing"
)

// Function BenchmarkMd5Checksum determines performance values for various
// number of bytes using MD5.
// Typical Results:
//   pkg: github.com/abitofhelp/Checksums/checksum
//   BenchmarkMd5Checksum/Testing_1_bytes-8         	10000000	       170 ns/op
//   BenchmarkMd5Checksum/Testing_10_bytes-8        	10000000	       170 ns/op
//   BenchmarkMd5Checksum/Testing_100_bytes-8       	 5000000	       285 ns/op
//   BenchmarkMd5Checksum/Testing_1000_bytes-8      	 1000000	      1811 ns/op
//   BenchmarkMd5Checksum/Testing_10000_bytes-8     	  100000	     17222 ns/op
//   BenchmarkMd5Checksum/Testing_100000_bytes-8    	   10000	    168173 ns/op
//   BenchmarkMd5Checksum/Testing_1000000_bytes-8   	    1000	   1663159 ns/op
//   BenchmarkMd5Checksum/Testing_10000000_bytes-8  	     100	  17200860 ns/op
//   BenchmarkMd5Checksum/Testing_100000000_bytes-8 	      10	 166912634 ns/op
func BenchmarkMd5Checksum(b *testing.B) {

	// Various buffer sizes for the test
	lengths := []uint64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

	for _, length := range lengths {
		b.Run(fmt.Sprintf("Testing_%d_bytes", length), func(b *testing.B) {
			doMd5Test(uint64(length), b)
		})
	}
}

// Function doMd5Test performs the actual test for MD5.
func doMd5Test(length uint64, b *testing.B) {
	// Build a buffer of the correct length using random data.
	content := make([]byte, length)
	rand.Read(content)

	// Exclude the setup time in the performance metrics.
	b.ResetTimer()

	// Determine the benchmark.  The testing framework will determine
	// how many iterations are required.
	for i := 0; i < b.N; i++ {
		md5.Sum(content)
	}
}
