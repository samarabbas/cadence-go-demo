package activities

import (
	"context"

	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

func init() {
	activity.RegisterWithOptions(deposit, activity.RegisterOptions{Name: "deposit"})
}

func deposit(ctx context.Context, accountId, referenceId string, amount int) error {
	logger := activity.GetLogger(ctx)
	logger.Info("deposit requested",
		zap.String("AccountId", accountId),
		zap.String("ReferenceId", referenceId),
		zap.Int("Amount", amount))

	/*err := errors.New("banking service is down")
	logger.Error("banking service is down",
		zap.Error(err),
		zap.String("AccountId", accountId),
		zap.String("ReferenceId", referenceId),
		zap.Int("Amount", amount))
	return err*/
	return nil
}
