package executions

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/parodos-dev/parodos-service/pkg/workflows"
)

type CallWorkflowCommand struct {
	WorkflowName      string
	WorkflowNamespace string
	Parameters        map[string]interface{}
}

func NewCallWorkflowCommand(workflowName, workflowNamespace string, parameters map[string]interface{}) *CallWorkflowCommand {
	return &CallWorkflowCommand{
		WorkflowName:      workflowName,
		WorkflowNamespace: workflowNamespace,
		Parameters:        parameters,
	}
}

type callWorkflowHandler struct {
	workflowService workflows.WorkflowService
}

func NewCallWorkflowHandler() *callWorkflowHandler {
	return &callWorkflowHandler{}
}

func (h *callWorkflowHandler) Execute(ctx context.Context, command CallWorkflowCommand) (*http.Response, error) {

	wf, err := h.workflowService.GetByName(command.WorkflowNamespace, command.WorkflowName)
	if err != nil {
		return nil, err
	}

	host, err := wf.GetHost()
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(command.Parameters)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, host.String(), bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
