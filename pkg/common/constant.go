package common

const (
	GoChatServicePath = "/apis/v1/go-chat"
)

type DynamicType int64

func (d DynamicType) Int64() int64 {
	return int64(d)
}

const (
	DynamicTypeImage DynamicType = iota + 1 //图文动态
	DynamicTypeVideo                        //视频动态
)

type Visibility int64

func (v Visibility) Int64() int64 {
	return int64(v)
}

const (
	VisibilityPublic  Visibility = iota + 1 //公开的
	VisibilityPrivate                       //私有的
	VisibilityFriend                        //好友可见
	VisibilityFan                           //粉丝可见
)

type DynamicStatus int64

func (d DynamicStatus) Int64() int64 {
	return int64(d)
}

const (
	DynamicStatusPublished   DynamicStatus = iota + 1 //已发布
	DynamicStatusUnPublished                          //未发布
	DynamicStatusUnderReview                          //审核中
	DynamicStatusReviewPass                           //审核未通过
)
