////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a BSD-style  license that can be found in the LICENSE file.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"crypto/sha256"
	"log"
	"time"
)

func main() {
	doit()
}

var ave = time.Duration(0)

func doit() {
	defer timeTrack(time.Now(), "Checksum")
	for i := 0; i < 1000; i++ {
		start := time.Now()
		sha256.Sum256([]byte("I am a secret; Hide me!"))
		elapsed := time.Since(start)
		ave += elapsed
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
	log.Printf("On average, %s took %s", name, ave/1000.0)
}
