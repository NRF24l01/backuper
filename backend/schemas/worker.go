package schemas

type WorkerCreateRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}

type WorkerCreateResponse struct {
	WorkerUUID string `json:"worker_uuid"`
	Token      string `json:"worker_token"`
}