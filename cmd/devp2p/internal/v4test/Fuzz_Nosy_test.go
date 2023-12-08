package v4test

import (
	"testing"

	"github.com/ethereum/go-ethereum/internal/utesting"
	"github.com/ethereum/go-ethereum/p2p/discover/v4wire"
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

func Fuzz_Nosy_pingWithJunk_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *pingWithJunk
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		req.Kind()
	})
}

func Fuzz_Nosy_pingWithJunk_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *pingWithJunk
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		req.Name()
	})
}

func Fuzz_Nosy_pingWrongType_Kind__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *pingWrongType
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		req.Kind()
	})
}

func Fuzz_Nosy_pingWrongType_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *pingWrongType
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		req.Name()
	})
}

func Fuzz_Nosy_testenv_checkPingPong__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remote string, listen1 string, listen2 string, pingHash []byte) {
		te := newTestEnv(remote, listen1, listen2)
		te.checkPingPong(pingHash)
	})
}

// skipping Fuzz_Nosy_testenv_checkPong__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v4wire.Packet

func Fuzz_Nosy_testenv_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remote string, listen1 string, listen2 string) {
		te := newTestEnv(remote, listen1, listen2)
		te.close()
	})
}

// skipping Fuzz_Nosy_testenv_localEndpoint__ because parameters include func, chan, or unsupported interface: net.PacketConn

// skipping Fuzz_Nosy_testenv_read__ because parameters include func, chan, or unsupported interface: net.PacketConn

func Fuzz_Nosy_testenv_remoteEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remote string, listen1 string, listen2 string) {
		te := newTestEnv(remote, listen1, listen2)
		te.remoteEndpoint()
	})
}

// skipping Fuzz_Nosy_testenv_send__ because parameters include func, chan, or unsupported interface: net.PacketConn

func Fuzz_Nosy_BasicFindnode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		BasicFindnode(t1)
	})
}

func Fuzz_Nosy_BasicPing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		BasicPing(t1)
	})
}

func Fuzz_Nosy_BondThenPingWithWrongFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		BondThenPingWithWrongFrom(t1)
	})
}

func Fuzz_Nosy_FindnodeAmplificationInvalidPongHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		FindnodeAmplificationInvalidPongHash(t1)
	})
}

func Fuzz_Nosy_FindnodeAmplificationWrongIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		FindnodeAmplificationWrongIP(t1)
	})
}

func Fuzz_Nosy_FindnodePastExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		FindnodePastExpiration(t1)
	})
}

func Fuzz_Nosy_FindnodeWithoutEndpointProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		FindnodeWithoutEndpointProof(t1)
	})
}

func Fuzz_Nosy_PingExtraData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		PingExtraData(t1)
	})
}

func Fuzz_Nosy_PingExtraDataWrongFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		PingExtraDataWrongFrom(t1)
	})
}

func Fuzz_Nosy_PingPastExpiration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		PingPastExpiration(t1)
	})
}

func Fuzz_Nosy_PingWrongFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		PingWrongFrom(t1)
	})
}

func Fuzz_Nosy_PingWrongTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		PingWrongTo(t1)
	})
}

func Fuzz_Nosy_UnsolicitedNeighbors__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		UnsolicitedNeighbors(t1)
	})
}

func Fuzz_Nosy_WrongPacketType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		WrongPacketType(t1)
	})
}

func Fuzz_Nosy_bond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *utesting.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var te *testenv
		fill_err = tp.Fill(&te)
		if fill_err != nil {
			return
		}
		if t1 == nil || te == nil {
			return
		}

		bond(t1, te)
	})
}

func Fuzz_Nosy_contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns []v4wire.Node
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var key v4wire.Pubkey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		contains(ns, key)
	})
}
