package bigreq

/*
	This file was generated by bigreq.xml on May 10 2012 4:20:27pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the BIG-REQUESTS extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 12, "BIG-REQUESTS").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named BIG-REQUESTS could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["BIG-REQUESTS"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["BIG-REQUESTS"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["BIG-REQUESTS"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["BIG-REQUESTS"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["BIG-REQUESTS"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Request Enable
// size: 4
type EnableCookie struct {
	*xgb.Cookie
}

func Enable(c *xgb.Conn) EnableCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(enableRequest(c), cookie)
	return EnableCookie{cookie}
}

func EnableUnchecked(c *xgb.Conn) EnableCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(enableRequest(c), cookie)
	return EnableCookie{cookie}
}

// Request reply for Enable
// size: 12
type EnableReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	MaximumRequestLength uint32
}

// Waits and reads reply data from request Enable
func (cook EnableCookie) Reply() (*EnableReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return enableReply(buf), nil
}

// Read reply into structure from buffer for Enable
func enableReply(buf []byte) *EnableReply {
	v := new(EnableReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.MaximumRequestLength = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for Enable
func enableRequest(c *xgb.Conn) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["BIG-REQUESTS"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}
