package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDrdsInstanceRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDrdsInstanceRequest) Invoke(client goaliyun.Client) (*DescribeDrdsInstanceResponse, error) {
	resp := &DescribeDrdsInstanceResponse{}
	err := client.Request("drds", "DescribeDrdsInstance", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDrdsInstanceResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeDrdsInstanceData
}

type DescribeDrdsInstanceData struct {
	DrdsInstanceId     goaliyun.String
	Type               goaliyun.String
	RegionId           goaliyun.String
	ZoneId             goaliyun.String
	Description        goaliyun.String
	NetworkType        goaliyun.String
	Status             goaliyun.String
	CreateTime         goaliyun.Integer
	Version            goaliyun.Integer
	Specification      goaliyun.String
	VpcCloudInstanceId goaliyun.String
	Vips               DescribeDrdsInstanceVipList
}

type DescribeDrdsInstanceVip struct {
	IP        goaliyun.String
	Port      goaliyun.String
	Type      goaliyun.String
	VpcId     goaliyun.String
	VswitchId goaliyun.String
}

type DescribeDrdsInstanceVipList []DescribeDrdsInstanceVip

func (list *DescribeDrdsInstanceVipList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstanceVip)
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
