package v5wire

import (
	"bytes"
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/p2p/enode"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_Codec_Decode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var inputData []byte
		fill_err = tp.Fill(&inputData)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Decode(inputData, addr)
	})
}

// skipping Fuzz_Nosy_Codec_Encode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

func Fuzz_Nosy_Codec_EncodeRaw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var head Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var msgdata []byte
		fill_err = tp.Fill(&msgdata)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.EncodeRaw(id, head, msgdata)
	})
}

func Fuzz_Nosy_Codec_decodeHandshake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var fromAddr string
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		var head *Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if c == nil || head == nil {
			return
		}

		c.decodeHandshake(fromAddr, head)
	})
}

func Fuzz_Nosy_Codec_decodeHandshakeAuthData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var head *Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if c == nil || head == nil {
			return
		}

		c.decodeHandshakeAuthData(head)
	})
}

func Fuzz_Nosy_Codec_decodeHandshakeMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var fromAddr string
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		var head *Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var headerData []byte
		fill_err = tp.Fill(&headerData)
		if fill_err != nil {
			return
		}
		var msgData []byte
		fill_err = tp.Fill(&msgData)
		if fill_err != nil {
			return
		}
		if c == nil || head == nil {
			return
		}

		c.decodeHandshakeMessage(fromAddr, head, headerData, msgData)
	})
}

func Fuzz_Nosy_Codec_decodeHandshakeRecord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var local *enode.Node
		fill_err = tp.Fill(&local)
		if fill_err != nil {
			return
		}
		var wantID enode.ID
		fill_err = tp.Fill(&wantID)
		if fill_err != nil {
			return
		}
		var remote []byte
		fill_err = tp.Fill(&remote)
		if fill_err != nil {
			return
		}
		if c == nil || local == nil {
			return
		}

		c.decodeHandshakeRecord(local, wantID, remote)
	})
}

func Fuzz_Nosy_Codec_decodeMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var fromAddr string
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		var head *Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var headerData []byte
		fill_err = tp.Fill(&headerData)
		if fill_err != nil {
			return
		}
		var msgData []byte
		fill_err = tp.Fill(&msgData)
		if fill_err != nil {
			return
		}
		if c == nil || head == nil {
			return
		}

		c.decodeMessage(fromAddr, head, headerData, msgData)
	})
}

func Fuzz_Nosy_Codec_decodeWhoareyou__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var head *Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var headerData []byte
		fill_err = tp.Fill(&headerData)
		if fill_err != nil {
			return
		}
		if c == nil || head == nil {
			return
		}

		c.decodeWhoareyou(head, headerData)
	})
}

func Fuzz_Nosy_Codec_decryptMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		var nonce []byte
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		var headerData []byte
		fill_err = tp.Fill(&headerData)
		if fill_err != nil {
			return
		}
		var readKey []byte
		fill_err = tp.Fill(&readKey)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.decryptMessage(input, nonce, headerData, readKey)
	})
}

func Fuzz_Nosy_Codec_encodeHandshakeHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var toID enode.ID
		fill_err = tp.Fill(&toID)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var challenge *Whoareyou
		fill_err = tp.Fill(&challenge)
		if fill_err != nil {
			return
		}
		if c == nil || challenge == nil {
			return
		}

		c.encodeHandshakeHeader(toID, addr, challenge)
	})
}

func Fuzz_Nosy_Codec_encodeMessageHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var toID enode.ID
		fill_err = tp.Fill(&toID)
		if fill_err != nil {
			return
		}
		var s *session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if c == nil || s == nil {
			return
		}

		c.encodeMessageHeader(toID, s)
	})
}

func Fuzz_Nosy_Codec_encodeRandom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var toID enode.ID
		fill_err = tp.Fill(&toID)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.encodeRandom(toID)
	})
}

func Fuzz_Nosy_Codec_encodeWhoareyou__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var toID enode.ID
		fill_err = tp.Fill(&toID)
		if fill_err != nil {
			return
		}
		var packet *Whoareyou
		fill_err = tp.Fill(&packet)
		if fill_err != nil {
			return
		}
		if c == nil || packet == nil {
			return
		}

		c.encodeWhoareyou(toID, packet)
	})
}

// skipping Fuzz_Nosy_Codec_encryptMessage__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

func Fuzz_Nosy_Codec_makeHandshakeAuth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var toID enode.ID
		fill_err = tp.Fill(&toID)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var challenge *Whoareyou
		fill_err = tp.Fill(&challenge)
		if fill_err != nil {
			return
		}
		if c == nil || challenge == nil {
			return
		}

		c.makeHandshakeAuth(toID, addr, challenge)
	})
}

func Fuzz_Nosy_Codec_makeHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var toID enode.ID
		fill_err = tp.Fill(&toID)
		if fill_err != nil {
			return
		}
		var flag byte
		fill_err = tp.Fill(&flag)
		if fill_err != nil {
			return
		}
		var authsizeExtra int
		fill_err = tp.Fill(&authsizeExtra)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.makeHeader(toID, flag, authsizeExtra)
	})
}

func Fuzz_Nosy_Codec_writeHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var head *Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if c == nil || head == nil {
			return
		}

		c.writeHeaders(head)
	})
}

// skipping Fuzz_Nosy_Findnode_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Findnode_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Findnode
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Kind()
	})
}

func Fuzz_Nosy_Findnode_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Findnode
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Name()
	})
}

func Fuzz_Nosy_Findnode_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Findnode
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestID()
	})
}

func Fuzz_Nosy_Findnode_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Findnode
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id []byte
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SetRequestID(id)
	})
}

func Fuzz_Nosy_Header_mask__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Header
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var destID enode.ID
		fill_err = tp.Fill(&destID)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.mask(destID)
	})
}

// skipping Fuzz_Nosy_Nodes_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Nodes_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Nodes
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_Nodes_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Nodes
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Nodes_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Nodes
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestID()
	})
}

func Fuzz_Nosy_Nodes_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Nodes
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id []byte
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SetRequestID(id)
	})
}

// skipping Fuzz_Nosy_Ping_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Ping_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Ping
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_Ping_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Ping
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Ping_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Ping
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestID()
	})
}

func Fuzz_Nosy_Ping_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Ping
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id []byte
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SetRequestID(id)
	})
}

// skipping Fuzz_Nosy_Pong_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Pong_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Pong
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_Pong_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Pong
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Pong_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Pong
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestID()
	})
}

func Fuzz_Nosy_Pong_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *Pong
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id []byte
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SetRequestID(id)
	})
}

func Fuzz_Nosy_SessionCache_deleteHandshake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SessionCache
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.deleteHandshake(id, addr)
	})
}

func Fuzz_Nosy_SessionCache_getHandshake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SessionCache
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.getHandshake(id, addr)
	})
}

func Fuzz_Nosy_SessionCache_handshakeGC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SessionCache
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.handshakeGC()
	})
}

func Fuzz_Nosy_SessionCache_nextNonce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SessionCache
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var s *session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if sc == nil || s == nil {
			return
		}

		sc.nextNonce(s)
	})
}

func Fuzz_Nosy_SessionCache_readKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SessionCache
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.readKey(id, addr)
	})
}

func Fuzz_Nosy_SessionCache_session__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SessionCache
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if sc == nil {
			return
		}

		sc.session(id, addr)
	})
}

func Fuzz_Nosy_SessionCache_storeNewSession__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SessionCache
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var s *session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if sc == nil || s == nil {
			return
		}

		sc.storeNewSession(id, addr, s)
	})
}

func Fuzz_Nosy_SessionCache_storeSentHandshake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sc *SessionCache
		fill_err = tp.Fill(&sc)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var challenge *Whoareyou
		fill_err = tp.Fill(&challenge)
		if fill_err != nil {
			return
		}
		if sc == nil || challenge == nil {
			return
		}

		sc.storeSentHandshake(id, addr, challenge)
	})
}

func Fuzz_Nosy_StaticHeader_checkValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *StaticHeader
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var packetLen int
		fill_err = tp.Fill(&packetLen)
		if fill_err != nil {
			return
		}
		var protocolID [6]byte
		fill_err = tp.Fill(&protocolID)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.checkValid(packetLen, protocolID)
	})
}

// skipping Fuzz_Nosy_TalkRequest_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_TalkRequest_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *TalkRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_TalkRequest_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *TalkRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_TalkRequest_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TalkRequest
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestID()
	})
}

func Fuzz_Nosy_TalkRequest_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TalkRequest
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id []byte
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SetRequestID(id)
	})
}

// skipping Fuzz_Nosy_TalkResponse_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_TalkResponse_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *TalkResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_TalkResponse_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *TalkResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_TalkResponse_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TalkResponse
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.RequestID()
	})
}

func Fuzz_Nosy_TalkResponse_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *TalkResponse
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var id []byte
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SetRequestID(id)
	})
}

// skipping Fuzz_Nosy_Unknown_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Unknown_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Unknown
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_Unknown_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Unknown
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Unknown_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Unknown
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RequestID()
	})
}

func Fuzz_Nosy_Unknown_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Unknown
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 []byte
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SetRequestID(_x2)
	})
}

// skipping Fuzz_Nosy_Whoareyou_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Whoareyou_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Whoareyou
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Kind()
	})
}

func Fuzz_Nosy_Whoareyou_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Whoareyou
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_Whoareyou_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Whoareyou
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.RequestID()
	})
}

func Fuzz_Nosy_Whoareyou_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Whoareyou
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 []byte
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SetRequestID(_x2)
	})
}

func Fuzz_Nosy_session_keysFlipped__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *session
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.keysFlipped()
	})
}

// skipping Fuzz_Nosy_Packet_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Packet_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ptype byte, body []byte) {
		_x1, err := DecodeMessage(ptype, body)
		if err != nil {
			return
		}
		_x1.Kind()
	})
}

func Fuzz_Nosy_Packet_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ptype byte, body []byte) {
		_x1, err := DecodeMessage(ptype, body)
		if err != nil {
			return
		}
		_x1.Name()
	})
}

func Fuzz_Nosy_Packet_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ptype byte, body []byte) {
		_x1, err := DecodeMessage(ptype, body)
		if err != nil {
			return
		}
		_x1.RequestID()
	})
}

func Fuzz_Nosy_Packet_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ptype byte, body []byte, _x3 []byte) {
		_x1, err := DecodeMessage(ptype, body)
		if err != nil {
			return
		}
		_x1.SetRequestID(_x3)
	})
}

func Fuzz_Nosy_s256raw_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 s256raw
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.ENRKey()
	})
}

func Fuzz_Nosy_EncodePubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PublicKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		EncodePubkey(key)
	})
}

// skipping Fuzz_Nosy_IsInvalidHeader__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_bytesCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *bytes.Buffer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		bytesCopy(r)
	})
}

func Fuzz_Nosy_decryptGCM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte, nonce []byte, ct []byte, authData []byte) {
		decryptGCM(key, nonce, ct, authData)
	})
}

func Fuzz_Nosy_ecdh__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var privkey *ecdsa.PrivateKey
		fill_err = tp.Fill(&privkey)
		if fill_err != nil {
			return
		}
		var pubkey *ecdsa.PublicKey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if privkey == nil || pubkey == nil {
			return
		}

		ecdh(privkey, pubkey)
	})
}

func Fuzz_Nosy_encryptGCM__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dest []byte, key []byte, nonce []byte, plaintext []byte, authData []byte) {
		encryptGCM(dest, key, nonce, plaintext, authData)
	})
}

// skipping Fuzz_Nosy_idNonceHash__ because parameters include func, chan, or unsupported interface: hash.Hash

// skipping Fuzz_Nosy_makeIDSignature__ because parameters include func, chan, or unsupported interface: hash.Hash
