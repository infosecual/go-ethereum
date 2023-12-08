package node

import (
	"context"
	"net/http"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
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

func Fuzz_Nosy_Config_ExtRPCEnabled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ExtRPCEnabled()
	})
}

func Fuzz_Nosy_Config_GetKeyStoreDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.GetKeyStoreDir()
	})
}

func Fuzz_Nosy_Config_HTTPEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.HTTPEndpoint()
	})
}

func Fuzz_Nosy_Config_IPCEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.IPCEndpoint()
	})
}

func Fuzz_Nosy_Config_KeyDirConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.KeyDirConfig()
	})
}

func Fuzz_Nosy_Config_NodeDB__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.NodeDB()
	})
}

func Fuzz_Nosy_Config_NodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.NodeKey()
	})
}

func Fuzz_Nosy_Config_NodeName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.NodeName()
	})
}

func Fuzz_Nosy_Config_ResolvePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ResolvePath(path)
	})
}

func Fuzz_Nosy_Config_WSEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.WSEndpoint()
	})
}

func Fuzz_Nosy_Config_checkLegacyFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.checkLegacyFile(path)
	})
}

func Fuzz_Nosy_Config_checkLegacyFiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.checkLegacyFiles()
	})
}

func Fuzz_Nosy_Config_instanceDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.instanceDir()
	})
}

func Fuzz_Nosy_Config_name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *Config
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.name()
	})
}

func Fuzz_Nosy_Node_AccountManager__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.AccountManager()
	})
}

func Fuzz_Nosy_Node_Attach__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.Attach()
	})
}

func Fuzz_Nosy_Node_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.Close()
	})
}

func Fuzz_Nosy_Node_Config__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.Config()
	})
}

func Fuzz_Nosy_Node_DataDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.DataDir()
	})
}

func Fuzz_Nosy_Node_EventMux__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.EventMux()
	})
}

func Fuzz_Nosy_Node_HTTPAuthEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.HTTPAuthEndpoint()
	})
}

func Fuzz_Nosy_Node_HTTPEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.HTTPEndpoint()
	})
}

func Fuzz_Nosy_Node_IPCEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.IPCEndpoint()
	})
}

func Fuzz_Nosy_Node_InstanceDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.InstanceDir()
	})
}

func Fuzz_Nosy_Node_KeyStoreDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.KeyStoreDir()
	})
}

func Fuzz_Nosy_Node_OpenDatabase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var cache int
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var handles int
		fill_err = tp.Fill(&handles)
		if fill_err != nil {
			return
		}
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.OpenDatabase(name, cache, handles, namespace, readonly)
	})
}

func Fuzz_Nosy_Node_OpenDatabaseWithFreezer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var cache int
		fill_err = tp.Fill(&cache)
		if fill_err != nil {
			return
		}
		var handles int
		fill_err = tp.Fill(&handles)
		if fill_err != nil {
			return
		}
		var ancient string
		fill_err = tp.Fill(&ancient)
		if fill_err != nil {
			return
		}
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var readonly bool
		fill_err = tp.Fill(&readonly)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.OpenDatabaseWithFreezer(name, cache, handles, ancient, namespace, readonly)
	})
}

func Fuzz_Nosy_Node_RPCHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.RPCHandler()
	})
}

func Fuzz_Nosy_Node_RegisterAPIs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var apis []rpc.API
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.RegisterAPIs(apis)
	})
}

// skipping Fuzz_Nosy_Node_RegisterHandler__ because parameters include func, chan, or unsupported interface: net/http.Handler

// skipping Fuzz_Nosy_Node_RegisterLifecycle__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/node.Lifecycle

func Fuzz_Nosy_Node_RegisterProtocols__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var protocols []p2p.Protocol
		fill_err = tp.Fill(&protocols)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.RegisterProtocols(protocols)
	})
}

func Fuzz_Nosy_Node_ResolveAncient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var ancient string
		fill_err = tp.Fill(&ancient)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.ResolveAncient(name, ancient)
	})
}

func Fuzz_Nosy_Node_ResolvePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var x string
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.ResolvePath(x)
	})
}

func Fuzz_Nosy_Node_Server__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.Server()
	})
}

func Fuzz_Nosy_Node_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.Start()
	})
}

func Fuzz_Nosy_Node_WSAuthEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.WSAuthEndpoint()
	})
}

func Fuzz_Nosy_Node_WSEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.WSEndpoint()
	})
}

func Fuzz_Nosy_Node_Wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.Wait()
	})
}

func Fuzz_Nosy_Node_apis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.apis()
	})
}

func Fuzz_Nosy_Node_closeDataDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.closeDataDir()
	})
}

func Fuzz_Nosy_Node_closeDatabases__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.closeDatabases()
	})
}

// skipping Fuzz_Nosy_Node_doClose__ because parameters include func, chan, or unsupported interface: []error

func Fuzz_Nosy_Node_getAPIs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.getAPIs()
	})
}

func Fuzz_Nosy_Node_obtainJWTSecret__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var cliParam string
		fill_err = tp.Fill(&cliParam)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.obtainJWTSecret(cliParam)
	})
}

func Fuzz_Nosy_Node_openDataDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.openDataDir()
	})
}

func Fuzz_Nosy_Node_openEndpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.openEndpoints()
	})
}

func Fuzz_Nosy_Node_startInProc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var apis []rpc.API
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.startInProc(apis)
	})
}

func Fuzz_Nosy_Node_startRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.startRPC()
	})
}

func Fuzz_Nosy_Node_stopInProc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.stopInProc()
	})
}

func Fuzz_Nosy_Node_stopRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.stopRPC()
	})
}

// skipping Fuzz_Nosy_Node_stopServices__ because parameters include func, chan, or unsupported interface: []github.com/ethereum/go-ethereum/node.Lifecycle

// skipping Fuzz_Nosy_Node_wrapDatabase__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/ethdb.Database

func Fuzz_Nosy_Node_wsServerForPort__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conf *Config
		fill_err = tp.Fill(&conf)
		if fill_err != nil {
			return
		}
		var port int
		fill_err = tp.Fill(&port)
		if fill_err != nil {
			return
		}
		var authenticated bool
		fill_err = tp.Fill(&authenticated)
		if fill_err != nil {
			return
		}
		if conf == nil {
			return
		}

		n, err := New(conf)
		if err != nil {
			return
		}
		n.wsServerForPort(port, authenticated)
	})
}

func Fuzz_Nosy_StopError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *StopError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_adminAPI_AddPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.AddPeer(url)
	})
}

func Fuzz_Nosy_adminAPI_AddTrustedPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.AddTrustedPeer(url)
	})
}

func Fuzz_Nosy_adminAPI_Datadir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Datadir()
	})
}

func Fuzz_Nosy_adminAPI_NodeInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.NodeInfo()
	})
}

func Fuzz_Nosy_adminAPI_PeerEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.PeerEvents(ctx)
	})
}

func Fuzz_Nosy_adminAPI_Peers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.Peers()
	})
}

func Fuzz_Nosy_adminAPI_RemovePeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.RemovePeer(url)
	})
}

func Fuzz_Nosy_adminAPI_RemoveTrustedPeer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var url string
		fill_err = tp.Fill(&url)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.RemoveTrustedPeer(url)
	})
}

func Fuzz_Nosy_adminAPI_StartHTTP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var host *string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var port *int
		fill_err = tp.Fill(&port)
		if fill_err != nil {
			return
		}
		var cors *string
		fill_err = tp.Fill(&cors)
		if fill_err != nil {
			return
		}
		var apis *string
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}
		var vhosts *string
		fill_err = tp.Fill(&vhosts)
		if fill_err != nil {
			return
		}
		if api == nil || host == nil || port == nil || cors == nil || apis == nil || vhosts == nil {
			return
		}

		api.StartHTTP(host, port, cors, apis, vhosts)
	})
}

func Fuzz_Nosy_adminAPI_StartRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var host *string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var port *int
		fill_err = tp.Fill(&port)
		if fill_err != nil {
			return
		}
		var cors *string
		fill_err = tp.Fill(&cors)
		if fill_err != nil {
			return
		}
		var apis *string
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}
		var vhosts *string
		fill_err = tp.Fill(&vhosts)
		if fill_err != nil {
			return
		}
		if api == nil || host == nil || port == nil || cors == nil || apis == nil || vhosts == nil {
			return
		}

		api.StartRPC(host, port, cors, apis, vhosts)
	})
}

func Fuzz_Nosy_adminAPI_StartWS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		var host *string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var port *int
		fill_err = tp.Fill(&port)
		if fill_err != nil {
			return
		}
		var allowedOrigins *string
		fill_err = tp.Fill(&allowedOrigins)
		if fill_err != nil {
			return
		}
		var apis *string
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}
		if api == nil || host == nil || port == nil || allowedOrigins == nil || apis == nil {
			return
		}

		api.StartWS(host, port, allowedOrigins, apis)
	})
}

func Fuzz_Nosy_adminAPI_StopHTTP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.StopHTTP()
	})
}

func Fuzz_Nosy_adminAPI_StopRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.StopRPC()
	})
}

func Fuzz_Nosy_adminAPI_StopWS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var api *adminAPI
		fill_err = tp.Fill(&api)
		if fill_err != nil {
			return
		}
		if api == nil {
			return
		}

		api.StopWS()
	})
}

func Fuzz_Nosy_closeTrackingDB_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *closeTrackingDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Close()
	})
}

func Fuzz_Nosy_gzipResponseWriter_Flush__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *gzipResponseWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Flush()
	})
}

func Fuzz_Nosy_gzipResponseWriter_Header__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *gzipResponseWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Header()
	})
}

func Fuzz_Nosy_gzipResponseWriter_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *gzipResponseWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.Write(b)
	})
}

func Fuzz_Nosy_gzipResponseWriter_WriteHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *gzipResponseWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var status int
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.WriteHeader(status)
	})
}

func Fuzz_Nosy_gzipResponseWriter_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *gzipResponseWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.close()
	})
}

func Fuzz_Nosy_gzipResponseWriter_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *gzipResponseWriter
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.init()
	})
}

// skipping Fuzz_Nosy_httpServer_ServeHTTP__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_httpServer_disableRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.disableRPC()
	})
}

func Fuzz_Nosy_httpServer_disableWS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.disableWS()
	})
}

func Fuzz_Nosy_httpServer_doStop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.doStop()
	})
}

func Fuzz_Nosy_httpServer_enableRPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var apis []rpc.API
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}
		var config httpConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.enableRPC(apis, config)
	})
}

func Fuzz_Nosy_httpServer_enableWS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var apis []rpc.API
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}
		var config wsConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.enableWS(apis, config)
	})
}

func Fuzz_Nosy_httpServer_listenAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.listenAddr()
	})
}

func Fuzz_Nosy_httpServer_rpcAllowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.rpcAllowed()
	})
}

func Fuzz_Nosy_httpServer_setListenAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var host string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var port int
		fill_err = tp.Fill(&port)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.setListenAddr(host, port)
	})
}

func Fuzz_Nosy_httpServer_start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.start()
	})
}

func Fuzz_Nosy_httpServer_stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.stop()
	})
}

func Fuzz_Nosy_httpServer_stopWS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.stopWS()
	})
}

func Fuzz_Nosy_httpServer_wsAllowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *httpServer
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.wsAllowed()
	})
}

func Fuzz_Nosy_ipcServer_start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var is *ipcServer
		fill_err = tp.Fill(&is)
		if fill_err != nil {
			return
		}
		var apis []rpc.API
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}
		if is == nil {
			return
		}

		is.start(apis)
	})
}

func Fuzz_Nosy_ipcServer_stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var is *ipcServer
		fill_err = tp.Fill(&is)
		if fill_err != nil {
			return
		}
		if is == nil {
			return
		}

		is.stop()
	})
}

// skipping Fuzz_Nosy_jwtHandler_ServeHTTP__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_virtualHostHandler_ServeHTTP__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

func Fuzz_Nosy_web3API_ClientVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *web3API
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.ClientVersion()
	})
}

func Fuzz_Nosy_web3API_Sha3__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *web3API
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var input hexutil.Bytes
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Sha3(input)
	})
}

// skipping Fuzz_Nosy_Lifecycle_Start__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/node.Lifecycle

// skipping Fuzz_Nosy_Lifecycle_Stop__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/node.Lifecycle

func Fuzz_Nosy_CheckTimeouts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var timeouts *rpc.HTTPTimeouts
		fill_err = tp.Fill(&timeouts)
		if fill_err != nil {
			return
		}
		if timeouts == nil {
			return
		}

		CheckTimeouts(timeouts)
	})
}

func Fuzz_Nosy_DefaultIPCEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, clientIdentifier string) {
		DefaultIPCEndpoint(clientIdentifier)
	})
}

// skipping Fuzz_Nosy_StartHTTPEndpoint__ because parameters include func, chan, or unsupported interface: net/http.Handler

func Fuzz_Nosy_checkModuleAvailability__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var modules []string
		fill_err = tp.Fill(&modules)
		if fill_err != nil {
			return
		}
		var apis []rpc.API
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}

		checkModuleAvailability(modules, apis)
	})
}

func Fuzz_Nosy_checkPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		checkPath(r, path)
	})
}

// skipping Fuzz_Nosy_containsLifecycle__ because parameters include func, chan, or unsupported interface: []github.com/ethereum/go-ethereum/node.Lifecycle

func Fuzz_Nosy_isNonEmptyDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dir string) {
		isNonEmptyDir(dir)
	})
}

func Fuzz_Nosy_isWebsocket__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *http.Request
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		isWebsocket(r)
	})
}
