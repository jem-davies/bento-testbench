package processors

import (
	"context"

	"github.com/warpstreamlabs/bento/public/service"
)

func metaTestProcSpec() *service.ConfigSpec {
	return service.NewConfigSpec()
}

func init() {
	err := service.RegisterProcessor("meta_test", metaTestProcSpec(), func(conf *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {
		return newMTP(conf, mgr)
	})
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------

type mtp struct{}

func newMTP(conf *service.ParsedConfig, mgr *service.Resources) (*mtp, error) {
	return &mtp{}, nil
}

//------------------------------------------------------------------------------

func (mtp *mtp) Process(ctx context.Context, msg *service.Message) (service.MessageBatch, error) {
	m := msg.Copy()
	m.MetaSetMut("test", map[string]string{"key1": "value1", "key2": "value2"})
	return service.MessageBatch{m}, nil
}

func (mtp *mtp) Close(ctx context.Context) error {
	return nil
}
