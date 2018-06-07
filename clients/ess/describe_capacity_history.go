package ess

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCapacityHistoryRequest struct {
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ScalingGroupId       string `position:"Query" name:"ScalingGroupId"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	EndTime              string `position:"Query" name:"EndTime"`
	StartTime            string `position:"Query" name:"StartTime"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCapacityHistoryRequest) Invoke(client goaliyun.Client) (*DescribeCapacityHistoryResponse, error) {
	resp := &DescribeCapacityHistoryResponse{}
	err := client.Request("ess", "DescribeCapacityHistory", "2014-08-28", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCapacityHistoryResponse struct {
	TotalCount           goaliyun.Integer
	PageNumber           goaliyun.Integer
	PageSize             goaliyun.Integer
	CapacityHistoryItems DescribeCapacityHistoryCapacityHistoryModelList
}

type DescribeCapacityHistoryCapacityHistoryModel struct {
	ScalingGroupId      goaliyun.String
	TotalCapacity       goaliyun.Integer
	AttachedCapacity    goaliyun.Integer
	AutoCreatedCapacity goaliyun.Integer
	Timestamp           goaliyun.String
}

type DescribeCapacityHistoryCapacityHistoryModelList []DescribeCapacityHistoryCapacityHistoryModel

func (list *DescribeCapacityHistoryCapacityHistoryModelList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCapacityHistoryCapacityHistoryModel)
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
