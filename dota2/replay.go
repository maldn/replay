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

func (r *Replay) ReadPacket() (m proto.Message, err error) {

	//fmt.Printf("size:%d\tpos:%d",r.size,r.pos)
	if r.pos >= r.size {
		//fmt.Printf("end of replay.. size:%d\tpos:%d",r.size,r.pos)
		err_msg := fmt.Sprintf("end of replay.. size:%d\tpos:%d", r.size, r.pos)
		err = errors.New(err_msg)
		return nil, err
	}

	// no idea if its a bug in go-protobuf or valves stuff
	// valve defines a msgType of -1 as DEM_Error but seems like the encoding is Uvarint..
	// strange..
	msgType, size := binary.Uvarint(r.buf[r.pos:])
	r.pos += size
	if r.LogLevel >= 3 {
		fmt.Printf("%s: msgType: %#x(%d) size:%d\n", getPacketName(uint64(msgType)), msgType, msgType, size)
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
	m = getType(uint64(msgType))
	//check if we know the type
	if m != nil {
		err = proto.Unmarshal(r.buf[r.pos:r.pos+int(pktLen)], m)
		if err != nil {
			log.Fatalf("type:%d\npos:%d\npktLen:%d\n%v", msgType, r.pos, pktLen, err)
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
	return m, err
}

//returns name from demo.proto:enum EDemoCommands
func getPacketName(i uint64) (name string) {
	name = demo.EDemoCommands_name[int32(i)]
	if name == "" {
		name = "Unknown Packet"
	}
	return name
}

// naming scheme from demo.proto:enum EDemoCommands
// no idea what they were smoking. probably just grown/legacy
//DEM_FileHeader -> CDemoFileHeader 
func getType(i uint64) proto.Message {
	switch i {
	//undefined in valves .proto
	//-1:&CDemoError{},
	case 0:
		return &demo.CDemoStop{}
	case 1:
		return &demo.CDemoFileHeader{}
	case 2:
		return &demo.CDemoFileInfo{}
	case 3:
		return &demo.CDemoSyncTick{}
	case 4:
		return &demo.CDemoSendTables{}
	case 5:
		return &demo.CDemoClassInfo{}
	case 6:
		return &demo.CDemoStringTables{}
	case 7:
		return &demo.CDemoPacket{}

	//undefined in valves .proto
	//8:&CDemoSignonPacket{},

	case 9:
		return &demo.CDemoConsoleCmd{}
	case 10:
		return &demo.CDemoCustomData{}
	case 11:
		return &demo.CDemoCustomDataCallbacks{}
	case 12:
		return &demo.CDemoUserCmd{}
	case 13:
		return &demo.CDemoFullPacket{}

		//undefined in valves .proto
		//14:&CDemoMax{},

		// pseudo message, indication compression
		// if messageID is > 0x70 binary OR it and decompress into that
		//0x70:&CDemoIsCompressed{},
	}
	return nil
}
