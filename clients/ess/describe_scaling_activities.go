package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeScalingActivitiesRequest struct {
	ScalingActivityId9   string `position:"Query" name:"ScalingActivityId.9"`
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ScalingActivityId5   string `position:"Query" name:"ScalingActivityId.5"`
	ScalingActivityId6   string `position:"Query" name:"ScalingActivityId.6"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	ScalingActivityId7   string `position:"Query" name:"ScalingActivityId.7"`
	ScalingActivityId8   string `position:"Query" name:"ScalingActivityId.8"`
	ScalingActivityId1   string `position:"Query" name:"ScalingActivityId.1"`
	ScalingActivityId2   string `position:"Query" name:"ScalingActivityId.2"`
	ScalingActivityId3   string `position:"Query" name:"ScalingActivityId.3"`
	ScalingActivityId4   string `position:"Query" name:"ScalingActivityId.4"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	StatusCode           string `position:"Query" name:"StatusCode"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ScalingActivityId11  string `position:"Query" name:"ScalingActivityId.11"`
	ScalingActivityId10  string `position:"Query" name:"ScalingActivityId.10"`
	ScalingActivityId13  string `position:"Query" name:"ScalingActivityId.13"`
	ScalingActivityId12  string `position:"Query" name:"ScalingActivityId.12"`
	ScalingActivityId15  string `position:"Query" name:"ScalingActivityId.15"`
	ScalingActivityId14  string `position:"Query" name:"ScalingActivityId.14"`
	ScalingActivityId17  string `position:"Query" name:"ScalingActivityId.17"`
	ScalingActivityId16  string `position:"Query" name:"ScalingActivityId.16"`
	ScalingActivityId19  string `position:"Query" name:"ScalingActivityId.19"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingActivityId18  string `position:"Query" name:"ScalingActivityId.18"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	ScalingActivityId20  string `position:"Query" name:"ScalingActivityId.20"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeScalingActivitiesRequest) Invoke(client goaliyun.Client) (*DescribeScalingActivitiesResponse, error) {
	resp := &DescribeScalingActivitiesResponse{}
	err := client.Request("ess", "DescribeScalingActivities", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeScalingActivitiesResponse struct {
	TotalCount        goaliyun.Integer
	PageNumber        goaliyun.Integer
	PageSize          goaliyun.Integer
	RequestId         goaliyun.String
	ScalingActivities DescribeScalingActivitiesScalingActivityList
}

type DescribeScalingActivitiesScalingActivity struct {
	ScalingActivityId   goaliyun.String
	ScalingGroupId      goaliyun.String
	Description         goaliyun.String
	Cause               goaliyun.String
	StartTime           goaliyun.String
	EndTime             goaliyun.String
	Progress            goaliyun.Integer
	StatusCode          goaliyun.String
	StatusMessage       goaliyun.String
	TotalCapacity       goaliyun.String
	AttachedCapacity    goaliyun.String
	AutoCreatedCapacity goaliyun.String
}

type DescribeScalingActivitiesScalingActivityList []DescribeScalingActivitiesScalingActivity

func (list *DescribeScalingActivitiesScalingActivityList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeScalingActivitiesScalingActivity)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
