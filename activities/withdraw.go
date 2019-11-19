package activities

import (
	"context"

	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

func init() {
	activity.RegisterWithOptions(withdraw, activity.RegisterOptions{Name: "withdraw"})
}

func withdraw(ctx context.Context, accountId, referenceId string, amount int) error {
	logger := activity.GetLogger(ctx)
	logger.Info("withdrawal requested",
		zap.String("AccountId", accountId),
		zap.String("ReferenceId", referenceId),
		zap.Int("Amount", amount))

	return nil
}
