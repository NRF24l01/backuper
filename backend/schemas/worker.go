package schemas

type WorkerCreateRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}

type WorkerCreateResponse struct {
	WorkerUUID string `json:"worker_uuid"`
	Token      string `json:"worker_token"`
}

type WorkerStatus struct {
	WorkerUUID string `json:"worker_uuid"`
	Name       string `json:"name"`
	CreatedAt  int64  `json:"created_at"`
	LastOnline int64  `json:"last_online"`
}

type WorkerListResponse struct {
	Workers []WorkerStatus `json:"workers"`
}