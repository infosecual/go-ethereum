package ethtest

import (
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/utesting"
	"github.com/ethereum/go-ethereum/p2p"
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

func Fuzz_Nosy_Chain_ForkID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainfile string, genesis string) {
		c, err := loadChain(chainfile, genesis)
		if err != nil {
			return
		}
		c.ForkID()
	})
}

func Fuzz_Nosy_Chain_GetHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesis string
		fill_err = tp.Fill(&genesis)
		if fill_err != nil {
			return
		}
		var req *GetBlockHeaders
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		c, err := loadChain(chainfile, genesis)
		if err != nil {
			return
		}
		c.GetHeaders(req)
	})
}

func Fuzz_Nosy_Chain_Head__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainfile string, genesis string) {
		c, err := loadChain(chainfile, genesis)
		if err != nil {
			return
		}
		c.Head()
	})
}

func Fuzz_Nosy_Chain_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainfile string, genesis string) {
		c, err := loadChain(chainfile, genesis)
		if err != nil {
			return
		}
		c.Len()
	})
}

func Fuzz_Nosy_Chain_RootAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainfile string, genesis string, height int) {
		c, err := loadChain(chainfile, genesis)
		if err != nil {
			return
		}
		c.RootAt(height)
	})
}

func Fuzz_Nosy_Chain_Shorten__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainfile string, genesis string, height int) {
		c, err := loadChain(chainfile, genesis)
		if err != nil {
			return
		}
		c.Shorten(height)
	})
}

func Fuzz_Nosy_Chain_TD__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainfile string, genesis string) {
		c, err := loadChain(chainfile, genesis)
		if err != nil {
			return
		}
		c.TD()
	})
}

func Fuzz_Nosy_Chain_TotalDifficultyAt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chainfile string, genesis string, height int) {
		c, err := loadChain(chainfile, genesis)
		if err != nil {
			return
		}
		c.TotalDifficultyAt(height)
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

func Fuzz_Nosy_Conn_ReadSnap__(f *testing.F) {
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
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ReadSnap(id)
	})
}

// skipping Fuzz_Nosy_Conn_Write__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/devp2p/internal/ethtest.Message

func Fuzz_Nosy_Conn_handshake__(f *testing.F) {
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

		c.handshake()
	})
}

func Fuzz_Nosy_Conn_headersRequest__(f *testing.F) {
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
		var request *GetBlockHeaders
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		var chain *Chain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var reqID uint64
		fill_err = tp.Fill(&reqID)
		if fill_err != nil {
			return
		}
		if c == nil || request == nil || chain == nil {
			return
		}

		c.headersRequest(request, chain, reqID)
	})
}

func Fuzz_Nosy_Conn_negotiateEthProtocol__(f *testing.F) {
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
		var caps []p2p.Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.negotiateEthProtocol(caps)
	})
}

func Fuzz_Nosy_Conn_peer__(f *testing.F) {
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
		var chain *Chain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var status *Status
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if c == nil || chain == nil || status == nil {
			return
		}

		c.peer(chain, status)
	})
}

func Fuzz_Nosy_Conn_readAndServe__(f *testing.F) {
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
		var chain *Chain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if c == nil || chain == nil {
			return
		}

		c.readAndServe(chain, timeout)
	})
}

// skipping Fuzz_Nosy_Conn_snapRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/devp2p/internal/ethtest.Message

func Fuzz_Nosy_Conn_statusExchange__(f *testing.F) {
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
		var chain *Chain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var status *Status
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if c == nil || chain == nil || status == nil {
			return
		}

		c.statusExchange(chain, status)
	})
}

func Fuzz_Nosy_Conn_waitForResponse__(f *testing.F) {
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
		var chain *Chain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		var requestID uint64
		fill_err = tp.Fill(&requestID)
		if fill_err != nil {
			return
		}
		if c == nil || chain == nil {
			return
		}

		c.waitForResponse(chain, timeout, requestID)
	})
}

func Fuzz_Nosy_Error_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Error
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Code()
	})
}

func Fuzz_Nosy_Error_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Error
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_Error_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Error
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ReqID()
	})
}

func Fuzz_Nosy_Error_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Error
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.String()
	})
}

func Fuzz_Nosy_Error_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Error
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Unwrap()
	})
}

func Fuzz_Nosy_Suite_EthTests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.EthTests()
	})
}

func Fuzz_Nosy_Suite_SnapTests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.SnapTests()
	})
}

func Fuzz_Nosy_Suite_TestBlockHashAnnounce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestBlockHashAnnounce(t4)
	})
}

func Fuzz_Nosy_Suite_TestBroadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestBroadcast(t4)
	})
}

func Fuzz_Nosy_Suite_TestGetBlockBodies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestGetBlockBodies(t4)
	})
}

func Fuzz_Nosy_Suite_TestGetBlockHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestGetBlockHeaders(t4)
	})
}

func Fuzz_Nosy_Suite_TestLargeAnnounce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestLargeAnnounce(t4)
	})
}

func Fuzz_Nosy_Suite_TestLargeTxRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestLargeTxRequest(t4)
	})
}

func Fuzz_Nosy_Suite_TestMaliciousHandshake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestMaliciousHandshake(t4)
	})
}

func Fuzz_Nosy_Suite_TestMaliciousStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestMaliciousStatus(t4)
	})
}

func Fuzz_Nosy_Suite_TestMaliciousTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestMaliciousTx(t4)
	})
}

func Fuzz_Nosy_Suite_TestNewPooledTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestNewPooledTxs(t4)
	})
}

func Fuzz_Nosy_Suite_TestOldAnnounce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestOldAnnounce(t4)
	})
}

func Fuzz_Nosy_Suite_TestSameRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestSameRequestID(t4)
	})
}

func Fuzz_Nosy_Suite_TestSimultaneousRequests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestSimultaneousRequests(t4)
	})
}

func Fuzz_Nosy_Suite_TestSnapGetAccountRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestSnapGetAccountRange(t4)
	})
}

func Fuzz_Nosy_Suite_TestSnapGetByteCodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestSnapGetByteCodes(t4)
	})
}

func Fuzz_Nosy_Suite_TestSnapGetStorageRanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestSnapGetStorageRanges(t4)
	})
}

func Fuzz_Nosy_Suite_TestSnapStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestSnapStatus(t4)
	})
}

func Fuzz_Nosy_Suite_TestSnapTrieNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestSnapTrieNodes(t4)
	})
}

func Fuzz_Nosy_Suite_TestStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestStatus(t4)
	})
}

func Fuzz_Nosy_Suite_TestTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestTransaction(t4)
	})
}

func Fuzz_Nosy_Suite_TestZeroRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.TestZeroRequestID(t4)
	})
}

func Fuzz_Nosy_Suite_createSendAndRecvConns__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.createSendAndRecvConns()
	})
}

func Fuzz_Nosy_Suite_dial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.dial()
	})
}

func Fuzz_Nosy_Suite_dialSnap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.dialSnap()
	})
}

func Fuzz_Nosy_Suite_hashAnnounce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.hashAnnounce()
	})
}

func Fuzz_Nosy_Suite_maliciousHandshakes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.maliciousHandshakes(t4)
	})
}

func Fuzz_Nosy_Suite_maliciousStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var conn *Conn
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if dest == nil || conn == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.maliciousStatus(conn)
	})
}

func Fuzz_Nosy_Suite_oldAnnounce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.oldAnnounce()
	})
}

func Fuzz_Nosy_Suite_sendMaliciousTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.sendMaliciousTxs(t4)
	})
}

func Fuzz_Nosy_Suite_sendNextBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.sendNextBlock()
	})
}

func Fuzz_Nosy_Suite_sendSuccessfulTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.sendSuccessfulTxs(t4)
	})
}

func Fuzz_Nosy_Suite_snapGetAccountRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var tc *accRangeTest
		fill_err = tp.Fill(&tc)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil || tc == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.snapGetAccountRange(t4, tc)
	})
}

func Fuzz_Nosy_Suite_snapGetByteCodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var tc *byteCodesTest
		fill_err = tp.Fill(&tc)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil || tc == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.snapGetByteCodes(t4, tc)
	})
}

func Fuzz_Nosy_Suite_snapGetStorageRanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var tc *stRangesTest
		fill_err = tp.Fill(&tc)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil || tc == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.snapGetStorageRanges(t4, tc)
	})
}

func Fuzz_Nosy_Suite_snapGetTrieNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var t4 *utesting.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var tc *trieNodesTest
		fill_err = tp.Fill(&tc)
		if fill_err != nil {
			return
		}
		if dest == nil || t4 == nil || tc == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.snapGetTrieNodes(t4, tc)
	})
}

func Fuzz_Nosy_Suite_testAnnounce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var sendConn *Conn
		fill_err = tp.Fill(&sendConn)
		if fill_err != nil {
			return
		}
		var receiveConn *Conn
		fill_err = tp.Fill(&receiveConn)
		if fill_err != nil {
			return
		}
		var blockAnnouncement *NewBlock
		fill_err = tp.Fill(&blockAnnouncement)
		if fill_err != nil {
			return
		}
		if dest == nil || sendConn == nil || receiveConn == nil || blockAnnouncement == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.testAnnounce(sendConn, receiveConn, blockAnnouncement)
	})
}

func Fuzz_Nosy_Suite_waitAnnounce__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var conn *Conn
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var blockAnnouncement *NewBlock
		fill_err = tp.Fill(&blockAnnouncement)
		if fill_err != nil {
			return
		}
		if dest == nil || conn == nil || blockAnnouncement == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.waitAnnounce(conn, blockAnnouncement)
	})
}

func Fuzz_Nosy_Suite_waitForBlockImport__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var genesisfile string
		fill_err = tp.Fill(&genesisfile)
		if fill_err != nil {
			return
		}
		var conn *Conn
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if dest == nil || conn == nil || block == nil {
			return
		}

		s, err := NewSuite(dest, chainfile, genesisfile)
		if err != nil {
			return
		}
		s.waitForBlockImport(conn, block)
	})
}

func Fuzz_Nosy_AccountRange_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg AccountRange
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_AccountRange_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg AccountRange
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_BlockBodies_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg BlockBodies
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_BlockBodies_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg BlockBodies
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_BlockHeaders_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg BlockHeaders
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_BlockHeaders_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg BlockHeaders
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_ByteCodes_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg ByteCodes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_ByteCodes_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg ByteCodes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_Disconnect_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Disconnect
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_Disconnect_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Disconnect
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_GetAccountRange_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetAccountRange
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_GetAccountRange_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetAccountRange
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_GetBlockBodies_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetBlockBodies
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_GetBlockBodies_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetBlockBodies
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_GetBlockHeaders_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetBlockHeaders
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_GetBlockHeaders_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetBlockHeaders
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_GetByteCodes_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetByteCodes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_GetByteCodes_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetByteCodes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_GetPooledTransactions_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetPooledTransactions
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_GetPooledTransactions_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetPooledTransactions
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_GetStorageRanges_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetStorageRanges
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_GetStorageRanges_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetStorageRanges
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_GetTrieNodes_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetTrieNodes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_GetTrieNodes_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg GetTrieNodes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_Hello_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Hello
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_Hello_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Hello
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

// skipping Fuzz_Nosy_Message_Code__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/devp2p/internal/ethtest.Message

// skipping Fuzz_Nosy_Message_ReqID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/devp2p/internal/ethtest.Message

func Fuzz_Nosy_NewBlock_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg NewBlock
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_NewBlock_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg NewBlock
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_NewBlockHashes_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg NewBlockHashes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_NewBlockHashes_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg NewBlockHashes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_NewPooledTransactionHashes_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg NewPooledTransactionHashes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_NewPooledTransactionHashes_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg NewPooledTransactionHashes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_NewPooledTransactionHashes66_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg NewPooledTransactionHashes66
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_NewPooledTransactionHashes66_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg NewPooledTransactionHashes66
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_Ping_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Ping
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_Ping_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Ping
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_Pong_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Pong
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_Pong_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Pong
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_PooledTransactions_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg PooledTransactions
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_PooledTransactions_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg PooledTransactions
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_Status_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Status
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_Status_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Status
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_StorageRanges_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg StorageRanges
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_StorageRanges_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg StorageRanges
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_Transactions_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Transactions
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_Transactions_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Transactions
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_TrieNodes_Code__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg TrieNodes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Code()
	})
}

func Fuzz_Nosy_TrieNodes_ReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg TrieNodes
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.ReqID()
	})
}

func Fuzz_Nosy_blocksFromFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainfile string
		fill_err = tp.Fill(&chainfile)
		if fill_err != nil {
			return
		}
		var gblock *types.Block
		fill_err = tp.Fill(&gblock)
		if fill_err != nil {
			return
		}
		if gblock == nil {
			return
		}

		blocksFromFile(chainfile, gblock)
	})
}

func Fuzz_Nosy_compareReceivedTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var recvTxs []common.Hash
		fill_err = tp.Fill(&recvTxs)
		if fill_err != nil {
			return
		}
		var txs []*types.Transaction
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}

		compareReceivedTxs(recvTxs, txs)
	})
}

func Fuzz_Nosy_decodeNibbles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, nibbles []byte, bytes []byte) {
		decodeNibbles(nibbles, bytes)
	})
}

func Fuzz_Nosy_generateTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Suite
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var numTxs int
		fill_err = tp.Fill(&numTxs)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		generateTxs(s, numTxs)
	})
}

func Fuzz_Nosy_hasTerm__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s []byte) {
		hasTerm(s)
	})
}

func Fuzz_Nosy_headersMatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var expected []*types.Header
		fill_err = tp.Fill(&expected)
		if fill_err != nil {
			return
		}
		var headers []*types.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		headersMatch(expected, headers)
	})
}

func Fuzz_Nosy_hexToCompact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hex []byte) {
		hexToCompact(hex)
	})
}

func Fuzz_Nosy_keybytesToHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, str []byte) {
		keybytesToHex(str)
	})
}

func Fuzz_Nosy_largeBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, megabytes int) {
		largeBuffer(megabytes)
	})
}

func Fuzz_Nosy_largeString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, megabytes int) {
		largeString(megabytes)
	})
}
