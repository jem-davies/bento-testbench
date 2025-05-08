package processors

import (
	"context"
	"errors"

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
	return nil, errors.New("THIS IS AN ERROR")
}

func (mtp *mtp) Close(ctx context.Context) error {
	return nil
}
