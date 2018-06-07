package cbn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeCenAttachedChildInstancesRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	RegionId             string `position:"Query" name:"RegionId"`
}

func (req *DescribeCenAttachedChildInstancesRequest) Invoke(client goaliyun.Client) (*DescribeCenAttachedChildInstancesResponse, error) {
	resp := &DescribeCenAttachedChildInstancesResponse{}
	err := client.Request("cbn", "DescribeCenAttachedChildInstances", "2017-09-12", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeCenAttachedChildInstancesResponse struct {
	RequestId      goaliyun.String
	TotalCount     goaliyun.Integer
	PageNumber     goaliyun.Integer
	PageSize       goaliyun.Integer
	ChildInstances DescribeCenAttachedChildInstancesChildInstanceList
}

type DescribeCenAttachedChildInstancesChildInstance struct {
	CenId                 goaliyun.String
	ChildInstanceId       goaliyun.String
	ChildInstanceType     goaliyun.String
	ChildInstanceRegionId goaliyun.String
	ChildInstanceOwnerId  goaliyun.Integer
	Status                goaliyun.String
}

type DescribeCenAttachedChildInstancesChildInstanceList []DescribeCenAttachedChildInstancesChildInstance

func (list *DescribeCenAttachedChildInstancesChildInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCenAttachedChildInstancesChildInstance)
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
