package mhfpacket

import (
	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/byteframe"
)

// MsgMhfGetUdMyPoint represents the MSG_MHF_GET_UD_MY_POINT
type MsgMhfGetUdMyPoint struct {
	AckHandle uint32
}

// Opcode returns the ID associated with this packet type.
func (m *MsgMhfGetUdMyPoint) Opcode() network.PacketID {
	return network.MSG_MHF_GET_UD_MY_POINT
}

// Parse parses the packet from binary
func (m *MsgMhfGetUdMyPoint) Parse(bf *byteframe.ByteFrame) error {
	m.AckHandle = bf.ReadUint32()
	return nil
}

// Build builds a binary packet from the current data.
func (m *MsgMhfGetUdMyPoint) Build(bf *byteframe.ByteFrame) error {
	panic("Not implemented")
}
