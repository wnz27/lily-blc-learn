package v2

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	logging "github.com/ipfs/go-log/v2"

	"github.com/filecoin-project/lily/chain/indexer"
	"github.com/filecoin-project/lily/chain/indexer/v2/load"
	"github.com/filecoin-project/lily/chain/indexer/v2/load/persistable"
	"github.com/filecoin-project/lily/chain/indexer/v2/transform"
	"github.com/filecoin-project/lily/model"
	"github.com/filecoin-project/lily/tasks"
)

var log = logging.Logger("indexmanager")

type Manager struct {
	indexer *TipSetIndexer
	stuff   *ThingIDK
	api     tasks.DataSource
	strg    model.Storage
}

func NewIndexManager(strg model.Storage, api tasks.DataSource, tasks []string) (*Manager, error) {
	stuff, err := GetTransformersForTasks(tasks...)
	if err != nil {
		return nil, err
	}
	return &Manager{
		indexer: NewTipSetIndexer(api, stuff.Tasks, 1024),
		stuff:   stuff,
		api:     api,
		strg:    strg,
	}, nil
}

func (m *Manager) TipSet(ctx context.Context, ts *types.TipSet, options ...indexer.Option) (bool, error) {
	start := time.Now()
	results, err := m.indexer.TipSet(ctx, ts)
	if err != nil {
		return false, err
	}

	transformer, consumer, err := m.startRouters(ctx,
		m.stuff.Transformers,
		[]load.Handler{&persistable.PersistableResultConsumer{Strg: m.strg}},
	)
	if err != nil {
		return false, err
	}

	// TODO handle the error case here, remove the panic in the goroutine
	// - a simple solution would be to collect all transformer results and then send them to the consumer.
	//	 this will prevent partial persistence at the cost of more memory.
	go func() {
		for res := range results {
			if len(res.State().Data) > 0 {
				if err := transformer.Route(ctx, res); err != nil {
					panic(err)
				}
			}
		}
		if err := transformer.Stop(); err != nil {
			panic(err)
		}
	}()
	for res := range transformer.Results() {
		if err := consumer.Route(ctx, res); err != nil {
			return false, err
		}
	}
	if err := consumer.Stop(); err != nil {
		return false, err
	}
	log.Infow("index complete", "duration", time.Since(start))
	return true, nil
}

type Transformer interface {
	Route(ctx context.Context, data transform.IndexState) error
	Results() chan transform.Result
	Stop() error
}

type Loader interface {
	Route(ctx context.Context, data transform.Result) error
	Stop() error
}

func (m *Manager) startRouters(ctx context.Context, handlers []transform.Handler, consumers []load.Handler) (Transformer, Loader, error) {
	tr, err := transform.NewRouter(m.stuff.Tasks, handlers...)
	if err != nil {
		return nil, nil, err
	}
	tr.Start(ctx, m.api)

	lr, err := load.NewRouter(consumers...)
	if err != nil {
		return nil, nil, err
	}
	lr.Start(ctx)

	return tr, lr, nil
}