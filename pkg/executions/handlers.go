package executions

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parodos-dev/parodos-service/pkg/workflows"
	"github.com/swaggo/swag/example/celler/httputil"
)

type executionHandler struct {
	callWorkflowHandler *callWorkflowHandlerCommand
}

func NewExecutionHandler(workflowService workflows.WorkflowService) *executionHandler {
	res := &executionHandler{
		callWorkflowHandler: NewCallWorkflowHandlerCommand(workflowService),
	}
	return res
}

// // @Summary 	Execute a workflow
// // @Description return the execution id.
// // @Accept 		json
// // @Produce  	json
// // @Success 	200 {array} Group "ok"
// // @Failure		400  {object}  httputil.HTTPError
// // @Failure		404  {object}  httputil.HTTPError
// // @Failure		500  {object}  httputil.HTTPError
// // @Router 		/execute [get]

func (handler *executionHandler) execute(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("body cannot be readed"))
		return
	}
	var params map[string]interface{}
	if err := json.Unmarshal(jsonData, params); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, fmt.Errorf("body cannot be readed"))
	}

	command := NewCallWorkflowCommand(
		ctx.Param("name"),
		ctx.Param("group"),
		params)

	res, err := handler.callWorkflowHandler.Execute(context.TODO(), command)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, fmt.Errorf("body cannot be readed"))
		return
	}

	ctx.JSON(200, res.Body)

}
