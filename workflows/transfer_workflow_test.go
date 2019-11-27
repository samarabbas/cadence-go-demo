package workflows

import (
	"testing"

	_ "github.com/samarabbas/cadence-go-demo/activities"
	"github.com/stretchr/testify/assert"
	"go.uber.org/cadence/testsuite"
)

func TestSuccess(t *testing.T) {
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
