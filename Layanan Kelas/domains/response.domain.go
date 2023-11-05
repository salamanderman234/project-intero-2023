package domain

type BasicResponse struct {
	Message string        `json:"message"`
	Datas   any           `json:"detail,omitempty"`
	Errors  []ErrorDetail `json:"errors,omitempty"`
}

type ErrorDetail struct {
	Field   string `json:"field"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
