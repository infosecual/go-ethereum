package bloombits

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
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

func Fuzz_Nosy_Generator_AddBloom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sections uint
		fill_err = tp.Fill(&sections)
		if fill_err != nil {
			return
		}
		var index uint
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var bloom types.Bloom
		fill_err = tp.Fill(&bloom)
		if fill_err != nil {
			return
		}

		b, err := NewGenerator(sections)
		if err != nil {
			return
		}
		b.AddBloom(index, bloom)
	})
}

func Fuzz_Nosy_Generator_Bitset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sections uint, idx uint) {
		b, err := NewGenerator(sections)
		if err != nil {
			return
		}
		b.Bitset(idx)
	})
}

// skipping Fuzz_Nosy_Matcher_Start__ because parameters include func, chan, or unsupported interface: chan uint64

func Fuzz_Nosy_Matcher_addScheduler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sectionSize uint64
		fill_err = tp.Fill(&sectionSize)
		if fill_err != nil {
			return
		}
		var filters [][][]byte
		fill_err = tp.Fill(&filters)
		if fill_err != nil {
			return
		}
		var idx uint
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}

		m := NewMatcher(sectionSize, filters)
		m.addScheduler(idx)
	})
}

// skipping Fuzz_Nosy_Matcher_distributor__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/core/bloombits.request

func Fuzz_Nosy_Matcher_run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sectionSize uint64
		fill_err = tp.Fill(&sectionSize)
		if fill_err != nil {
			return
		}
		var filters [][][]byte
		fill_err = tp.Fill(&filters)
		if fill_err != nil {
			return
		}
		var begin uint64
		fill_err = tp.Fill(&begin)
		if fill_err != nil {
			return
		}
		var end uint64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		var buffer int
		fill_err = tp.Fill(&buffer)
		if fill_err != nil {
			return
		}
		var session *MatcherSession
		fill_err = tp.Fill(&session)
		if fill_err != nil {
			return
		}
		if session == nil {
			return
		}

		m := NewMatcher(sectionSize, filters)
		m.run(begin, end, buffer, session)
	})
}

// skipping Fuzz_Nosy_Matcher_subMatch__ because parameters include func, chan, or unsupported interface: chan *github.com/ethereum/go-ethereum/core/bloombits.partialMatches

func Fuzz_Nosy_MatcherSession_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MatcherSession
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

func Fuzz_Nosy_MatcherSession_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MatcherSession
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Error()
	})
}

// skipping Fuzz_Nosy_MatcherSession_Multiplex__ because parameters include func, chan, or unsupported interface: chan chan *github.com/ethereum/go-ethereum/core/bloombits.Retrieval

func Fuzz_Nosy_MatcherSession_allocateRetrieval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MatcherSession
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.allocateRetrieval()
	})
}

func Fuzz_Nosy_MatcherSession_allocateSections__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MatcherSession
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var bit uint
		fill_err = tp.Fill(&bit)
		if fill_err != nil {
			return
		}
		var count int
		fill_err = tp.Fill(&count)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.allocateSections(bit, count)
	})
}

func Fuzz_Nosy_MatcherSession_deliverSections__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MatcherSession
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var bit uint
		fill_err = tp.Fill(&bit)
		if fill_err != nil {
			return
		}
		var sections []uint64
		fill_err = tp.Fill(&sections)
		if fill_err != nil {
			return
		}
		var bitsets [][]byte
		fill_err = tp.Fill(&bitsets)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.deliverSections(bit, sections, bitsets)
	})
}

func Fuzz_Nosy_MatcherSession_pendingSections__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MatcherSession
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var bit uint
		fill_err = tp.Fill(&bit)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.pendingSections(bit)
	})
}

func Fuzz_Nosy_scheduler_deliver__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var idx uint
		fill_err = tp.Fill(&idx)
		if fill_err != nil {
			return
		}
		var sections []uint64
		fill_err = tp.Fill(&sections)
		if fill_err != nil {
			return
		}
		var d3 [][]byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}

		s := newScheduler(idx)
		s.deliver(sections, d3)
	})
}

func Fuzz_Nosy_scheduler_reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, idx uint) {
		s := newScheduler(idx)
		s.reset()
	})
}

// skipping Fuzz_Nosy_scheduler_run__ because parameters include func, chan, or unsupported interface: chan uint64

// skipping Fuzz_Nosy_scheduler_scheduleDeliveries__ because parameters include func, chan, or unsupported interface: chan uint64

// skipping Fuzz_Nosy_scheduler_scheduleRequests__ because parameters include func, chan, or unsupported interface: chan uint64
