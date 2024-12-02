package v1

type BaseSearchReq struct {
	PageNum   int      `json:"pageNum"`
	PageSize  int      `json:"pageSize"`
	Order     []string `json:"order"`
	SearchKey string   `json:"searchKey"`
}
