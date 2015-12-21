package mysqlproto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandshakeV10FullPacket(t *testing.T) {
	data := []byte{
		0x4a, 0x00, 0x00, 0x00, 0x0a, 0x35, 0x2e, 0x36, 0x2e, 0x32, 0x35,
		0x00, 0x9e, 0x2e, 0x00, 0x00, 0x4f, 0x61, 0x7b, 0x65, 0x68, 0x5c,
		0x73, 0x4e, 0x00, 0xff, 0xf7, 0x21, 0x02, 0x00, 0x7f, 0x80, 0x15,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x69,
		0x48, 0x6a, 0x5d, 0x73, 0x4a, 0x55, 0x50, 0x70, 0x64, 0x24, 0x25,
		0x00, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69,
		0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x00,
	}
	stream := newBuffer(data)
	packet, err := ReadHandshakeV10(NewStream(stream))
	assert.Nil(t, err)
	assert.Equal(t, packet.ProtocolVersion, byte(0x0a))
	assert.Equal(t, packet.ServerVersion, "5.6.25")
	assert.Equal(t, packet.ConnectionID, [4]byte{0x9e, 0x2e, 0x00, 0x00})
	assert.Equal(t, packet.AuthPluginData, []byte(`Oa{eh\sNiHj]sJUPpd$%`))
	assert.Equal(t, packet.CapabilityFlags, uint32(0x807ff7ff))
	assert.Equal(t, packet.CharacterSet, byte(0x21))
	assert.Equal(t, packet.StatusFlags, [2]byte{0x02, 0x00})
	assert.Equal(t, packet.AuthPluginName, "mysql_native_password")
}

func TestNewHandshakeV10ShortPacket(t *testing.T) {
	buf := newBuffer([]byte{
		0x17, 0x00, 0x00, 0x00, 0x0a, 0x35, 0x2e, 0x36, 0x2e, 0x32, 0x35,
		0x00, 0x9e, 0x2e, 0x00, 0x00, 0x4f, 0x61, 0x7b, 0x65, 0x68, 0x5c,
		0x73, 0x4e, 0x00, 0xff, 0xf7,
	})
	packet, err := ReadHandshakeV10(NewStream(buf))
	assert.Nil(t, err)
	assert.Equal(t, packet.ProtocolVersion, byte(0x0a))
	assert.Equal(t, packet.ServerVersion, "5.6.25")
	assert.Equal(t, packet.ConnectionID, [4]byte{0x9e, 0x2e, 0x00, 0x00})
	assert.Equal(t, packet.AuthPluginData, []byte(`Oa{eh\sN`))
	assert.Equal(t, packet.CapabilityFlags, uint32(0xf7ff))
	assert.Equal(t, packet.CharacterSet, byte(0x00))
	assert.Equal(t, packet.StatusFlags, [2]byte{0x00, 0x00})
	assert.Equal(t, packet.AuthPluginName, "")
}
