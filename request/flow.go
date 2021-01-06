package request

type CreateFlow struct{
	FlowName     string    `json:"flow_name"`
	FlowType     int       `json:"flow_type"`
	Description  string    `json:"description"`
}

type UpdateFlow struct{
	FlowId       int       `json:"flow_id"`
	FlowName     string    `json:"flow_name"`
	FlowType     int       `json:"flow_type"`
	Description  string    `json:"description"`
}

type CreateFlowInstance struct{
	FlowId       int       `json:"flow_id"`
	FlowInfo     string    `json:"flow_info"`
	Applicant    string    `json:"applicant"`
	ApplyReason  string    `json:"apply_reason"`
}

type UpdateFlowInstance struct{
	InstanceId   int `json:"instance_id"`
	FlowInfo     string    `json:"flow_info"`
	ApplyReason  string    `json:"apply_reason"`
}