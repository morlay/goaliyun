package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDrdsInstanceNetInfoForInnerRequest struct {
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDrdsInstanceNetInfoForInnerRequest) Invoke(client goaliyun.Client) (*DescribeDrdsInstanceNetInfoForInnerResponse, error) {
	resp := &DescribeDrdsInstanceNetInfoForInnerResponse{}
	err := client.Request("drds", "DescribeDrdsInstanceNetInfoForInner", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDrdsInstanceNetInfoForInnerResponse struct {
	RequestId      goaliyun.String
	Success        bool
	DrdsInstanceId goaliyun.String
	NetworkType    goaliyun.String
	NetInfos       DescribeDrdsInstanceNetInfoForInnerNetInfoList
}

type DescribeDrdsInstanceNetInfoForInnerNetInfo struct {
	IP       goaliyun.String
	Port     goaliyun.String
	Type     goaliyun.String
	IsForVpc bool
}

type DescribeDrdsInstanceNetInfoForInnerNetInfoList []DescribeDrdsInstanceNetInfoForInnerNetInfo

func (list *DescribeDrdsInstanceNetInfoForInnerNetInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDrdsInstanceNetInfoForInnerNetInfo)
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
