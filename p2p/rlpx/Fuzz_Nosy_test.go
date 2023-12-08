package rlpx

import (
	"crypto/ecdsa"
	"io"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto/ecies"
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

func Fuzz_Nosy_Conn_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Close()
	})
}

func Fuzz_Nosy_Conn_Handshake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		if c == nil || prv == nil {
			return
		}

		c.Handshake(prv)
	})
}

func Fuzz_Nosy_Conn_InitWithSecrets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var sec Secrets
		fill_err = tp.Fill(&sec)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.InitWithSecrets(sec)
	})
}

func Fuzz_Nosy_Conn_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Read()
	})
}

func Fuzz_Nosy_Conn_SetDeadline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var time time.Time
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SetDeadline(time)
	})
}

func Fuzz_Nosy_Conn_SetReadDeadline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var time time.Time
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SetReadDeadline(time)
	})
}

func Fuzz_Nosy_Conn_SetSnappy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var snappy bool
		fill_err = tp.Fill(&snappy)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SetSnappy(snappy)
	})
}

func Fuzz_Nosy_Conn_SetWriteDeadline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var time time.Time
		fill_err = tp.Fill(&time)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.SetWriteDeadline(time)
	})
}

func Fuzz_Nosy_Conn_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var code uint64
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Write(code, d3)
	})
}

func Fuzz_Nosy_handshakeState_handleAuthMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handshakeState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var msg *authMsgV4
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		if h == nil || msg == nil || prv == nil {
			return
		}

		h.handleAuthMsg(msg, prv)
	})
}

func Fuzz_Nosy_handshakeState_handleAuthResp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handshakeState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var msg *authRespV4
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if h == nil || msg == nil {
			return
		}

		h.handleAuthResp(msg)
	})
}

func Fuzz_Nosy_handshakeState_makeAuthMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handshakeState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		if h == nil || prv == nil {
			return
		}

		h.makeAuthMsg(prv)
	})
}

func Fuzz_Nosy_handshakeState_makeAuthResp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handshakeState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.makeAuthResp()
	})
}

// skipping Fuzz_Nosy_handshakeState_readMsg__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_handshakeState_runInitiator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handshakeState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var conn io.ReadWriter
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		var remote *ecdsa.PublicKey
		fill_err = tp.Fill(&remote)
		if fill_err != nil {
			return
		}
		if h == nil || prv == nil || remote == nil {
			return
		}

		h.runInitiator(conn, prv, remote)
	})
}

func Fuzz_Nosy_handshakeState_runRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handshakeState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var conn io.ReadWriter
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		if h == nil || prv == nil {
			return
		}

		h.runRecipient(conn, prv)
	})
}

// skipping Fuzz_Nosy_handshakeState_sealEIP8__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_handshakeState_secrets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handshakeState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var auth []byte
		fill_err = tp.Fill(&auth)
		if fill_err != nil {
			return
		}
		var authResp []byte
		fill_err = tp.Fill(&authResp)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.secrets(auth, authResp)
	})
}

func Fuzz_Nosy_handshakeState_staticSharedSecret__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handshakeState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		if h == nil || prv == nil {
			return
		}

		h.staticSharedSecret(prv)
	})
}

func Fuzz_Nosy_hashMAC_compute__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *hashMAC
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var sum1 []byte
		fill_err = tp.Fill(&sum1)
		if fill_err != nil {
			return
		}
		var seed []byte
		fill_err = tp.Fill(&seed)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.compute(sum1, seed)
	})
}

func Fuzz_Nosy_hashMAC_computeFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *hashMAC
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var framedata []byte
		fill_err = tp.Fill(&framedata)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.computeFrame(framedata)
	})
}

func Fuzz_Nosy_hashMAC_computeHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *hashMAC
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var header []byte
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.computeHeader(header)
	})
}

func Fuzz_Nosy_readBuffer_grow__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *readBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.grow(n)
	})
}

func Fuzz_Nosy_readBuffer_read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *readBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.read(r, n)
	})
}

func Fuzz_Nosy_readBuffer_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *readBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.reset()
	})
}

func Fuzz_Nosy_sessionState_readFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *sessionState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var conn io.Reader
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.readFrame(conn)
	})
}

func Fuzz_Nosy_sessionState_writeFrame__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *sessionState
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var conn io.Writer
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var code uint64
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.writeFrame(conn, code, d4)
	})
}

func Fuzz_Nosy_writeBuffer_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *writeBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Write(d2)
	})
}

func Fuzz_Nosy_writeBuffer_appendZero__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *writeBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.appendZero(n)
	})
}

func Fuzz_Nosy_writeBuffer_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *writeBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.reset()
	})
}

func Fuzz_Nosy_exportPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pub *ecies.PublicKey
		fill_err = tp.Fill(&pub)
		if fill_err != nil {
			return
		}
		if pub == nil {
			return
		}

		exportPubkey(pub)
	})
}

func Fuzz_Nosy_growslice__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, wantLength int) {
		growslice(b, wantLength)
	})
}

func Fuzz_Nosy_putUint24__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32, b []byte) {
		putUint24(v, b)
	})
}

func Fuzz_Nosy_readUint24__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		readUint24(b)
	})
}

func Fuzz_Nosy_xor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, one []byte, other []byte) {
		xor(one, other)
	})
}
