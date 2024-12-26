package v1

type BaseDeleteReq struct {
	ID int64 `json:"id" dc:"活动id"`
}

type BaseListReq struct {
	PageNum   int      `json:"pageNum"`
	PageSize  int      `json:"pageSize"`
	Order     []string `json:"order"`
	SearchKey string   `json:"searchKey"`
}
