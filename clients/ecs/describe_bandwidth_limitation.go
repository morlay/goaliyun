package ecs

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeBandwidthLimitationRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceType         string `position:"Query" name:"InstanceType"`
	InstanceChargeType   string `position:"Query" name:"InstanceChargeType"`
	ResourceId           string `position:"Query" name:"ResourceId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OperationType        string `position:"Query" name:"OperationType"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SpotStrategy         string `position:"Query" name:"SpotStrategy"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeBandwidthLimitationRequest) Invoke(client goaliyun.Client) (*DescribeBandwidthLimitationResponse, error) {
	resp := &DescribeBandwidthLimitationResponse{}
	err := client.Request("ecs", "DescribeBandwidthLimitation", "2014-05-26", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeBandwidthLimitationResponse struct {
	RequestId  goaliyun.String
	Bandwidths DescribeBandwidthLimitationBandwidthList
}

type DescribeBandwidthLimitationBandwidth struct {
	InternetChargeType goaliyun.String
	Min                goaliyun.Integer
	Max                goaliyun.Integer
	Unit               goaliyun.String
}

type DescribeBandwidthLimitationBandwidthList []DescribeBandwidthLimitationBandwidth

func (list *DescribeBandwidthLimitationBandwidthList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeBandwidthLimitationBandwidth)
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
