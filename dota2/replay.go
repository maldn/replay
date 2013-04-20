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
	PktCount  map[string]int
	Packets   []*Packet `json:"packets`
}

type GameMessage struct {
	Name string        `json:"name"`
	Id   uint64        `json:"msg_id"`
	Data proto.Message `json:"data"`
}

type Packet struct {
	Name           string        `json:"name"`
	Id             uint64        `json:"msg_id"`
	Tick           uint64        `json:"tick"`
	Pos            int           `json:"position"`
	Size           int           `json:"size"`
	Compressed     bool          `json:"compressed"`
	CompressedSize int           `json:"compressed_size"`
	Messages       []GameMessage `json:"messages"`
	Data           proto.Message `json:"data"`
}

func (p *Packet) String() (s string) {
	j, err := json.MarshalIndent(p, "", "  ")
	s = string(j)
	if err != nil {
		return fmt.Sprintf("%v", p)
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
	r.Packets = make([]*Packet, 0, 1000000)
	return r, err
}
func (r *Replay) Parse() (err error) {
	for {
		//TODO: just return, no need for an error
		if r.pos >= r.size {
			err_msg := fmt.Sprintf("end of replay.. size:%d\tpos:%d\n", r.size, r.pos)
			err = errors.New(err_msg)
			return err
		}

		p, size, err := readPacket(r.buf[r.pos:])
		if err != nil {
			return err
		}
		p.Pos = r.pos
		r.pos += size

		//gather some stats about distribution of packets
		r.PktCount[getPacketName(p.Id)]++

		//handle messages that contain other packets
		switch p.Id {
		// DEM_Packet and DEM_SignonPacket have the same structure
		case 7, 8:
			//cast from proto.Message to concrete type, to access .Data
			pkt := p.Data.(*CDemoPacket)
			// we dont want to output the compressed/packed data
			p.Data = nil
			packets, err := ParseGamePacket(pkt.Data)
			p.Messages = packets
			if err != nil {
				return err
			}
		//FullPacket
		case 13:
			// im lazy, so stringTables is data and Packet is decoded to Messages
			pkt := p.Data.(*CDemoFullPacket)
			// we dont want to output the compressed/packed data
			p.Data = pkt.StringTable
			packets, err := ParseGamePacket(pkt.Packet.Data)
			p.Messages = packets

			//tables, err := ParseGamePacket(pkt.StringTable)
			//p.Messages = append(p.Messages, tables...)
			if err != nil {
				return err
			}
		}

		// used to exclude unwanted (huge) messages.
		switch p.Id {
		case 4: //SendTables
		case 5: //classInfo
		case 6: //stringTables
		//case 7://Packet
		//case 13: //fullPacket

		default:
			r.Packets = append(r.Packets, p)

		}
	}
	return
}

	// DOTA_UM_AddUnitToSelection = 		   64;
	// DOTA_UM_AIDebugLine =				   65;
	// DOTA_UM_ChatEvent =  				   66;
	// DOTA_UM_CombatHeroPositions =		   67;
	// DOTA_UM_CombatLogData =  			   68;
	// DOTA_UM_CombatLogShowDeath = 		   70;
	// DOTA_UM_CreateLinearProjectile =	   71;
	// DOTA_UM_DestroyLinearProjectile =	   72;
	// DOTA_UM_DodgeTrackingProjectiles =	   73;
	// DOTA_UM_GlobalLightColor =   		   74;
	// DOTA_UM_GlobalLightDirection =   	   75;
	// DOTA_UM_InvalidCommand = 			   76;
	// DOTA_UM_LocationPing =   			   77;
	// DOTA_UM_MapLine =					   78;
	// DOTA_UM_MiniKillCamInfo =			   79;
	// DOTA_UM_MinimapDebugPoint =  		   80;
	// DOTA_UM_MinimapEvent =   			   81;
	// DOTA_UM_NevermoreRequiem =   		   82;
	// DOTA_UM_OverheadEvent =  			   83;
	// DOTA_UM_SetNextAutobuyItem = 		   84;
	// DOTA_UM_SharedCooldown = 			   85;
	// DOTA_UM_SpectatorPlayerClick =   	   86;
	// DOTA_UM_TutorialTipInfo =			   87;
	// DOTA_UM_UnitEvent =  				   88;
	// DOTA_UM_ParticleManager	= 			   89;
	// DOTA_UM_BotChat =				   	   90;
	// DOTA_UM_HudError = 					   91;
	// DOTA_UM_ItemPurchased =				   92;
	// DOTA_UM_Ping =						   93;
	// DOTA_UM_ItemFound =					   94;


func ParseGamePacket(buf []byte) (gameMessages []GameMessage, err error) {
	var pos int = 0
	gameMessageType, size := binary.Uvarint(buf)
	pos = size
	gameMessageSize, size := binary.Uvarint(buf[pos:])
	pos += size

	gameMessage := GameMessage{}
	gameMessage.Data = nil
	//fmt.Printf("#####gamePacketType: %v\n",gamePacketType)
	gameMessage.Id = gameMessageType


	switch gameMessageType {
	case 0:
		gameMessage.Data = &CNETMsg_NOP{}
	case 1:
		gameMessage.Data = &CNETMsg_Disconnect{}
	case 2:
		gameMessage.Data = &CNETMsg_File{}
	case 3:
		gameMessage.Data = &CNETMsg_SplitScreenUser{}
	case 4:
		gameMessage.Data = &CNETMsg_Tick{}
	case 5:
		gameMessage.Data = &CNETMsg_StringCmd{}
	case 6:
		gameMessage.Data = &CNETMsg_SetConVar{}
	case 7:
		gameMessage.Data = &CNETMsg_SignonState{}
	case 8:// first message from server about game; map etc
		gameMessage.Data = &CSVCMsg_ServerInfo{}
	case 9:// sends a sendtable description for a game class
		gameMessage.Data = &CSVCMsg_SendTable{}
		case 10:// Info about classes (first byte is a CLASSINFO_ define).
		gameMessage.Data = &CSVCMsg_ClassInfo{}
	case 11:// tells client if server paused or unpaused
		gameMessage.Data = &CSVCMsg_SetPause{}
		case 12:// inits shared string tables
		gameMessage.Data = &CSVCMsg_CreateStringTable{}
		case 13:// updates a string table
		gameMessage.Data = &CSVCMsg_UpdateStringTable{}
		case 14:// inits used voice codecs & quality
		gameMessage.Data = &CSVCMsg_VoiceInit{}
		case 15:// Voicestream data from the server
		gameMessage.Data = &CSVCMsg_VoiceData{}
		case 16:// print text to console
		gameMessage.Data = &CSVCMsg_Print{}
		case 17:// starts playing sound
		gameMessage.Data = &CSVCMsg_Sounds{}
		case 18:// sets entity as point of view
		gameMessage.Data = &CSVCMsg_SetView{}
		case 19:// sets/corrects players viewangle
		gameMessage.Data = &CSVCMsg_FixAngle{}
		case 20:// adjusts crosshair in auto aim mode to lock on traget
		gameMessage.Data = &CSVCMsg_CrosshairAngle{}
		case 21:// add a static decal to the world BSP
		gameMessage.Data = &CSVCMsg_BSPDecal{}
		case 22:// split screen style message
		gameMessage.Data = &CSVCMsg_SplitScreen{}
		case 23:// a game specific message 
		gameMessage.Data = &CSVCMsg_UserMessage{}
		
		case 24:// a message for an entity
		//FIXME this is undefined in valves .proto
		//gameMessage.Data = &CSVCMsg_EntityMessage{}
		case 25:// global game event fired
		gameMessage.Data = &CSVCMsg_GameEvent{}
		case 26:// non-delta compressed entities
		gameMessage.Data = &CSVCMsg_PacketEntities{}
		case 27:// non-reliable event object
		gameMessage.Data = &CSVCMsg_TempEntities{}
		case 28:// only sound indices for now
		gameMessage.Data = &CSVCMsg_Prefetch{}
		case 29:// display a menu from a plugin
		gameMessage.Data = &CSVCMsg_Menu{}
		case 30:// list of known games events and fields
		gameMessage.Data = &CSVCMsg_GameEventList{}
		case 31:// Server wants to know the value of a cvar on the client
		gameMessage.Data = &CSVCMsg_GetCvarValue{}

		// no unknown messages in my test-replays so far
	default:
		log.Printf("unknown gameMessage %v",gameMessageType)
	}
	
	// generic way to set Name to type of .Data
	//saves some lines of switch statement
	gameMessage.Name = fmt.Sprintf("%T", gameMessage.Data)

	err = proto.Unmarshal(buf[pos:pos+int(gameMessageSize)], gameMessage.Data)
	
	//special handling to extract embedded messages
	switch gameMessageType {
		case 23: //CSVCMsg_UserMessage
		// we unpack UserMessages directly
		// too bad we can't reassign userMessage.MsgData (its type []byte)
		um := gameMessage.Data.(*CSVCMsg_UserMessage)
		d,_ := ParseUserMessage(um.MsgData, int(*um.MsgType))
		gameMessage.Data = d
	}

	gameMessages = append(gameMessages, gameMessage)
	return
}

type UserMessage struct {
	// we include CSVCMsg_UserMessage so we are as close as possible to 'original' type
	// but we want the decoded message directly in here
	// we cant use CSVCMsg_UserMessage.MsgData as the types mismatch ([]byte vs proto.Message)
	Msg proto.Message
	MsgType int
	CSVCMsg_UserMessage
}

func ParseUserMessage(buf []byte, msgType int) (msg *UserMessage, err error) {
	msg = &UserMessage{MsgType:msgType}

	switch msgType {
	case 66:
		msg.Msg = &CDOTAUserMsg_ChatEvent{}
	case 92:

	}
	err := proto.Unmarshal(buf, msg.Msg)
	if err != nil {
		log.Fatalf("chat: %v\n%v\n", msg.Msg,err)
	}
	return
}

func readPacket(buf []byte) (p *Packet, pos int, err error) {
	// read type, tick number and size from stream/buf
	// then decode the apropriate message and return new/next position in stream
	p = &Packet{}
	pos = 0
	pktType, varint_len := binary.Uvarint(buf)
	p.Id = pktType
	pos += varint_len

	tick, varint_len := binary.Uvarint(buf[pos:])
	p.Tick = tick
	pos += varint_len

	pktLen, varint_len := binary.Uvarint(buf[pos:])
	pos += varint_len
	p.CompressedSize = int(pktLen)
	// we overwrite this later in case of a compressed packet
	p.Size = int(pktLen)

	// raw packet data
	data := buf[pos : pos+p.CompressedSize]

	pkt, compressed := getType(p.Id)
	p.Compressed = compressed

	if p.Compressed == true {
		// we really only care about actual message IDs, not the compression indicator
		p.Id = p.Id &^ 0x70

		decodedLen, err := snappy.DecodedLen(data)
		decoded := make([]byte, decodedLen)
		data, err := snappy.Decode(decoded, data)
		//p, s_p := binary.Uvarint(decoded[1:])
		//data = decoded[0:]
		if err != nil {
			log.Fatal(err)
		}
		// set Size to decompressed size, we still have CompressedSize (was set above)
		p.Size = decodedLen

		err = proto.Unmarshal(data, pkt)
		if err != nil {
			fmt.Printf("### %s\n", getPacketName(p.Id))
			log.Fatalf("compress decode error type:%d\npos:%d\npktLen:%d\n%v", p.Id, pos, p.CompressedSize, err)
		}

	} else {
		err = proto.Unmarshal(data, pkt)
		if err != nil {
			log.Fatalf("type:%d\npos:%d\npktLen:%d\n%v", p.Id, pos, p.CompressedSize, err)
		}
	}
	pos += p.CompressedSize
	p.Name = fmt.Sprintf("%T", pkt)
	p.Data = pkt
	return
}

//returns name from demo.proto:enum EDemoCommands
func getPacketName(i uint64) (name string) {
	name = "Unknown Packet"
	if name = EDemoCommands_name[int32(i)]; name == "" {
		//maybe it's compressed
		if name = EDemoCommands_name[int32(i&^0x70)]; name != "" {
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
func getType(msgType uint64) (m proto.Message, compressed bool) {
	if (msgType & 0x70) == 0x70 {
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
		//m= &CDemoSignonPacket{}
		//same structure
		m = &CDemoPacket{}
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
