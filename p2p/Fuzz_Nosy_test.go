package p2p

import (
	"context"
	"crypto/ecdsa"
	"net"
	"testing"

	"github.com/ethereum/go-ethereum/common/mclock"
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

func Fuzz_Nosy_MsgPipeRW_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *MsgPipeRW
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Close()
	})
}

func Fuzz_Nosy_MsgPipeRW_ReadMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *MsgPipeRW
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReadMsg()
	})
}

func Fuzz_Nosy_MsgPipeRW_WriteMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *MsgPipeRW
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.WriteMsg(msg)
	})
}

func Fuzz_Nosy_Peer_Caps__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.Caps()
	})
}

func Fuzz_Nosy_Peer_Disconnect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		var reason DiscReason
		fill_err = tp.Fill(&reason)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.Disconnect(reason)
	})
}

func Fuzz_Nosy_Peer_Fullname__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.Fullname()
	})
}

func Fuzz_Nosy_Peer_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.ID()
	})
}

func Fuzz_Nosy_Peer_Inbound__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.Inbound()
	})
}

func Fuzz_Nosy_Peer_Info__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.Info()
	})
}

func Fuzz_Nosy_Peer_LocalAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.LocalAddr()
	})
}

func Fuzz_Nosy_Peer_Log__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.Log()
	})
}

func Fuzz_Nosy_Peer_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.Name()
	})
}

func Fuzz_Nosy_Peer_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.Node()
	})
}

func Fuzz_Nosy_Peer_RemoteAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.RemoteAddr()
	})
}

func Fuzz_Nosy_Peer_RunningCap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var versions []uint
		fill_err = tp.Fill(&versions)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.RunningCap(protocol, versions)
	})
}

func Fuzz_Nosy_Peer_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.String()
	})
}

func Fuzz_Nosy_Peer_getProto__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		var code uint64
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.getProto(code)
	})
}

func Fuzz_Nosy_Peer_handle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.handle(msg)
	})
}

func Fuzz_Nosy_Peer_pingLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.pingLoop()
	})
}

// skipping Fuzz_Nosy_Peer_readLoop__ because parameters include func, chan, or unsupported interface: chan<- error

func Fuzz_Nosy_Peer_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}
		var pipe *MsgPipeRW
		fill_err = tp.Fill(&pipe)
		if fill_err != nil {
			return
		}
		if pipe == nil {
			return
		}

		p := NewPeerPipe(id, name, caps, pipe)
		p.run()
	})
}

// skipping Fuzz_Nosy_Peer_startProtocols__ because parameters include func, chan, or unsupported interface: <-chan struct{}

func Fuzz_Nosy_Server_AddPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var node *enode.Node
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if srv == nil || node == nil {
			return
		}

		srv.AddPeer(node)
	})
}

func Fuzz_Nosy_Server_AddTrustedPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var node *enode.Node
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if srv == nil || node == nil {
			return
		}

		srv.AddTrustedPeer(node)
	})
}

func Fuzz_Nosy_Server_LocalNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.LocalNode()
	})
}

func Fuzz_Nosy_Server_NodeInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.NodeInfo()
	})
}

func Fuzz_Nosy_Server_PeerCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.PeerCount()
	})
}

func Fuzz_Nosy_Server_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.Peers()
	})
}

func Fuzz_Nosy_Server_PeersInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.PeersInfo()
	})
}

func Fuzz_Nosy_Server_RemovePeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var node *enode.Node
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if srv == nil || node == nil {
			return
		}

		srv.RemovePeer(node)
	})
}

func Fuzz_Nosy_Server_RemoveTrustedPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var node *enode.Node
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if srv == nil || node == nil {
			return
		}

		srv.RemoveTrustedPeer(node)
	})
}

func Fuzz_Nosy_Server_Self__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.Self()
	})
}

// skipping Fuzz_Nosy_Server_SetupConn__ because parameters include func, chan, or unsupported interface: net.Conn

func Fuzz_Nosy_Server_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.Start()
	})
}

func Fuzz_Nosy_Server_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.Stop()
	})
}

// skipping Fuzz_Nosy_Server_SubscribeEvents__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/p2p.PeerEvent

func Fuzz_Nosy_Server_addPeerChecks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var peers map[enode.ID]*Peer
		fill_err = tp.Fill(&peers)
		if fill_err != nil {
			return
		}
		var inboundCount int
		fill_err = tp.Fill(&inboundCount)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if srv == nil || c == nil {
			return
		}

		srv.addPeerChecks(peers, inboundCount, c)
	})
}

func Fuzz_Nosy_Server_checkInboundConn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var remoteIP net.IP
		fill_err = tp.Fill(&remoteIP)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.checkInboundConn(remoteIP)
	})
}

// skipping Fuzz_Nosy_Server_checkpoint__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/p2p.conn

func Fuzz_Nosy_Server_consumePortMappingRequests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.consumePortMappingRequests()
	})
}

// skipping Fuzz_Nosy_Server_doPeerOp__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.peerOpFunc

func Fuzz_Nosy_Server_launchPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if srv == nil || c == nil {
			return
		}

		srv.launchPeer(c)
	})
}

func Fuzz_Nosy_Server_listenLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.listenLoop()
	})
}

func Fuzz_Nosy_Server_maxDialedConns__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.maxDialedConns()
	})
}

func Fuzz_Nosy_Server_maxInboundConns__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.maxInboundConns()
	})
}

func Fuzz_Nosy_Server_portMappingLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.portMappingLoop()
	})
}

func Fuzz_Nosy_Server_postHandshakeChecks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var peers map[enode.ID]*Peer
		fill_err = tp.Fill(&peers)
		if fill_err != nil {
			return
		}
		var inboundCount int
		fill_err = tp.Fill(&inboundCount)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if srv == nil || c == nil {
			return
		}

		srv.postHandshakeChecks(peers, inboundCount, c)
	})
}

func Fuzz_Nosy_Server_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.run()
	})
}

func Fuzz_Nosy_Server_runPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var p *Peer
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if srv == nil || p == nil {
			return
		}

		srv.runPeer(p)
	})
}

func Fuzz_Nosy_Server_setupConn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var flags connFlag
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var dialDest *enode.Node
		fill_err = tp.Fill(&dialDest)
		if fill_err != nil {
			return
		}
		if srv == nil || c == nil || dialDest == nil {
			return
		}

		srv.setupConn(c, flags, dialDest)
	})
}

func Fuzz_Nosy_Server_setupDialScheduler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.setupDialScheduler()
	})
}

func Fuzz_Nosy_Server_setupDiscovery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.setupDiscovery()
	})
}

func Fuzz_Nosy_Server_setupListening__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.setupListening()
	})
}

func Fuzz_Nosy_Server_setupLocalNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.setupLocalNode()
	})
}

func Fuzz_Nosy_Server_setupPortMapping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.setupPortMapping()
	})
}

func Fuzz_Nosy_Server_setupUDPListening__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var srv *Server
		fill_err = tp.Fill(&srv)
		if fill_err != nil {
			return
		}
		if srv == nil {
			return
		}

		srv.setupUDPListening()
	})
}

func Fuzz_Nosy_conn_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.String()
	})
}

func Fuzz_Nosy_conn_is__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var f2 connFlag
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.is(f2)
	})
}

func Fuzz_Nosy_conn_set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var f2 connFlag
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		var val bool
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.set(f2, val)
	})
}

func Fuzz_Nosy_dialScheduler_addStatic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if d == nil || n == nil {
			return
		}

		d.addStatic(n)
	})
}

func Fuzz_Nosy_dialScheduler_addToStaticPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var task *dialTask
		fill_err = tp.Fill(&task)
		if fill_err != nil {
			return
		}
		if d == nil || task == nil {
			return
		}

		d.addToStaticPool(task)
	})
}

func Fuzz_Nosy_dialScheduler_checkDial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if d == nil || n == nil {
			return
		}

		d.checkDial(n)
	})
}

func Fuzz_Nosy_dialScheduler_expireHistory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.expireHistory()
	})
}

func Fuzz_Nosy_dialScheduler_freeDialSlots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.freeDialSlots()
	})
}

func Fuzz_Nosy_dialScheduler_logStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.logStats()
	})
}

// skipping Fuzz_Nosy_dialScheduler_loop__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enode.Iterator

func Fuzz_Nosy_dialScheduler_peerAdded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if d == nil || c == nil {
			return
		}

		d.peerAdded(c)
	})
}

func Fuzz_Nosy_dialScheduler_peerRemoved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var c *conn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if d == nil || c == nil {
			return
		}

		d.peerRemoved(c)
	})
}

// skipping Fuzz_Nosy_dialScheduler_readNodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enode.Iterator

func Fuzz_Nosy_dialScheduler_rearmHistoryTimer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.rearmHistoryTimer()
	})
}

func Fuzz_Nosy_dialScheduler_removeFromStaticPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var idx int
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.removeFromStaticPool(idx)
	})
}

func Fuzz_Nosy_dialScheduler_removeStatic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if d == nil || n == nil {
			return
		}

		d.removeStatic(n)
	})
}

func Fuzz_Nosy_dialScheduler_startDial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var task *dialTask
		fill_err = tp.Fill(&task)
		if fill_err != nil {
			return
		}
		if d == nil || task == nil {
			return
		}

		d.startDial(task)
	})
}

func Fuzz_Nosy_dialScheduler_startStaticDials__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.startStaticDials(n)
	})
}

func Fuzz_Nosy_dialScheduler_stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.stop()
	})
}

func Fuzz_Nosy_dialScheduler_updateStaticPool__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.updateStaticPool(id)
	})
}

func Fuzz_Nosy_dialTask_String__(f *testing.F) {
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
		var flags connFlag
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		t1 := newDialTask(dest, flags)
		t1.String()
	})
}

func Fuzz_Nosy_dialTask_dial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 *enode.Node
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var flags connFlag
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var d4 *enode.Node
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		if d1 == nil || d == nil || d4 == nil {
			return
		}

		t1 := newDialTask(d1, flags)
		t1.dial(d, d4)
	})
}

func Fuzz_Nosy_dialTask_needResolve__(f *testing.F) {
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
		var flags connFlag
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		t1 := newDialTask(dest, flags)
		t1.needResolve()
	})
}

func Fuzz_Nosy_dialTask_resolve__(f *testing.F) {
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
		var flags connFlag
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if dest == nil || d == nil {
			return
		}

		t1 := newDialTask(dest, flags)
		t1.resolve(d)
	})
}

func Fuzz_Nosy_dialTask_run__(f *testing.F) {
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
		var flags connFlag
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var d *dialScheduler
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if dest == nil || d == nil {
			return
		}

		t1 := newDialTask(dest, flags)
		t1.run(d)
	})
}

func Fuzz_Nosy_eofSignal_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *eofSignal
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Read(buf)
	})
}

func Fuzz_Nosy_expHeap_Pop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *expHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.Pop()
	})
}

// skipping Fuzz_Nosy_expHeap_Push__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_expHeap_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *expHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var item string
		fill_err = tp.Fill(&item)
		if fill_err != nil {
			return
		}
		var exp mclock.AbsTime
		fill_err = tp.Fill(&exp)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.add(item, exp)
	})
}

// skipping Fuzz_Nosy_expHeap_expire__ because parameters include func, chan, or unsupported interface: func(string)

func Fuzz_Nosy_expHeap_nextExpiry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *expHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.nextExpiry()
	})
}

func Fuzz_Nosy_meteredConn_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *meteredConn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Read(b)
	})
}

func Fuzz_Nosy_meteredConn_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *meteredConn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.Write(b)
	})
}

func Fuzz_Nosy_msgEventer_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev *msgEventer
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if ev == nil {
			return
		}

		ev.Close()
	})
}

func Fuzz_Nosy_msgEventer_ReadMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev *msgEventer
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		if ev == nil {
			return
		}

		ev.ReadMsg()
	})
}

func Fuzz_Nosy_msgEventer_WriteMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ev *msgEventer
		fill_err = tp.Fill(&ev)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if ev == nil {
			return
		}

		ev.WriteMsg(msg)
	})
}

func Fuzz_Nosy_peerError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pe *peerError
		fill_err = tp.Fill(&pe)
		if fill_err != nil {
			return
		}
		if pe == nil {
			return
		}

		pe.Error()
	})
}

func Fuzz_Nosy_protoRW_ReadMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rw *protoRW
		fill_err = tp.Fill(&rw)
		if fill_err != nil {
			return
		}
		if rw == nil {
			return
		}

		rw.ReadMsg()
	})
}

func Fuzz_Nosy_protoRW_WriteMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rw *protoRW
		fill_err = tp.Fill(&rw)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if rw == nil {
			return
		}

		rw.WriteMsg(msg)
	})
}

func Fuzz_Nosy_rlpxTransport_ReadMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *rlpxTransport
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.ReadMsg()
	})
}

func Fuzz_Nosy_rlpxTransport_WriteMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *rlpxTransport
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.WriteMsg(msg)
	})
}

// skipping Fuzz_Nosy_rlpxTransport_close__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_rlpxTransport_doEncHandshake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *rlpxTransport
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var prv *ecdsa.PrivateKey
		fill_err = tp.Fill(&prv)
		if fill_err != nil {
			return
		}
		if t1 == nil || prv == nil {
			return
		}

		t1.doEncHandshake(prv)
	})
}

func Fuzz_Nosy_rlpxTransport_doProtoHandshake__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *rlpxTransport
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var our *protoHandshake
		fill_err = tp.Fill(&our)
		if fill_err != nil {
			return
		}
		if t1 == nil || our == nil {
			return
		}

		t1.doProtoHandshake(our)
	})
}

func Fuzz_Nosy_sharedUDPConn_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *sharedUDPConn
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Close()
	})
}

func Fuzz_Nosy_sharedUDPConn_ReadFromUDP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *sharedUDPConn
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ReadFromUDP(b)
	})
}

func Fuzz_Nosy_Cap_Cmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cap Cap
		fill_err = tp.Fill(&cap)
		if fill_err != nil {
			return
		}
		var other Cap
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}

		cap.Cmp(other)
	})
}

func Fuzz_Nosy_Cap_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cap Cap
		fill_err = tp.Fill(&cap)
		if fill_err != nil {
			return
		}

		cap.String()
	})
}

func Fuzz_Nosy_DiscReason_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DiscReason
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.Error()
	})
}

func Fuzz_Nosy_DiscReason_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DiscReason
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		d.String()
	})
}

// skipping Fuzz_Nosy_Msg_Decode__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Msg_Discard__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Discard()
	})
}

func Fuzz_Nosy_Msg_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.String()
	})
}

func Fuzz_Nosy_Msg_Time__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		msg.Time()
	})
}

// skipping Fuzz_Nosy_MsgReader_ReadMsg__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.MsgReader

// skipping Fuzz_Nosy_MsgWriter_WriteMsg__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.MsgWriter

// skipping Fuzz_Nosy_NodeDialer_Dial__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.NodeDialer

func Fuzz_Nosy_Protocol_cap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p Protocol
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		p.cap()
	})
}

func Fuzz_Nosy_connFlag_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 connFlag
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}

		f1.String()
	})
}

func Fuzz_Nosy_dialConfig_withDefaults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg dialConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		cfg.withDefaults()
	})
}

func Fuzz_Nosy_expHeap_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h expHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		h.Len()
	})
}

func Fuzz_Nosy_expHeap_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h expHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		h.Less(i, j)
	})
}

func Fuzz_Nosy_expHeap_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h expHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j int
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		h.Swap(i, j)
	})
}

func Fuzz_Nosy_expHeap_contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h expHeap
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var item string
		fill_err = tp.Fill(&item)
		if fill_err != nil {
			return
		}

		h.contains(item)
	})
}

// skipping Fuzz_Nosy_nodeResolver_Resolve__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.nodeResolver

func Fuzz_Nosy_tcpDialer_Dial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 tcpDialer
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var dest *enode.Node
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}
		if dest == nil {
			return
		}

		t1.Dial(ctx, dest)
	})
}

// skipping Fuzz_Nosy_transport_close__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.transport

// skipping Fuzz_Nosy_transport_doEncHandshake__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.transport

// skipping Fuzz_Nosy_transport_doProtoHandshake__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.transport

func Fuzz_Nosy_countMatchingProtocols__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var protocols []Protocol
		fill_err = tp.Fill(&protocols)
		if fill_err != nil {
			return
		}
		var caps []Cap
		fill_err = tp.Fill(&caps)
		if fill_err != nil {
			return
		}

		countMatchingProtocols(protocols, caps)
	})
}

// skipping Fuzz_Nosy_markDialError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_matchProtocols__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p.MsgReadWriter
