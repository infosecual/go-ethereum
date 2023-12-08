package main

import (
	"io"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/ethereum/go-ethereum/p2p/dnsdisc"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/rlp"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"github.com/urfave/cli/v2"
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

func Fuzz_Nosy_cloudflareClient_checkZone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := newCloudflareClient(ctx)
		c.checkZone(name)
	})
}

func Fuzz_Nosy_cloudflareClient_deploy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var t3 *dnsdisc.Tree
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if ctx == nil || t3 == nil {
			return
		}

		c := newCloudflareClient(ctx)
		c.deploy(name, t3)
	})
}

func Fuzz_Nosy_cloudflareClient_uploadRecords__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var records map[string]string
		fill_err = tp.Fill(&records)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := newCloudflareClient(ctx)
		c.uploadRecords(name, records)
	})
}

func Fuzz_Nosy_crawler_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *crawler
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		var nthreads int
		fill_err = tp.Fill(&nthreads)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.run(timeout, nthreads)
	})
}

// skipping Fuzz_Nosy_crawler_runIterator__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/p2p/enode.Iterator

func Fuzz_Nosy_crawler_updateNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *crawler
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if c == nil || n == nil {
			return
		}

		c.updateNode(n)
	})
}

func Fuzz_Nosy_route53Client_checkZone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := newRoute53Client(ctx)
		c.checkZone(name)
	})
}

func Fuzz_Nosy_route53Client_collectRecords__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := newRoute53Client(ctx)
		c.collectRecords(name)
	})
}

func Fuzz_Nosy_route53Client_computeChanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var records map[string]string
		fill_err = tp.Fill(&records)
		if fill_err != nil {
			return
		}
		var existing map[string]recordSet
		fill_err = tp.Fill(&existing)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := newRoute53Client(ctx)
		c.computeChanges(name, records, existing)
	})
}

func Fuzz_Nosy_route53Client_deleteDomain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := newRoute53Client(ctx)
		c.deleteDomain(name)
	})
}

func Fuzz_Nosy_route53Client_deploy__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var t3 *dnsdisc.Tree
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		if ctx == nil || t3 == nil {
			return
		}

		c := newRoute53Client(ctx)
		c.deploy(name, t3)
	})
}

func Fuzz_Nosy_route53Client_findZoneID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := newRoute53Client(ctx)
		c.findZoneID(name)
	})
}

func Fuzz_Nosy_route53Client_submitChanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var changes []types.Change
		fill_err = tp.Fill(&changes)
		if fill_err != nil {
			return
		}
		var comment string
		fill_err = tp.Fill(&comment)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		c := newRoute53Client(ctx)
		c.submitChanges(changes, comment)
	})
}

func Fuzz_Nosy_nodeSet_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var nodes []*enode.Node
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		ns := loadNodesJSON(file)
		ns.add(nodes...)
	})
}

func Fuzz_Nosy_nodeSet_nodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		ns := loadNodesJSON(file)
		ns.nodes()
	})
}

func Fuzz_Nosy_nodeSet_topN__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string, n int) {
		ns := loadNodesJSON(file)
		ns.topN(n)
	})
}

func Fuzz_Nosy_nodeSet_verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, file string) {
		ns := loadNodesJSON(file)
		ns.verify()
	})
}

// skipping Fuzz_Nosy_resolver_RequestENR__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/cmd/devp2p.resolver

func Fuzz_Nosy_changeCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch types.Change
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		changeCount(ch)
	})
}

func Fuzz_Nosy_changeSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ch types.Change
		fill_err = tp.Fill(&ch)
		if fill_err != nil {
			return
		}

		changeSize(ch)
	})
}

// skipping Fuzz_Nosy_commandHasFlag__ because parameters include func, chan, or unsupported interface: github.com/urfave/cli/v2.Flag

func Fuzz_Nosy_decodeRecordBase64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		decodeRecordBase64(b)
	})
}

func Fuzz_Nosy_decodeRecordHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		decodeRecordHex(b)
	})
}

func Fuzz_Nosy_directoryName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string) {
		directoryName(dir)
	})
}

func Fuzz_Nosy_dumpNodeURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
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

		dumpNodeURL(out, n)
	})
}

func Fuzz_Nosy_dumpRecord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
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

		dumpRecord(out, r)
	})
}

// skipping Fuzz_Nosy_dumpRecordKV__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_exit__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_formatAttrIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v rlp.RawValue
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		formatAttrIP(v)
	})
}

func Fuzz_Nosy_formatAttrRaw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v rlp.RawValue
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		formatAttrRaw(v)
	})
}

func Fuzz_Nosy_formatAttrString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v rlp.RawValue
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		formatAttrString(v)
	})
}

func Fuzz_Nosy_formatAttrUint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v rlp.RawValue
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		formatAttrUint(v)
	})
}

func Fuzz_Nosy_isSubdomain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string, domain string) {
		isSubdomain(name, domain)
	})
}

func Fuzz_Nosy_loadTreeDefinitionForExport__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string) {
		loadTreeDefinitionForExport(dir)
	})
}

func Fuzz_Nosy_makeDeletionChanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var records map[string]recordSet
		fill_err = tp.Fill(&records)
		if fill_err != nil {
			return
		}
		var keep map[string]string
		fill_err = tp.Fill(&keep)
		if fill_err != nil {
			return
		}

		makeDeletionChanges(records, keep)
	})
}

func Fuzz_Nosy_makeDiscoveryConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		makeDiscoveryConfig(ctx)
	})
}

func Fuzz_Nosy_parseBootnodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		parseBootnodes(ctx)
	})
}

func Fuzz_Nosy_parseExtAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, spec string) {
		parseExtAddr(spec)
	})
}

func Fuzz_Nosy_parseFilterLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		parseFilterLimit(args)
	})
}

func Fuzz_Nosy_parseFilters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var args []string
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		parseFilters(args)
	})
}

func Fuzz_Nosy_showAttributeCounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns nodeSet
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}

		showAttributeCounts(ns)
	})
}

func Fuzz_Nosy_sortChanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var changes []types.Change
		fill_err = tp.Fill(&changes)
		if fill_err != nil {
			return
		}

		sortChanges(changes)
	})
}

func Fuzz_Nosy_splitChanges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var changes []types.Change
		fill_err = tp.Fill(&changes)
		if fill_err != nil {
			return
		}
		var sizeLimit int
		fill_err = tp.Fill(&sizeLimit)
		if fill_err != nil {
			return
		}
		var countLimit int
		fill_err = tp.Fill(&countLimit)
		if fill_err != nil {
			return
		}

		splitChanges(changes, sizeLimit, countLimit)
	})
}

func Fuzz_Nosy_splitTXT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, value string) {
		splitTXT(value)
	})
}

func Fuzz_Nosy_startV4__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		startV4(ctx)
	})
}

func Fuzz_Nosy_startV5__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx *cli.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctx == nil {
			return
		}

		startV5(ctx)
	})
}

func Fuzz_Nosy_treeDefinitionFiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, directory string) {
		treeDefinitionFiles(directory)
	})
}

func Fuzz_Nosy_writeNodesJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var nodes nodeSet
		fill_err = tp.Fill(&nodes)
		if fill_err != nil {
			return
		}

		writeNodesJSON(file, nodes)
	})
}

func Fuzz_Nosy_writeTXTJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var file string
		fill_err = tp.Fill(&file)
		if fill_err != nil {
			return
		}
		var txt map[string]string
		fill_err = tp.Fill(&txt)
		if fill_err != nil {
			return
		}

		writeTXTJSON(file, txt)
	})
}

func Fuzz_Nosy_writeTreeMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var directory string
		fill_err = tp.Fill(&directory)
		if fill_err != nil {
			return
		}
		var def *dnsDefinition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}
		if def == nil {
			return
		}

		writeTreeMetadata(directory, def)
	})
}

func Fuzz_Nosy_writeTreeNodes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var directory string
		fill_err = tp.Fill(&directory)
		if fill_err != nil {
			return
		}
		var def *dnsDefinition
		fill_err = tp.Fill(&def)
		if fill_err != nil {
			return
		}
		if def == nil {
			return
		}

		writeTreeNodes(directory, def)
	})
}
