package ecs

import (
	"github.com/morlay/goaliyun"
)

type DescribeEventDetailRequest struct {
	EventId              string `position:"Query" name:"EventId"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeEventDetailRequest) Invoke(client goaliyun.Client) (*DescribeEventDetailResponse, error) {
	resp := &DescribeEventDetailResponse{}
	err := client.Request("ecs", "DescribeEventDetail", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeEventDetailResponse struct {
	RequestId     goaliyun.String
	ResourceId    goaliyun.String
	EventType     goaliyun.String
	EventCategory goaliyun.String
	Status        goaliyun.String
	SupportModify goaliyun.String
	PlanTime      goaliyun.String
	ExpireTime    goaliyun.String
	EventId       goaliyun.String
	StartTime     goaliyun.String
	EndTime       goaliyun.String
	EffectTime    goaliyun.String
	LimitTime     goaliyun.String
	Mark          goaliyun.String
}
