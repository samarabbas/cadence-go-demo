package workflows

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"testing"

	_ "github.com/samarabbas/cadence-go-demo/activities"
	"github.com/stretchr/testify/assert"
	"go.uber.org/cadence/testsuite"
)

func TestWorkflowWithRealActivities(t *testing.T) {
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()
	request := AccountTransferRequest{
		FromAccountId: "account1",
		ToAccountId:   "account2",
		ReferenceId:   "reference1",
		Amount:        1000,
	}
	env.ExecuteWorkflow(transferWorkflow, request)
	assert.NoError(t, env.GetWorkflowError())
}

func TestWorkflowWithdrawFailure(t *testing.T) {
	suite := testsuite.WorkflowTestSuite{}
	env := suite.NewTestWorkflowEnvironment()
	env.OnActivity("withdraw", mock.Anything, "account1", "reference1", 1000).
		Return(errors.New("simulated failure")).
		Times(100)
	request := AccountTransferRequest{
		FromAccountId: "account1",
		ToAccountId:   "account2",
		ReferenceId:   "reference1",
		Amount:        1000,
	}
	env.ExecuteWorkflow(transferWorkflow, request)
	assert.Error(t, env.GetWorkflowError())
	assert.Equal(t, "simulated failure", env.GetWorkflowError().Error())
}
