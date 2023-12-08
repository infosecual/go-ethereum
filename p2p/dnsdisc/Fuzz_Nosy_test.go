package dnsdisc

import (
	"context"
	"crypto/ecdsa"
	"testing"

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

func Fuzz_Nosy_Client_NewIterator__(f *testing.F) {
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
		var urls []string
		fill_err = tp.Fill(&urls)
		if fill_err != nil {
			return
		}

		c := NewClient(cfg)
		c.NewIterator(urls...)
	})
}

func Fuzz_Nosy_Client_SyncTree__(f *testing.F) {
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
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}

		c := NewClient(cfg)
		c.SyncTree(url)
	})
}

func Fuzz_Nosy_Client_doResolveEntry__(f *testing.F) {
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
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var domain string
		fill_err = tp.Fill(&domain)
		if fill_err != nil {
			return
		}
		var hash string
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		c := NewClient(cfg)
		c.doResolveEntry(ctx, domain, hash)
	})
}

func Fuzz_Nosy_Client_newRandomIterator__(f *testing.F) {
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

		c := NewClient(cfg)
		c.newRandomIterator()
	})
}

func Fuzz_Nosy_Client_resolveEntry__(f *testing.F) {
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
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var domain string
		fill_err = tp.Fill(&domain)
		if fill_err != nil {
			return
		}
		var hash string
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		c := NewClient(cfg)
		c.resolveEntry(ctx, domain, hash)
	})
}

func Fuzz_Nosy_Client_resolveRoot__(f *testing.F) {
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
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		if loc == nil {
			return
		}

		c := NewClient(cfg)
		c.resolveRoot(ctx, loc)
	})
}

func Fuzz_Nosy_Tree_Links__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var seq uint
		fill_err = tp.Fill(&seq)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var links []string
		fill_err = tp.Fill(&links)
		if fill_err != nil {
			return
		}

		t1, err := MakeTree(seq, nodes, links)
		if err != nil {
			return
		}
		t1.Links()
	})
}

func Fuzz_Nosy_Tree_Nodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var seq uint
		fill_err = tp.Fill(&seq)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var links []string
		fill_err = tp.Fill(&links)
		if fill_err != nil {
			return
		}

		t1, err := MakeTree(seq, nodes, links)
		if err != nil {
			return
		}
		t1.Nodes()
	})
}

func Fuzz_Nosy_Tree_Seq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var seq uint
		fill_err = tp.Fill(&seq)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var links []string
		fill_err = tp.Fill(&links)
		if fill_err != nil {
			return
		}

		t1, err := MakeTree(seq, nodes, links)
		if err != nil {
			return
		}
		t1.Seq()
	})
}

func Fuzz_Nosy_Tree_SetSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var seq uint
		fill_err = tp.Fill(&seq)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var links []string
		fill_err = tp.Fill(&links)
		if fill_err != nil {
			return
		}
		var pubkey *ecdsa.PublicKey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var signature string
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if pubkey == nil {
			return
		}

		t1, err := MakeTree(seq, nodes, links)
		if err != nil {
			return
		}
		t1.SetSignature(pubkey, signature)
	})
}

func Fuzz_Nosy_Tree_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var seq uint
		fill_err = tp.Fill(&seq)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var links []string
		fill_err = tp.Fill(&links)
		if fill_err != nil {
			return
		}
		var key *ecdsa.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var domain string
		fill_err = tp.Fill(&domain)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		t1, err := MakeTree(seq, nodes, links)
		if err != nil {
			return
		}
		t1.Sign(key, domain)
	})
}

func Fuzz_Nosy_Tree_Signature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var seq uint
		fill_err = tp.Fill(&seq)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var links []string
		fill_err = tp.Fill(&links)
		if fill_err != nil {
			return
		}

		t1, err := MakeTree(seq, nodes, links)
		if err != nil {
			return
		}
		t1.Signature()
	})
}

func Fuzz_Nosy_Tree_ToTXT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var seq uint
		fill_err = tp.Fill(&seq)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}
		var links []string
		fill_err = tp.Fill(&links)
		if fill_err != nil {
			return
		}
		var domain string
		fill_err = tp.Fill(&domain)
		if fill_err != nil {
			return
		}

		t1, err := MakeTree(seq, nodes, links)
		if err != nil {
			return
		}
		t1.ToTXT(domain)
	})
}

// skipping Fuzz_Nosy_Tree_build__ because parameters include func, chan, or unsupported interface: []github.com/ethereum/go-ethereum/p2p/dnsdisc.entry

func Fuzz_Nosy_branchEntry_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *branchEntry
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

func Fuzz_Nosy_clientTree_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.String()
	})
}

func Fuzz_Nosy_clientTree_canSyncRandom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.canSyncRandom()
	})
}

func Fuzz_Nosy_clientTree_gcLinks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.gcLinks()
	})
}

func Fuzz_Nosy_clientTree_nextScheduledRootCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.nextScheduledRootCheck()
	})
}

func Fuzz_Nosy_clientTree_rootUpdateDue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.rootUpdateDue()
	})
}

func Fuzz_Nosy_clientTree_slowdownRootUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.slowdownRootUpdate(ctx)
	})
}

// skipping Fuzz_Nosy_clientTree_syncAll__ because parameters include func, chan, or unsupported interface: map[string]github.com/ethereum/go-ethereum/p2p/dnsdisc.entry

func Fuzz_Nosy_clientTree_syncNextLink__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.syncNextLink(ctx)
	})
}

func Fuzz_Nosy_clientTree_syncNextRandomENR__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.syncNextRandomENR(ctx)
	})
}

func Fuzz_Nosy_clientTree_syncRandom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.syncRandom(ctx)
	})
}

func Fuzz_Nosy_clientTree_updateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if c == nil || lc == nil || loc == nil {
			return
		}

		ct := newClientTree(c, lc, loc)
		ct.updateRoot(ctx)
	})
}

func Fuzz_Nosy_enrEntry_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *enrEntry
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

func Fuzz_Nosy_linkCache_addLink__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to string
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if lc == nil {
			return
		}

		lc.addLink(from, to)
	})
}

func Fuzz_Nosy_linkCache_isReferenced__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var r string
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if lc == nil {
			return
		}

		lc.isReferenced(r)
	})
}

func Fuzz_Nosy_linkCache_resetLinks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lc *linkCache
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var keep map[string]struct{}
		fill_err = tp.Fill(&keep)
		if fill_err != nil {
			return
		}
		if lc == nil {
			return
		}

		lc.resetLinks(from, keep)
	})
}

func Fuzz_Nosy_linkEntry_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, e string) {
		e1, err := parseLink(e)
		if err != nil {
			return
		}
		e1.String()
	})
}

func Fuzz_Nosy_randomIterator_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
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

func Fuzz_Nosy_randomIterator_Next__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
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

func Fuzz_Nosy_randomIterator_Node__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
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

func Fuzz_Nosy_randomIterator_addTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.addTree(url)
	})
}

func Fuzz_Nosy_randomIterator_nextNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.nextNode()
	})
}

func Fuzz_Nosy_randomIterator_pickTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.pickTree()
	})
}

func Fuzz_Nosy_randomIterator_rebuildTrees__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.rebuildTrees()
	})
}

func Fuzz_Nosy_randomIterator_syncableTrees__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.syncableTrees()
	})
}

func Fuzz_Nosy_randomIterator_waitForRootUpdates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var it *randomIterator
		fill_err = tp.Fill(&it)
		if fill_err != nil {
			return
		}
		var trees []*clientTree
		fill_err = tp.Fill(&trees)
		if fill_err != nil {
			return
		}
		if it == nil {
			return
		}

		it.waitForRootUpdates(trees)
	})
}

func Fuzz_Nosy_rootEntry_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txt string
		fill_err = tp.Fill(&txt)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		if loc == nil {
			return
		}

		e, err := parseAndVerifyRoot(txt, loc)
		if err != nil {
			return
		}
		e.String()
	})
}

func Fuzz_Nosy_rootEntry_sigHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txt string
		fill_err = tp.Fill(&txt)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		if loc == nil {
			return
		}

		e, err := parseAndVerifyRoot(txt, loc)
		if err != nil {
			return
		}
		e.sigHash()
	})
}

func Fuzz_Nosy_rootEntry_verifySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txt string
		fill_err = tp.Fill(&txt)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		var pubkey *ecdsa.PublicKey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		if loc == nil || pubkey == nil {
			return
		}

		e, err := parseAndVerifyRoot(txt, loc)
		if err != nil {
			return
		}
		e.verifySignature(pubkey)
	})
}

func Fuzz_Nosy_subtreeSync_done__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		var root string
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var link bool
		fill_err = tp.Fill(&link)
		if fill_err != nil {
			return
		}
		if c == nil || loc == nil {
			return
		}

		ts := newSubtreeSync(c, loc, root, link)
		ts.done()
	})
}

// skipping Fuzz_Nosy_subtreeSync_resolveAll__ because parameters include func, chan, or unsupported interface: map[string]github.com/ethereum/go-ethereum/p2p/dnsdisc.entry

func Fuzz_Nosy_subtreeSync_resolveNext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var loc *linkEntry
		fill_err = tp.Fill(&loc)
		if fill_err != nil {
			return
		}
		var root string
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var link bool
		fill_err = tp.Fill(&link)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash string
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if c == nil || loc == nil {
			return
		}

		ts := newSubtreeSync(c, loc, root, link)
		ts.resolveNext(ctx, hash)
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

// skipping Fuzz_Nosy_Resolver_LookupTXT__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/dnsdisc.Resolver

func Fuzz_Nosy_entryError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err entryError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_nameError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err nameError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_ParseURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, url string) {
		ParseURL(url)
	})
}

func Fuzz_Nosy_isValidHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		isValidHash(s)
	})
}

func Fuzz_Nosy_removeHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h []string
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		removeHash(h, index)
	})
}

func Fuzz_Nosy_sortByID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		sortByID(nodes)
	})
}

// skipping Fuzz_Nosy_subdomain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/dnsdisc.entry

func Fuzz_Nosy_truncateHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		truncateHash(hash)
	})
}
