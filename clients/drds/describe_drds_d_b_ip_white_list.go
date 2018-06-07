package drds

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

type DescribeDrdsDBIpWhiteListRequest struct {
	DbName         string `position:"Query" name:"DbName"`
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	GroupName      string `position:"Query" name:"GroupName"`
	RegionId       string `position:"Query" name:"RegionId"`
}

func (req *DescribeDrdsDBIpWhiteListRequest) Invoke(client goaliyun.Client) (*DescribeDrdsDBIpWhiteListResponse, error) {
	resp := &DescribeDrdsDBIpWhiteListResponse{}
	err := client.Request("drds", "DescribeDrdsDBIpWhiteList", "2017-10-16", req.RegionId, "", "").Do(req, resp)
	return resp, err
}

type DescribeDrdsDBIpWhiteListResponse struct {
	RequestId goaliyun.String
	Success   bool
	Data      DescribeDrdsDBIpWhiteListData
}

type DescribeDrdsDBIpWhiteListData struct {
	IpWhiteList DescribeDrdsDBIpWhiteListIpWhiteListList
}

type DescribeDrdsDBIpWhiteListIpWhiteListList []goaliyun.String

func (list *DescribeDrdsDBIpWhiteListIpWhiteListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]goaliyun.String)
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
