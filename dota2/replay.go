/**
 * Copyright (c) 2013 Malte Böhme
 * Licensed under the MIT license.
 */

package dota2

import (
	//"bytes"
	"code.google.com/p/goprotobuf/proto"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	//"os"
	//"encoding/hex"
	"errors"
	"code.google.com/p/snappy-go/snappy"
	"./proto" // generated files `protoc --go_out=generated *.proto`
)

type Replay struct {
	buf       []byte
	size, pos int // replay size, current position
	LogLevel  int
}

func NewReplay(filename string) (r *Replay, err error) {
	r = new(Replay)
	r.buf, err = ioutil.ReadFile(filename)
	r.size = len(r.buf)
	//FIXME: validate header
	//header is a 8 byte string + 32bit size of replay file.
	r.pos = 12
	return r, err
}

func (r *Replay) ReadPacket() (msgType uint64, m proto.Message, err error) {

	if r.pos >= r.size {
		err_msg := fmt.Sprintf("end of replay.. size:%d\tpos:%d\n", r.size, r.pos)
		err = errors.New(err_msg)
		// braindead inline-signaling, since a value of 0 might mean DEM_Stop packet
		return 1<<16-1, nil, err
	}

	// no idea if its a bug in go-protobuf or valves stuff
	// valve defines a msgType of -1 as DEM_Error but seems like the encoding is Uvarint..
	// strange..
	msgType, size := binary.Uvarint(r.buf[r.pos:])
	r.pos += size
	if r.LogLevel >= 3 {
		fmt.Printf("%s: msgType: %#x(%d) size:%d\n", getPacketName(msgType), msgType, msgType, size)
	}

	tick, size := binary.Uvarint(r.buf[r.pos:])
	r.pos += size
	if r.LogLevel >= 3 {
		fmt.Printf("tick: %#x(%d) len:%d\n", tick, tick, size)
	}

	pktLen, size := binary.Uvarint(r.buf[r.pos:])
	r.pos += size
	if r.LogLevel >= 3 {
		fmt.Printf("length: %d (%#x) %d\n", pktLen, pktLen, size)
	}

	//cast/return correct type
	m, compressed := getType(msgType)
	//check if we know the type
	if m != nil {
		data := r.buf[r.pos:r.pos+int(pktLen)]

		if r.LogLevel >= 4 && compressed == true {
			fmt.Printf("compressed: %d %d\n", msgType, msgType &^ 0x70)

		}
		if compressed == true {
			s, err := snappy.DecodedLen(data)
			decoded := make([]byte,s)
			data, err := snappy.Decode(decoded, data)
			//p, s_p := binary.Uvarint(decoded[1:])
			//data = decoded[0:]
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Printf("p:%v s_p:%v\n",p,s_p)
			//fmt.Printf("comp-size:\n %v\n s:%d len(data):%d\n", hex.Dump(data[0:120]), s, len(data))
			//if tick != 82704 {
			//if msgType == 116 {
				err = proto.Unmarshal(data, m)
				if err != nil {
					fmt.Printf("### %s\n",getPacketName(msgType))
					log.Fatalf("compress decode error type:%d\npos:%d\npktLen:%d\n%v", msgType, r.pos, pktLen, err)
				}
			//}
		} else {
			err = proto.Unmarshal(data, m)
			if err != nil {
				log.Fatalf("type:%d\npos:%d\npktLen:%d\n%v", msgType, r.pos, pktLen, err)
			}
		}
		
	} else {
		if r.LogLevel >= 2 {
			fmt.Printf("Unknown msgType:%d with pktLen %d at pos %d\n", msgType, pktLen, r.pos)
		}
	}
	if r.LogLevel >= 4 {
		fmt.Printf("%v\n\n", m)
	}
	r.pos += int(pktLen)
	return msgType, m, err
}

//returns name from demo.proto:enum EDemoCommands
func getPacketName(i uint64) (name string) {
	name = "Unknown Packet"
	//name = demo.EDemoCommands_name[int32(i)]
	if name = demo.EDemoCommands_name[int32(i)]; name == "" {
		//maybe it's compressed
		if name = demo.EDemoCommands_name[int32(i &^ 0x70)]; name != "" {
			name = "Compressed " + name
		}
	}
	return name
}

// naming scheme from demo.proto:enum EDemoCommands
// no idea what they were smoking. probably just grown/legacy
//DEM_FileHeader -> CDemoFileHeader 
func getType(msgType uint64) (m proto.Message, compressed bool) {
	if (msgType & 0x70)==0x70 {
		msgType = msgType &^ 0x70
		compressed = true
	}
	switch msgType {
	//undefined in valves .proto
	//-1:&CDemoError{},
	case 0:
		m = &demo.CDemoStop{}
	case 1:
		m = &demo.CDemoFileHeader{}
	case 2:
		m = &demo.CDemoFileInfo{}
	case 3:
		m = &demo.CDemoSyncTick{}
	case 4:
		m = &demo.CDemoSendTables{}
	case 5:
		m = &demo.CDemoClassInfo{}
	case 6:
		m = &demo.CDemoStringTables{}
	case 7:
		m = &demo.CDemoPacket{}

	//undefined in valves .proto
	//8:&CDemoSignonPacket{},

	case 9:
		m = &demo.CDemoConsoleCmd{}
	case 10:
		m = &demo.CDemoCustomData{}
	case 11:
		m = &demo.CDemoCustomDataCallbacks{}
	case 12:
		m = &demo.CDemoUserCmd{}
	case 13:
		m = &demo.CDemoFullPacket{}

		//undefined in valves .proto
		//14:&CDemoMax{},

		// pseudo message, indication compression
		// if messageID is > 0x70 binary OR it and decompress into that
		//0x70:&CDemoIsCompressed{},
	}
	return
}
