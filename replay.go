/**
 * Copyright (c) 2013 Malte Böhme
 * Licensed under the MIT license.
 */

package main

import (
	"./dota2"
	"flag"
	"fmt"
	"log"
	"os"
	"encoding/json"
	//"reflect"
	//"encoding/binary"
	//"./dota2/proto"
	"time"
)

type jsonPkt struct {
	// i dont like Uppercase keys
	Name string `json:"name"`
	Id uint64 `json:"id"`
	Data interface{}`json:"data"`
}

//var filename = flag.String("filename", "", "replay filename")
var logLevel = flag.Int("log", 1, "log level")
func main() {
	//meassure how long parsing took
	t0 := time.Now()

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("need (only) replay-file as argument")
		os.Exit(1)
	}
	filename := flag.Arg(0)
	replay, err := dota2.NewReplay(filename)
	replay.LogLevel = *logLevel
	if err != nil {
		log.Fatalf("cannot open replay '%s'", filename)
	}
	err = replay.Parse()
	pkts, _ := json.MarshalIndent(replay.Packets,"","  ")
	fmt.Printf("%v\nreplay.pkts:%s\n",err,pkts)
	t1 := time.Now()
	fmt.Printf("Parsing took %v.\n", t1.Sub(t0))
	//return
	//fmt.Printf("pos:%v\n",replay.pos)
	// var p dota2.Packet
	// for {
	// 	p, err := replay.ReadPacket()
	// 	if err != nil {
	// 		//most likely we have reached the end of the replay


	// 		stats, _ := json.MarshalIndent(replay.PktCount,"","  ")
	// 		fmt.Printf("replay.PktCount:%s\n",stats)

	// 		//some statistics on what and how many packets were parsed
	// 		t1 := time.Now()
	// 		fmt.Printf("Parsing took %v.\n", t1.Sub(t0))

	// 		log.Fatal(err)
	// 		break
	// 	}
		
		//example for outputting json
		// type 7 is DEM_Packet
		// var log_t uint64 = 7
		// if p.Id == log_t && p != nil {
		// 	j, _ := json.MarshalIndent(p,"","  ")
		// 	fmt.Printf("%s\n", j) // json
			//v := reflect.ValueOf(m).Elem()
			//b := v.FieldByName("data").Bytes()
			//foo,_ :=binary.Uvarint(b)
			//fmt.Printf("%#x\n", (*demo.CDemoPacket).(m).GetData())
			//fmt.Printf("%v\n", m) // protobuf repr
			//return
			// example of using reflect to get fields not specified in type proto.Message
			// this will be useful when reverse-engineering unknows fields
			//v := reflect.ValueOf(m).Elem()
			//b := v.FieldByName("XXX_unrecognized").Bytes()
			//foo,_ :=binary.Uvarint(b)
			//fmt.Printf("XXX_unrecognized: %#x\n", b)
		//}
	//}

	//never reached, just to make sure we use m and please the go compiler
	//fmt.Printf("m:%v\n", p)
}
