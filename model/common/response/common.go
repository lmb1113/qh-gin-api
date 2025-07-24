package response

type PageResult struct {
	Extra    interface{} `json:"extra"`
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}
