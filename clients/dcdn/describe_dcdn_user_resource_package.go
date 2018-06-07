package dcdn

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDcdnUserResourcePackageRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	RegionId      string `position:"Query" name:"RegionId"`
}

func (req *DescribeDcdnUserResourcePackageRequest) Invoke(client goaliyun.Client) (*DescribeDcdnUserResourcePackageResponse, error) {
	resp := &DescribeDcdnUserResourcePackageResponse{}
	err := client.Request("dcdn", "DescribeDcdnUserResourcePackage", "2018-01-15", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDcdnUserResourcePackageResponse struct {
	RequestId            goaliyun.String
	ResourcePackageInfos DescribeDcdnUserResourcePackageResourcePackageInfoList
}

type DescribeDcdnUserResourcePackageResourcePackageInfo struct {
	CurrCapacity  goaliyun.String
	InitCapacity  goaliyun.String
	CommodityCode goaliyun.String
	DisplayName   goaliyun.String
	InstanceId    goaliyun.String
	Status        goaliyun.String
}

type DescribeDcdnUserResourcePackageResourcePackageInfoList []DescribeDcdnUserResourcePackageResourcePackageInfo

func (list *DescribeDcdnUserResourcePackageResourcePackageInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnUserResourcePackageResourcePackageInfo)
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
