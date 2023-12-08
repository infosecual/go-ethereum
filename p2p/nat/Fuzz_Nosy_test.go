package nat

import (
	"net"
	"testing"
	"time"

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

func Fuzz_Nosy_autodisc_AddMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *autodisc
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var extport int
		fill_err = tp.Fill(&extport)
		if fill_err != nil {
			return
		}
		var intport int
		fill_err = tp.Fill(&intport)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var lifetime time.Duration
		fill_err = tp.Fill(&lifetime)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.AddMapping(protocol, extport, intport, name, lifetime)
	})
}

func Fuzz_Nosy_autodisc_DeleteMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *autodisc
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var extport int
		fill_err = tp.Fill(&extport)
		if fill_err != nil {
			return
		}
		var intport int
		fill_err = tp.Fill(&intport)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.DeleteMapping(protocol, extport, intport)
	})
}

func Fuzz_Nosy_autodisc_ExternalIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *autodisc
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ExternalIP()
	})
}

func Fuzz_Nosy_autodisc_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *autodisc
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.String()
	})
}

func Fuzz_Nosy_autodisc_wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *autodisc
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.wait()
	})
}

func Fuzz_Nosy_pmp_AddMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *pmp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var extport int
		fill_err = tp.Fill(&extport)
		if fill_err != nil {
			return
		}
		var intport int
		fill_err = tp.Fill(&intport)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var lifetime time.Duration
		fill_err = tp.Fill(&lifetime)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.AddMapping(protocol, extport, intport, name, lifetime)
	})
}

func Fuzz_Nosy_pmp_DeleteMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *pmp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var extport int
		fill_err = tp.Fill(&extport)
		if fill_err != nil {
			return
		}
		var intport int
		fill_err = tp.Fill(&intport)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.DeleteMapping(protocol, extport, intport)
	})
}

func Fuzz_Nosy_pmp_ExternalIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *pmp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ExternalIP()
	})
}

func Fuzz_Nosy_pmp_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *pmp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.String()
	})
}

func Fuzz_Nosy_upnp_AddMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *upnp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var extport int
		fill_err = tp.Fill(&extport)
		if fill_err != nil {
			return
		}
		var intport int
		fill_err = tp.Fill(&intport)
		if fill_err != nil {
			return
		}
		var desc string
		fill_err = tp.Fill(&desc)
		if fill_err != nil {
			return
		}
		var lifetime time.Duration
		fill_err = tp.Fill(&lifetime)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.AddMapping(protocol, extport, intport, desc, lifetime)
	})
}

func Fuzz_Nosy_upnp_DeleteMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *upnp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var extport int
		fill_err = tp.Fill(&extport)
		if fill_err != nil {
			return
		}
		var intport int
		fill_err = tp.Fill(&intport)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.DeleteMapping(protocol, extport, intport)
	})
}

func Fuzz_Nosy_upnp_ExternalIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *upnp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ExternalIP()
	})
}

func Fuzz_Nosy_upnp_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *upnp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.String()
	})
}

func Fuzz_Nosy_upnp_addAnyPortMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *upnp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var extport int
		fill_err = tp.Fill(&extport)
		if fill_err != nil {
			return
		}
		var intport int
		fill_err = tp.Fill(&intport)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var desc string
		fill_err = tp.Fill(&desc)
		if fill_err != nil {
			return
		}
		var lifetimeS uint32
		fill_err = tp.Fill(&lifetimeS)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.addAnyPortMapping(protocol, extport, intport, ip, desc, lifetimeS)
	})
}

func Fuzz_Nosy_upnp_internalAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *upnp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.internalAddress()
	})
}

func Fuzz_Nosy_upnp_natEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *upnp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.natEnabled()
	})
}

func Fuzz_Nosy_upnp_randomPort__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *upnp
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.randomPort()
	})
}

// skipping Fuzz_Nosy_upnp_withRateLimit__ because parameters include func, chan, or unsupported interface: func() error

func Fuzz_Nosy_ExtIP_AddMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 ExtIP
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 int
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}
		var _x5 string
		fill_err = tp.Fill(&_x5)
		if fill_err != nil {
			return
		}
		var _x6 time.Duration
		fill_err = tp.Fill(&_x6)
		if fill_err != nil {
			return
		}

		_x1.AddMapping(_x2, _x3, _x4, _x5, _x6)
	})
}

func Fuzz_Nosy_ExtIP_DeleteMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 ExtIP
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 int
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var _x4 int
		fill_err = tp.Fill(&_x4)
		if fill_err != nil {
			return
		}

		_x1.DeleteMapping(_x2, _x3, _x4)
	})
}

func Fuzz_Nosy_ExtIP_ExternalIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n ExtIP
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.ExternalIP()
	})
}

func Fuzz_Nosy_ExtIP_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n ExtIP
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		n.String()
	})
}

func Fuzz_Nosy_Interface_AddMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var extport int
		fill_err = tp.Fill(&extport)
		if fill_err != nil {
			return
		}
		var intport int
		fill_err = tp.Fill(&intport)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var lifetime time.Duration
		fill_err = tp.Fill(&lifetime)
		if fill_err != nil {
			return
		}

		_x1 := discoverPMP()
		_x1.AddMapping(protocol, extport, intport, name, lifetime)
	})
}

func Fuzz_Nosy_Interface_DeleteMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, protocol string, extport int, intport int) {
		_x1 := discoverPMP()
		_x1.DeleteMapping(protocol, extport, intport)
	})
}

// skipping Fuzz_Nosy_upnpClient_AddPortMapping__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/nat.upnpClient

// skipping Fuzz_Nosy_upnpClient_DeletePortMapping__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/nat.upnpClient

// skipping Fuzz_Nosy_upnpClient_GetExternalIPAddress__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/nat.upnpClient

// skipping Fuzz_Nosy_upnpClient_GetNATRSIPStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/nat.upnpClient

// skipping Fuzz_Nosy_Map__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/nat.Interface

// skipping Fuzz_Nosy_discover__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/p2p/nat.upnp
