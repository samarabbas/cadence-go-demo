package workflows

import (
	"time"

	"go.uber.org/cadence"
	"go.uber.org/cadence/workflow"

	"github.com/samarabbas/cadence-go-demo/common"
)

func init() {
	workflow.RegisterWithOptions(transferWorkflow, workflow.RegisterOptions{Name: "transfer"})
}

type (
	AccountTransferRequest struct {
		FromAccountId string
		ToAccountId   string
		ReferenceId   string
		Amount        int
	}
)

func transferWorkflow(ctx workflow.Context, transferRequest AccountTransferRequest) error {
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

	err := workflow.ExecuteActivity(ctx, "withdraw",
		transferRequest.FromAccountId,
		transferRequest.ReferenceId,
		transferRequest.Amount).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, "deposit",
		transferRequest.ToAccountId,
		transferRequest.ReferenceId,
		transferRequest.Amount).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
