/**
 * Copyright (c) 2013 Malte Böhme
 * Licensed under the MIT license.
 */

package dota2

import (
	//replay "demo.pb" // generated files `protoc --go_out=generated *.proto`
	//replay "./proto/demo.pb"
	"code.google.com/p/goprotobuf/proto"
	"code.google.com/p/snappy-go/snappy"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	//"bytes"
	//"encoding/hex"
	//"os"
)

type Replay struct {
	buf       []byte
	size, pos int // replay size, current position
	LogLevel  int
	PktCount map[string]int
}

type Packet struct {
	Name string `json:"name"`
	Id uint64 `json:"msg_id"`
	Tick uint64 `json:"tick"`
	Pos int `json:"position"`
	Size int `json:"size"`
	CompressedSize int `json:"compressed_size"`
	Data interface{}`json:"data"`
}
func (p *Packet) String() (s string) {
	j, err := json.MarshalIndent(p,"","  ")
	s = string(j)
	if err != nil {
		return fmt.Sprintf("%v",p)
	}
	return s

}
func NewReplay(filename string) (r *Replay, err error) {
	r = new(Replay)
	r.buf, err = ioutil.ReadFile(filename)
	r.size = len(r.buf)
	//FIXME: validate header
	//header is a 8 byte string + 32bit size of replay file.
	r.pos = 12
	r.PktCount = make(map[string]int)
	return r, err
}

func (r *Replay) ReadPacket() (p *Packet, err error) {
	p = &Packet{}
	if r.pos >= r.size {
		err_msg := fmt.Sprintf("end of replay.. size:%d\tpos:%d\n", r.size, r.pos)
		err = errors.New(err_msg)
		return nil, err
	}


	// no idea if its a bug in go-protobuf or valves stuff
	// valve defines a msgType of -1 as DEM_Error but seems like the encoding is Uvarint..
	// strange..
	msgType, size := binary.Uvarint(r.buf[r.pos:])
	r.pos += size
	if r.LogLevel >= 3 {
		fmt.Printf("%s: msgType: %#x(%d) size:%d\n", getPacketName(msgType), msgType, msgType, size)
	}
	// we overwrite this in case of a compressed packet
	p.Id = msgType

	tick, size := binary.Uvarint(r.buf[r.pos:])
	r.pos += size
	if r.LogLevel >= 3 {
		fmt.Printf("tick: %#x(%d) len:%d\n", tick, tick, size)
	}
	p.Tick = tick

	pktLen, size := binary.Uvarint(r.buf[r.pos:])
	r.pos += size
	if r.LogLevel >= 3 {
		fmt.Printf("length: %d (%#x) %d\n", pktLen, pktLen, size)
	}
	p.CompressedSize =int(pktLen)
	// we overwrite this in case of a compressed packet
	p.Size = int(pktLen)

	//cast/return correct type
	m, compressed := r.getType(msgType)
	//check if we know the type
	if m != nil {
		data := r.buf[r.pos:r.pos+int(pktLen)]


		if compressed == true {
			if r.LogLevel >= 4 {
				fmt.Printf("compressed: %d %d\n", msgType, msgType &^ 0x70)
			}
			// we really only care about actual message IDs, not the compression indicator
			p.Id = msgType &^ 0x70

			s, err := snappy.DecodedLen(data)
			decoded := make([]byte,s)
			data, err := snappy.Decode(decoded, data)
			//p, s_p := binary.Uvarint(decoded[1:])
			//data = decoded[0:]
			if err != nil {
				log.Fatal(err)
			}
			// set Size to decompressed size, we still have CompressedSize (was set above)
			p.Size = s
			
			err = proto.Unmarshal(data, m)
			if err != nil {
				fmt.Printf("### %s\n",getPacketName(msgType))
				log.Fatalf("compress decode error type:%d\npos:%d\npktLen:%d\n%v", msgType, r.pos, pktLen, err)
			}

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
	//gather some stats about distribution of packets
	r.PktCount[getPacketName(msgType)]++
	

	p.Name = getPacketName(p.Id)
	p.Data = m
	p.Pos = r.pos

	r.pos += int(pktLen)
	return p, err

}

//returns name from demo.proto:enum EDemoCommands
func getPacketName(i uint64) (name string) {
	name = "Unknown Packet"
	if name = EDemoCommands_name[int32(i)]; name == "" {
		//maybe it's compressed
		if name = EDemoCommands_name[int32(i &^ 0x70)]; name != "" {
			name = "Compressed " + name
		}
	}
	return name
}

// naming scheme from demo.proto:enum EDemoCommands
// no idea what they were smoking. probably just grown/legacy
//DEM_FileHeader -> FileHeader
//
// no need to cache objects here, allocation is fast
// allocating 26111 DemoPackets takes 2 ms

// i tried hard to make this switch go away, but there is no way with go and protocol buffers
// to get to the type of a message from its msgID
// go's reflection is unable to inspect and get types from packages. you have to know the types beforehand
// and valve didnt encoded the msgID, tick and msgSize in a protobuf-message
// so that it would be possible to embed the messages in another protobuf message
// ... sucks :-)
func (r *Replay) getType(msgType uint64) (m proto.Message, compressed bool) {
	if (msgType & 0x70)==0x70 {
		msgType = msgType &^ 0x70
		compressed = true
	}
	switch msgType {
	//undefined in valves .proto
	//-1:&Error{},
	case 0:
		m = &CDemoStop{}
	case 1:
		m = &CDemoFileHeader{}
	case 2:
		m = &CDemoFileInfo{}
	case 3:
		m = &CDemoSyncTick{}
	case 4:
		m = &CDemoSendTables{}
	case 5:
		m = &CDemoClassInfo{}
	case 6:
		m = &CDemoStringTables{}
	case 7:
		m = &CDemoPacket{}
		
	//undefined in valves .proto
	// but by me :)
	case 8:
		m= &CDemoSignonPacket{}
	case 9:
		m = &CDemoConsoleCmd{}
	case 10:
		m = &CDemoCustomData{}
	case 11:
		m = &CDemoCustomDataCallbacks{}
	case 12:
		m = &CDemoUserCmd{}
	case 13:
		m = &CDemoFullPacket{}

		//undefined in valves .proto
		//14:&Max{},

		// pseudo message, indication compression
		// if messageID has the bits 0x70 set, binary OR it and decompress into that
		//0x70:&IsCompressed{},
	}
	return
}
