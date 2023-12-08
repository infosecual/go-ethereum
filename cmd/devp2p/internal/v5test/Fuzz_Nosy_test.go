package v5test

import (
	"testing"

	"github.com/ethereum/go-ethereum/internal/utesting"
	"github.com/ethereum/go-ethereum/p2p/enr"
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

func Fuzz_Nosy_Suite_AllTests__(f *testing.F) {
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
		if s == nil {
			return
		}

		s.AllTests()
	})
}

func Fuzz_Nosy_Suite_TestFindnodeResults__(f *testing.F) {
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
		var t2 *utesting.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if s == nil || t2 == nil {
			return
		}

		s.TestFindnodeResults(t2)
	})
}

func Fuzz_Nosy_Suite_TestFindnodeZeroDistance__(f *testing.F) {
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
		var t2 *utesting.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if s == nil || t2 == nil {
			return
		}

		s.TestFindnodeZeroDistance(t2)
	})
}

func Fuzz_Nosy_Suite_TestPing__(f *testing.F) {
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
		var t2 *utesting.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if s == nil || t2 == nil {
			return
		}

		s.TestPing(t2)
	})
}

func Fuzz_Nosy_Suite_TestPingHandshakeInterrupted__(f *testing.F) {
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
		var t2 *utesting.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if s == nil || t2 == nil {
			return
		}

		s.TestPingHandshakeInterrupted(t2)
	})
}

func Fuzz_Nosy_Suite_TestPingLargeRequestID__(f *testing.F) {
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
		var t2 *utesting.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if s == nil || t2 == nil {
			return
		}

		s.TestPingLargeRequestID(t2)
	})
}

func Fuzz_Nosy_Suite_TestPingMultiIP__(f *testing.F) {
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
		var t2 *utesting.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if s == nil || t2 == nil {
			return
		}

		s.TestPingMultiIP(t2)
	})
}

func Fuzz_Nosy_Suite_TestTalkRequest__(f *testing.F) {
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
		var t2 *utesting.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if s == nil || t2 == nil {
			return
		}

		s.TestTalkRequest(t2)
	})
}

// skipping Fuzz_Nosy_Suite_listen1__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/devp2p/internal/v5test.logger

// skipping Fuzz_Nosy_Suite_listen2__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/devp2p/internal/v5test.logger

func Fuzz_Nosy_bystander_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bn *bystander
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}
		if bn == nil {
			return
		}

		bn.close()
	})
}

func Fuzz_Nosy_bystander_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bn *bystander
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}
		if bn == nil {
			return
		}

		bn.id()
	})
}

func Fuzz_Nosy_bystander_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bn *bystander
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}
		if bn == nil {
			return
		}

		bn.loop()
	})
}

func Fuzz_Nosy_bystander_notifyAdded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bn *bystander
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}
		if bn == nil {
			return
		}

		bn.notifyAdded()
	})
}

func Fuzz_Nosy_conn_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tc *conn
		fill_err = tp.Fill(&tc)
		if fill_err != nil {
			return
		}
		if tc == nil {
			return
		}

		tc.close()
	})
}

// skipping Fuzz_Nosy_conn_findnode__ because parameters include func, chan, or unsupported interface: net.PacketConn

func Fuzz_Nosy_conn_listen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tc *conn
		fill_err = tp.Fill(&tc)
		if fill_err != nil {
			return
		}
		var ip string
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if tc == nil {
			return
		}

		tc.listen(ip)
	})
}

// skipping Fuzz_Nosy_conn_logf__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_conn_nextReqID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tc *conn
		fill_err = tp.Fill(&tc)
		if fill_err != nil {
			return
		}
		if tc == nil {
			return
		}

		tc.nextReqID()
	})
}

// skipping Fuzz_Nosy_conn_read__ because parameters include func, chan, or unsupported interface: net.PacketConn

// skipping Fuzz_Nosy_conn_reqresp__ because parameters include func, chan, or unsupported interface: net.PacketConn

// skipping Fuzz_Nosy_conn_setEndpoint__ because parameters include func, chan, or unsupported interface: net.PacketConn

// skipping Fuzz_Nosy_conn_write__ because parameters include func, chan, or unsupported interface: net.PacketConn

// skipping Fuzz_Nosy_readError_AppendLogInfo__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_readError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *readError
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Error()
	})
}

func Fuzz_Nosy_readError_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *readError
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

func Fuzz_Nosy_readError_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *readError
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

func Fuzz_Nosy_readError_RequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *readError
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

func Fuzz_Nosy_readError_SetRequestID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *readError
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var _x2 []byte
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SetRequestID(_x2)
	})
}

func Fuzz_Nosy_readError_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *readError
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Unwrap()
	})
}

// skipping Fuzz_Nosy_logger_Logf__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/devp2p/internal/v5test.logger

// skipping Fuzz_Nosy_checkPong__ because parameters include func, chan, or unsupported interface: net.PacketConn

func Fuzz_Nosy_checkRecords__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var records []*enr.Record
		fill_err = tp.Fill(&records)
		if fill_err != nil {
			return
		}

		checkRecords(records)
	})
}

func Fuzz_Nosy_containsUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ints []uint
		fill_err = tp.Fill(&ints)
		if fill_err != nil {
			return
		}
		var x uint
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		containsUint(ints, x)
	})
}
