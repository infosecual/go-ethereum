package nodestate

import (
	"testing"
	"time"

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

// skipping Fuzz_Nosy_NodeStateMachine_AddLogMetrics__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/metrics.Meter

func Fuzz_Nosy_NodeStateMachine_AddTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var flags Flags
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.AddTimeout(n, flags, timeout)
	})
}

// skipping Fuzz_Nosy_NodeStateMachine_ForEach__ because parameters include func, chan, or unsupported interface: func(n *github.com/ethereum/go-ethereum/p2p/enode.Node, state github.com/ethereum/go-ethereum/p2p/nodestate.Flags)

func Fuzz_Nosy_NodeStateMachine_GetField__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var field Field
		fill_err = tp.Fill(&field)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.GetField(n, field)
	})
}

func Fuzz_Nosy_NodeStateMachine_GetNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.GetNode(id)
	})
}

func Fuzz_Nosy_NodeStateMachine_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.GetState(n)
	})
}

// skipping Fuzz_Nosy_NodeStateMachine_Operation__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_NodeStateMachine_Persist__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.Persist(n)
	})
}

// skipping Fuzz_Nosy_NodeStateMachine_SetField__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_NodeStateMachine_SetFieldSub__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_NodeStateMachine_SetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var setFlags Flags
		fill_err = tp.Fill(&setFlags)
		if fill_err != nil {
			return
		}
		var resetFlags Flags
		fill_err = tp.Fill(&resetFlags)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.SetState(n, setFlags, resetFlags, timeout)
	})
}

func Fuzz_Nosy_NodeStateMachine_SetStateSub__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var setFlags Flags
		fill_err = tp.Fill(&setFlags)
		if fill_err != nil {
			return
		}
		var resetFlags Flags
		fill_err = tp.Fill(&resetFlags)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.SetStateSub(n, setFlags, resetFlags, timeout)
	})
}

func Fuzz_Nosy_NodeStateMachine_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.Start()
	})
}

func Fuzz_Nosy_NodeStateMachine_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.Stop()
	})
}

// skipping Fuzz_Nosy_NodeStateMachine_SubscribeField__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/nodestate.FieldCallback

// skipping Fuzz_Nosy_NodeStateMachine_SubscribeState__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/p2p/nodestate.StateCallback

func Fuzz_Nosy_NodeStateMachine_addTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var mask bitMask
		fill_err = tp.Fill(&mask)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.addTimeout(n, mask, timeout)
	})
}

func Fuzz_Nosy_NodeStateMachine_checkStarted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.checkStarted()
	})
}

func Fuzz_Nosy_NodeStateMachine_decodeNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.decodeNode(id, d3)
	})
}

func Fuzz_Nosy_NodeStateMachine_deleteNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.deleteNode(id)
	})
}

func Fuzz_Nosy_NodeStateMachine_fieldIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var field Field
		fill_err = tp.Fill(&field)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.fieldIndex(field)
	})
}

func Fuzz_Nosy_NodeStateMachine_loadFromDb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.loadFromDb()
	})
}

func Fuzz_Nosy_NodeStateMachine_newNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.newNode(n)
	})
}

func Fuzz_Nosy_NodeStateMachine_offlineCallbacks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var start bool
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.offlineCallbacks(start)
	})
}

func Fuzz_Nosy_NodeStateMachine_opCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.opCheck()
	})
}

func Fuzz_Nosy_NodeStateMachine_opFinish__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.opFinish()
	})
}

func Fuzz_Nosy_NodeStateMachine_opStart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.opStart()
	})
}

func Fuzz_Nosy_NodeStateMachine_removeTimeouts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var node *nodeInfo
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		var mask bitMask
		fill_err = tp.Fill(&mask)
		if fill_err != nil {
			return
		}
		if ns == nil || node == nil {
			return
		}

		ns.removeTimeouts(node, mask)
	})
}

func Fuzz_Nosy_NodeStateMachine_saveNode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var id enode.ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var node *nodeInfo
		fill_err = tp.Fill(&node)
		if fill_err != nil {
			return
		}
		if ns == nil || node == nil {
			return
		}

		ns.saveNode(id, node)
	})
}

func Fuzz_Nosy_NodeStateMachine_saveToDb__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.saveToDb()
	})
}

// skipping Fuzz_Nosy_NodeStateMachine_setField__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_NodeStateMachine_setState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var setFlags Flags
		fill_err = tp.Fill(&setFlags)
		if fill_err != nil {
			return
		}
		var resetFlags Flags
		fill_err = tp.Fill(&resetFlags)
		if fill_err != nil {
			return
		}
		var timeout time.Duration
		fill_err = tp.Fill(&timeout)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.setState(n, setFlags, resetFlags, timeout)
	})
}

func Fuzz_Nosy_NodeStateMachine_stateMask__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var flags Flags
		fill_err = tp.Fill(&flags)
		if fill_err != nil {
			return
		}
		if ns == nil {
			return
		}

		ns.stateMask(flags)
	})
}

func Fuzz_Nosy_NodeStateMachine_updateEnode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ns *NodeStateMachine
		fill_err = tp.Fill(&ns)
		if fill_err != nil {
			return
		}
		var n *enode.Node
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if ns == nil || n == nil {
			return
		}

		ns.updateEnode(n)
	})
}

// skipping Fuzz_Nosy_Setup_NewField__ because parameters include func, chan, or unsupported interface: reflect.Type

func Fuzz_Nosy_Setup_NewFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Setup
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.NewFlag(name)
	})
}

// skipping Fuzz_Nosy_Setup_NewPersistentField__ because parameters include func, chan, or unsupported interface: reflect.Type

func Fuzz_Nosy_Setup_NewPersistentFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Setup
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.NewPersistentFlag(name)
	})
}

func Fuzz_Nosy_Setup_OfflineFlag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Setup
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.OfflineFlag()
	})
}

func Fuzz_Nosy_Flags_And__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var b Flags
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		a := MergeFlags(list...)
		a.And(b)
	})
}

func Fuzz_Nosy_Flags_AndNot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var b Flags
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		a := MergeFlags(list...)
		a.AndNot(b)
	})
}

func Fuzz_Nosy_Flags_Equals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var b Flags
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		a := MergeFlags(list...)
		a.Equals(b)
	})
}

func Fuzz_Nosy_Flags_HasAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var b Flags
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		a := MergeFlags(list...)
		a.HasAll(b)
	})
}

func Fuzz_Nosy_Flags_HasNone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var b Flags
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		a := MergeFlags(list...)
		a.HasNone(b)
	})
}

func Fuzz_Nosy_Flags_IsEmpty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}

		a := MergeFlags(list...)
		a.IsEmpty()
	})
}

func Fuzz_Nosy_Flags_Or__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var b Flags
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		a := MergeFlags(list...)
		a.Or(b)
	})
}

func Fuzz_Nosy_Flags_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}

		f1 := MergeFlags(list...)
		f1.String()
	})
}

func Fuzz_Nosy_Flags_Xor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var list []Flags
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}
		var b Flags
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		a := MergeFlags(list...)
		a.Xor(b)
	})
}

func Fuzz_Nosy_dummyIdentity_NodeAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id dummyIdentity
		fill_err = tp.Fill(&id)
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

		id.NodeAddr(r)
	})
}

func Fuzz_Nosy_dummyIdentity_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id dummyIdentity
		fill_err = tp.Fill(&id)
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

		id.Verify(r, sig)
	})
}
