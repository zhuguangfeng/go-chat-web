package dao

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	dtoV1 "github.com/zhuguangfeng/go-chat/dto/v1"
	"github.com/zhuguangfeng/go-chat/model"
	"strconv"
)

type ActivityEsDao interface {
	InputActivity(ctx context.Context, activity model.ActivityEs) error
	SearchActivity(ctx context.Context, searchReq dtoV1.SearchActivityReq) ([]model.ActivityEs, error)
}

type OlivereActivityEsDao struct {
	esCli *elastic.Client
}

func NewActivityEsDao(esCli *elastic.Client) ActivityEsDao {
	return &OlivereActivityEsDao{
		esCli: esCli,
	}
}

func (dao *OlivereActivityEsDao) InputActivity(ctx context.Context, activity model.ActivityEs) error {
	_, err := dao.esCli.Index().Index(model.ActivityIndexName).Id(strconv.Itoa(int(activity.ID))).BodyJson(activity).Do(ctx)
	return err
}

func (dao *OlivereActivityEsDao) SearchActivity(ctx context.Context, searchReq dtoV1.SearchActivityReq) ([]model.ActivityEs, error) {
	titleQuery := elastic.NewMatchQuery("title", searchReq.SearchKey)
	descQuery := elastic.NewMatchQuery("desc", searchReq.SearchKey)
	//or 查询
	or := elastic.NewBoolQuery().Should(titleQuery, descQuery)

	//addressQuery := elastic.NewMatchQuery("address", searchReq.Address)
	//ageRestrictQuery := elastic.NewTermQuery("ageRestrict", searchReq.AgeRestrict)
	//genderRestrictQuery := elastic.NewTermQuery("genderRestrict", searchReq.GenderRestrict)
	//costRestrictQuery := elastic.NewTermQuery("costRestrict", searchReq.CostRestrict)
	//visibilityQuery := elastic.NewTermQuery("visibility", searchReq.Visibility)
	//categoryQuery := elastic.NewTermQuery("category", searchReq.Category)
	//startTimeQuery := elastic.NewRangeQuery("startTime").Lte(searchReq.StartTime)
	//endTimeQuery := elastic.NewRangeQuery("startTime").Gte(searchReq.EndTime)

	//and查询
	//and := elastic.NewBoolQuery().Must(addressQuery, ageRestrictQuery, genderRestrictQuery, costRestrictQuery, visibilityQuery, categoryQuery, startTimeQuery, endTimeQuery, or)

	resp, err := dao.esCli.Search(model.ActivityIndexName).Query(or).Do(ctx)

	if err != nil {
		return nil, err
	}

	var res []model.ActivityEs

	for _, hit := range resp.Hits.Hits {
		var art model.ActivityEs
		err := json.Unmarshal(hit.Source, &art)
		if err != nil {
			return nil, err
		}
		res = append(res, art)
	}
	return res, nil
}
