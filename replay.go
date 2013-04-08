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
	"encoding/json"
	//"reflect"
	//"encoding/binary"
)

//var filename = flag.String("filename", "", "replay filename")
var logLevel = flag.Int("log", 1, "log level")
func main() {
	
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

	//fmt.Printf("pos:%v\n",replay.pos)
	var m proto.Message
	for {
		msgType, m, err := replay.ReadPacket()
		if err != nil {
			//most likely we have reached the end of the replay
			log.Fatal(err)
			break
		}
		//example for outputting json
		if msgType == 7 && m != nil {
			j, _ := json.MarshalIndent(m,"","  ")
			fmt.Printf("%s\n", j) // json
			//fmt.Printf("%v\n", m) // protobuf repr
			return
			// example of using reflect to get fields not specified in type proto.Message
			// this will be useful when reverse-engineering unknows fields
			//v := reflect.ValueOf(m).Elem()
			//b := v.FieldByName("XXX_unrecognized").Bytes()
			//foo,_ :=binary.Uvarint(b)
			//fmt.Printf("XXX_unrecognized: %#x\n", b)
		}
	}

	//never reached, just to make sure we use m and please the go compiler
	fmt.Printf("m:%v\n", m)
}
