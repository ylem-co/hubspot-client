package hubspotclient

type CreateTicketRequest struct {
	Properties CreateTicketRequestProperties `json:"properties"`
}

type CreateTicketRequestProperties struct {
	Pipeline         string `json:"hs_pipeline"`
	PipelineStage    string `json:"hs_pipeline_stage"`
	HsTicketPriority string `json:"hs_ticket_priority"`
	HsOwnerId        string `json:"hubspot_owner_id"`
	Subject          string `json:"subject"`
	Content          string `json:"content"`
}
