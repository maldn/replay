// Code generated by protoc-gen-go.
// source: usermessages.proto
// DO NOT EDIT!

package dota2

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"
import dota21 "netmessages.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type EBaseUserMessages int32

const (
	EBaseUserMessages_UM_AchievementEvent   EBaseUserMessages = 1
	EBaseUserMessages_UM_CloseCaption       EBaseUserMessages = 2
	EBaseUserMessages_UM_CloseCaptionDirect EBaseUserMessages = 3
	EBaseUserMessages_UM_CurrentTimescale   EBaseUserMessages = 4
	EBaseUserMessages_UM_DesiredTimescale   EBaseUserMessages = 5
	EBaseUserMessages_UM_Fade               EBaseUserMessages = 6
	EBaseUserMessages_UM_GameTitle          EBaseUserMessages = 7
	EBaseUserMessages_UM_Geiger             EBaseUserMessages = 8
	EBaseUserMessages_UM_HintText           EBaseUserMessages = 9
	EBaseUserMessages_UM_HudMsg             EBaseUserMessages = 10
	EBaseUserMessages_UM_HudText            EBaseUserMessages = 11
	EBaseUserMessages_UM_KeyHintText        EBaseUserMessages = 12
	EBaseUserMessages_UM_MessageText        EBaseUserMessages = 13
	EBaseUserMessages_UM_RequestState       EBaseUserMessages = 14
	EBaseUserMessages_UM_ResetHUD           EBaseUserMessages = 15
	EBaseUserMessages_UM_Rumble             EBaseUserMessages = 16
	EBaseUserMessages_UM_SayText            EBaseUserMessages = 17
	EBaseUserMessages_UM_SayText2           EBaseUserMessages = 18
	EBaseUserMessages_UM_SayTextChannel     EBaseUserMessages = 19
	EBaseUserMessages_UM_Shake              EBaseUserMessages = 20
	EBaseUserMessages_UM_ShakeDir           EBaseUserMessages = 21
	EBaseUserMessages_UM_StatsCrawlMsg      EBaseUserMessages = 22
	EBaseUserMessages_UM_StatsSkipState     EBaseUserMessages = 23
	EBaseUserMessages_UM_TextMsg            EBaseUserMessages = 24
	EBaseUserMessages_UM_Tilt               EBaseUserMessages = 25
	EBaseUserMessages_UM_Train              EBaseUserMessages = 26
	EBaseUserMessages_UM_VGUIMenu           EBaseUserMessages = 27
	EBaseUserMessages_UM_VoiceMask          EBaseUserMessages = 28
	EBaseUserMessages_UM_VoiceSubtitle      EBaseUserMessages = 29
	EBaseUserMessages_UM_SendAudio          EBaseUserMessages = 30
	EBaseUserMessages_UM_MAX_BASE           EBaseUserMessages = 63
)

var EBaseUserMessages_name = map[int32]string{
	1:  "UM_AchievementEvent",
	2:  "UM_CloseCaption",
	3:  "UM_CloseCaptionDirect",
	4:  "UM_CurrentTimescale",
	5:  "UM_DesiredTimescale",
	6:  "UM_Fade",
	7:  "UM_GameTitle",
	8:  "UM_Geiger",
	9:  "UM_HintText",
	10: "UM_HudMsg",
	11: "UM_HudText",
	12: "UM_KeyHintText",
	13: "UM_MessageText",
	14: "UM_RequestState",
	15: "UM_ResetHUD",
	16: "UM_Rumble",
	17: "UM_SayText",
	18: "UM_SayText2",
	19: "UM_SayTextChannel",
	20: "UM_Shake",
	21: "UM_ShakeDir",
	22: "UM_StatsCrawlMsg",
	23: "UM_StatsSkipState",
	24: "UM_TextMsg",
	25: "UM_Tilt",
	26: "UM_Train",
	27: "UM_VGUIMenu",
	28: "UM_VoiceMask",
	29: "UM_VoiceSubtitle",
	30: "UM_SendAudio",
	63: "UM_MAX_BASE",
}
var EBaseUserMessages_value = map[string]int32{
	"UM_AchievementEvent":   1,
	"UM_CloseCaption":       2,
	"UM_CloseCaptionDirect": 3,
	"UM_CurrentTimescale":   4,
	"UM_DesiredTimescale":   5,
	"UM_Fade":               6,
	"UM_GameTitle":          7,
	"UM_Geiger":             8,
	"UM_HintText":           9,
	"UM_HudMsg":             10,
	"UM_HudText":            11,
	"UM_KeyHintText":        12,
	"UM_MessageText":        13,
	"UM_RequestState":       14,
	"UM_ResetHUD":           15,
	"UM_Rumble":             16,
	"UM_SayText":            17,
	"UM_SayText2":           18,
	"UM_SayTextChannel":     19,
	"UM_Shake":              20,
	"UM_ShakeDir":           21,
	"UM_StatsCrawlMsg":      22,
	"UM_StatsSkipState":     23,
	"UM_TextMsg":            24,
	"UM_Tilt":               25,
	"UM_Train":              26,
	"UM_VGUIMenu":           27,
	"UM_VoiceMask":          28,
	"UM_VoiceSubtitle":      29,
	"UM_SendAudio":          30,
	"UM_MAX_BASE":           63,
}

func (x EBaseUserMessages) Enum() *EBaseUserMessages {
	p := new(EBaseUserMessages)
	*p = x
	return p
}
func (x EBaseUserMessages) String() string {
	return proto.EnumName(EBaseUserMessages_name, int32(x))
}
func (x EBaseUserMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *EBaseUserMessages) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EBaseUserMessages_value, data, "EBaseUserMessages")
	if err != nil {
		return err
	}
	*x = EBaseUserMessages(value)
	return nil
}

type CUserMsg_AchievementEvent struct {
	Achievement      *uint32 `protobuf:"varint,1,opt,name=achievement" json:"achievement,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_AchievementEvent) Reset()         { *m = CUserMsg_AchievementEvent{} }
func (m *CUserMsg_AchievementEvent) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_AchievementEvent) ProtoMessage()    {}

func (m *CUserMsg_AchievementEvent) GetAchievement() uint32 {
	if m != nil && m.Achievement != nil {
		return *m.Achievement
	}
	return 0
}

type CUserMsg_CloseCaption struct {
	Hash             *uint32  `protobuf:"fixed32,1,opt,name=hash" json:"hash,omitempty"`
	Duration         *float32 `protobuf:"fixed32,2,opt,name=duration" json:"duration,omitempty"`
	FromPlayer       *bool    `protobuf:"varint,3,opt,name=from_player" json:"from_player,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CUserMsg_CloseCaption) Reset()         { *m = CUserMsg_CloseCaption{} }
func (m *CUserMsg_CloseCaption) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_CloseCaption) ProtoMessage()    {}

func (m *CUserMsg_CloseCaption) GetHash() uint32 {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return 0
}

func (m *CUserMsg_CloseCaption) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *CUserMsg_CloseCaption) GetFromPlayer() bool {
	if m != nil && m.FromPlayer != nil {
		return *m.FromPlayer
	}
	return false
}

type CUserMsg_CurrentTimescale struct {
	Current          *float32 `protobuf:"fixed32,1,opt,name=current" json:"current,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CUserMsg_CurrentTimescale) Reset()         { *m = CUserMsg_CurrentTimescale{} }
func (m *CUserMsg_CurrentTimescale) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_CurrentTimescale) ProtoMessage()    {}

func (m *CUserMsg_CurrentTimescale) GetCurrent() float32 {
	if m != nil && m.Current != nil {
		return *m.Current
	}
	return 0
}

type CUserMsg_DesiredTimescale struct {
	Desired          *float32 `protobuf:"fixed32,1,opt,name=desired" json:"desired,omitempty"`
	Duration         *float32 `protobuf:"fixed32,2,opt,name=duration" json:"duration,omitempty"`
	Interpolator     *uint32  `protobuf:"varint,3,opt,name=interpolator" json:"interpolator,omitempty"`
	StartBlendTime   *float32 `protobuf:"fixed32,4,opt,name=start_blend_time" json:"start_blend_time,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CUserMsg_DesiredTimescale) Reset()         { *m = CUserMsg_DesiredTimescale{} }
func (m *CUserMsg_DesiredTimescale) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_DesiredTimescale) ProtoMessage()    {}

func (m *CUserMsg_DesiredTimescale) GetDesired() float32 {
	if m != nil && m.Desired != nil {
		return *m.Desired
	}
	return 0
}

func (m *CUserMsg_DesiredTimescale) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *CUserMsg_DesiredTimescale) GetInterpolator() uint32 {
	if m != nil && m.Interpolator != nil {
		return *m.Interpolator
	}
	return 0
}

func (m *CUserMsg_DesiredTimescale) GetStartBlendTime() float32 {
	if m != nil && m.StartBlendTime != nil {
		return *m.StartBlendTime
	}
	return 0
}

type CUserMsg_Fade struct {
	Duration         *uint32 `protobuf:"varint,1,opt,name=duration" json:"duration,omitempty"`
	HoldTime         *uint32 `protobuf:"varint,2,opt,name=hold_time" json:"hold_time,omitempty"`
	Flags            *uint32 `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	Color            *uint32 `protobuf:"fixed32,4,opt,name=color" json:"color,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_Fade) Reset()         { *m = CUserMsg_Fade{} }
func (m *CUserMsg_Fade) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_Fade) ProtoMessage()    {}

func (m *CUserMsg_Fade) GetDuration() uint32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *CUserMsg_Fade) GetHoldTime() uint32 {
	if m != nil && m.HoldTime != nil {
		return *m.HoldTime
	}
	return 0
}

func (m *CUserMsg_Fade) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *CUserMsg_Fade) GetColor() uint32 {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return 0
}

type CUserMsg_Shake struct {
	Command          *uint32  `protobuf:"varint,1,opt,name=command" json:"command,omitempty"`
	Amplitude        *float32 `protobuf:"fixed32,2,opt,name=amplitude" json:"amplitude,omitempty"`
	Frequency        *float32 `protobuf:"fixed32,3,opt,name=frequency" json:"frequency,omitempty"`
	Duration         *float32 `protobuf:"fixed32,4,opt,name=duration" json:"duration,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CUserMsg_Shake) Reset()         { *m = CUserMsg_Shake{} }
func (m *CUserMsg_Shake) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_Shake) ProtoMessage()    {}

func (m *CUserMsg_Shake) GetCommand() uint32 {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return 0
}

func (m *CUserMsg_Shake) GetAmplitude() float32 {
	if m != nil && m.Amplitude != nil {
		return *m.Amplitude
	}
	return 0
}

func (m *CUserMsg_Shake) GetFrequency() float32 {
	if m != nil && m.Frequency != nil {
		return *m.Frequency
	}
	return 0
}

func (m *CUserMsg_Shake) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

type CUserMsg_ShakeDir struct {
	Shake            *CUserMsg_Shake    `protobuf:"bytes,1,opt,name=shake" json:"shake,omitempty"`
	Direction        *dota21.CMsgVector `protobuf:"bytes,2,opt,name=direction" json:"direction,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CUserMsg_ShakeDir) Reset()         { *m = CUserMsg_ShakeDir{} }
func (m *CUserMsg_ShakeDir) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_ShakeDir) ProtoMessage()    {}

func (m *CUserMsg_ShakeDir) GetShake() *CUserMsg_Shake {
	if m != nil {
		return m.Shake
	}
	return nil
}

func (m *CUserMsg_ShakeDir) GetDirection() *dota21.CMsgVector {
	if m != nil {
		return m.Direction
	}
	return nil
}

type CUserMsg_Tilt struct {
	Command          *uint32            `protobuf:"varint,1,opt,name=command" json:"command,omitempty"`
	EaseInOut        *bool              `protobuf:"varint,2,opt,name=ease_in_out" json:"ease_in_out,omitempty"`
	Angle            *dota21.CMsgVector `protobuf:"bytes,3,opt,name=angle" json:"angle,omitempty"`
	Duration         *float32           `protobuf:"fixed32,4,opt,name=duration" json:"duration,omitempty"`
	Time             *float32           `protobuf:"fixed32,5,opt,name=time" json:"time,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CUserMsg_Tilt) Reset()         { *m = CUserMsg_Tilt{} }
func (m *CUserMsg_Tilt) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_Tilt) ProtoMessage()    {}

func (m *CUserMsg_Tilt) GetCommand() uint32 {
	if m != nil && m.Command != nil {
		return *m.Command
	}
	return 0
}

func (m *CUserMsg_Tilt) GetEaseInOut() bool {
	if m != nil && m.EaseInOut != nil {
		return *m.EaseInOut
	}
	return false
}

func (m *CUserMsg_Tilt) GetAngle() *dota21.CMsgVector {
	if m != nil {
		return m.Angle
	}
	return nil
}

func (m *CUserMsg_Tilt) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *CUserMsg_Tilt) GetTime() float32 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type CUserMsg_SayText struct {
	Client           *uint32 `protobuf:"varint,1,opt,name=client" json:"client,omitempty"`
	Text             *string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Chat             *bool   `protobuf:"varint,3,opt,name=chat" json:"chat,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_SayText) Reset()         { *m = CUserMsg_SayText{} }
func (m *CUserMsg_SayText) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_SayText) ProtoMessage()    {}

func (m *CUserMsg_SayText) GetClient() uint32 {
	if m != nil && m.Client != nil {
		return *m.Client
	}
	return 0
}

func (m *CUserMsg_SayText) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *CUserMsg_SayText) GetChat() bool {
	if m != nil && m.Chat != nil {
		return *m.Chat
	}
	return false
}

type CUserMsg_SayText2 struct {
	Client           *uint32 `protobuf:"varint,1,opt,name=client" json:"client,omitempty"`
	Chat             *bool   `protobuf:"varint,2,opt,name=chat" json:"chat,omitempty"`
	Format           *string `protobuf:"bytes,3,opt,name=format" json:"format,omitempty"`
	Prefix           *string `protobuf:"bytes,4,opt,name=prefix" json:"prefix,omitempty"`
	Text             *string `protobuf:"bytes,5,opt,name=text" json:"text,omitempty"`
	Location         *string `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_SayText2) Reset()         { *m = CUserMsg_SayText2{} }
func (m *CUserMsg_SayText2) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_SayText2) ProtoMessage()    {}

func (m *CUserMsg_SayText2) GetClient() uint32 {
	if m != nil && m.Client != nil {
		return *m.Client
	}
	return 0
}

func (m *CUserMsg_SayText2) GetChat() bool {
	if m != nil && m.Chat != nil {
		return *m.Chat
	}
	return false
}

func (m *CUserMsg_SayText2) GetFormat() string {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return ""
}

func (m *CUserMsg_SayText2) GetPrefix() string {
	if m != nil && m.Prefix != nil {
		return *m.Prefix
	}
	return ""
}

func (m *CUserMsg_SayText2) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *CUserMsg_SayText2) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

type CUserMsg_HudMsg struct {
	Channel          *uint32  `protobuf:"varint,1,opt,name=channel" json:"channel,omitempty"`
	X                *float32 `protobuf:"fixed32,2,opt,name=x" json:"x,omitempty"`
	Y                *float32 `protobuf:"fixed32,3,opt,name=y" json:"y,omitempty"`
	Color1           *uint32  `protobuf:"varint,4,opt,name=color1" json:"color1,omitempty"`
	Color2           *uint32  `protobuf:"varint,5,opt,name=color2" json:"color2,omitempty"`
	Effect           *uint32  `protobuf:"varint,6,opt,name=effect" json:"effect,omitempty"`
	FadeInTime       *float32 `protobuf:"fixed32,7,opt,name=fade_in_time" json:"fade_in_time,omitempty"`
	FadeOutTime      *float32 `protobuf:"fixed32,8,opt,name=fade_out_time" json:"fade_out_time,omitempty"`
	HoldTime         *float32 `protobuf:"fixed32,9,opt,name=hold_time" json:"hold_time,omitempty"`
	FxTime           *float32 `protobuf:"fixed32,10,opt,name=fx_time" json:"fx_time,omitempty"`
	Message          *string  `protobuf:"bytes,11,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CUserMsg_HudMsg) Reset()         { *m = CUserMsg_HudMsg{} }
func (m *CUserMsg_HudMsg) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_HudMsg) ProtoMessage()    {}

func (m *CUserMsg_HudMsg) GetChannel() uint32 {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetColor1() uint32 {
	if m != nil && m.Color1 != nil {
		return *m.Color1
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetColor2() uint32 {
	if m != nil && m.Color2 != nil {
		return *m.Color2
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetEffect() uint32 {
	if m != nil && m.Effect != nil {
		return *m.Effect
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetFadeInTime() float32 {
	if m != nil && m.FadeInTime != nil {
		return *m.FadeInTime
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetFadeOutTime() float32 {
	if m != nil && m.FadeOutTime != nil {
		return *m.FadeOutTime
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetHoldTime() float32 {
	if m != nil && m.HoldTime != nil {
		return *m.HoldTime
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetFxTime() float32 {
	if m != nil && m.FxTime != nil {
		return *m.FxTime
	}
	return 0
}

func (m *CUserMsg_HudMsg) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type CUserMsg_HudText struct {
	Message          *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_HudText) Reset()         { *m = CUserMsg_HudText{} }
func (m *CUserMsg_HudText) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_HudText) ProtoMessage()    {}

func (m *CUserMsg_HudText) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type CUserMsg_TextMsg struct {
	Dest             *uint32  `protobuf:"varint,1,opt,name=dest" json:"dest,omitempty"`
	Param            []string `protobuf:"bytes,2,rep,name=param" json:"param,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CUserMsg_TextMsg) Reset()         { *m = CUserMsg_TextMsg{} }
func (m *CUserMsg_TextMsg) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_TextMsg) ProtoMessage()    {}

func (m *CUserMsg_TextMsg) GetDest() uint32 {
	if m != nil && m.Dest != nil {
		return *m.Dest
	}
	return 0
}

func (m *CUserMsg_TextMsg) GetParam() []string {
	if m != nil {
		return m.Param
	}
	return nil
}

type CUserMsg_GameTitle struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_GameTitle) Reset()         { *m = CUserMsg_GameTitle{} }
func (m *CUserMsg_GameTitle) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_GameTitle) ProtoMessage()    {}

type CUserMsg_ResetHUD struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_ResetHUD) Reset()         { *m = CUserMsg_ResetHUD{} }
func (m *CUserMsg_ResetHUD) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_ResetHUD) ProtoMessage()    {}

type CUserMsg_SendAudio struct {
	Stop             *bool   `protobuf:"varint,2,opt,name=stop" json:"stop,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_SendAudio) Reset()         { *m = CUserMsg_SendAudio{} }
func (m *CUserMsg_SendAudio) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_SendAudio) ProtoMessage()    {}

func (m *CUserMsg_SendAudio) GetStop() bool {
	if m != nil && m.Stop != nil {
		return *m.Stop
	}
	return false
}

func (m *CUserMsg_SendAudio) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type CUserMsg_VoiceMask struct {
	AudiblePlayersMask []int32 `protobuf:"varint,1,rep,name=audible_players_mask" json:"audible_players_mask,omitempty"`
	PlayerModEnabled   *bool   `protobuf:"varint,2,opt,name=player_mod_enabled" json:"player_mod_enabled,omitempty"`
	XXX_unrecognized   []byte  `json:"-"`
}

func (m *CUserMsg_VoiceMask) Reset()         { *m = CUserMsg_VoiceMask{} }
func (m *CUserMsg_VoiceMask) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_VoiceMask) ProtoMessage()    {}

func (m *CUserMsg_VoiceMask) GetAudiblePlayersMask() []int32 {
	if m != nil {
		return m.AudiblePlayersMask
	}
	return nil
}

func (m *CUserMsg_VoiceMask) GetPlayerModEnabled() bool {
	if m != nil && m.PlayerModEnabled != nil {
		return *m.PlayerModEnabled
	}
	return false
}

type CUserMsg_RequestState struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_RequestState) Reset()         { *m = CUserMsg_RequestState{} }
func (m *CUserMsg_RequestState) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_RequestState) ProtoMessage()    {}

type CUserMsg_HintText struct {
	Message          *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_HintText) Reset()         { *m = CUserMsg_HintText{} }
func (m *CUserMsg_HintText) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_HintText) ProtoMessage()    {}

func (m *CUserMsg_HintText) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type CUserMsg_KeyHintText struct {
	Messages         []string `protobuf:"bytes,1,rep,name=messages" json:"messages,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CUserMsg_KeyHintText) Reset()         { *m = CUserMsg_KeyHintText{} }
func (m *CUserMsg_KeyHintText) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_KeyHintText) ProtoMessage()    {}

func (m *CUserMsg_KeyHintText) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

type CUserMsg_StatsCrawlMsg struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_StatsCrawlMsg) Reset()         { *m = CUserMsg_StatsCrawlMsg{} }
func (m *CUserMsg_StatsCrawlMsg) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_StatsCrawlMsg) ProtoMessage()    {}

type CUserMsg_StatsSkipState struct {
	NumSkips         *int32 `protobuf:"varint,1,opt,name=num_skips" json:"num_skips,omitempty"`
	NumPlayers       *int32 `protobuf:"varint,2,opt,name=num_players" json:"num_players,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_StatsSkipState) Reset()         { *m = CUserMsg_StatsSkipState{} }
func (m *CUserMsg_StatsSkipState) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_StatsSkipState) ProtoMessage()    {}

func (m *CUserMsg_StatsSkipState) GetNumSkips() int32 {
	if m != nil && m.NumSkips != nil {
		return *m.NumSkips
	}
	return 0
}

func (m *CUserMsg_StatsSkipState) GetNumPlayers() int32 {
	if m != nil && m.NumPlayers != nil {
		return *m.NumPlayers
	}
	return 0
}

type CUserMsg_VoiceSubtitle struct {
	EntIndex         *int32 `protobuf:"varint,1,opt,name=ent_index" json:"ent_index,omitempty"`
	Menu             *int32 `protobuf:"varint,2,opt,name=menu" json:"menu,omitempty"`
	Item             *int32 `protobuf:"varint,3,opt,name=item" json:"item,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_VoiceSubtitle) Reset()         { *m = CUserMsg_VoiceSubtitle{} }
func (m *CUserMsg_VoiceSubtitle) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_VoiceSubtitle) ProtoMessage()    {}

func (m *CUserMsg_VoiceSubtitle) GetEntIndex() int32 {
	if m != nil && m.EntIndex != nil {
		return *m.EntIndex
	}
	return 0
}

func (m *CUserMsg_VoiceSubtitle) GetMenu() int32 {
	if m != nil && m.Menu != nil {
		return *m.Menu
	}
	return 0
}

func (m *CUserMsg_VoiceSubtitle) GetItem() int32 {
	if m != nil && m.Item != nil {
		return *m.Item
	}
	return 0
}

type CUserMsg_VGUIMenu struct {
	Name             *string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Show             *bool                     `protobuf:"varint,2,opt,name=show" json:"show,omitempty"`
	Keys             []*CUserMsg_VGUIMenu_Keys `protobuf:"bytes,3,rep,name=keys" json:"keys,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *CUserMsg_VGUIMenu) Reset()         { *m = CUserMsg_VGUIMenu{} }
func (m *CUserMsg_VGUIMenu) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_VGUIMenu) ProtoMessage()    {}

func (m *CUserMsg_VGUIMenu) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CUserMsg_VGUIMenu) GetShow() bool {
	if m != nil && m.Show != nil {
		return *m.Show
	}
	return false
}

func (m *CUserMsg_VGUIMenu) GetKeys() []*CUserMsg_VGUIMenu_Keys {
	if m != nil {
		return m.Keys
	}
	return nil
}

type CUserMsg_VGUIMenu_Keys struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value            *string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_VGUIMenu_Keys) Reset()         { *m = CUserMsg_VGUIMenu_Keys{} }
func (m *CUserMsg_VGUIMenu_Keys) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_VGUIMenu_Keys) ProtoMessage()    {}

func (m *CUserMsg_VGUIMenu_Keys) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CUserMsg_VGUIMenu_Keys) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type CUserMsg_Geiger struct {
	Range            *int32 `protobuf:"varint,1,opt,name=range" json:"range,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_Geiger) Reset()         { *m = CUserMsg_Geiger{} }
func (m *CUserMsg_Geiger) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_Geiger) ProtoMessage()    {}

func (m *CUserMsg_Geiger) GetRange() int32 {
	if m != nil && m.Range != nil {
		return *m.Range
	}
	return 0
}

type CUserMsg_Rumble struct {
	Index            *int32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Data             *int32 `protobuf:"varint,2,opt,name=data" json:"data,omitempty"`
	Flags            *int32 `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_Rumble) Reset()         { *m = CUserMsg_Rumble{} }
func (m *CUserMsg_Rumble) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_Rumble) ProtoMessage()    {}

func (m *CUserMsg_Rumble) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *CUserMsg_Rumble) GetData() int32 {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return 0
}

func (m *CUserMsg_Rumble) GetFlags() int32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

type CUserMsg_Train struct {
	Train            *int32 `protobuf:"varint,1,opt,name=train" json:"train,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CUserMsg_Train) Reset()         { *m = CUserMsg_Train{} }
func (m *CUserMsg_Train) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_Train) ProtoMessage()    {}

func (m *CUserMsg_Train) GetTrain() int32 {
	if m != nil && m.Train != nil {
		return *m.Train
	}
	return 0
}

type CUserMsg_SayTextChannel struct {
	Player           *int32  `protobuf:"varint,1,opt,name=player" json:"player,omitempty"`
	Channel          *int32  `protobuf:"varint,2,opt,name=channel" json:"channel,omitempty"`
	Text             *string `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_SayTextChannel) Reset()         { *m = CUserMsg_SayTextChannel{} }
func (m *CUserMsg_SayTextChannel) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_SayTextChannel) ProtoMessage()    {}

func (m *CUserMsg_SayTextChannel) GetPlayer() int32 {
	if m != nil && m.Player != nil {
		return *m.Player
	}
	return 0
}

func (m *CUserMsg_SayTextChannel) GetChannel() int32 {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return 0
}

func (m *CUserMsg_SayTextChannel) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type CUserMsg_MessageText struct {
	Color            *uint32 `protobuf:"varint,1,opt,name=color" json:"color,omitempty"`
	Text             *string `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CUserMsg_MessageText) Reset()         { *m = CUserMsg_MessageText{} }
func (m *CUserMsg_MessageText) String() string { return proto.CompactTextString(m) }
func (*CUserMsg_MessageText) ProtoMessage()    {}

func (m *CUserMsg_MessageText) GetColor() uint32 {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return 0
}

func (m *CUserMsg_MessageText) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func init() {
	proto.RegisterEnum("dota2.EBaseUserMessages", EBaseUserMessages_name, EBaseUserMessages_value)
}
