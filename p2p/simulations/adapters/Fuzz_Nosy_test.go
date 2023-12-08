package adapters

import (
	"context"
	"net"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/p2p/enode"
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

func Fuzz_Nosy_ExecAdapter_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, baseDir string) {
		e := NewExecAdapter(baseDir)
		e.Name()
	})
}

func Fuzz_Nosy_ExecAdapter_NewNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var baseDir string
		fill_err = tp.Fill(&baseDir)
		if fill_err != nil {
			return
		}
		var config *NodeConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		e := NewExecAdapter(baseDir)
		e.NewNode(config)
	})
}

func Fuzz_Nosy_ExecNode_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Addr()
	})
}

func Fuzz_Nosy_ExecNode_Client__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Client()
	})
}

func Fuzz_Nosy_ExecNode_NodeInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.NodeInfo()
	})
}

func Fuzz_Nosy_ExecNode_ServeRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var clientConn *websocket.Conn
		fill_err = tp.Fill(&clientConn)
		if fill_err != nil {
			return
		}
		if n == nil || clientConn == nil {
			return
		}

		n.ServeRPC(clientConn)
	})
}

func Fuzz_Nosy_ExecNode_Snapshots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Snapshots()
	})
}

func Fuzz_Nosy_ExecNode_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var snapshots map[string][]byte
		fill_err = tp.Fill(&snapshots)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Start(snapshots)
	})
}

func Fuzz_Nosy_ExecNode_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Stop()
	})
}

func Fuzz_Nosy_ExecNode_execCommand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.execCommand()
	})
}

func Fuzz_Nosy_ExecNode_waitForStartupJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *ExecNode
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.waitForStartupJSON(ctx)
	})
}

func Fuzz_Nosy_NodeConfig_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		n := RandomNodeConfig()
		n.UnmarshalJSON(d1)
	})
}

func Fuzz_Nosy_NodeConfig_initEnode__(f *testing.F) {
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
		var tcpport int
		fill_err = tp.Fill(&tcpport)
		if fill_err != nil {
			return
		}
		var udpport int
		fill_err = tp.Fill(&udpport)
		if fill_err != nil {
			return
		}

		n := RandomNodeConfig()
		n.initEnode(ip, tcpport, udpport)
	})
}

func Fuzz_Nosy_SimAdapter_Dial__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SimAdapter
		fill_err = tp.Fill(&s)
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
		if s == nil || dest == nil {
			return
		}

		s.Dial(ctx, dest)
	})
}

func Fuzz_Nosy_SimAdapter_DialRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SimAdapter
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.DialRPC(id)
	})
}

func Fuzz_Nosy_SimAdapter_GetNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SimAdapter
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetNode(id)
	})
}

func Fuzz_Nosy_SimAdapter_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SimAdapter
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Name()
	})
}

func Fuzz_Nosy_SimAdapter_NewNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SimAdapter
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var config *NodeConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if s == nil || config == nil {
			return
		}

		s.NewNode(config)
	})
}

func Fuzz_Nosy_SimNode_Addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Addr()
	})
}

func Fuzz_Nosy_SimNode_Client__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Client()
	})
}

func Fuzz_Nosy_SimNode_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Close()
	})
}

func Fuzz_Nosy_SimNode_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Node()
	})
}

func Fuzz_Nosy_SimNode_NodeInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.NodeInfo()
	})
}

func Fuzz_Nosy_SimNode_ServeRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		var conn *websocket.Conn
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		if sn == nil || conn == nil {
			return
		}

		sn.ServeRPC(conn)
	})
}

func Fuzz_Nosy_SimNode_Server__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Server()
	})
}

func Fuzz_Nosy_SimNode_Service__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Service(name)
	})
}

func Fuzz_Nosy_SimNode_ServiceMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.ServiceMap()
	})
}

func Fuzz_Nosy_SimNode_Services__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Services()
	})
}

func Fuzz_Nosy_SimNode_Snapshots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Snapshots()
	})
}

func Fuzz_Nosy_SimNode_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		var snapshots map[string][]byte
		fill_err = tp.Fill(&snapshots)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Start(snapshots)
	})
}

func Fuzz_Nosy_SimNode_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sn *SimNode
		fill_err = tp.Fill(&sn)
		if fill_err != nil {
			return
		}
		if sn == nil {
			return
		}

		sn.Stop()
	})
}

// skipping Fuzz_Nosy_SimNode_SubscribeEvents__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/p2p.PeerEvent

func Fuzz_Nosy_wsRPCDialer_DialRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *wsRPCDialer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.DialRPC(id)
	})
}

// skipping Fuzz_Nosy_Node_Addr__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.Node

// skipping Fuzz_Nosy_Node_Client__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.Node

// skipping Fuzz_Nosy_Node_NodeInfo__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.Node

// skipping Fuzz_Nosy_Node_ServeRPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.Node

// skipping Fuzz_Nosy_Node_Snapshots__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.Node

// skipping Fuzz_Nosy_Node_Start__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.Node

// skipping Fuzz_Nosy_Node_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.Node

// skipping Fuzz_Nosy_NodeAdapter_Name__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.NodeAdapter

// skipping Fuzz_Nosy_NodeAdapter_NewNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.NodeAdapter

// skipping Fuzz_Nosy_RPCDialer_DialRPC__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.RPCDialer

func Fuzz_Nosy_SnapshotAPI_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api SnapshotAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}

		api.Snapshot()
	})
}

// skipping Fuzz_Nosy_RegisterLifecycles__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/simulations/adapters.LifecycleConstructors

func Fuzz_Nosy_wsCopy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wg *sync.WaitGroup
		fill_err = tp.Fill(&wg)
		if fill_err != nil {
			return
		}
		var src *websocket.Conn
		fill_err = tp.Fill(&src)
		if fill_err != nil {
			return
		}
		var dst *websocket.Conn
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if wg == nil || src == nil || dst == nil {
			return
		}

		wsCopy(wg, src, dst)
	})
}
