package mr_smart

import (
	"encoding/binary"
	"github.com/ywengineer/g-util/util"
	"github.com/ywengineer/mr.smart/codec"
	"github.com/ywengineer/mr.smart/config"
	"testing"
)

func TestServer(t *testing.T) {
	srv, err := NewSmartServer(&config.ValueLoader{
		Conf: &config.Conf{Network: "tcp", Address: ":12345", Workers: 1, WorkerLoadBalance: 2},
	}, []ChannelInitializer{
		WithByteOrder(func() binary.ByteOrder {
			return binary.LittleEndian
		}),
		WithCodec(func() codec.Codec {
			return &codec.JSONCodec{}
		}),
	})
	if err != nil {
		t.Errorf("create smart server failed. %v", err)
		t.FailNow()
	}
	_, err = srv.Serve()
	if err != nil {
		t.Errorf("start smart server failed. %v", err)
		t.FailNow()
	}
	t.Log("smart server was started.")
	_ = <-util.WatchQuitSignal()

	t.Logf("smart server was stopped. %v", srv.Shutdown())
}