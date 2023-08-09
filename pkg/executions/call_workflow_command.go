package executions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/parodos-dev/parodos-service/pkg/workflows"
)

type CallWorkflowCommand struct {
	WorkflowName      string
	WorkflowNamespace string
	Parameters        map[string]interface{}
}

func (c *CallWorkflowCommand) GetParamsPayload() ([]byte, error) {
	return json.Marshal(c.Parameters)
}

func NewCallWorkflowCommand(workflowName, workflowNamespace string, parameters map[string]interface{}) *CallWorkflowCommand {
	return &CallWorkflowCommand{
		WorkflowName:      workflowName,
		WorkflowNamespace: workflowNamespace,
		Parameters:        parameters,
	}
}

type callWorkflowHandlerCommand struct {
	workflowService workflows.WorkflowService
	httpClient      *http.Client
}

func NewCallWorkflowHandlerCommand(workflowService workflows.WorkflowService) *callWorkflowHandlerCommand {
	return &callWorkflowHandlerCommand{
		workflowService: workflowService,
		httpClient:      &http.Client{},
	}
}

func (h *callWorkflowHandlerCommand) Execute(ctx context.Context, command *CallWorkflowCommand) (*http.Response, error) {

	wf, err := h.workflowService.GetByName(command.WorkflowNamespace, command.WorkflowName)
	if err != nil {
		return nil, err
	}

	if err = wf.ValidateInputData(command.Parameters); err != nil {
		return nil, err
	}

	host, err := wf.GetHost()
	if err != nil {
		return nil, fmt.Errorf("Cannot get the correct host for the workflow: %v", err)
	}

	payload, err := command.GetParamsPayload()
	if err != nil {
		return nil, fmt.Errorf("Cannot validate the payload for the workflow: %v", err)

	}

	return h.sendRequest(host, payload)
}

func (h *callWorkflowHandlerCommand) sendRequest(host *url.URL, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, host.String(), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return h.httpClient.Do(req)
}
