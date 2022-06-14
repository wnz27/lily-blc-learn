package vm

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"golang.org/x/sync/errgroup"

	"github.com/filecoin-project/lily/lens"
	"github.com/filecoin-project/lily/lens/util"
	"github.com/filecoin-project/lily/model"
	messagemodel "github.com/filecoin-project/lily/model/messages"
	visormodel "github.com/filecoin-project/lily/model/visor"
	tasks "github.com/filecoin-project/lily/tasks"
	messages "github.com/filecoin-project/lily/tasks/messages"
)

type Task struct {
	node tasks.DataSource
}

func NewTask(node tasks.DataSource) *Task {
	return &Task{node: node}
}

func (t *Task) ProcessTipSets(ctx context.Context, current *types.TipSet, executed *types.TipSet) (model.Persistable, *visormodel.ProcessingReport, error) {
	ctx, span := otel.Tracer("").Start(ctx, "ProcessTipSets")
	if span.IsRecording() {
		span.SetAttributes(
			attribute.String("current", current.String()),
			attribute.Int64("current_height", int64(current.Height())),
			attribute.String("executed", executed.String()),
			attribute.Int64("executed_height", int64(executed.Height())),
			attribute.String("processor", "internal_parsed_message"),
		)
	}
	defer span.End()

	// execute in parallel as both operations are slow
	grp, _ := errgroup.WithContext(ctx)
	var mex []*lens.MessageExecution
	grp.Go(func() error {
		var err error
		mex, err = t.node.MessageExecutions(ctx, current, executed)
		if err != nil {
			return fmt.Errorf("getting messages executions for tipset: %w", err)
		}
		return nil
	})

	var getActorCode func(a address.Address) (cid.Cid, bool)
	grp.Go(func() error {
		var err error
		getActorCode, err = util.MakeGetActorCodeFunc(ctx, t.node.Store(), current, executed)
		if err != nil {
			return fmt.Errorf("failed to make actor code query function: %w", err)
		}
		return nil
	})

	report := &visormodel.ProcessingReport{
		Height:    int64(current.Height()),
		StateRoot: current.ParentState().String(),
	}

	// if either fail, report error and bail
	if err := grp.Wait(); err != nil {
		report.ErrorsDetected = err
		return nil, report, nil
	}

	var (
		vmMessageResults = make(messagemodel.VmMessageList, 0, len(mex))
		errorsDetected   = make([]*messages.MessageError, 0)
	)
	for _, parent := range mex {
		select {
		case <-ctx.Done():
			return nil, nil, fmt.Errorf("context done: %w", ctx.Err())
		default:
		}

		for _, child := range messages.GetChildMessagesOf(parent) {
			// Cid() computes a CID, so only call it once
			childCid := child.Message.Cid()
			toCode, ok := getActorCode(child.Message.To)
			if !ok {
				errorsDetected = append(errorsDetected, &messages.MessageError{
					Cid:   parent.Cid,
					Error: fmt.Errorf("failed to get to actor code for message: %s", childCid).Error(),
				})
				continue
			}
			toName, _, err := util.ActorNameAndFamilyFromCode(parent.ToActorCode)
			if err != nil {
				// TODO what do we do if there is an error? Continue with unknown family names or abort?
				errorsDetected = append(errorsDetected, &messages.MessageError{
					Cid:   parent.Cid,
					Error: fmt.Errorf("failed get message (%s) to actor name and family: %w", parent.Cid, err).Error(),
				})
				continue
			}
			meta, err := util.MethodParamsReturnForMessage(child, toCode)
			if err != nil {
				errorsDetected = append(errorsDetected, &messages.MessageError{
					Cid:   parent.Cid,
					Error: fmt.Errorf("failed get child message (%s) metadata: %w", childCid, err).Error(),
				})
				continue
			}
			vmMessageResults = append(vmMessageResults, &messagemodel.VmMessage{
				Height:    int64(parent.Height),
				StateRoot: parent.StateRoot.String(),
				Parent:    parent.Cid.String(),
				Cid:       childCid.String(),
				From:      child.Message.From.String(),
				To:        child.Message.To.String(),
				Value:     child.Message.Value.String(),
				GasUsed:   child.Receipt.GasUsed,
				ExitCode:  int64(child.Receipt.ExitCode),
				ActorName: toName,
				Method:    meta.MethodName,
				Params:    meta.Params,
				Return:    meta.Return,
			})
		}
	}

	if len(errorsDetected) != 0 {
		report.ErrorsDetected = errorsDetected
	}
	return vmMessageResults, report, nil
}
