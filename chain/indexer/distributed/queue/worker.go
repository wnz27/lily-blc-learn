package queue

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	logging "github.com/ipfs/go-log/v2"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/filecoin-project/lily/chain/indexer/distributed"
	"github.com/filecoin-project/lily/chain/indexer/distributed/queue/tasks"
	"github.com/filecoin-project/lily/metrics"
)

var log = logging.Logger("lily/distributed/worker")

type AsynqWorker struct {
	done chan struct{}

	name     string
	server   *distributed.TipSetWorker
	handlers []TaskHandler
}
type TaskHandler interface {
	Type() string
	Handler() asynq.HandlerFunc
}

func NewAsynqWorker(name string, server *distributed.TipSetWorker, handlers ...TaskHandler) *AsynqWorker {
	return &AsynqWorker{
		name:     name,
		server:   server,
		handlers: handlers,
	}
}

func (t *AsynqWorker) Run(ctx context.Context) error {
	t.done = make(chan struct{})
	defer close(t.done)

	mux := asynq.NewServeMux()
	for _, handler := range t.handlers {
		log.Infow("registered task handler", "type", handler.Type())
		mux.HandleFunc(handler.Type(), handler.Handler())
	}

	t.server.ServerConfig.Logger = log.With("name", t.name)
	t.server.ServerConfig.ErrorHandler = &WorkerErrorHandler{}

	stats.Record(ctx, metrics.TipSetWorkerConcurrency.M(int64(t.server.ServerConfig.Concurrency)))
	for queueName, priority := range t.server.ServerConfig.Queues {
		if err := stats.RecordWithTags(ctx,
			[]tag.Mutator{tag.Upsert(metrics.QueueName, queueName)},
			metrics.TipSetWorkerQueuePriority.M(int64(priority))); err != nil {
			return err
		}
	}

	server := asynq.NewServer(t.server.RedisConfig, t.server.ServerConfig)
	if err := server.Start(mux); err != nil {
		return err
	}
	<-ctx.Done()
	server.Shutdown()
	return nil
}

func (t *AsynqWorker) Done() <-chan struct{} {
	return t.done
}

type WorkerErrorHandler struct{}

func (w *WorkerErrorHandler) HandleError(ctx context.Context, task *asynq.Task, err error) {
	switch task.Type() {
	case tasks.TypeIndexTipSet:
		var p tasks.IndexTipSetPayload
		if err := json.Unmarshal(task.Payload(), &p); err != nil {
			log.Errorw("failed to decode task type (developer error?)", "error", err)
			return
		}
		if p.HasTraceCarrier() {
			if sc := p.TraceCarrier.AsSpanContext(); sc.IsValid() {
				ctx = trace.ContextWithRemoteSpanContext(ctx, sc)
				trace.SpanFromContext(ctx).RecordError(err)
			}
		}
		log.Errorw("task failed", zap.Inline(p), "type", task.Type(), "error", err)
	case tasks.TypeGapFillTipSet:
		var p tasks.GapFillTipSetPayload
		if err := json.Unmarshal(task.Payload(), &p); err != nil {
			log.Errorw("failed to decode task type (developer error?)", "error", err)
			return
		}
		if p.HasTraceCarrier() {
			if sc := p.TraceCarrier.AsSpanContext(); sc.IsValid() {
				ctx = trace.ContextWithRemoteSpanContext(ctx, sc)
				trace.SpanFromContext(ctx).RecordError(err)
			}
		}
		log.Errorw("task failed", zap.Inline(p), "type", task.Type(), "error", err)
	}
}
