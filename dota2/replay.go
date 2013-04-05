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

	"./demo" // generated files `protoc --go_out=. *.proto`
)

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

type Replay struct {
	buf                 []byte
	size, pos, logLevel int // replay size, current position
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
	msgType, size := binary.Uvarint(r.buf[r.pos:])
	fmt.Printf("------\nmsgType: %#x(%d) size:%d\n", msgType, msgType, size)
	fmt.Printf("\t-> %s\n", demo.EDemoCommands_name[int32(msgType)])
	r.pos += size

	tick, size := binary.Uvarint(r.buf[r.pos:])
	fmt.Printf("tick: %#x(%d) len:%d\n", tick, tick, size)
	r.pos += size

	pktLen, size := binary.Uvarint(r.buf[r.pos:])
	fmt.Printf("length: %d (%#x) %d\n", pktLen, pktLen, size)
	r.pos += size

	//cast/return correct type
	m = getType(msgType)
	//check if we know the type
	if m != nil {
		err = proto.Unmarshal(r.buf[r.pos:r.pos+int(pktLen)], m)
		if err != nil {
			log.Fatalf("type:%d\npos:%d\npktLen:%d\n%v", msgType, r.pos, pktLen, err)
		}
	} else {
		fmt.Printf("Unknown msgType:%d with pktLen %d at pos %d\n", msgType, pktLen, r.pos)
	}
	r.pos += int(pktLen)
	return m, err
}