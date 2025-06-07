package structs

type Pagination struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Sort     []Sort      `json:"sort,omitempty"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}
