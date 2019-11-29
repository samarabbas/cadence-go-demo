package workflows

import (
	"time"

	"go.uber.org/cadence"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"

	"github.com/samarabbas/cadence-go-demo/common"
)

func init() {
	workflow.RegisterWithOptions(batchTransferWorkflow, workflow.RegisterOptions{Name: "batch-transfer"})
}

type (
	WithdrawSignal struct {
		FromAccountId string
		ReferenceId   string
		Amount        int
	}

	BatchTransferRequest struct {
		ToAccountId string
		ReferenceId string
		BatchSize   int
	}

	batchState struct {
		references map[string]struct{} // used for deduping signals
		balance    int
		count      int
	}
)

func newBatchState() *batchState {
	state := &batchState{}
	state.references = make(map[string]struct{})

	return state
}

func (s *batchState) transfer(ctx workflow.Context, request BatchTransferRequest) error {
	logger := workflow.GetLogger(ctx)

	workflow.SetQueryHandler(ctx, "get-count", func() (int, error) {
		return s.count, nil
	})

	workflow.SetQueryHandler(ctx, "get-balance", func() (int, error) {
		return s.balance, nil
	})

	withdrawSignalCh := workflow.GetSignalChannel(ctx, "withdraw")

	for s.count < request.BatchSize {
		var withdrawSignal WithdrawSignal
		withdrawSignalCh.Receive(ctx, &withdrawSignal)
		logger.Info("withdraw signal received", zap.String("FromAccountId", withdrawSignal.FromAccountId))
		if _, ok := s.references[withdrawSignal.ReferenceId]; !ok {
			s.references[withdrawSignal.ReferenceId] = struct{}{}

			err := workflow.ExecuteActivity(ctx, "withdraw",
				withdrawSignal.FromAccountId,
				withdrawSignal.ReferenceId,
				withdrawSignal.Amount).Get(ctx, nil)

			if err != nil {
				return err
			}

			s.balance += withdrawSignal.Amount
			s.count++
		}
	}

	logger.Info("all withdrawals completed")

	err := workflow.ExecuteActivity(ctx, "deposit",
		request.ToAccountId,
		request.ReferenceId,
		s.balance).Get(ctx, nil)

	if err != nil {
		return err
	}

	logger.Info("deposit completed")

	return nil
}

func batchTransferWorkflow(ctx workflow.Context, request BatchTransferRequest) error {
	ao := workflow.ActivityOptions{
		TaskList:               common.ActivityTaskList,
		ScheduleToStartTimeout: 10 * time.Minute,
		StartToCloseTimeout:    5 * time.Second,
		RetryPolicy: &cadence.RetryPolicy{
			InitialInterval:    time.Second,
			MaximumInterval:    10 * time.Second,
			ExpirationInterval: 10 * time.Minute,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	logger.Info("batch transfer workflow started")

	state := newBatchState()
	return state.transfer(ctx, request)
}
