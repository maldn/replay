/**
 * Copyright (c) 2013 Malte Böhme
 * Licensed under the MIT license.
 */

package main

import (
	"./dota2"
	"code.google.com/p/goprotobuf/proto"
	"flag"
	"fmt"
	"log"
	"os"
)

//var filename = flag.String("filename", "", "replay filename")
func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("need (only) replay-file as argument")
		os.Exit(1)
	}
	filename := flag.Arg(0)
	replay, err := dota2.NewReplay(filename)
	replay.LogLevel = 2
	if err != nil {
		log.Fatalf("cannot open replay '%s'", filename)
	}

	//fmt.Printf("pos:%v\n",replay.pos)
	var m proto.Message
	for {
		m, err = replay.ReadPacket()
		if err != nil {
			//most likely we have reached the end of the replay
			log.Fatal(err)
			break
		}
		//fmt.Printf("m:%v\n", m)
	}
	//never reached, just to make sure we use m and please the go compiler
	fmt.Printf("m:%v\n", m)
}
