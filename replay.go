package main

import (
	"./dota2"
	"code.google.com/p/goprotobuf/proto"
	"flag"
	"fmt"
	"log"
)

//var filename = flag.String("filename", "", "replay filename")
func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Printf("need (only) replay-file as argument")
	}
	filename := flag.Arg(0)
	replay, err := dota2.NewReplay(filename)
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
