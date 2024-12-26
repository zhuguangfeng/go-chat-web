package common

const (
	GoChatServicePath = "/apis/v1/go-chat"
)

type DynamicType uint

func (d DynamicType) Uint() int64 {
	return int64(d)
}

const (
	DynamicTypeImage DynamicType = iota + 1 //图文动态
	DynamicTypeVideo                        //视频动态
)

type Visibility uint

func (v Visibility) Uint() int64 {
	return int64(v)
}

const (
	VisibilityPublic  Visibility = iota + 1 //公开的
	VisibilityPrivate                       //私有的
	VisibilityFriend                        //好友可见
	VisibilityFan                           //粉丝可见
)

type DynamicStatus uint

func (d DynamicStatus) Uint() int64 {
	return int64(d)
}

const (
	DynamicStatusPublished   DynamicStatus = iota + 1 //已发布
	DynamicStatusUnPublished                          //未发布
	DynamicStatusUnderReview                          //审核中
	DynamicStatusReviewPass                           //审核未通过
)

//type ActivitySignupStatus uint
//
//func (a ActivitySignupStatus) Uint() uint {
//	return uint(a)
//}
//
//const (
//	ActivitySignupStatusPendingReview ActivitySignupStatus = iota + 1
//	ActivitySignupStatusReviewPass
//)
