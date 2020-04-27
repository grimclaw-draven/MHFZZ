package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetGuildScoutList represents the MSG_MHF_GET_GUILD_SCOUT_LIST
type MsgMhfGetGuildScoutList struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetGuildScoutList) Opcode() network.PacketID {
	return network.MSG_MHF_GET_GUILD_SCOUT_LIST
}

// Parse parses the packet from binary
func (m *MsgMhfGetGuildScoutList) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetGuildScoutList) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
