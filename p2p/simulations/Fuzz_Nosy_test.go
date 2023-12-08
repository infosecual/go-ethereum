package simulations

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/simulations/adapters"
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

func Fuzz_Nosy_Client_ConnectNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string, nodeID string, peerID string) {
		c := NewClient(url)
		c.ConnectNode(nodeID, peerID)
	})
}

func Fuzz_Nosy_Client_CreateNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var config *adapters.NodeConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		c := NewClient(url)
		c.CreateNode(config)
	})
}

func Fuzz_Nosy_Client_CreateSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		c := NewClient(url)
		c.CreateSnapshot()
	})
}

func Fuzz_Nosy_Client_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string, path string) {
		c := NewClient(url)
		c.Delete(path)
	})
}

func Fuzz_Nosy_Client_DisconnectNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string, nodeID string, peerID string) {
		c := NewClient(url)
		c.DisconnectNode(nodeID, peerID)
	})
}

// skipping Fuzz_Nosy_Client_Get__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Client_GetNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		c := NewClient(url)
		c.GetNetwork()
	})
}

func Fuzz_Nosy_Client_GetNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string, nodeID string) {
		c := NewClient(url)
		c.GetNode(nodeID)
	})
}

func Fuzz_Nosy_Client_GetNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		c := NewClient(url)
		c.GetNodes()
	})
}

func Fuzz_Nosy_Client_LoadSnapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var snap *Snapshot
		fill_err = tp.Fill(&snap)
		if fill_err != nil {
			return
		}
		if snap == nil {
			return
		}

		c := NewClient(url)
		c.LoadSnapshot(snap)
	})
}

// skipping Fuzz_Nosy_Client_Post__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Client_RPCClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var nodeID string
		fill_err = tp.Fill(&nodeID)
		if fill_err != nil {
			return
		}

		c := NewClient(url)
		c.RPCClient(ctx, nodeID)
	})
}

// skipping Fuzz_Nosy_Client_Send__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Client_StartNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		c := NewClient(url)
		c.StartNetwork()
	})
}

func Fuzz_Nosy_Client_StartNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string, nodeID string) {
		c := NewClient(url)
		c.StartNode(nodeID)
	})
}

func Fuzz_Nosy_Client_StopNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		c := NewClient(url)
		c.StopNetwork()
	})
}

func Fuzz_Nosy_Client_StopNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string, nodeID string) {
		c := NewClient(url)
		c.StopNode(nodeID)
	})
}

// skipping Fuzz_Nosy_Client_SubscribeNetwork__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/p2p/simulations.Event

func Fuzz_Nosy_Conn_String__(f *testing.F) {
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

		c.String()
	})
}

func Fuzz_Nosy_Conn_nodesUp__(f *testing.F) {
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

		c.nodesUp()
	})
}

func Fuzz_Nosy_Event_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Event
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

func Fuzz_Nosy_Msg_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Msg
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_Network_Config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Config()
	})
}

func Fuzz_Nosy_Network_Connect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Connect(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_ConnectNodesChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.ConnectNodesChain(ids)
	})
}

func Fuzz_Nosy_Network_ConnectNodesFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.ConnectNodesFull(ids)
	})
}

func Fuzz_Nosy_Network_ConnectNodesRing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.ConnectNodesRing(ids)
	})
}

func Fuzz_Nosy_Network_ConnectNodesStar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		var center enode.ID
		fill_err = tp.Fill(&center)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.ConnectNodesStar(ids, center)
	})
}

func Fuzz_Nosy_Network_ConnectToLastNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.ConnectToLastNode(id)
	})
}

func Fuzz_Nosy_Network_ConnectToRandomNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.ConnectToRandomNode(id)
	})
}

func Fuzz_Nosy_Network_DidConnect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var one enode.ID
		fill_err = tp.Fill(&one)
		if fill_err != nil {
			return
		}
		var other enode.ID
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.DidConnect(one, other)
	})
}

func Fuzz_Nosy_Network_DidDisconnect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var one enode.ID
		fill_err = tp.Fill(&one)
		if fill_err != nil {
			return
		}
		var other enode.ID
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.DidDisconnect(one, other)
	})
}

func Fuzz_Nosy_Network_DidReceive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var sender enode.ID
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var receiver enode.ID
		fill_err = tp.Fill(&receiver)
		if fill_err != nil {
			return
		}
		var proto string
		fill_err = tp.Fill(&proto)
		if fill_err != nil {
			return
		}
		var code uint64
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.DidReceive(sender, receiver, proto, code)
	})
}

func Fuzz_Nosy_Network_DidSend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var sender enode.ID
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var receiver enode.ID
		fill_err = tp.Fill(&receiver)
		if fill_err != nil {
			return
		}
		var proto string
		fill_err = tp.Fill(&proto)
		if fill_err != nil {
			return
		}
		var code uint64
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.DidSend(sender, receiver, proto, code)
	})
}

func Fuzz_Nosy_Network_Disconnect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Disconnect(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_Events__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Events()
	})
}

func Fuzz_Nosy_Network_GetConn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetConn(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_GetNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetNode(id)
	})
}

func Fuzz_Nosy_Network_GetNodeByName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetNodeByName(name)
	})
}

func Fuzz_Nosy_Network_GetNodeIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetNodeIDs(excludeIDs...)
	})
}

func Fuzz_Nosy_Network_GetNodeIDsByProperty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var property string
		fill_err = tp.Fill(&property)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetNodeIDsByProperty(property)
	})
}

func Fuzz_Nosy_Network_GetNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetNodes(excludeIDs...)
	})
}

func Fuzz_Nosy_Network_GetNodesByID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var nodeIDs []enode.ID
		fill_err = tp.Fill(&nodeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetNodesByID(nodeIDs)
	})
}

func Fuzz_Nosy_Network_GetNodesByProperty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var property string
		fill_err = tp.Fill(&property)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetNodesByProperty(property)
	})
}

func Fuzz_Nosy_Network_GetOrCreateConn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetOrCreateConn(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_GetRandomDownNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetRandomDownNode(excludeIDs...)
	})
}

func Fuzz_Nosy_Network_GetRandomNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetRandomNode(excludeIDs...)
	})
}

func Fuzz_Nosy_Network_GetRandomUpNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.GetRandomUpNode(excludeIDs...)
	})
}

func Fuzz_Nosy_Network_InitConn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.InitConn(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_Load__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var snap *Snapshot
		fill_err = tp.Fill(&snap)
		if fill_err != nil {
			return
		}
		if net == nil || snap == nil {
			return
		}

		net.Load(snap)
	})
}

func Fuzz_Nosy_Network_NewNodeWithConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var conf *adapters.NodeConfig
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if net == nil || conf == nil {
			return
		}

		net.NewNodeWithConfig(conf)
	})
}

func Fuzz_Nosy_Network_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Reset()
	})
}

func Fuzz_Nosy_Network_Shutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Shutdown()
	})
}

func Fuzz_Nosy_Network_Snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Snapshot()
	})
}

func Fuzz_Nosy_Network_SnapshotWithServices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var addServices []string
		fill_err = tp.Fill(&addServices)
		if fill_err != nil {
			return
		}
		var removeServices []string
		fill_err = tp.Fill(&removeServices)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.SnapshotWithServices(addServices, removeServices)
	})
}

func Fuzz_Nosy_Network_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Start(id)
	})
}

func Fuzz_Nosy_Network_StartAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.StartAll()
	})
}

func Fuzz_Nosy_Network_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.Stop(id)
	})
}

func Fuzz_Nosy_Network_StopAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.StopAll()
	})
}

// skipping Fuzz_Nosy_Network_Subscribe__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/p2p/simulations.Event

func Fuzz_Nosy_Network_connect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.connect(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_connectNodesChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.connectNodesChain(ids)
	})
}

func Fuzz_Nosy_Network_connectNotConnected__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.connectNotConnected(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_executeConnEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var e *Event
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if net == nil || e == nil {
			return
		}

		net.executeConnEvent(e)
	})
}

func Fuzz_Nosy_Network_executeControlEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var event *Event
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if net == nil || event == nil {
			return
		}

		net.executeControlEvent(event)
	})
}

func Fuzz_Nosy_Network_executeNodeEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var e *Event
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if net == nil || e == nil {
			return
		}

		net.executeNodeEvent(e)
	})
}

func Fuzz_Nosy_Network_getConn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getConn(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_getDownNodeIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getDownNodeIDs()
	})
}

func Fuzz_Nosy_Network_getNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getNode(id)
	})
}

func Fuzz_Nosy_Network_getNodeByName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getNodeByName(name)
	})
}

func Fuzz_Nosy_Network_getNodeIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getNodeIDs(excludeIDs)
	})
}

func Fuzz_Nosy_Network_getNodeIDsByProperty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var property string
		fill_err = tp.Fill(&property)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getNodeIDsByProperty(property)
	})
}

func Fuzz_Nosy_Network_getNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getNodes(excludeIDs)
	})
}

func Fuzz_Nosy_Network_getNodesByID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var nodeIDs []enode.ID
		fill_err = tp.Fill(&nodeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getNodesByID(nodeIDs)
	})
}

func Fuzz_Nosy_Network_getNodesByProperty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var property string
		fill_err = tp.Fill(&property)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getNodesByProperty(property)
	})
}

func Fuzz_Nosy_Network_getOrCreateConn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getOrCreateConn(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_getRandomNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getRandomNode(ids, excludeIDs)
	})
}

func Fuzz_Nosy_Network_getRandomUpNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getRandomUpNode(excludeIDs...)
	})
}

func Fuzz_Nosy_Network_getUpNodeIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.getUpNodeIDs()
	})
}

func Fuzz_Nosy_Network_initConn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var oneID enode.ID
		fill_err = tp.Fill(&oneID)
		if fill_err != nil {
			return
		}
		var otherID enode.ID
		fill_err = tp.Fill(&otherID)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.initConn(oneID, otherID)
	})
}

func Fuzz_Nosy_Network_snapshot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var addServices []string
		fill_err = tp.Fill(&addServices)
		if fill_err != nil {
			return
		}
		var removeServices []string
		fill_err = tp.Fill(&removeServices)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.snapshot(addServices, removeServices)
	})
}

func Fuzz_Nosy_Network_startWithSnapshots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var snapshots map[string][]byte
		fill_err = tp.Fill(&snapshots)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		net.startWithSnapshots(id, snapshots)
	})
}

// skipping Fuzz_Nosy_Network_watchPeerEvents__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/p2p.PeerEvent

func Fuzz_Nosy_Node_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.ID()
	})
}

func Fuzz_Nosy_Node_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.MarshalJSON()
	})
}

func Fuzz_Nosy_Node_NodeInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
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

func Fuzz_Nosy_Node_SetUp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var up bool
		fill_err = tp.Fill(&up)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.SetUp(up)
	})
}

func Fuzz_Nosy_Node_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
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

func Fuzz_Nosy_Node_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var raw []byte
		fill_err = tp.Fill(&raw)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.UnmarshalJSON(raw)
	})
}

func Fuzz_Nosy_Node_Up__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Up()
	})
}

func Fuzz_Nosy_Node_copy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.copy()
	})
}

func Fuzz_Nosy_NoopService_APIs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NoopService
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.APIs()
	})
}

func Fuzz_Nosy_NoopService_Protocols__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NoopService
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Protocols()
	})
}

func Fuzz_Nosy_NoopService_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NoopService
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Start()
	})
}

func Fuzz_Nosy_NoopService_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *NoopService
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Stop()
	})
}

// skipping Fuzz_Nosy_Server_ConnectNode__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_CreateNode__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_CreateSnapshot__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_DELETE__ because parameters include func, chan, or unsupported interface: net/http.HandlerFunc

// skipping Fuzz_Nosy_Server_DisconnectNode__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_GET__ because parameters include func, chan, or unsupported interface: net/http.HandlerFunc

// skipping Fuzz_Nosy_Server_GetMockers__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_GetNetwork__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_GetNode__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_GetNodes__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_JSON__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_LoadSnapshot__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_NodeRPC__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_OPTIONS__ because parameters include func, chan, or unsupported interface: net/http.HandlerFunc

// skipping Fuzz_Nosy_Server_Options__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_POST__ because parameters include func, chan, or unsupported interface: net/http.HandlerFunc

// skipping Fuzz_Nosy_Server_ResetNetwork__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_ServeHTTP__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_StartMocker__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_StartNetwork__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_StartNode__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_StopMocker__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_StopNetwork__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_StopNode__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_StreamNetworkEvents__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_wrapHandler__ because parameters include func, chan, or unsupported interface: net/http.HandlerFunc

func Fuzz_Nosy_Simulation_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var step *Step
		fill_err = tp.Fill(&step)
		if fill_err != nil {
			return
		}
		if network == nil || step == nil {
			return
		}

		s := NewSimulation(network)
		s.Run(ctx, step)
	})
}

func Fuzz_Nosy_Simulation_watchNetwork__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var network *Network
		fill_err = tp.Fill(&network)
		if fill_err != nil {
			return
		}
		var result *StepResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if network == nil || result == nil {
			return
		}

		s := NewSimulation(network)
		s.watchNetwork(result)
	})
}

func Fuzz_Nosy_MsgFilters_Match__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var filterParam string
		fill_err = tp.Fill(&filterParam)
		if fill_err != nil {
			return
		}
		var msg *Msg
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m, err := NewMsgFilters(filterParam)
		if err != nil {
			return
		}
		m.Match(msg)
	})
}

func Fuzz_Nosy_ConnLabel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var source enode.ID
		fill_err = tp.Fill(&source)
		if fill_err != nil {
			return
		}
		var t2 enode.ID
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}

		ConnLabel(source, t2)
	})
}

func Fuzz_Nosy_LookupMocker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, mockerType string) {
		LookupMocker(mockerType)
	})
}

func Fuzz_Nosy_VerifyChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		if t1 == nil || net == nil {
			return
		}

		VerifyChain(t1, net, ids)
	})
}

func Fuzz_Nosy_VerifyFull__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		if t1 == nil || net == nil {
			return
		}

		VerifyFull(t1, net, ids)
	})
}

func Fuzz_Nosy_VerifyRing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		if t1 == nil || net == nil {
			return
		}

		VerifyRing(t1, net, ids)
	})
}

func Fuzz_Nosy_VerifyStar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		var centerIndex int
		fill_err = tp.Fill(&centerIndex)
		if fill_err != nil {
			return
		}
		if t1 == nil || net == nil {
			return
		}

		VerifyStar(t1, net, ids, centerIndex)
	})
}

// skipping Fuzz_Nosy_boot__ because parameters include func, chan, or unsupported interface: chan struct{}

func Fuzz_Nosy_connectNodesInRing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var net *Network
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var nodeCount int
		fill_err = tp.Fill(&nodeCount)
		if fill_err != nil {
			return
		}
		if net == nil {
			return
		}

		connectNodesInRing(net, nodeCount)
	})
}

func Fuzz_Nosy_filterIDs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ids []enode.ID
		fill_err = tp.Fill(&ids)
		if fill_err != nil {
			return
		}
		var excludeIDs []enode.ID
		fill_err = tp.Fill(&excludeIDs)
		if fill_err != nil {
			return
		}

		filterIDs(ids, excludeIDs)
	})
}

// skipping Fuzz_Nosy_probabilistic__ because parameters include func, chan, or unsupported interface: chan struct{}

// skipping Fuzz_Nosy_startStop__ because parameters include func, chan, or unsupported interface: chan struct{}
