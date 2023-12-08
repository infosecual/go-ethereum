package enode

import (
	"crypto/ecdsa"
	"io"
	"net"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/p2p/netutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
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

func Fuzz_Nosy_DB_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.Close()
	})
}

func Fuzz_Nosy_DB_DeleteNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.DeleteNode(id)
	})
}

func Fuzz_Nosy_DB_FindFails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.FindFails(id, ip)
	})
}

func Fuzz_Nosy_DB_FindFailsV5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.FindFailsV5(id, ip)
	})
}

func Fuzz_Nosy_DB_LastPingReceived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.LastPingReceived(id, ip)
	})
}

func Fuzz_Nosy_DB_LastPongReceived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.LastPongReceived(id, ip)
	})
}

func Fuzz_Nosy_DB_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.Node(id)
	})
}

func Fuzz_Nosy_DB_NodeSeq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.NodeSeq(id)
	})
}

func Fuzz_Nosy_DB_QuerySeeds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var maxAge time.Duration
		fill_err = tp.Fill(&maxAge)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.QuerySeeds(n, maxAge)
	})
}

func Fuzz_Nosy_DB_Resolve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
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

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.Resolve(n)
	})
}

func Fuzz_Nosy_DB_UpdateFindFails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var fails int
		fill_err = tp.Fill(&fails)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.UpdateFindFails(id, ip, fails)
	})
}

func Fuzz_Nosy_DB_UpdateFindFailsV5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var fails int
		fill_err = tp.Fill(&fails)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.UpdateFindFailsV5(id, ip, fails)
	})
}

func Fuzz_Nosy_DB_UpdateLastPingReceived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var instance time.Time
		fill_err = tp.Fill(&instance)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.UpdateLastPingReceived(id, ip, instance)
	})
}

func Fuzz_Nosy_DB_UpdateLastPongReceived__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var instance time.Time
		fill_err = tp.Fill(&instance)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.UpdateLastPongReceived(id, ip, instance)
	})
}

func Fuzz_Nosy_DB_UpdateNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var node *Node
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if node == nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.UpdateNode(node)
	})
}

func Fuzz_Nosy_DB_ensureExpirer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.ensureExpirer()
	})
}

func Fuzz_Nosy_DB_expireNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.expireNodes()
	})
}

func Fuzz_Nosy_DB_expirer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.expirer()
	})
}

func Fuzz_Nosy_DB_fetchInt64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, key []byte) {
		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.fetchInt64(key)
	})
}

func Fuzz_Nosy_DB_fetchUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, key []byte) {
		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.fetchUint64(key)
	})
}

func Fuzz_Nosy_DB_localSeq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.localSeq(id)
	})
}

func Fuzz_Nosy_DB_storeInt64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, key []byte, n int64) {
		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.storeInt64(key, n)
	})
}

func Fuzz_Nosy_DB_storeLocalSeq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}

		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.storeLocalSeq(id, n)
	})
}

func Fuzz_Nosy_DB_storeUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string, key []byte, n uint64) {
		db, err := newPersistentDB(path)
		if err != nil {
			return
		}
		db.storeUint64(key, n)
	})
}

// skipping Fuzz_Nosy_FairMix_AddSource__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enode.Iterator

func Fuzz_Nosy_FairMix_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		m := NewFairMix(timeout)
		m.Close()
	})
}

func Fuzz_Nosy_FairMix_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		m := NewFairMix(timeout)
		m.Next()
	})
}

func Fuzz_Nosy_FairMix_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		m := NewFairMix(timeout)
		m.Node()
	})
}

func Fuzz_Nosy_FairMix_deleteSource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		var s *mixSource
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		m := NewFairMix(timeout)
		m.deleteSource(s)
	})
}

func Fuzz_Nosy_FairMix_nextFromAny__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		m := NewFairMix(timeout)
		m.nextFromAny()
	})
}

func Fuzz_Nosy_FairMix_pickSource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}

		m := NewFairMix(timeout)
		m.pickSource()
	})
}

// skipping Fuzz_Nosy_FairMix_runSource__ because parameters include func, chan, or unsupported interface: chan struct{}

func Fuzz_Nosy_ID_UnmarshalText__(f *testing.F) {
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
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		n := PubkeyToIDV4(key)
		n.UnmarshalText(text)
	})
}

func Fuzz_Nosy_LocalNode_Database__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.Database()
	})
}

// skipping Fuzz_Nosy_LocalNode_Delete__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.Entry

func Fuzz_Nosy_LocalNode_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.ID()
	})
}

func Fuzz_Nosy_LocalNode_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.Node()
	})
}

func Fuzz_Nosy_LocalNode_Seq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.Seq()
	})
}

// skipping Fuzz_Nosy_LocalNode_Set__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.Entry

func Fuzz_Nosy_LocalNode_SetFallbackIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.SetFallbackIP(ip)
	})
}

func Fuzz_Nosy_LocalNode_SetFallbackUDP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var port int
		fill_err = tp.Fill(&port)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.SetFallbackUDP(port)
	})
}

func Fuzz_Nosy_LocalNode_SetStaticIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.SetStaticIP(ip)
	})
}

func Fuzz_Nosy_LocalNode_UDPContact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var toaddr *net.UDPAddr
		fill_err = tp.Fill(&toaddr)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil || toaddr == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.UDPContact(toaddr)
	})
}

func Fuzz_Nosy_LocalNode_UDPEndpointStatement__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var fromaddr *net.UDPAddr
		fill_err = tp.Fill(&fromaddr)
		if fill_err != nil {
			return
		}
		var endpoint *net.UDPAddr
		fill_err = tp.Fill(&endpoint)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil || fromaddr == nil || endpoint == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.UDPEndpointStatement(fromaddr, endpoint)
	})
}

func Fuzz_Nosy_LocalNode_bumpSeq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.bumpSeq()
	})
}

// skipping Fuzz_Nosy_LocalNode_delete__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.Entry

func Fuzz_Nosy_LocalNode_endpointForIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.endpointForIP(ip)
	})
}

func Fuzz_Nosy_LocalNode_invalidate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.invalidate()
	})
}

// skipping Fuzz_Nosy_LocalNode_set__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.Entry

func Fuzz_Nosy_LocalNode_sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.sign()
	})
}

func Fuzz_Nosy_LocalNode_updateEndpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil || key == nil {
			return
		}

		ln := NewLocalNode(db, key)
		ln.updateEndpoints()
	})
}

func Fuzz_Nosy_Node_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.ID()
	})
}

func Fuzz_Nosy_Node_IP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.IP()
	})
}

func Fuzz_Nosy_Node_Incomplete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.Incomplete()
	})
}

// skipping Fuzz_Nosy_Node_Load__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enr.Entry

func Fuzz_Nosy_Node_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.MarshalText()
	})
}

func Fuzz_Nosy_Node_Pubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.Pubkey()
	})
}

func Fuzz_Nosy_Node_Record__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.Record()
	})
}

func Fuzz_Nosy_Node_Seq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.Seq()
	})
}

func Fuzz_Nosy_Node_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.String()
	})
}

func Fuzz_Nosy_Node_TCP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.TCP()
	})
}

func Fuzz_Nosy_Node_UDP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.UDP()
	})
}

func Fuzz_Nosy_Node_URLv4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.URLv4()
	})
}

func Fuzz_Nosy_Node_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte, text []byte) {
		n := mustDecodeNode(id, d2)
		n.UnmarshalText(text)
	})
}

func Fuzz_Nosy_Node_ValidateComplete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id []byte, d2 []byte) {
		n := mustDecodeNode(id, d2)
		n.ValidateComplete()
	})
}

func Fuzz_Nosy_Secp256k1_DecodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Secp256k1
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var s *rlp.Stream
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if v == nil || s == nil {
			return
		}

		v.DecodeRLP(s)
	})
}

func Fuzz_Nosy_filterIter_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *filterIter
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Next()
	})
}

func Fuzz_Nosy_lnEndpoint_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *lnEndpoint
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.get()
	})
}

func Fuzz_Nosy_sliceIter_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *sliceIter
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

func Fuzz_Nosy_sliceIter_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *sliceIter
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

func Fuzz_Nosy_sliceIter_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *sliceIter
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

func Fuzz_Nosy_ID_Bytes__(f *testing.F) {
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

		n := PubkeyToIDV4(key)
		n.Bytes()
	})
}

func Fuzz_Nosy_ID_GoString__(f *testing.F) {
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

		n := PubkeyToIDV4(key)
		n.GoString()
	})
}

func Fuzz_Nosy_ID_MarshalText__(f *testing.F) {
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

		n := PubkeyToIDV4(key)
		n.MarshalText()
	})
}

func Fuzz_Nosy_ID_String__(f *testing.F) {
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

		n := PubkeyToIDV4(key)
		n.String()
	})
}

func Fuzz_Nosy_ID_TerminalString__(f *testing.F) {
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

		n := PubkeyToIDV4(key)
		n.TerminalString()
	})
}

func Fuzz_Nosy_Iterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nodes []*Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		_x1 := CycleNodes(nodes)
		_x1.Close()
	})
}

func Fuzz_Nosy_Iterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nodes []*Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		_x1 := CycleNodes(nodes)
		_x1.Next()
	})
}

func Fuzz_Nosy_Iterator_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nodes []*Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		_x1 := CycleNodes(nodes)
		_x1.Node()
	})
}

func Fuzz_Nosy_NullID_NodeAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NullID
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var r *enr.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		_x1.NodeAddr(r)
	})
}

func Fuzz_Nosy_NullID_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 NullID
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var r *enr.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		_x1.Verify(r, sig)
	})
}

func Fuzz_Nosy_Secp256k1_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v Secp256k1
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ENRKey()
	})
}

func Fuzz_Nosy_Secp256k1_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v Secp256k1
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var w io.Writer
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		v.EncodeRLP(w)
	})
}

func Fuzz_Nosy_V4ID_NodeAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 V4ID
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var r *enr.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		_x1.NodeAddr(r)
	})
}

func Fuzz_Nosy_V4ID_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 V4ID
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var r *enr.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		_x1.Verify(r, sig)
	})
}

func Fuzz_Nosy_s256raw_ENRKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 s256raw
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.ENRKey()
	})
}

func Fuzz_Nosy_v4CompatID_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 v4CompatID
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var r *enr.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var sig []byte
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		_x1.Verify(r, sig)
	})
}

func Fuzz_Nosy_DistCmp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 ID
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var a ID
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b ID
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		DistCmp(t1, a, b)
	})
}

func Fuzz_Nosy_LogDist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a ID
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var b ID
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		LogDist(a, b)
	})
}

// skipping Fuzz_Nosy_ReadNodes__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/enode.Iterator

func Fuzz_Nosy_deleteRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *leveldb.DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var prefix []byte
		fill_err = tp.Fill(&prefix)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		deleteRange(db, prefix)
	})
}

func Fuzz_Nosy_isNewV4__(f *testing.F) {
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

		isNewV4(n)
	})
}

func Fuzz_Nosy_localItemKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var field string
		fill_err = tp.Fill(&field)
		if fill_err != nil {
			return
		}

		localItemKey(id, field)
	})
}

func Fuzz_Nosy_nodeItemKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var field string
		fill_err = tp.Fill(&field)
		if fill_err != nil {
			return
		}

		nodeItemKey(id, ip, field)
	})
}

func Fuzz_Nosy_nodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		nodeKey(id)
	})
}

func Fuzz_Nosy_predictAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *netutil.IPTracker
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		predictAddr(t1)
	})
}

func Fuzz_Nosy_signV4Compat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *enr.Record
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var pubkey *ecdsa.PublicKey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if r == nil || pubkey == nil {
			return
		}

		signV4Compat(r, pubkey)
	})
}

func Fuzz_Nosy_splitNodeItemKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		splitNodeItemKey(key)
	})
}

func Fuzz_Nosy_splitNodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		splitNodeKey(key)
	})
}

func Fuzz_Nosy_v5Key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var ip net.IP
		fill_err = tp.Fill(&ip)
		if fill_err != nil {
			return
		}
		var field string
		fill_err = tp.Fill(&field)
		if fill_err != nil {
			return
		}

		v5Key(id, ip, field)
	})
}
