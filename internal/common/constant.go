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

type ActivityStatus uint

func (a ActivityStatus) Uint() uint {
	return uint(a)
}

const (
	ActivityStatusPendingReview ActivityStatus = iota + 1 //待审核
	ActivityStatusReviewPass                              //审核失败
	ActivityStatusSignUp                                  //报名中 == 审核通过
	ActivityStatusCancel                                  //已取消
	ActivityStatusStart                                   //已开始
	ActivityStatusEnd                                     //已结束
)

type ReviewStatus uint

func (r ReviewStatus) Uint() uint {
	return uint(r)
}

const (
	ReviewStatusPendingReview ReviewStatus = iota + 1 //待审核
	ReviewStatusReviewCancel                          //审核取消
	ReviewStatusSuccess                               //审核通过
	ReviewStatusPass                                  //审核拒绝
)
