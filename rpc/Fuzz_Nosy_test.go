package rpc

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"

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

func Fuzz_Nosy_BlockNumber_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bn *BlockNumber
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if bn == nil {
			return
		}

		bn.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_BlockNumberOrHash_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockNr BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}

		bnh := BlockNumberOrHashWithNumber(blockNr)
		bnh.Hash()
	})
}

func Fuzz_Nosy_BlockNumberOrHash_Number__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockNr BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}

		bnh := BlockNumberOrHashWithNumber(blockNr)
		bnh.Number()
	})
}

func Fuzz_Nosy_BlockNumberOrHash_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockNr BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}

		bnh := BlockNumberOrHashWithNumber(blockNr)
		bnh.String()
	})
}

func Fuzz_Nosy_BlockNumberOrHash_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blockNr BlockNumber
		fill_err = tp.Fill(&blockNr)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		bnh := BlockNumberOrHashWithNumber(blockNr)
		bnh.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_Client_BatchCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var b []BatchElem
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		c, err := DialIO(ctx, in, out)
		if err != nil {
			return
		}
		c.BatchCall(b)
	})
}

func Fuzz_Nosy_Client_BatchCallContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var b []BatchElem
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		c, err := DialIO(c1, in, out)
		if err != nil {
			return
		}
		c.BatchCallContext(c4, b)
	})
}

// skipping Fuzz_Nosy_Client_Call__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Client_CallContext__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Client_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		c, err := DialIO(ctx, in, out)
		if err != nil {
			return
		}
		c.Close()
	})
}

// skipping Fuzz_Nosy_Client_EthSubscribe__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Client_Notify__ because parameters include func, chan, or unsupported interface: []interface{}

// skipping Fuzz_Nosy_Client_RegisterName__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Client_SetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		c, err := DialIO(ctx, in, out)
		if err != nil {
			return
		}
		c.SetHeader(key, value)
	})
}

// skipping Fuzz_Nosy_Client_ShhSubscribe__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Client_Subscribe__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Client_SupportedModules__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		c, err := DialIO(ctx, in, out)
		if err != nil {
			return
		}
		c.SupportedModules()
	})
}

func Fuzz_Nosy_Client_SupportsSubscriptions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		c, err := DialIO(ctx, in, out)
		if err != nil {
			return
		}
		c.SupportsSubscriptions()
	})
}

// skipping Fuzz_Nosy_Client_dispatch__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.ServerCodec

func Fuzz_Nosy_Client_drainRead__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		c, err := DialIO(ctx, in, out)
		if err != nil {
			return
		}
		c.drainRead()
	})
}

// skipping Fuzz_Nosy_Client_newClientConn__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.ServerCodec

// skipping Fuzz_Nosy_Client_newMessage__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_Client_nextID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}

		c, err := DialIO(ctx, in, out)
		if err != nil {
			return
		}
		c.nextID()
	})
}

// skipping Fuzz_Nosy_Client_read__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.ServerCodec

func Fuzz_Nosy_Client_reconnect__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}

		c, err := DialIO(c1, in, out)
		if err != nil {
			return
		}
		c.reconnect(c4)
	})
}

// skipping Fuzz_Nosy_Client_send__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Client_sendBatchHTTP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var in io.Reader
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}
		var out io.Writer
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		var c4 context.Context
		fill_err = tp.Fill(&c4)
		if fill_err != nil {
			return
		}
		var op *requestOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var msgs []*jsonrpcMessage
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}
		if op == nil {
			return
		}

		c, err := DialIO(c1, in, out)
		if err != nil {
			return
		}
		c.sendBatchHTTP(c4, op, msgs)
	})
}

// skipping Fuzz_Nosy_Client_sendHTTP__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Client_write__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_ClientSubscription_Err__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var channel reflect.Value
		fill_err = tp.Fill(&channel)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		sub := newClientSubscription(c, namespace, channel)
		sub.Err()
	})
}

func Fuzz_Nosy_ClientSubscription_Unsubscribe__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var channel reflect.Value
		fill_err = tp.Fill(&channel)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		sub := newClientSubscription(c, namespace, channel)
		sub.Unsubscribe()
	})
}

// skipping Fuzz_Nosy_ClientSubscription_close__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_ClientSubscription_deliver__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var channel reflect.Value
		fill_err = tp.Fill(&channel)
		if fill_err != nil {
			return
		}
		var result json.RawMessage
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		sub := newClientSubscription(c, namespace, channel)
		sub.deliver(result)
	})
}

func Fuzz_Nosy_ClientSubscription_forward__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var channel reflect.Value
		fill_err = tp.Fill(&channel)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		sub := newClientSubscription(c, namespace, channel)
		sub.forward()
	})
}

func Fuzz_Nosy_ClientSubscription_requestUnsubscribe__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var channel reflect.Value
		fill_err = tp.Fill(&channel)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		sub := newClientSubscription(c, namespace, channel)
		sub.requestUnsubscribe()
	})
}

func Fuzz_Nosy_ClientSubscription_run__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var channel reflect.Value
		fill_err = tp.Fill(&channel)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		sub := newClientSubscription(c, namespace, channel)
		sub.run()
	})
}

func Fuzz_Nosy_ClientSubscription_unmarshal__(f *testing.F) {
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
		var namespace string
		fill_err = tp.Fill(&namespace)
		if fill_err != nil {
			return
		}
		var channel reflect.Value
		fill_err = tp.Fill(&channel)
		if fill_err != nil {
			return
		}
		var result json.RawMessage
		fill_err = tp.Fill(&result)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		sub := newClientSubscription(c, namespace, channel)
		sub.unmarshal(result)
	})
}

func Fuzz_Nosy_Notifier_Closed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Notifier
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.Closed()
	})
}

func Fuzz_Nosy_Notifier_CreateSubscription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Notifier
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.CreateSubscription()
	})
}

// skipping Fuzz_Nosy_Notifier_Notify__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_Notifier_activate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Notifier
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.activate()
	})
}

// skipping Fuzz_Nosy_Notifier_send__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_Notifier_takeSubscription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var n *Notifier
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if n == nil {
			return
		}

		n.takeSubscription()
	})
}

func Fuzz_Nosy_RPCService_Modules__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *RPCService
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Modules()
	})
}

// skipping Fuzz_Nosy_Server_RegisterName__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Server_ServeCodec__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.ServerCodec

// skipping Fuzz_Nosy_Server_ServeHTTP__ because parameters include func, chan, or unsupported interface: net/http.ResponseWriter

// skipping Fuzz_Nosy_Server_ServeListener__ because parameters include func, chan, or unsupported interface: net.Listener

func Fuzz_Nosy_Server_SetBatchLimits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, itemLimit int, maxResponseSize int) {
		s := NewServer()
		s.SetBatchLimits(itemLimit, maxResponseSize)
	})
}

func Fuzz_Nosy_Server_WebsocketHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedOrigins []string
		fill_err = tp.Fill(&allowedOrigins)
		if fill_err != nil {
			return
		}

		s := NewServer()
		s.WebsocketHandler(allowedOrigins)
	})
}

// skipping Fuzz_Nosy_Server_serveSingleRequest__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.ServerCodec

// skipping Fuzz_Nosy_Server_trackCodec__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.ServerCodec

// skipping Fuzz_Nosy_Server_untrackCodec__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.ServerCodec

func Fuzz_Nosy_Subscription_Err__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Subscription
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.Err()
	})
}

func Fuzz_Nosy_Subscription_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Subscription
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarshalJSON()
	})
}

// skipping Fuzz_Nosy_batchCallBuffer_doWrite__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.jsonWriter

func Fuzz_Nosy_batchCallBuffer_nextCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *batchCallBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.nextCall()
	})
}

func Fuzz_Nosy_batchCallBuffer_pushResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *batchCallBuffer
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var answer *jsonrpcMessage
		fill_err = tp.Fill(&answer)
		if fill_err != nil {
			return
		}
		if b == nil || answer == nil {
			return
		}

		b.pushResponse(answer)
	})
}

// skipping Fuzz_Nosy_batchCallBuffer_respondWithError__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.jsonWriter

// skipping Fuzz_Nosy_batchCallBuffer_write__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.jsonWriter

func Fuzz_Nosy_callback_call__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var receiver reflect.Value
		fill_err = tp.Fill(&receiver)
		if fill_err != nil {
			return
		}
		var fn reflect.Value
		fill_err = tp.Fill(&fn)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		var args []reflect.Value
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}

		c := newCallback(receiver, fn)
		c.call(ctx, method, args)
	})
}

func Fuzz_Nosy_callback_makeArgTypes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var receiver reflect.Value
		fill_err = tp.Fill(&receiver)
		if fill_err != nil {
			return
		}
		var fn reflect.Value
		fill_err = tp.Fill(&fn)
		if fill_err != nil {
			return
		}

		c := newCallback(receiver, fn)
		c.makeArgTypes()
	})
}

func Fuzz_Nosy_clientConfig_initHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *clientConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.initHeaders()
	})
}

func Fuzz_Nosy_clientConfig_setHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *clientConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		cfg.setHeader(key, value)
	})
}

// skipping Fuzz_Nosy_clientConn_close__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_handler_addRequestOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var op *requestOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		if h == nil || op == nil {
			return
		}

		h.addRequestOp(op)
	})
}

func Fuzz_Nosy_handler_addSubscriptions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var nn []*Notifier
		fill_err = tp.Fill(&nn)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.addSubscriptions(nn)
	})
}

// skipping Fuzz_Nosy_handler_cancelAllRequests__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_handler_cancelServerSubscriptions__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_handler_close__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_handler_handleBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var msgs []*jsonrpcMessage
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.handleBatch(msgs)
	})
}

func Fuzz_Nosy_handler_handleCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var cp *callProc
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if h == nil || cp == nil || msg == nil {
			return
		}

		h.handleCall(cp, msg)
	})
}

func Fuzz_Nosy_handler_handleCallMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx *callProc
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if h == nil || ctx == nil || msg == nil {
			return
		}

		h.handleCallMsg(ctx, msg)
	})
}

func Fuzz_Nosy_handler_handleMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if h == nil || msg == nil {
			return
		}

		h.handleMsg(msg)
	})
}

func Fuzz_Nosy_handler_handleNonBatchCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var cp *callProc
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if h == nil || cp == nil || msg == nil {
			return
		}

		h.handleNonBatchCall(cp, msg)
	})
}

// skipping Fuzz_Nosy_handler_handleResponses__ because parameters include func, chan, or unsupported interface: func(*github.com/ethereum/go-ethereum/rpc.jsonrpcMessage)

func Fuzz_Nosy_handler_handleSubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var cp *callProc
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if h == nil || cp == nil || msg == nil {
			return
		}

		h.handleSubscribe(cp, msg)
	})
}

func Fuzz_Nosy_handler_handleSubscriptionResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if h == nil || msg == nil {
			return
		}

		h.handleSubscriptionResult(msg)
	})
}

func Fuzz_Nosy_handler_removeRequestOp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var op *requestOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		if h == nil || op == nil {
			return
		}

		h.removeRequestOp(op)
	})
}

func Fuzz_Nosy_handler_respondWithBatchTooLarge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var cp *callProc
		fill_err = tp.Fill(&cp)
		if fill_err != nil {
			return
		}
		var batch []*jsonrpcMessage
		fill_err = tp.Fill(&batch)
		if fill_err != nil {
			return
		}
		if h == nil || cp == nil {
			return
		}

		h.respondWithBatchTooLarge(cp, batch)
	})
}

func Fuzz_Nosy_handler_runMethod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var callb *callback
		fill_err = tp.Fill(&callb)
		if fill_err != nil {
			return
		}
		var args []reflect.Value
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if h == nil || msg == nil || callb == nil {
			return
		}

		h.runMethod(ctx, msg, callb, args)
	})
}

// skipping Fuzz_Nosy_handler_startCallProc__ because parameters include func, chan, or unsupported interface: func(*github.com/ethereum/go-ethereum/rpc.callProc)

func Fuzz_Nosy_handler_unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *handler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id ID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.unsubscribe(ctx, id)
	})
}

func Fuzz_Nosy_httpConn_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *httpConn
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.close()
	})
}

func Fuzz_Nosy_httpConn_closed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *httpConn
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.closed()
	})
}

// skipping Fuzz_Nosy_httpConn_doRequest__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_httpConn_peerInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *httpConn
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.peerInfo()
	})
}

func Fuzz_Nosy_httpConn_readBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *httpConn
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.readBatch()
	})
}

func Fuzz_Nosy_httpConn_remoteAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hc *httpConn
		fill_err = tp.Fill(&hc)
		if fill_err != nil {
			return
		}
		if hc == nil {
			return
		}

		hc.remoteAddr()
	})
}

// skipping Fuzz_Nosy_httpConn_writeJSON__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_httpServerConn_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *httpServerConn
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

func Fuzz_Nosy_httpServerConn_RemoteAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *httpServerConn
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.RemoteAddr()
	})
}

func Fuzz_Nosy_httpServerConn_SetWriteDeadline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *httpServerConn
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var _x2 time.Time
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.SetWriteDeadline(_x2)
	})
}

func Fuzz_Nosy_internalServerError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *internalServerError
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

func Fuzz_Nosy_internalServerError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *internalServerError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorCode()
	})
}

func Fuzz_Nosy_invalidMessageError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *invalidMessageError
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

func Fuzz_Nosy_invalidMessageError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *invalidMessageError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorCode()
	})
}

func Fuzz_Nosy_invalidParamsError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *invalidParamsError
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

func Fuzz_Nosy_invalidParamsError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *invalidParamsError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorCode()
	})
}

func Fuzz_Nosy_invalidRequestError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *invalidRequestError
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

func Fuzz_Nosy_invalidRequestError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *invalidRequestError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorCode()
	})
}

func Fuzz_Nosy_jsonCodec_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *jsonCodec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.close()
	})
}

func Fuzz_Nosy_jsonCodec_closed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *jsonCodec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.closed()
	})
}

func Fuzz_Nosy_jsonCodec_peerInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *jsonCodec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.peerInfo()
	})
}

func Fuzz_Nosy_jsonCodec_readBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *jsonCodec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.readBatch()
	})
}

func Fuzz_Nosy_jsonCodec_remoteAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *jsonCodec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.remoteAddr()
	})
}

// skipping Fuzz_Nosy_jsonCodec_writeJSON__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_jsonError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *jsonError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_jsonError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *jsonError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.ErrorCode()
	})
}

func Fuzz_Nosy_jsonError_ErrorData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err *jsonError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}
		if err == nil {
			return
		}

		err.ErrorData()
	})
}

func Fuzz_Nosy_jsonrpcMessage_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.String()
	})
}

// skipping Fuzz_Nosy_jsonrpcMessage_errorResponse__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_jsonrpcMessage_hasValidID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.hasValidID()
	})
}

func Fuzz_Nosy_jsonrpcMessage_hasValidVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.hasValidVersion()
	})
}

func Fuzz_Nosy_jsonrpcMessage_isCall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.isCall()
	})
}

func Fuzz_Nosy_jsonrpcMessage_isNotification__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.isNotification()
	})
}

func Fuzz_Nosy_jsonrpcMessage_isResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.isResponse()
	})
}

func Fuzz_Nosy_jsonrpcMessage_isSubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.isSubscribe()
	})
}

func Fuzz_Nosy_jsonrpcMessage_isUnsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.isUnsubscribe()
	})
}

func Fuzz_Nosy_jsonrpcMessage_namespace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *jsonrpcMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		msg.namespace()
	})
}

// skipping Fuzz_Nosy_jsonrpcMessage_response__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_methodNotFoundError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *methodNotFoundError
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

func Fuzz_Nosy_methodNotFoundError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *methodNotFoundError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorCode()
	})
}

func Fuzz_Nosy_parseError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *parseError
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

func Fuzz_Nosy_parseError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *parseError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorCode()
	})
}

func Fuzz_Nosy_requestOp_wait__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var op *requestOp
		fill_err = tp.Fill(&op)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var c *Client
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if op == nil || c == nil {
			return
		}

		op.wait(ctx, c)
	})
}

func Fuzz_Nosy_serviceRegistry_callback__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *serviceRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.callback(method)
	})
}

// skipping Fuzz_Nosy_serviceRegistry_registerName__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_serviceRegistry_subscription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *serviceRegistry
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var service string
		fill_err = tp.Fill(&service)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.subscription(service, name)
	})
}

func Fuzz_Nosy_subscriptionNotFoundError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *subscriptionNotFoundError
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

func Fuzz_Nosy_subscriptionNotFoundError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *subscriptionNotFoundError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.ErrorCode()
	})
}

func Fuzz_Nosy_websocketCodec_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wc *websocketCodec
		fill_err = tp.Fill(&wc)
		if fill_err != nil {
			return
		}
		if wc == nil {
			return
		}

		wc.close()
	})
}

func Fuzz_Nosy_websocketCodec_peerInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wc *websocketCodec
		fill_err = tp.Fill(&wc)
		if fill_err != nil {
			return
		}
		if wc == nil {
			return
		}

		wc.peerInfo()
	})
}

func Fuzz_Nosy_websocketCodec_pingLoop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var wc *websocketCodec
		fill_err = tp.Fill(&wc)
		if fill_err != nil {
			return
		}
		if wc == nil {
			return
		}

		wc.pingLoop()
	})
}

// skipping Fuzz_Nosy_websocketCodec_writeJSON__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_BlockNumber_Int64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bn BlockNumber
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}

		bn.Int64()
	})
}

func Fuzz_Nosy_BlockNumber_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bn BlockNumber
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}

		bn.MarshalText()
	})
}

func Fuzz_Nosy_BlockNumber_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bn BlockNumber
		fill_err = tp.Fill(&bn)
		if fill_err != nil {
			return
		}

		bn.String()
	})
}

func Fuzz_Nosy_ClientOption_applyOption__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value string
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		var _x3 *clientConfig
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x3 == nil {
			return
		}

		_x1 := WithHeader(key, value)
		_x1.applyOption(_x3)
	})
}

// skipping Fuzz_Nosy_Conn_SetWriteDeadline__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.Conn

// skipping Fuzz_Nosy_ConnRemoteAddr_RemoteAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.ConnRemoteAddr

// skipping Fuzz_Nosy_DataError_Error__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.DataError

// skipping Fuzz_Nosy_DataError_ErrorData__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.DataError

// skipping Fuzz_Nosy_Error_Error__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.Error

// skipping Fuzz_Nosy_Error_ErrorCode__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.Error

func Fuzz_Nosy_HTTPError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err HTTPError
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_ServerCodec_close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conn *websocket.Conn
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var host string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var req http.Header
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var readLimit int64
		fill_err = tp.Fill(&readLimit)
		if fill_err != nil {
			return
		}
		if conn == nil {
			return
		}

		_x1 := newWebsocketCodec(conn, host, req, readLimit)
		_x1.close()
	})
}

func Fuzz_Nosy_ServerCodec_peerInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conn *websocket.Conn
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var host string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var req http.Header
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var readLimit int64
		fill_err = tp.Fill(&readLimit)
		if fill_err != nil {
			return
		}
		if conn == nil {
			return
		}

		_x1 := newWebsocketCodec(conn, host, req, readLimit)
		_x1.peerInfo()
	})
}

func Fuzz_Nosy_ServerCodec_readBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var conn *websocket.Conn
		fill_err = tp.Fill(&conn)
		if fill_err != nil {
			return
		}
		var host string
		fill_err = tp.Fill(&host)
		if fill_err != nil {
			return
		}
		var req http.Header
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var readLimit int64
		fill_err = tp.Fill(&readLimit)
		if fill_err != nil {
			return
		}
		if conn == nil {
			return
		}

		_x1 := newWebsocketCodec(conn, host, req, readLimit)
		_x1.readBatch()
	})
}

// skipping Fuzz_Nosy_deadlineCloser_SetWriteDeadline__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.deadlineCloser

func Fuzz_Nosy_idForLog_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var id idForLog
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		id.String()
	})
}

// skipping Fuzz_Nosy_jsonWriter_closed__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.jsonWriter

// skipping Fuzz_Nosy_jsonWriter_remoteAddr__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.jsonWriter

// skipping Fuzz_Nosy_jsonWriter_writeJSON__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.jsonWriter

func Fuzz_Nosy_notificationsUnsupportedError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e notificationsUnsupportedError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_notificationsUnsupportedError_ErrorCode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e notificationsUnsupportedError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.ErrorCode()
	})
}

// skipping Fuzz_Nosy_notificationsUnsupportedError_Is__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_optionFunc_applyOption__ because parameters include func, chan, or unsupported interface: github.com/ethereum/go-ethereum/rpc.optionFunc

func Fuzz_Nosy_stdioConn_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var io stdioConn
		fill_err = tp.Fill(&io)
		if fill_err != nil {
			return
		}

		io.Close()
	})
}

func Fuzz_Nosy_stdioConn_Read__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var io stdioConn
		fill_err = tp.Fill(&io)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		io.Read(b)
	})
}

func Fuzz_Nosy_stdioConn_RemoteAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var io stdioConn
		fill_err = tp.Fill(&io)
		if fill_err != nil {
			return
		}

		io.RemoteAddr()
	})
}

func Fuzz_Nosy_stdioConn_SetWriteDeadline__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var io stdioConn
		fill_err = tp.Fill(&io)
		if fill_err != nil {
			return
		}
		var t2 time.Time
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}

		io.SetWriteDeadline(t2)
	})
}

func Fuzz_Nosy_stdioConn_Write__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var io stdioConn
		fill_err = tp.Fill(&io)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		io.Write(b)
	})
}

func Fuzz_Nosy_wsHandshakeError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e wsHandshakeError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_ClientFromContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		ClientFromContext(ctx)
	})
}

func Fuzz_Nosy_ContextRequestTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		ContextRequestTimeout(ctx)
	})
}

func Fuzz_Nosy_NotifierFromContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		NotifierFromContext(ctx)
	})
}

func Fuzz_Nosy_StartIPCEndpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ipcEndpoint string
		fill_err = tp.Fill(&ipcEndpoint)
		if fill_err != nil {
			return
		}
		var apis []API
		fill_err = tp.Fill(&apis)
		if fill_err != nil {
			return
		}

		StartIPCEndpoint(ipcEndpoint, apis)
	})
}

func Fuzz_Nosy_formatName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, name string) {
		formatName(name)
	})
}

func Fuzz_Nosy_isBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var raw json.RawMessage
		fill_err = tp.Fill(&raw)
		if fill_err != nil {
			return
		}

		isBatch(raw)
	})
}

// skipping Fuzz_Nosy_isErrorType__ because parameters include func, chan, or unsupported interface: reflect.Type

// skipping Fuzz_Nosy_isPubSub__ because parameters include func, chan, or unsupported interface: reflect.Type

// skipping Fuzz_Nosy_isSubscriptionType__ because parameters include func, chan, or unsupported interface: reflect.Type

// skipping Fuzz_Nosy_originIsAllowed__ because parameters include func, chan, or unsupported interface: github.com/deckarep/golang-set/v2.Set[string]

// skipping Fuzz_Nosy_parseArgumentArray__ because parameters include func, chan, or unsupported interface: []reflect.Type

func Fuzz_Nosy_parseMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var raw json.RawMessage
		fill_err = tp.Fill(&raw)
		if fill_err != nil {
			return
		}

		parseMessage(raw)
	})
}

func Fuzz_Nosy_parseOriginURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, origin string) {
		parseOriginURL(origin)
	})
}

// skipping Fuzz_Nosy_parsePositionalArguments__ because parameters include func, chan, or unsupported interface: []reflect.Type

func Fuzz_Nosy_parseSubscriptionName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rawArgs json.RawMessage
		fill_err = tp.Fill(&rawArgs)
		if fill_err != nil {
			return
		}

		parseSubscriptionName(rawArgs)
	})
}

func Fuzz_Nosy_ruleAllowsOrigin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, allowedOrigin string, browserOrigin string) {
		ruleAllowsOrigin(allowedOrigin, browserOrigin)
	})
}

func Fuzz_Nosy_suitableCallbacks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var receiver reflect.Value
		fill_err = tp.Fill(&receiver)
		if fill_err != nil {
			return
		}

		suitableCallbacks(receiver)
	})
}

func Fuzz_Nosy_updateServeTimeHistogram__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		var success bool
		fill_err = tp.Fill(&success)
		if fill_err != nil {
			return
		}
		var elapsed time.Duration
		fill_err = tp.Fill(&elapsed)
		if fill_err != nil {
			return
		}

		updateServeTimeHistogram(method, success, elapsed)
	})
}

func Fuzz_Nosy_validateRequest__(f *testing.F) {
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

		validateRequest(r)
	})
}

func Fuzz_Nosy_wsClientHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, endpoint string, origin string) {
		wsClientHeaders(endpoint, origin)
	})
}

func Fuzz_Nosy_wsHandshakeValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedOrigins []string
		fill_err = tp.Fill(&allowedOrigins)
		if fill_err != nil {
			return
		}

		wsHandshakeValidator(allowedOrigins)
	})
}
