package netutil

import (
	"net"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/mclock"
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

func Fuzz_Nosy_DistinctNetSet_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DistinctNetSet
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Add(ip)
	})
}

func Fuzz_Nosy_DistinctNetSet_Remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DistinctNetSet
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Remove(ip)
	})
}

func Fuzz_Nosy_DistinctNetSet_key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *DistinctNetSet
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.key(ip)
	})
}

func Fuzz_Nosy_IPTracker_AddContact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var window time.Duration
		fill_err = tp.Fill(&window)
		if fill_err != nil {
			return
		}
		var contactWindow time.Duration
		fill_err = tp.Fill(&contactWindow)
		if fill_err != nil {
			return
		}
		var minStatements int
		fill_err = tp.Fill(&minStatements)
		if fill_err != nil {
			return
		}
		var host string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}

		it := NewIPTracker(window, contactWindow, minStatements)
		it.AddContact(host)
	})
}

func Fuzz_Nosy_IPTracker_AddStatement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var window time.Duration
		fill_err = tp.Fill(&window)
		if fill_err != nil {
			return
		}
		var contactWindow time.Duration
		fill_err = tp.Fill(&contactWindow)
		if fill_err != nil {
			return
		}
		var minStatements int
		fill_err = tp.Fill(&minStatements)
		if fill_err != nil {
			return
		}
		var host string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var endpoint string
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}

		it := NewIPTracker(window, contactWindow, minStatements)
		it.AddStatement(host, endpoint)
	})
}

func Fuzz_Nosy_IPTracker_PredictEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var window time.Duration
		fill_err = tp.Fill(&window)
		if fill_err != nil {
			return
		}
		var contactWindow time.Duration
		fill_err = tp.Fill(&contactWindow)
		if fill_err != nil {
			return
		}
		var minStatements int
		fill_err = tp.Fill(&minStatements)
		if fill_err != nil {
			return
		}

		it := NewIPTracker(window, contactWindow, minStatements)
		it.PredictEndpoint()
	})
}

func Fuzz_Nosy_IPTracker_PredictFullConeNAT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var window time.Duration
		fill_err = tp.Fill(&window)
		if fill_err != nil {
			return
		}
		var contactWindow time.Duration
		fill_err = tp.Fill(&contactWindow)
		if fill_err != nil {
			return
		}
		var minStatements int
		fill_err = tp.Fill(&minStatements)
		if fill_err != nil {
			return
		}

		it := NewIPTracker(window, contactWindow, minStatements)
		it.PredictFullConeNAT()
	})
}

func Fuzz_Nosy_IPTracker_gcContact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var window time.Duration
		fill_err = tp.Fill(&window)
		if fill_err != nil {
			return
		}
		var contactWindow time.Duration
		fill_err = tp.Fill(&contactWindow)
		if fill_err != nil {
			return
		}
		var minStatements int
		fill_err = tp.Fill(&minStatements)
		if fill_err != nil {
			return
		}
		var now mclock.AbsTime
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}

		it := NewIPTracker(window, contactWindow, minStatements)
		it.gcContact(now)
	})
}

func Fuzz_Nosy_IPTracker_gcStatements__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var window time.Duration
		fill_err = tp.Fill(&window)
		if fill_err != nil {
			return
		}
		var contactWindow time.Duration
		fill_err = tp.Fill(&contactWindow)
		if fill_err != nil {
			return
		}
		var minStatements int
		fill_err = tp.Fill(&minStatements)
		if fill_err != nil {
			return
		}
		var now mclock.AbsTime
		fill_err = tp.Fill(&now)
		if fill_err != nil {
			return
		}

		it := NewIPTracker(window, contactWindow, minStatements)
		it.gcStatements(now)
	})
}

func Fuzz_Nosy_Netlist_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string, cidr string) {
		l, err := ParseNetlist(s)
		if err != nil {
			return
		}
		l.Add(cidr)
	})
}

func Fuzz_Nosy_Netlist_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s string
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}

		l, err := ParseNetlist(s)
		if err != nil {
			return
		}
		l.Contains(ip)
	})
}

// skipping Fuzz_Nosy_Netlist_UnmarshalTOML__ because parameters include func, chan, or unsupported interface: func(interface{}) error

func Fuzz_Nosy_DistinctNetSet_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s DistinctNetSet
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}

		s.Contains(ip)
	})
}

func Fuzz_Nosy_DistinctNetSet_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s DistinctNetSet
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.Len()
	})
}

func Fuzz_Nosy_DistinctNetSet_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s DistinctNetSet
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.String()
	})
}

func Fuzz_Nosy_Netlist_MarshalTOML__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		l, err := ParseNetlist(s)
		if err != nil {
			return
		}
		l.MarshalTOML()
	})
}

func Fuzz_Nosy_IsLAN__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}

		IsLAN(ip)
	})
}

func Fuzz_Nosy_IsSpecialNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}

		IsSpecialNetwork(ip)
	})
}

// skipping Fuzz_Nosy_IsTemporaryError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_IsTimeout__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_SameNet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bits uint
		fill_err = tp.Fill(&bits)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var other net.IP
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}

		SameNet(bits, ip, other)
	})
}

// skipping Fuzz_Nosy_isPacketTooBig__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_sameNet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bits uint
		fill_err = tp.Fill(&bits)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var other net.IP
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}

		sameNet(bits, ip, other)
	})
}
