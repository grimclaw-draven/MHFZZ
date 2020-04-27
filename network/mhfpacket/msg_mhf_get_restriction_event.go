package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetRestrictionEvent represents the MSG_MHF_GET_RESTRICTION_EVENT
type MsgMhfGetRestrictionEvent struct{}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetRestrictionEvent) Opcode() network.PacketID {
	return network.MSG_MHF_GET_RESTRICTION_EVENT
}

// Parse parses the packet from binary
func (m *MsgMhfGetRestrictionEvent) Parse(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetRestrictionEvent) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
