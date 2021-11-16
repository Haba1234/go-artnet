package packet

import "github.com/Haba1234/go-artnet/packet/code"

var _ ArtNetPacket = &ArtTodRequestPacket{}

// ArtTodRequestPacket contains an ArtTodRequest Packet.
type ArtTodRequestPacket struct {
	Header

	// Filler1
	_ byte

	// Filler2
	_ byte

	// Spare
	_ [7]byte

	// Net is the top 7 bits of the 15 bit Port-Address to which this packet is destined
	Net uint8

	// Command
	_ byte

	// AddCount
	_ byte

	// Address
	_ [32]byte
}

// GetOpCode returns the OpCode parsed by validate method.
func (p *ArtTodRequestPacket) GetOpCode() code.OpCode {
	return p.OpCode
}

// MarshalBinary marshals an ArtAddressPacket into a byte slice.
func (p *ArtTodRequestPacket) MarshalBinary() ([]byte, error) {
	return marshalPacket(p)
}

// UnmarshalBinary unmarshal the contents of a byte slice into an ArtAddressPacket.
func (p *ArtTodRequestPacket) UnmarshalBinary(b []byte) error {
	return unmarshalPacket(p, b)
}

// validate is used to validate the Packet.
func (p *ArtTodRequestPacket) validate() error {
	if err := p.Header.validate(); err != nil {
		return err
	}
	if p.OpCode != code.OpTodRequest {
		return errInvalidOpCode
	}
	return nil
}

// finish is used to finish the Packet for sending.
func (p *ArtTodRequestPacket) finish() {
	p.Header.finish()
}
