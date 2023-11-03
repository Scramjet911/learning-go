package stage

type FirstStageInput struct {
	Solution string `json:"sol"`
}

type StageResponse struct {
	Status string `json:"status"`
	Hint   string `json:"hint"`
	Code   string `json:"code"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
