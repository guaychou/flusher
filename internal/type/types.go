package types

type ErrorResponse struct{
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type HealthResponse struct{
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
	Version string `json:"version,omitempty"`
}

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}