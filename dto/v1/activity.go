package v1

type SearchActivityReq struct {
	BaseSearchReq
	AgeRestrict    uint   `json:"ageRestrict"`
	GenderRestrict uint   `json:"genderRestrict"`
	CostRestrict   uint   `json:"CostRestrict"`
	Visibility     uint   `json:"visibility"`
	Address        string `json:"address"`
	Category       uint   `json:"category"`
	StartTime      uint   `json:"startTime"`
	EndTime        uint   `json:"EndTime"`
	Status         uint   `json:"status"`
}
