package dto

type HealthResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}
