package database

type PagedResult struct {
	TotalResults int         `json:"total-results"`
	PageIndex    int         `json:"page-index"`
	PageSize     int         `json:"page-size"`
	List         interface{} `json:"list"`
}
