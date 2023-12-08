package ethstats

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
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

func Fuzz_Nosy_Service_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Start()
	})
}

func Fuzz_Nosy_Service_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Stop()
	})
}

func Fuzz_Nosy_Service_assembleBlockStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if s == nil || block == nil {
			return
		}

		s.assembleBlockStats(block)
	})
}

func Fuzz_Nosy_Service_login__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var conn *connWrapper
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if s == nil || conn == nil {
			return
		}

		s.login(conn)
	})
}

// skipping Fuzz_Nosy_Service_loop__ because parameters include func, chan, or unsupported interface: chan github.com/ethereum/go-ethereum/core.ChainHeadEvent

func Fuzz_Nosy_Service_readLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var conn *connWrapper
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if s == nil || conn == nil {
			return
		}

		s.readLoop(conn)
	})
}

func Fuzz_Nosy_Service_report__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var conn *connWrapper
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if s == nil || conn == nil {
			return
		}

		s.report(conn)
	})
}

func Fuzz_Nosy_Service_reportBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var conn *connWrapper
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var block *types.Block
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if s == nil || conn == nil || block == nil {
			return
		}

		s.reportBlock(conn, block)
	})
}

func Fuzz_Nosy_Service_reportHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var conn *connWrapper
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var list []uint64
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		if s == nil || conn == nil {
			return
		}

		s.reportHistory(conn, list)
	})
}

func Fuzz_Nosy_Service_reportLatency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var conn *connWrapper
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if s == nil || conn == nil {
			return
		}

		s.reportLatency(conn)
	})
}

func Fuzz_Nosy_Service_reportPending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var conn *connWrapper
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if s == nil || conn == nil {
			return
		}

		s.reportPending(conn)
	})
}

func Fuzz_Nosy_Service_reportStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var conn *connWrapper
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if s == nil || conn == nil {
			return
		}

		s.reportStats(conn)
	})
}

func Fuzz_Nosy_connWrapper_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conn *websocket.Conn
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if conn == nil {
			return
		}

		w := newConnectionWrapper(conn)
		w.Close()
	})
}

// skipping Fuzz_Nosy_connWrapper_ReadJSON__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_connWrapper_WriteJSON__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_backend_CurrentHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.backend

// skipping Fuzz_Nosy_backend_GetTd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.backend

// skipping Fuzz_Nosy_backend_HeaderByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.backend

// skipping Fuzz_Nosy_backend_Stats__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.backend

// skipping Fuzz_Nosy_backend_SubscribeChainHeadEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.backend

// skipping Fuzz_Nosy_backend_SubscribeNewTxsEvent__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.backend

// skipping Fuzz_Nosy_backend_SyncProgress__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.backend

// skipping Fuzz_Nosy_fullNodeBackend_BlockByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.fullNodeBackend

// skipping Fuzz_Nosy_fullNodeBackend_CurrentBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.fullNodeBackend

// skipping Fuzz_Nosy_fullNodeBackend_SuggestGasTipCap__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.fullNodeBackend

// skipping Fuzz_Nosy_miningNodeBackend_Miner__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethstats.miningNodeBackend

func Fuzz_Nosy_uncleStats_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s uncleStats
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.MarshalJSON()
	})
}

func Fuzz_Nosy_parseEthstatsURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		parseEthstatsURL(url)
	})
}
