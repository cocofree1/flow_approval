package request

type CreateNode struct {
	FlowId       int       `json:"flow_id"`
	NodeName     string    `json:"node_name"`
	FlowPos      int       `json:"flow_pos"`
	GroupId      int       `json:"group_id"`
	Description  string    `json:"description"`
}

type UpdateNode struct{
	NodeName     string      `json:"node_name"`
	FlowPos      int         `json:"flow_pos"`
	GroupId      int         `json:"group_id"`
	Description  string      `json:"description"`
}