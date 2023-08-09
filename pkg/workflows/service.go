package workflows

type WorkflowService interface {
	GetByName(namespace string, name string) (*Workflow, error)
}

type workflowServiceImpl struct{}

func NewWorkflowServiceImpl() *workflowServiceImpl {
	return &workflowServiceImpl{}
}

func (s *workflowServiceImpl) GetByName(namespace string, name string) (*Workflow, error) {
	return nil, nil
}
