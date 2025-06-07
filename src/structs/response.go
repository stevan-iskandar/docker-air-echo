package structs

type Response struct {
	Message string      `json:"message,omitempty"`
	Detail  string      `json:"detail,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
