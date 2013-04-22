/**
 * Copyright (c) 2013 Malte Böhme
 * Licensed under the MIT license.
 */

package main

import (
	"./dota2"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	//"reflect"
	//"encoding/binary"
	//"./dota2/proto"
	"runtime/pprof"
	"time"
)

var logLevel = flag.Int("log", 1, "log level")
var filename = flag.String("replay", "", "replay file")
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	//meassure how long parsing took
	t0 := time.Now()

	flag.Parse()
	if *filename == "" {
		log.Fatal("No Replay provided. use -replay=<file>.")
	}
	replay, err := dota2.NewReplay(*filename)
	replay.LogLevel = *logLevel
	if err != nil {
		log.Fatalf("cannot open replay '%s'", filename)
	}

	// encoding json takes almost all of the resources.
	// not much we can do here
	// on a side note: printing and not piping the output takes about 10s in itself
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	err = replay.Parse()
	pkts, _ := json.MarshalIndent(replay.Packets, "", "  ")
	fmt.Printf("%v\nreplay.pkts:%s\n", err, pkts)
	t1 := time.Now()
	log.Printf("Parsing took %v.\n", t1.Sub(t0))
}
