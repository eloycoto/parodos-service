package executions

import (
	"github.com/gin-gonic/gin"

	"github.com/parodos-dev/parodos-service/pkg/workflows"
)

func InitExecutions(router *gin.Engine, workflowService workflows.WorkflowService) {
	// workflowDefinitionHandler := newWorkflowDefinitionHandler()
	registerRouter(router, NewExecutionHandler(workflowService))
}

func registerRouter(router *gin.Engine, handler *executionHandler) {
	v1 := router.Group("/v1/executions")
	executions := v1.Group("/:group_id/workflow/:workflow_id")

	executions.POST("", handler.execute)
}
