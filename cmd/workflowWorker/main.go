package main

import (
	"github.com/uber-go/tally"
	"go.uber.org/cadence/worker"
	"go.uber.org/zap"

	"github.com/samarabbas/cadence-go-demo/common"

	_ "github.com/samarabbas/cadence-go-demo/workflows"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Zap logger created")
	scope := tally.NoopScope

	builder := common.NewBuilder(logger).
		SetDomain(common.Domain).
		SetHostPort(common.Host)
	service, err := builder.BuildServiceClient()
	if err != nil {
		panic(err)
	}

	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: scope,
		Logger:       logger,
	}

	worker := worker.New(service, common.Domain, common.WorkflowTaskList, workerOptions)
	worker.Start()
	// The workers are supposed to be long running process that should not exit.
	// Use select{} to block indefinitely for samples, you can quit by CMD+C.
	select {}
}
