package gocached

// General format of a packet
// Byte/     0       |       1       |       2       |       3       |
//    /              |               |               |               |
//   |0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|0 1 2 3 4 5 6 7|
//   +---------------+---------------+---------------+---------------+
//  0/ HEADER                                                        /
//   /                                                               /
//   /                                                               /
//   /                                                               /
//   +---------------+---------------+---------------+---------------+
// 24/ COMMAND-SPECIFIC EXTRAS (as needed)                           /
//  +/  (note length in the extras length header field)              /
//   +---------------+---------------+---------------+---------------+
//  m/ Key (as needed)                                               /
//  +/  (note length in key length header field)                     /
//   +---------------+---------------+---------------+---------------+
//  n/ Value (as needed)                                             /
//  +/  (note length is total body length header field, minus        /
//  +/   sum of the extras and key length body fields)               /
//   +---------------+---------------+---------------+---------------+
//   Total 24 bytes

// Header length is always 24 bytes
const HEADER_LENGTH = 24
const MAX_BODY_LENGTH = uint32(1 * 1e6)

// Magic Byte: for request and response packet
// http://code.google.com/p/memcached/wiki/BinaryProtocolRevamped#Magic_Byte
const (
	REQUEST  = 0x80
	RESPONSE = 0x81
)

// Response Status: a two byte field
// http://code.google.com/p/memcached/wiki/BinaryProtocolRevamped#Response_Status
const (
	SUCCESS                 = uint16(0x00)
	KEY_NOT_FOUND           = uint16(0x01)
	KEY_NOT_EXIST           = uint16(0x02)
	VALUE_TOO_LARGE         = uint16(0x03)
	INVALID_ARGUMENTS       = uint16(0x04)
	ITEM_NOT_STORED         = uint16(0x05)
	BAD_NUMERIC_VALUE       = uint16(0x06)
	INVALID_VBUCKET         = uint16(0x07)
	AUTHENTICATION_ERROR    = uint16(0x08)
	AUTHENTICATION_CONTINUE = uint16(0x09)
	UNKNOWN_COMMAND         = uint16(0x81)
	OUT_OF_MEMORY           = uint16(0x82)
	NOT_SUPPORTED           = uint16(0x83)
	INTERNAL_ERROR          = uint16(0x84)
	BUSY                    = uint16(0x85)
	TEMP_FAILURE            = uint16(0x86)
)

// Command Opcodes: a one-byte field
// http://code.google.com/p/memcached/wiki/BinaryProtocolRevamped#Command_Opcodes
const (
	GET                  = uint8(0x00)
	SET                  = uint8(0x01)
	ADD                  = uint8(0x02)
	REPLACE              = uint8(0x03)
	DELETE               = uint8(0x04)
	INCREMENT            = uint8(0x05)
	DECREMENT            = uint8(0x06)
	QUIT                 = uint8(0x07)
	FLUSH                = uint8(0x08)
	GETQ                 = uint8(0x09)
	NOOP                 = uint8(0x0a)
	VERSION              = uint8(0x0b)
	GETK                 = uint8(0x0c)
	GETKQ                = uint8(0x0d)
	APPEND               = uint8(0x0e)
	PREPEND              = uint8(0x0f)
	STAT                 = uint8(0x10)
	SETQ                 = uint8(0x11)
	ADDQ                 = uint8(0x12)
	REPLACEQ             = uint8(0x13)
	DELETEQ              = uint8(0x14)
	INCREMENTQ           = uint8(0x15)
	DECREMENTQ           = uint8(0x16)
	QUITQ                = uint8(0x17)
	FLUSHQ               = uint8(0x18)
	APPENDQ              = uint8(0x19)
	PREPENDQ             = uint8(0x1a)
	VERBOSITY            = uint8(0x1b)
	TOUCH                = uint8(0x1c)
	GAT                  = uint8(0x1d)
	GATQ                 = uint8(0x1d)
	SASL_LIST_MECHS      = uint8(0x20)
	SASL_AUTH            = uint8(0x21)
	SASL_STEP            = uint8(0x22)
	RGET                 = uint8(0x30)
	RSET                 = uint8(0x31)
	RSETQ                = uint8(0x32)
	RAPPEND              = uint8(0x33)
	RAPPENDQ             = uint8(0x34)
	RPREPEND             = uint8(0x35)
	RPREPENDQ            = uint8(0x36)
	RDELETE              = uint8(0x37)
	RDELETEQ             = uint8(0x38)
	RINCR                = uint8(0x39)
	RINCRQ               = uint8(0x3a)
	RDECR                = uint8(0x3b)
	RDECRQ               = uint8(0x3c)
	SET_VBUCKKET         = uint8(0x3d)
	GET_VBUCKKET         = uint8(0x3e)
	DEL_VBUCKKET         = uint8(0x3f)
	TAP_CONNECT          = uint8(0x40)
	TAP_MUTATION         = uint8(0x41)
	TAP_DELETE           = uint8(0x42)
	TAP_FLUSH            = uint8(0x43)
	TAP_OPAQUE           = uint8(0x44)
	TAP_VBUCKET_SET      = uint8(0x45)
	TAP_CHECKPOINT_START = uint8(0x46)
	TAP_CHECKPOINT_END   = uint8(0x47)
)

const (
	BACKFILL          = 0x01
	DUMP              = 0x02
	LIST_VBUCKETS     = 0x04
	TAKEOVER_VBUCKETS = 0x08
	SUPPORT_ACK       = 0x10
	REQUEST_KEYS_ONLY = 0x20
	CHECKPOINT        = 0x40
	REGISTERED_CLIENT = 0x80
)
