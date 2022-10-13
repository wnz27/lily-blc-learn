package extract

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/gammazero/workerpool"

	v2 "github.com/filecoin-project/lily/model/v2"
	"github.com/filecoin-project/lily/tasks"
)

type TipSetStateResult struct {
	Task      v2.ModelMeta
	StartTime time.Time
	Duration  time.Duration
	Models    []v2.LilyModel
	Error     *TipSetExtractorError
}

type TipSetExtractorError struct {
	Error error
}

func TipSetState(ctx context.Context, workers int, api tasks.DataSource, current, executed *types.TipSet, extractors map[v2.ModelMeta]v2.ExtractorFn, results chan *TipSetStateResult) error {
	pool := workerpool.New(workers)
	for task, extractor := range extractors {
		task := task
		extractor := extractor
		pool.Submit(func() {
			select {
			case <-ctx.Done():
				return
			default:
				start := time.Now()
				data, err := extractor(ctx, api, current, executed)
				log.Debugw("extracted model", "type", task.String(), "duration", time.Since(start))
				results <- &TipSetStateResult{
					Task:      task,
					StartTime: start,
					Duration:  time.Since(time.Now()),
					Models:    data,
					Error:     &TipSetExtractorError{Error: err},
				}
			}
		})
	}
	pool.StopWait()
	return nil
}
