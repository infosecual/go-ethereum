package discover

import (
	"context"
	"crypto/ecdsa"
	"net"
	"testing"

	"github.com/ethereum/go-ethereum/p2p/discover/v4wire"
	"github.com/ethereum/go-ethereum/p2p/discover/v5wire"
	"github.com/ethereum/go-ethereum/p2p/enode"
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

func Fuzz_Nosy_Table_Nodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.Nodes()
	})
}

func Fuzz_Nosy_Table_addIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var b *bucket
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if tab == nil || b == nil {
			return
		}

		tab.addIP(b, ip)
	})
}

func Fuzz_Nosy_Table_addReplacement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var b *bucket
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var n *node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if tab == nil || b == nil || n == nil {
			return
		}

		tab.addReplacement(b, n)
	})
}

func Fuzz_Nosy_Table_addSeenNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var n *node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if tab == nil || n == nil {
			return
		}

		tab.addSeenNode(n)
	})
}

func Fuzz_Nosy_Table_addVerifiedNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var n *node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if tab == nil || n == nil {
			return
		}

		tab.addVerifiedNode(n)
	})
}

func Fuzz_Nosy_Table_bucket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.bucket(id)
	})
}

func Fuzz_Nosy_Table_bucketAtDistance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var d int
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.bucketAtDistance(d)
	})
}

func Fuzz_Nosy_Table_bucketLen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.bucketLen(id)
	})
}

func Fuzz_Nosy_Table_bumpInBucket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var b *bucket
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var n *node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if tab == nil || b == nil || n == nil {
			return
		}

		tab.bumpInBucket(b, n)
	})
}

func Fuzz_Nosy_Table_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.close()
	})
}

func Fuzz_Nosy_Table_copyLiveNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.copyLiveNodes()
	})
}

func Fuzz_Nosy_Table_delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var node *node
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if tab == nil || node == nil {
			return
		}

		tab.delete(node)
	})
}

func Fuzz_Nosy_Table_deleteInBucket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var b *bucket
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var n *node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if tab == nil || b == nil || n == nil {
			return
		}

		tab.deleteInBucket(b, n)
	})
}

// skipping Fuzz_Nosy_Table_doRefresh__ because parameters include func, chan, or unsupported interface: chan struct{}

// skipping Fuzz_Nosy_Table_doRevalidate__ because parameters include func, chan, or unsupported interface: chan<- struct{}

func Fuzz_Nosy_Table_findnodeByID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var t2 enode.ID
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var nresults int
		fill_err = tp.Fill(&nresults)
		if fill_err != nil {
			return
		}
		var preferLive bool
		fill_err = tp.Fill(&preferLive)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.findnodeByID(t2, nresults, preferLive)
	})
}

func Fuzz_Nosy_Table_getNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.getNode(id)
	})
}

func Fuzz_Nosy_Table_isInitDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.isInitDone()
	})
}

func Fuzz_Nosy_Table_len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.len()
	})
}

func Fuzz_Nosy_Table_loadSeedNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.loadSeedNodes()
	})
}

func Fuzz_Nosy_Table_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.loop()
	})
}

func Fuzz_Nosy_Table_nextRefreshTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.nextRefreshTime()
	})
}

func Fuzz_Nosy_Table_nextRevalidateTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.nextRevalidateTime()
	})
}

func Fuzz_Nosy_Table_nodeToRevalidate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.nodeToRevalidate()
	})
}

func Fuzz_Nosy_Table_refresh__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.refresh()
	})
}

func Fuzz_Nosy_Table_removeIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var b *bucket
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if tab == nil || b == nil {
			return
		}

		tab.removeIP(b, ip)
	})
}

func Fuzz_Nosy_Table_replace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var b *bucket
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var last *node
		fill_err = tp.Fill(&last)
		if fill_err != nil {
			return
		}
		if tab == nil || b == nil || last == nil {
			return
		}

		tab.replace(b, last)
	})
}

func Fuzz_Nosy_Table_seedRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.seedRand()
	})
}

func Fuzz_Nosy_Table_self__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.self()
	})
}

func Fuzz_Nosy_Table_setFallbackNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tab *Table
		fill_err = tp.Fill(&tab)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		if tab == nil {
			return
		}

		tab.setFallbackNodes(nodes)
	})
}

func Fuzz_Nosy_UDPv4_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Close()
	})
}

func Fuzz_Nosy_UDPv4_LookupPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PublicKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if t1 == nil || key == nil {
			return
		}

		t1.LookupPubkey(key)
	})
}

func Fuzz_Nosy_UDPv4_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.Ping(n)
	})
}

func Fuzz_Nosy_UDPv4_RandomNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RandomNodes()
	})
}

func Fuzz_Nosy_UDPv4_RequestENR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.RequestENR(n)
	})
}

func Fuzz_Nosy_UDPv4_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.Resolve(n)
	})
}

func Fuzz_Nosy_UDPv4_Self__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Self()
	})
}

func Fuzz_Nosy_UDPv4_checkBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.checkBond(id, ip)
	})
}

func Fuzz_Nosy_UDPv4_ensureBond__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var toid enode.ID
		fill_err = tp.Fill(&toid)
		if fill_err != nil {
			return
		}
		var toaddr *net.UDPAddr
		fill_err = tp.Fill(&toaddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || toaddr == nil {
			return
		}

		t1.ensureBond(toid, toaddr)
	})
}

func Fuzz_Nosy_UDPv4_findnode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var toid enode.ID
		fill_err = tp.Fill(&toid)
		if fill_err != nil {
			return
		}
		var toaddr *net.UDPAddr
		fill_err = tp.Fill(&toaddr)
		if fill_err != nil {
			return
		}
		var t4 v4wire.Pubkey
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || toaddr == nil {
			return
		}

		t1.findnode(toid, toaddr, t4)
	})
}

func Fuzz_Nosy_UDPv4_handleENRRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var mac []byte
		fill_err = tp.Fill(&mac)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.handleENRRequest(h, from, fromID, mac)
	})
}

func Fuzz_Nosy_UDPv4_handleFindnode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var mac []byte
		fill_err = tp.Fill(&mac)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.handleFindnode(h, from, fromID, mac)
	})
}

func Fuzz_Nosy_UDPv4_handlePacket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if t1 == nil || from == nil {
			return
		}

		t1.handlePacket(from, buf)
	})
}

func Fuzz_Nosy_UDPv4_handlePing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var mac []byte
		fill_err = tp.Fill(&mac)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.handlePing(h, from, fromID, mac)
	})
}

// skipping Fuzz_Nosy_UDPv4_handleReply__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v4wire.Packet

func Fuzz_Nosy_UDPv4_lookupRandom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.lookupRandom()
	})
}

func Fuzz_Nosy_UDPv4_lookupSelf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.lookupSelf()
	})
}

func Fuzz_Nosy_UDPv4_loop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.loop()
	})
}

func Fuzz_Nosy_UDPv4_makePing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var toaddr *net.UDPAddr
		fill_err = tp.Fill(&toaddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || toaddr == nil {
			return
		}

		t1.makePing(toaddr)
	})
}

func Fuzz_Nosy_UDPv4_newLookup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var targetKey encPubkey
		fill_err = tp.Fill(&targetKey)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.newLookup(ctx, targetKey)
	})
}

func Fuzz_Nosy_UDPv4_newRandomLookup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.newRandomLookup(ctx)
	})
}

func Fuzz_Nosy_UDPv4_nodeFromRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var sender *net.UDPAddr
		fill_err = tp.Fill(&sender)
		if fill_err != nil {
			return
		}
		var rn v4wire.Node
		fill_err = tp.Fill(&rn)
		if fill_err != nil {
			return
		}
		if t1 == nil || sender == nil {
			return
		}

		t1.nodeFromRPC(sender, rn)
	})
}

func Fuzz_Nosy_UDPv4_ourEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.ourEndpoint()
	})
}

// skipping Fuzz_Nosy_UDPv4_pending__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.replyMatchFunc

func Fuzz_Nosy_UDPv4_ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.ping(n)
	})
}

// skipping Fuzz_Nosy_UDPv4_readLoop__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/p2p/discover.ReadPacket

// skipping Fuzz_Nosy_UDPv4_send__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v4wire.Packet

// skipping Fuzz_Nosy_UDPv4_sendPing__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_UDPv4_verifyENRRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromKey v4wire.Pubkey
		fill_err = tp.Fill(&fromKey)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.verifyENRRequest(h, from, fromID, fromKey)
	})
}

func Fuzz_Nosy_UDPv4_verifyENRResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromKey v4wire.Pubkey
		fill_err = tp.Fill(&fromKey)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.verifyENRResponse(h, from, fromID, fromKey)
	})
}

func Fuzz_Nosy_UDPv4_verifyFindnode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromKey v4wire.Pubkey
		fill_err = tp.Fill(&fromKey)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.verifyFindnode(h, from, fromID, fromKey)
	})
}

func Fuzz_Nosy_UDPv4_verifyNeighbors__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromKey v4wire.Pubkey
		fill_err = tp.Fill(&fromKey)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.verifyNeighbors(h, from, fromID, fromKey)
	})
}

func Fuzz_Nosy_UDPv4_verifyPing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromKey v4wire.Pubkey
		fill_err = tp.Fill(&fromKey)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.verifyPing(h, from, fromID, fromKey)
	})
}

func Fuzz_Nosy_UDPv4_verifyPong__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var h *packetHandlerV4
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromKey v4wire.Pubkey
		fill_err = tp.Fill(&fromKey)
		if fill_err != nil {
			return
		}
		if t1 == nil || h == nil || from == nil {
			return
		}

		t1.verifyPong(h, from, fromID, fromKey)
	})
}

// skipping Fuzz_Nosy_UDPv4_wrapPacket__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v4wire.Packet

func Fuzz_Nosy_UDPv4_write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv4
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var toaddr *net.UDPAddr
		fill_err = tp.Fill(&toaddr)
		if fill_err != nil {
			return
		}
		var toid enode.ID
		fill_err = tp.Fill(&toid)
		if fill_err != nil {
			return
		}
		var what string
		fill_err = tp.Fill(&what)
		if fill_err != nil {
			return
		}
		var packet []byte
		fill_err = tp.Fill(&packet)
		if fill_err != nil {
			return
		}
		if t1 == nil || toaddr == nil {
			return
		}

		t1.write(toaddr, toid, what, packet)
	})
}

func Fuzz_Nosy_UDPv5_AllNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.AllNodes()
	})
}

func Fuzz_Nosy_UDPv5_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Close()
	})
}

func Fuzz_Nosy_UDPv5_LocalNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.LocalNode()
	})
}

func Fuzz_Nosy_UDPv5_Lookup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var t2 enode.ID
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Lookup(t2)
	})
}

func Fuzz_Nosy_UDPv5_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.Ping(n)
	})
}

func Fuzz_Nosy_UDPv5_RandomNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RandomNodes()
	})
}

// skipping Fuzz_Nosy_UDPv5_RegisterTalkHandler__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.TalkRequestHandler

func Fuzz_Nosy_UDPv5_RequestENR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.RequestENR(n)
	})
}

func Fuzz_Nosy_UDPv5_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.Resolve(n)
	})
}

func Fuzz_Nosy_UDPv5_Self__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.Self()
	})
}

func Fuzz_Nosy_UDPv5_TalkRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var request []byte
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.TalkRequest(n, protocol, request)
	})
}

func Fuzz_Nosy_UDPv5_TalkRequestToID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr *net.UDPAddr
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var protocol string
		fill_err = tp.Fill(&protocol)
		if fill_err != nil {
			return
		}
		var request []byte
		fill_err = tp.Fill(&request)
		if fill_err != nil {
			return
		}
		if t1 == nil || addr == nil {
			return
		}

		t1.TalkRequestToID(id, addr, protocol, request)
	})
}

func Fuzz_Nosy_UDPv5_callDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var c *callV5
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if t1 == nil || c == nil {
			return
		}

		t1.callDone(c)
	})
}

// skipping Fuzz_Nosy_UDPv5_callToID__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

// skipping Fuzz_Nosy_UDPv5_callToNode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

func Fuzz_Nosy_UDPv5_collectTableNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var rip net.IP
		fill_err = tp.Fill(&rip)
		if fill_err != nil {
			return
		}
		var distances []uint
		fill_err = tp.Fill(&distances)
		if fill_err != nil {
			return
		}
		var limit int
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.collectTableNodes(rip, distances, limit)
	})
}

func Fuzz_Nosy_UDPv5_dispatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.dispatch()
	})
}

func Fuzz_Nosy_UDPv5_dispatchReadPacket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var from *net.UDPAddr
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var content []byte
		fill_err = tp.Fill(&content)
		if fill_err != nil {
			return
		}
		if t1 == nil || from == nil {
			return
		}

		t1.dispatchReadPacket(from, content)
	})
}

func Fuzz_Nosy_UDPv5_findnode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var distances []uint
		fill_err = tp.Fill(&distances)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.findnode(n, distances)
	})
}

func Fuzz_Nosy_UDPv5_getNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.getNode(id)
	})
}

// skipping Fuzz_Nosy_UDPv5_handle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

// skipping Fuzz_Nosy_UDPv5_handleCallResponse__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

func Fuzz_Nosy_UDPv5_handleFindnode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p *v5wire.Findnode
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromAddr *net.UDPAddr
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || p == nil || fromAddr == nil {
			return
		}

		t1.handleFindnode(p, fromID, fromAddr)
	})
}

func Fuzz_Nosy_UDPv5_handlePacket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var rawpacket []byte
		fill_err = tp.Fill(&rawpacket)
		if fill_err != nil {
			return
		}
		var fromAddr *net.UDPAddr
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || fromAddr == nil {
			return
		}

		t1.handlePacket(rawpacket, fromAddr)
	})
}

func Fuzz_Nosy_UDPv5_handlePing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p *v5wire.Ping
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromAddr *net.UDPAddr
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || p == nil || fromAddr == nil {
			return
		}

		t1.handlePing(p, fromID, fromAddr)
	})
}

func Fuzz_Nosy_UDPv5_handleUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p *v5wire.Unknown
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromAddr *net.UDPAddr
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || p == nil || fromAddr == nil {
			return
		}

		t1.handleUnknown(p, fromID, fromAddr)
	})
}

func Fuzz_Nosy_UDPv5_handleWhoareyou__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var p *v5wire.Whoareyou
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var fromAddr *net.UDPAddr
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || p == nil || fromAddr == nil {
			return
		}

		t1.handleWhoareyou(p, fromID, fromAddr)
	})
}

// skipping Fuzz_Nosy_UDPv5_initCall__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

func Fuzz_Nosy_UDPv5_lookupRandom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.lookupRandom()
	})
}

func Fuzz_Nosy_UDPv5_lookupSelf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.lookupSelf()
	})
}

func Fuzz_Nosy_UDPv5_lookupWorker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var destNode *node
		fill_err = tp.Fill(&destNode)
		if fill_err != nil {
			return
		}
		var t3 enode.ID
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if t1 == nil || destNode == nil {
			return
		}

		t1.lookupWorker(destNode, t3)
	})
}

func Fuzz_Nosy_UDPv5_matchWithCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var fromID enode.ID
		fill_err = tp.Fill(&fromID)
		if fill_err != nil {
			return
		}
		var nonce v5wire.Nonce
		fill_err = tp.Fill(&nonce)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.matchWithCall(fromID, nonce)
	})
}

func Fuzz_Nosy_UDPv5_newLookup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var t3 enode.ID
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.newLookup(ctx, t3)
	})
}

func Fuzz_Nosy_UDPv5_newRandomLookup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.newRandomLookup(ctx)
	})
}

func Fuzz_Nosy_UDPv5_ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || n == nil {
			return
		}

		t1.ping(n)
	})
}

func Fuzz_Nosy_UDPv5_readLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.readLoop()
	})
}

// skipping Fuzz_Nosy_UDPv5_send__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

func Fuzz_Nosy_UDPv5_sendCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var c *callV5
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if t1 == nil || c == nil {
			return
		}

		t1.sendCall(c)
	})
}

// skipping Fuzz_Nosy_UDPv5_sendFromAnotherThread__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

func Fuzz_Nosy_UDPv5_sendNextCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.sendNextCall(id)
	})
}

// skipping Fuzz_Nosy_UDPv5_sendResponse__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover/v5wire.Packet

func Fuzz_Nosy_UDPv5_startResponseTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var c *callV5
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if t1 == nil || c == nil {
			return
		}

		t1.startResponseTimeout(c)
	})
}

func Fuzz_Nosy_UDPv5_verifyResponseNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var c *callV5
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var r *enr.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var distances []uint
		fill_err = tp.Fill(&distances)
		if fill_err != nil {
			return
		}
		var seen map[enode.ID]struct{}
		fill_err = tp.Fill(&seen)
		if fill_err != nil {
			return
		}
		if t1 == nil || c == nil || r == nil {
			return
		}

		t1.verifyResponseNode(c, r, distances, seen)
	})
}

func Fuzz_Nosy_UDPv5_waitForNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *UDPv5
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var c *callV5
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var distances []uint
		fill_err = tp.Fill(&distances)
		if fill_err != nil {
			return
		}
		if t1 == nil || c == nil {
			return
		}

		t1.waitForNodes(c, distances)
	})
}

func Fuzz_Nosy_lookup_advance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *lookup
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.advance()
	})
}

// skipping Fuzz_Nosy_lookup_query__ because parameters include func, chan, or unsupported interface: chan<- []*github.com/ethereum/go-ethereum/p2p/discover.node

func Fuzz_Nosy_lookup_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *lookup
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.run()
	})
}

func Fuzz_Nosy_lookup_shutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *lookup
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.shutdown()
	})
}

func Fuzz_Nosy_lookup_slowdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *lookup
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.slowdown()
	})
}

func Fuzz_Nosy_lookup_startQueries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *lookup
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.startQueries()
	})
}

func Fuzz_Nosy_lookupIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *lookupIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Close()
	})
}

func Fuzz_Nosy_lookupIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *lookupIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Next()
	})
}

func Fuzz_Nosy_lookupIterator_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *lookupIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.Node()
	})
}

func Fuzz_Nosy_meteredUdpConn_ReadFromUDP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *meteredUdpConn
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

		c.ReadFromUDP(b)
	})
}

func Fuzz_Nosy_meteredUdpConn_WriteToUDP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *meteredUdpConn
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var addr *net.UDPAddr
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if c == nil || addr == nil {
			return
		}

		c.WriteToUDP(b, addr)
	})
}

func Fuzz_Nosy_node_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n1 := wrapNode(n)
		n1.String()
	})
}

func Fuzz_Nosy_node_addr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n1 := wrapNode(n)
		n1.addr()
	})
}

func Fuzz_Nosy_nodesByDistance_push__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *nodesByDistance
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var n *node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var maxElems int
		fill_err = tp.Fill(&maxElems)
		if fill_err != nil {
			return
		}
		if h == nil || n == nil {
			return
		}

		h.push(n, maxElems)
	})
}

func Fuzz_Nosy_talkSystem_handleRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var transport *UDPv5
		fill_err = tp.Fill(&transport)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var addr *net.UDPAddr
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var req *v5wire.TalkRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if transport == nil || addr == nil || req == nil {
			return
		}

		t1 := newTalkSystem(transport)
		t1.handleRequest(id, addr, req)
	})
}

// skipping Fuzz_Nosy_talkSystem_register__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.TalkRequestHandler

func Fuzz_Nosy_talkSystem_wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var transport *UDPv5
		fill_err = tp.Fill(&transport)
		if fill_err != nil {
			return
		}
		if transport == nil {
			return
		}

		t1 := newTalkSystem(transport)
		t1.wait()
	})
}

func Fuzz_Nosy_Config_withDefaults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		cfg.withDefaults()
	})
}

// skipping Fuzz_Nosy_UDPConn_Close__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.UDPConn

// skipping Fuzz_Nosy_UDPConn_LocalAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.UDPConn

// skipping Fuzz_Nosy_UDPConn_ReadFromUDP__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.UDPConn

// skipping Fuzz_Nosy_UDPConn_WriteToUDP__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.UDPConn

// skipping Fuzz_Nosy_codecV5_Decode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.codecV5

// skipping Fuzz_Nosy_codecV5_Encode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.codecV5

func Fuzz_Nosy_encPubkey_id__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PublicKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		e := encodePubkey(key)
		e.id()
	})
}

// skipping Fuzz_Nosy_transport_RequestENR__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.transport

// skipping Fuzz_Nosy_transport_Self__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.transport

// skipping Fuzz_Nosy_transport_lookupRandom__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.transport

// skipping Fuzz_Nosy_transport_lookupSelf__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.transport

// skipping Fuzz_Nosy_transport_ping__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/discover.transport

func Fuzz_Nosy_contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns []*node
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		contains(ns, id)
	})
}

func Fuzz_Nosy_containsUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x uint
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var xs []uint
		fill_err = tp.Fill(&xs)
		if fill_err != nil {
			return
		}

		containsUint(x, xs)
	})
}

func Fuzz_Nosy_deleteNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []*node
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var n *node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		deleteNode(list, n)
	})
}

func Fuzz_Nosy_lookupDistances__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 enode.ID
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var dest enode.ID
		fill_err = tp.Fill(&dest)
		if fill_err != nil {
			return
		}

		lookupDistances(t1, dest)
	})
}

func Fuzz_Nosy_min__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x int, y int) {
		min(x, y)
	})
}

func Fuzz_Nosy_packNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var reqid []byte
		fill_err = tp.Fill(&reqid)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		packNodes(reqid, nodes)
	})
}

func Fuzz_Nosy_pushNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []*node
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var n *node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var max int
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		pushNode(list, n, max)
	})
}

func Fuzz_Nosy_unwrapNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns []*node
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}

		unwrapNodes(ns)
	})
}

func Fuzz_Nosy_wrapNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns []*enode.Node
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}

		wrapNodes(ns)
	})
}
