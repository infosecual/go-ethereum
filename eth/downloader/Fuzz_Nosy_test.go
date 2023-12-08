package downloader

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/prque"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/event"
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

// skipping Fuzz_Nosy_Downloader_BeaconDevSync__ because parameters include func, chan, or unsupported interface: chan struct{}

func Fuzz_Nosy_Downloader_BeaconExtend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if d == nil || head == nil {
			return
		}

		d.BeaconExtend(mode, head)
	})
}

func Fuzz_Nosy_Downloader_BeaconSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var final *types.Header
		fill_err = tp.Fill(&final)
		if fill_err != nil {
			return
		}
		if d == nil || head == nil || final == nil {
			return
		}

		d.BeaconSync(mode, head, final)
	})
}

func Fuzz_Nosy_Downloader_Cancel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Cancel()
	})
}

// skipping Fuzz_Nosy_Downloader_DeliverSnapPacket__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/protocols/snap.Packet

func Fuzz_Nosy_Downloader_LegacySync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var head common.Hash
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var td *big.Int
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		var ttd *big.Int
		fill_err = tp.Fill(&ttd)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}
		if d == nil || td == nil || ttd == nil {
			return
		}

		d.LegacySync(id, head, td, ttd, mode)
	})
}

func Fuzz_Nosy_Downloader_Progress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Progress()
	})
}

// skipping Fuzz_Nosy_Downloader_RegisterPeer__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.Peer

// skipping Fuzz_Nosy_Downloader_SetBadBlockCallback__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.badBlockFn

func Fuzz_Nosy_Downloader_Terminate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Terminate()
	})
}

func Fuzz_Nosy_Downloader_UnregisterPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.UnregisterPeer(id)
	})
}

func Fuzz_Nosy_Downloader_beaconSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var final *types.Header
		fill_err = tp.Fill(&final)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if d == nil || head == nil || final == nil {
			return
		}

		d.beaconSync(mode, head, final, force)
	})
}

func Fuzz_Nosy_Downloader_cancel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.cancel()
	})
}

func Fuzz_Nosy_Downloader_commitPivotBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var result *fetchResult
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if d == nil || result == nil {
			return
		}

		d.commitPivotBlock(result)
	})
}

func Fuzz_Nosy_Downloader_commitSnapSyncData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var results []*fetchResult
		fill_err = tp.Fill(&results)
		if fill_err != nil {
			return
		}
		var stateSync *stateSync
		fill_err = tp.Fill(&stateSync)
		if fill_err != nil {
			return
		}
		if d == nil || stateSync == nil {
			return
		}

		d.commitSnapSyncData(results, stateSync)
	})
}

// skipping Fuzz_Nosy_Downloader_concurrentFetch__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

func Fuzz_Nosy_Downloader_fetchBeaconHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var from uint64
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.fetchBeaconHeaders(from)
	})
}

func Fuzz_Nosy_Downloader_fetchBodies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var from uint64
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var beaconMode bool
		fill_err = tp.Fill(&beaconMode)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.fetchBodies(from, beaconMode)
	})
}

func Fuzz_Nosy_Downloader_fetchHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if d == nil || p == nil {
			return
		}

		d.fetchHead(p)
	})
}

func Fuzz_Nosy_Downloader_fetchHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var from uint64
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var head uint64
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if d == nil || p == nil {
			return
		}

		d.fetchHeaders(p, from, head)
	})
}

func Fuzz_Nosy_Downloader_fetchHeadersByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var amount int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var skip int
		fill_err = tp.Fill(&skip)
		if fill_err != nil {
			return
		}
		var reverse bool
		fill_err = tp.Fill(&reverse)
		if fill_err != nil {
			return
		}
		if d == nil || p == nil {
			return
		}

		d.fetchHeadersByHash(p, hash, amount, skip, reverse)
	})
}

func Fuzz_Nosy_Downloader_fetchHeadersByNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		var amount int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var skip int
		fill_err = tp.Fill(&skip)
		if fill_err != nil {
			return
		}
		var reverse bool
		fill_err = tp.Fill(&reverse)
		if fill_err != nil {
			return
		}
		if d == nil || p == nil {
			return
		}

		d.fetchHeadersByNumber(p, number, amount, skip, reverse)
	})
}

func Fuzz_Nosy_Downloader_fetchReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var from uint64
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var beaconMode bool
		fill_err = tp.Fill(&beaconMode)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.fetchReceipts(from, beaconMode)
	})
}

func Fuzz_Nosy_Downloader_fillHeaderSkeleton__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var from uint64
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var skeleton []*types.Header
		fill_err = tp.Fill(&skeleton)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.fillHeaderSkeleton(from, skeleton)
	})
}

func Fuzz_Nosy_Downloader_findAncestor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var remoteHeader *types.Header
		fill_err = tp.Fill(&remoteHeader)
		if fill_err != nil {
			return
		}
		if d == nil || p == nil || remoteHeader == nil {
			return
		}

		d.findAncestor(p, remoteHeader)
	})
}

func Fuzz_Nosy_Downloader_findAncestorBinarySearch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}
		var remoteHeight uint64
		fill_err = tp.Fill(&remoteHeight)
		if fill_err != nil {
			return
		}
		var floor int64
		fill_err = tp.Fill(&floor)
		if fill_err != nil {
			return
		}
		if d == nil || p == nil {
			return
		}

		d.findAncestorBinarySearch(p, mode, remoteHeight, floor)
	})
}

func Fuzz_Nosy_Downloader_findAncestorSpanSearch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}
		var remoteHeight uint64
		fill_err = tp.Fill(&remoteHeight)
		if fill_err != nil {
			return
		}
		var localHeight uint64
		fill_err = tp.Fill(&localHeight)
		if fill_err != nil {
			return
		}
		var floor int64
		fill_err = tp.Fill(&floor)
		if fill_err != nil {
			return
		}
		if d == nil || p == nil {
			return
		}

		d.findAncestorSpanSearch(p, mode, remoteHeight, localHeight, floor)
	})
}

func Fuzz_Nosy_Downloader_findBeaconAncestor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.findBeaconAncestor()
	})
}

func Fuzz_Nosy_Downloader_getMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.getMode()
	})
}

func Fuzz_Nosy_Downloader_importBlockResults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var results []*fetchResult
		fill_err = tp.Fill(&results)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.importBlockResults(results)
	})
}

func Fuzz_Nosy_Downloader_processFullSyncContent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ttd *big.Int
		fill_err = tp.Fill(&ttd)
		if fill_err != nil {
			return
		}
		var beaconMode bool
		fill_err = tp.Fill(&beaconMode)
		if fill_err != nil {
			return
		}
		if d == nil || ttd == nil {
			return
		}

		d.processFullSyncContent(ttd, beaconMode)
	})
}

func Fuzz_Nosy_Downloader_processHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var origin uint64
		fill_err = tp.Fill(&origin)
		if fill_err != nil {
			return
		}
		var td *big.Int
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		var ttd *big.Int
		fill_err = tp.Fill(&ttd)
		if fill_err != nil {
			return
		}
		var beaconMode bool
		fill_err = tp.Fill(&beaconMode)
		if fill_err != nil {
			return
		}
		if d == nil || td == nil || ttd == nil {
			return
		}

		d.processHeaders(origin, td, ttd, beaconMode)
	})
}

func Fuzz_Nosy_Downloader_processSnapSyncContent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.processSnapSyncContent()
	})
}

func Fuzz_Nosy_Downloader_readHeaderRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var last *types.Header
		fill_err = tp.Fill(&last)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if d == nil || last == nil {
			return
		}

		d.readHeaderRange(last, count)
	})
}

func Fuzz_Nosy_Downloader_reportSnapSyncProgress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.reportSnapSyncProgress(force)
	})
}

func Fuzz_Nosy_Downloader_runStateSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var s *stateSync
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if d == nil || s == nil {
			return
		}

		d.runStateSync(s)
	})
}

// skipping Fuzz_Nosy_Downloader_spawnSync__ because parameters include func, chan, or unsupported interface: []func() error

func Fuzz_Nosy_Downloader_stateFetcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.stateFetcher()
	})
}

func Fuzz_Nosy_Downloader_syncState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.syncState(root)
	})
}

func Fuzz_Nosy_Downloader_syncWithPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var td *big.Int
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		var ttd *big.Int
		fill_err = tp.Fill(&ttd)
		if fill_err != nil {
			return
		}
		var beaconMode bool
		fill_err = tp.Fill(&beaconMode)
		if fill_err != nil {
			return
		}
		if d == nil || p == nil || td == nil || ttd == nil {
			return
		}

		d.syncWithPeer(p, hash, td, ttd, beaconMode)
	})
}

// skipping Fuzz_Nosy_Downloader_synchronise__ because parameters include func, chan, or unsupported interface: chan struct{}

// skipping Fuzz_Nosy_DownloaderAPI_SubscribeSyncStatus__ because parameters include func, chan, or unsupported interface: chan interface{}

func Fuzz_Nosy_DownloaderAPI_Syncing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var m *event.TypeMux
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if d == nil || m == nil {
			return
		}

		api := NewDownloaderAPI(d, m)
		api.Syncing(ctx)
	})
}

func Fuzz_Nosy_DownloaderAPI_eventLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var m *event.TypeMux
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if d == nil || m == nil {
			return
		}

		api := NewDownloaderAPI(d, m)
		api.eventLoop()
	})
}

func Fuzz_Nosy_SyncMode_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mode *SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if mode == nil {
			return
		}

		mode.UnmarshalText(text)
	})
}

func Fuzz_Nosy_SyncStatusSubscription_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SyncStatusSubscription
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Unsubscribe()
	})
}

func Fuzz_Nosy_beaconBackfiller_resume__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *beaconBackfiller
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.resume()
	})
}

func Fuzz_Nosy_beaconBackfiller_setMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *beaconBackfiller
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.setMode(mode)
	})
}

func Fuzz_Nosy_beaconBackfiller_suspend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *beaconBackfiller
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.suspend()
	})
}

func Fuzz_Nosy_bodyQueue_capacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *bodyQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var rtt time.Duration
		fill_err = tp.Fill(&rtt)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.capacity(peer, rtt)
	})
}

func Fuzz_Nosy_bodyQueue_deliver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *bodyQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var packet *eth.Response
		fill_err = tp.Fill(&packet)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil || packet == nil {
			return
		}

		q.deliver(peer, packet)
	})
}

func Fuzz_Nosy_bodyQueue_pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *bodyQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.pending()
	})
}

// skipping Fuzz_Nosy_bodyQueue_request__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/eth.Response

func Fuzz_Nosy_bodyQueue_reserve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *bodyQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var items int
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.reserve(peer, items)
	})
}

func Fuzz_Nosy_bodyQueue_unreserve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *bodyQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.unreserve(peer)
	})
}

func Fuzz_Nosy_bodyQueue_updateCapacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *bodyQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var items int
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		var span time.Duration
		fill_err = tp.Fill(&span)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.updateCapacity(peer, items, span)
	})
}

func Fuzz_Nosy_bodyQueue_waker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *bodyQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.waker()
	})
}

func Fuzz_Nosy_fetchResult_AllDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var fastSync bool
		fill_err = tp.Fill(&fastSync)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		f1 := newFetchResult(header, fastSync)
		f1.AllDone()
	})
}

func Fuzz_Nosy_fetchResult_Done__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var fastSync bool
		fill_err = tp.Fill(&fastSync)
		if fill_err != nil {
			return
		}
		var kind uint
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		f1 := newFetchResult(header, fastSync)
		f1.Done(kind)
	})
}

func Fuzz_Nosy_fetchResult_SetBodyDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var fastSync bool
		fill_err = tp.Fill(&fastSync)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		f1 := newFetchResult(header, fastSync)
		f1.SetBodyDone()
	})
}

func Fuzz_Nosy_fetchResult_SetReceiptsDone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var fastSync bool
		fill_err = tp.Fill(&fastSync)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		f1 := newFetchResult(header, fastSync)
		f1.SetReceiptsDone()
	})
}

func Fuzz_Nosy_headerQueue_capacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *headerQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var rtt time.Duration
		fill_err = tp.Fill(&rtt)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.capacity(peer, rtt)
	})
}

func Fuzz_Nosy_headerQueue_deliver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *headerQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var packet *eth.Response
		fill_err = tp.Fill(&packet)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil || packet == nil {
			return
		}

		q.deliver(peer, packet)
	})
}

func Fuzz_Nosy_headerQueue_pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *headerQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.pending()
	})
}

// skipping Fuzz_Nosy_headerQueue_request__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/eth.Response

func Fuzz_Nosy_headerQueue_reserve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *headerQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var items int
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.reserve(peer, items)
	})
}

func Fuzz_Nosy_headerQueue_unreserve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *headerQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.unreserve(peer)
	})
}

func Fuzz_Nosy_headerQueue_updateCapacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *headerQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var items int
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		var span time.Duration
		fill_err = tp.Fill(&span)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.updateCapacity(peer, items, span)
	})
}

func Fuzz_Nosy_headerQueue_waker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *headerQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.waker()
	})
}

func Fuzz_Nosy_peerCapacitySort_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *peerCapacitySort
		fill_err = tp.Fill(&ps)
		if fill_err != nil {
			return
		}
		if ps == nil {
			return
		}

		ps.Len()
	})
}

func Fuzz_Nosy_peerCapacitySort_Less__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *peerCapacitySort
		fill_err = tp.Fill(&ps)
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
		if ps == nil {
			return
		}

		ps.Less(i, j)
	})
}

func Fuzz_Nosy_peerCapacitySort_Swap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ps *peerCapacitySort
		fill_err = tp.Fill(&ps)
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
		if ps == nil {
			return
		}

		ps.Swap(i, j)
	})
}

func Fuzz_Nosy_peerConnection_BodyCapacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var targetRTT time.Duration
		fill_err = tp.Fill(&targetRTT)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.BodyCapacity(targetRTT)
	})
}

func Fuzz_Nosy_peerConnection_HeaderCapacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var targetRTT time.Duration
		fill_err = tp.Fill(&targetRTT)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.HeaderCapacity(targetRTT)
	})
}

func Fuzz_Nosy_peerConnection_Lacks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Lacks(hash)
	})
}

func Fuzz_Nosy_peerConnection_MarkLacking__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var hash common.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.MarkLacking(hash)
	})
}

func Fuzz_Nosy_peerConnection_ReceiptCapacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var targetRTT time.Duration
		fill_err = tp.Fill(&targetRTT)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ReceiptCapacity(targetRTT)
	})
}

func Fuzz_Nosy_peerConnection_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Reset()
	})
}

func Fuzz_Nosy_peerConnection_UpdateBodyRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var delivered int
		fill_err = tp.Fill(&delivered)
		if fill_err != nil {
			return
		}
		var elapsed time.Duration
		fill_err = tp.Fill(&elapsed)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.UpdateBodyRate(delivered, elapsed)
	})
}

func Fuzz_Nosy_peerConnection_UpdateHeaderRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var delivered int
		fill_err = tp.Fill(&delivered)
		if fill_err != nil {
			return
		}
		var elapsed time.Duration
		fill_err = tp.Fill(&elapsed)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.UpdateHeaderRate(delivered, elapsed)
	})
}

func Fuzz_Nosy_peerConnection_UpdateReceiptRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var delivered int
		fill_err = tp.Fill(&delivered)
		if fill_err != nil {
			return
		}
		var elapsed time.Duration
		fill_err = tp.Fill(&elapsed)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.UpdateReceiptRate(delivered, elapsed)
	})
}

func Fuzz_Nosy_peerSet_Peer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		ps := newPeerSet()
		ps.Peer(id)
	})
}

func Fuzz_Nosy_peerSet_Register__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		ps := newPeerSet()
		ps.Register(p)
	})
}

// skipping Fuzz_Nosy_peerSet_SubscribeEvents__ because parameters include func, chan, or unsupported interface: chan<- *github.com/ethereum/go-ethereum/eth/downloader.peeringEvent

func Fuzz_Nosy_peerSet_Unregister__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		ps := newPeerSet()
		ps.Unregister(id)
	})
}

func Fuzz_Nosy_queue_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.Close()
	})
}

func Fuzz_Nosy_queue_DeliverBodies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var txLists [][]*types.Transaction
		fill_err = tp.Fill(&txLists)
		if fill_err != nil {
			return
		}
		var txListHashes []common.Hash
		fill_err = tp.Fill(&txListHashes)
		if fill_err != nil {
			return
		}
		var uncleLists [][]*types.Header
		fill_err = tp.Fill(&uncleLists)
		if fill_err != nil {
			return
		}
		var uncleListHashes []common.Hash
		fill_err = tp.Fill(&uncleListHashes)
		if fill_err != nil {
			return
		}
		var withdrawalLists [][]*types.Withdrawal
		fill_err = tp.Fill(&withdrawalLists)
		if fill_err != nil {
			return
		}
		var withdrawalListHashes []common.Hash
		fill_err = tp.Fill(&withdrawalListHashes)
		if fill_err != nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.DeliverBodies(id, txLists, txListHashes, uncleLists, uncleListHashes, withdrawalLists, withdrawalListHashes)
	})
}

// skipping Fuzz_Nosy_queue_DeliverHeaders__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/downloader.headerTask

func Fuzz_Nosy_queue_DeliverReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var receiptList [][]*types.Receipt
		fill_err = tp.Fill(&receiptList)
		if fill_err != nil {
			return
		}
		var receiptListHashes []common.Hash
		fill_err = tp.Fill(&receiptListHashes)
		if fill_err != nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.DeliverReceipts(id, receiptList, receiptListHashes)
	})
}

func Fuzz_Nosy_queue_ExpireBodies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int, peer string) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.ExpireBodies(peer)
	})
}

func Fuzz_Nosy_queue_ExpireHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int, peer string) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.ExpireHeaders(peer)
	})
}

func Fuzz_Nosy_queue_ExpireReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int, peer string) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.ExpireReceipts(peer)
	})
}

func Fuzz_Nosy_queue_Idle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.Idle()
	})
}

func Fuzz_Nosy_queue_InFlightBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.InFlightBlocks()
	})
}

func Fuzz_Nosy_queue_InFlightReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.InFlightReceipts()
	})
}

func Fuzz_Nosy_queue_PendingBodies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.PendingBodies()
	})
}

func Fuzz_Nosy_queue_PendingHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.PendingHeaders()
	})
}

func Fuzz_Nosy_queue_PendingReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.PendingReceipts()
	})
}

func Fuzz_Nosy_queue_Prepare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.Prepare(offset, mode)
	})
}

func Fuzz_Nosy_queue_ReserveBodies__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.ReserveBodies(p, count)
	})
}

func Fuzz_Nosy_queue_ReserveHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.ReserveHeaders(p, count)
	})
}

func Fuzz_Nosy_queue_ReserveReceipts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.ReserveReceipts(p, count)
	})
}

func Fuzz_Nosy_queue_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b1 int, t2 int, b3 int, t4 int) {
		q := newQueue(b1, t2)
		q.Reset(b3, t4)
	})
}

func Fuzz_Nosy_queue_Results__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int, block bool) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.Results(block)
	})
}

func Fuzz_Nosy_queue_RetrieveHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.RetrieveHeaders()
	})
}

func Fuzz_Nosy_queue_Revoke__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int, peerID string) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.Revoke(peerID)
	})
}

func Fuzz_Nosy_queue_Schedule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var headers []*types.Header
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}
		var hashes []common.Hash
		fill_err = tp.Fill(&hashes)
		if fill_err != nil {
			return
		}
		var from uint64
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.Schedule(headers, hashes, from)
	})
}

func Fuzz_Nosy_queue_ScheduleSkeleton__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var from uint64
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var skeleton []*types.Header
		fill_err = tp.Fill(&skeleton)
		if fill_err != nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.ScheduleSkeleton(from, skeleton)
	})
}

func Fuzz_Nosy_queue_Stats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.Stats()
	})
}

// skipping Fuzz_Nosy_queue_deliver__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Timer

// skipping Fuzz_Nosy_queue_expire__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_queue_reserveHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockCacheLimit int
		fill_err = tp.Fill(&blockCacheLimit)
		if fill_err != nil {
			return
		}
		var thresholdInitialSize int
		fill_err = tp.Fill(&thresholdInitialSize)
		if fill_err != nil {
			return
		}
		var p *peerConnection
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		var taskPool map[common.Hash]*types.Header
		fill_err = tp.Fill(&taskPool)
		if fill_err != nil {
			return
		}
		var taskQueue *prque.Prque[int64, *types.Header]
		fill_err = tp.Fill(&taskQueue)
		if fill_err != nil {
			return
		}
		var pendPool map[string]*fetchRequest
		fill_err = tp.Fill(&pendPool)
		if fill_err != nil {
			return
		}
		var kind uint
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		if p == nil || taskQueue == nil {
			return
		}

		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.reserveHeaders(p, count, taskPool, taskQueue, pendPool, kind)
	})
}

func Fuzz_Nosy_queue_stats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockCacheLimit int, thresholdInitialSize int) {
		q := newQueue(blockCacheLimit, thresholdInitialSize)
		q.stats()
	})
}

func Fuzz_Nosy_receiptQueue_capacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *receiptQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var rtt time.Duration
		fill_err = tp.Fill(&rtt)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.capacity(peer, rtt)
	})
}

func Fuzz_Nosy_receiptQueue_deliver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *receiptQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var packet *eth.Response
		fill_err = tp.Fill(&packet)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil || packet == nil {
			return
		}

		q.deliver(peer, packet)
	})
}

func Fuzz_Nosy_receiptQueue_pending__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *receiptQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.pending()
	})
}

// skipping Fuzz_Nosy_receiptQueue_request__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/protocols/eth.Response

func Fuzz_Nosy_receiptQueue_reserve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *receiptQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var items int
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.reserve(peer, items)
	})
}

func Fuzz_Nosy_receiptQueue_unreserve__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *receiptQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.unreserve(peer)
	})
}

func Fuzz_Nosy_receiptQueue_updateCapacity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *receiptQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var items int
		fill_err = tp.Fill(&items)
		if fill_err != nil {
			return
		}
		var span time.Duration
		fill_err = tp.Fill(&span)
		if fill_err != nil {
			return
		}
		if q == nil || peer == nil {
			return
		}

		q.updateCapacity(peer, items, span)
	})
}

func Fuzz_Nosy_receiptQueue_waker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *receiptQueue
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.waker()
	})
}

func Fuzz_Nosy_resultStore_AddFetch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var size int
		fill_err = tp.Fill(&size)
		if fill_err != nil {
			return
		}
		var header *types.Header
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var fastSync bool
		fill_err = tp.Fill(&fastSync)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		r := newResultStore(size)
		r.AddFetch(header, fastSync)
	})
}

func Fuzz_Nosy_resultStore_GetCompleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, limit int) {
		r := newResultStore(size)
		r.GetCompleted(limit)
	})
}

func Fuzz_Nosy_resultStore_GetDeliverySlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, headerNumber uint64) {
		r := newResultStore(size)
		r.GetDeliverySlot(headerNumber)
	})
}

func Fuzz_Nosy_resultStore_HasCompletedItems__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
		r := newResultStore(size)
		r.HasCompletedItems()
	})
}

func Fuzz_Nosy_resultStore_Prepare__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, offset uint64) {
		r := newResultStore(size)
		r.Prepare(offset)
	})
}

func Fuzz_Nosy_resultStore_SetThrottleThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, threshold uint64) {
		r := newResultStore(size)
		r.SetThrottleThreshold(threshold)
	})
}

func Fuzz_Nosy_resultStore_countCompleted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int) {
		r := newResultStore(size)
		r.countCompleted()
	})
}

func Fuzz_Nosy_resultStore_getFetchResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, size int, headerNumber uint64) {
		r := newResultStore(size)
		r.getFetchResult(headerNumber)
	})
}

func Fuzz_Nosy_skeleton_Bounds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Bounds()
	})
}

func Fuzz_Nosy_skeleton_Header__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var number uint64
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Header(number)
	})
}

func Fuzz_Nosy_skeleton_Sync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var final *types.Header
		fill_err = tp.Fill(&final)
		if fill_err != nil {
			return
		}
		var force bool
		fill_err = tp.Fill(&force)
		if fill_err != nil {
			return
		}
		if s == nil || head == nil || final == nil {
			return
		}

		s.Sync(head, final, force)
	})
}

func Fuzz_Nosy_skeleton_Terminate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Terminate()
	})
}

// skipping Fuzz_Nosy_skeleton_assignTasks__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/eth/downloader.headerResponse

func Fuzz_Nosy_skeleton_cleanStales__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var filled *types.Header
		fill_err = tp.Fill(&filled)
		if fill_err != nil {
			return
		}
		if s == nil || filled == nil {
			return
		}

		s.cleanStales(filled)
	})
}

func Fuzz_Nosy_skeleton_executeTask__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var peer *peerConnection
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		var req *headerRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || peer == nil || req == nil {
			return
		}

		s.executeTask(peer, req)
	})
}

func Fuzz_Nosy_skeleton_initSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if s == nil || head == nil {
			return
		}

		s.initSync(head)
	})
}

func Fuzz_Nosy_skeleton_processNewHead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		var final *types.Header
		fill_err = tp.Fill(&final)
		if fill_err != nil {
			return
		}
		if s == nil || head == nil || final == nil {
			return
		}

		s.processNewHead(head, final)
	})
}

func Fuzz_Nosy_skeleton_processResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var res *headerResponse
		fill_err = tp.Fill(&res)
		if fill_err != nil {
			return
		}
		if s == nil || res == nil {
			return
		}

		s.processResponse(res)
	})
}

func Fuzz_Nosy_skeleton_revertRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *headerRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.revertRequest(req)
	})
}

func Fuzz_Nosy_skeleton_revertRequests__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var peer string
		fill_err = tp.Fill(&peer)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.revertRequests(peer)
	})
}

// skipping Fuzz_Nosy_skeleton_saveSyncStatus__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.KeyValueWriter

func Fuzz_Nosy_skeleton_scheduleRevertRequest__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var req *headerRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if s == nil || req == nil {
			return
		}

		s.scheduleRevertRequest(req)
	})
}

func Fuzz_Nosy_skeleton_startup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.startup()
	})
}

func Fuzz_Nosy_skeleton_sync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *skeleton
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var head *types.Header
		fill_err = tp.Fill(&head)
		if fill_err != nil {
			return
		}
		if s == nil || head == nil {
			return
		}

		s.sync(head)
	})
}

func Fuzz_Nosy_stateSync_Cancel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		s := newStateSync(d, root)
		s.Cancel()
	})
}

func Fuzz_Nosy_stateSync_Wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		s := newStateSync(d, root)
		s.Wait()
	})
}

func Fuzz_Nosy_stateSync_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Downloader
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var root common.Hash
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		s := newStateSync(d, root)
		s.run()
	})
}

// skipping Fuzz_Nosy_BlockChain_CurrentBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_CurrentSnapBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_GetBlockByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_HasBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_HasFastBlock__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_InsertChain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_InsertReceiptChain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_SnapSyncCommitHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_Snapshots__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_BlockChain_TrieDB__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.BlockChain

// skipping Fuzz_Nosy_LightChain_CurrentHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.LightChain

// skipping Fuzz_Nosy_LightChain_GetHeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.LightChain

// skipping Fuzz_Nosy_LightChain_GetTd__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.LightChain

// skipping Fuzz_Nosy_LightChain_HasHeader__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.LightChain

// skipping Fuzz_Nosy_LightChain_InsertHeaderChain__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.LightChain

// skipping Fuzz_Nosy_LightChain_SetHead__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.LightChain

// skipping Fuzz_Nosy_Peer_Head__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.Peer

// skipping Fuzz_Nosy_Peer_RequestBodies__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.Peer

// skipping Fuzz_Nosy_Peer_RequestHeadersByHash__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.Peer

// skipping Fuzz_Nosy_Peer_RequestHeadersByNumber__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.Peer

// skipping Fuzz_Nosy_Peer_RequestReceipts__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.Peer

func Fuzz_Nosy_SyncMode_IsValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}

		mode.IsValid()
	})
}

func Fuzz_Nosy_SyncMode_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}

		mode.MarshalText()
	})
}

func Fuzz_Nosy_SyncMode_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mode SyncMode
		fill_err = tp.Fill(&mode)
		if fill_err != nil {
			return
		}

		mode.String()
	})
}

// skipping Fuzz_Nosy_backfiller_resume__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.backfiller

// skipping Fuzz_Nosy_backfiller_suspend__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.backfiller

// skipping Fuzz_Nosy_typedQueue_capacity__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

// skipping Fuzz_Nosy_typedQueue_deliver__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

// skipping Fuzz_Nosy_typedQueue_pending__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

// skipping Fuzz_Nosy_typedQueue_request__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

// skipping Fuzz_Nosy_typedQueue_reserve__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

// skipping Fuzz_Nosy_typedQueue_unreserve__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

// skipping Fuzz_Nosy_typedQueue_updateCapacity__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

// skipping Fuzz_Nosy_typedQueue_waker__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/eth/downloader.typedQueue

func Fuzz_Nosy_calculateRequestSpan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, remoteHeight uint64, localHeight uint64) {
		calculateRequestSpan(remoteHeight, localHeight)
	})
}

func Fuzz_Nosy_splitAroundPivot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pivot uint64
		fill_err = tp.Fill(&pivot)
		if fill_err != nil {
			return
		}
		var results []*fetchResult
		fill_err = tp.Fill(&results)
		if fill_err != nil {
			return
		}

		splitAroundPivot(pivot, results)
	})
}
