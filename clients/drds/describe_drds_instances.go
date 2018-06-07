package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDrdsInstancesRequest struct {
	Type     string `position:"Query" name:"Type"`
	RegionId string `position:"Query" name:"RegionId"`
}

func (req *DescribeDrdsInstancesRequest) Invoke(client goaliyun.Client) (*DescribeDrdsInstancesResponse, error) {
	resp := &DescribeDrdsInstancesResponse{}
	err := client.Request("drds", "DescribeDrdsInstances", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDrdsInstancesResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeDrdsInstancesInstanceList
}

type DescribeDrdsInstancesInstance struct {
	DrdsInstanceId     goaliyun.String
	Type               goaliyun.String
	RegionId           goaliyun.String
	ZoneId             goaliyun.String
	Description        goaliyun.String
	NetworkType        goaliyun.String
	Status             goaliyun.String
	CreateTime         goaliyun.Integer
	Version            goaliyun.Integer
	VpcCloudInstanceId goaliyun.String
	Vips               DescribeDrdsInstancesVipList
}

type DescribeDrdsInstancesVip struct {
	IP        goaliyun.String
	Port      goaliyun.String
	Type      goaliyun.String
	VpcId     goaliyun.String
	VswitchId goaliyun.String
}

type DescribeDrdsInstancesInstanceList []DescribeDrdsInstancesInstance

func (list *DescribeDrdsInstancesInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstancesInstance)
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

type DescribeDrdsInstancesVipList []DescribeDrdsInstancesVip

func (list *DescribeDrdsInstancesVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstancesVip)
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
